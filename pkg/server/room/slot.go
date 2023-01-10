package room

import (
	"context"

	"github.com/google/uuid"
	bookingpb "be/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"

	"be/internal/model"
	"be/pkg/auth"
	"be/pkg/server/pbs"
)

func (s Service) AddRoomSlots(ctx context.Context,
	req *bookingpb.AddRoomSlotsRequest,
) (*bookingpb.AddRoomSlotsResponse, error) {
	slotsFromPB, err := getSlots(req.GetSlots())
	if err != nil {
		return nil, err
	}

	if len(slotsFromPB) > 0 {
		slots, err := s.store.Room().AddSlots(ctx, slotsFromPB)
		if err != nil {
			return nil, err
		}

		slotIDs := make([]string, 0, len(slots))

		for _, slot := range slots {
			slotIDs = append(slotIDs, slot.ID.String())
		}

		return &bookingpb.AddRoomSlotsResponse{Uuids: slotIDs}, nil
	}

	return &bookingpb.AddRoomSlotsResponse{}, nil
}

func (s Service) UpdateRoomSlot(
	ctx context.Context,
	req *bookingpb.UpdateRoomSlotRequest,
) (*bookingpb.UpdateRoomSlotResponse, error) {
	updateValidSlot, err := validUpdateSlotRequest(req)
	if err != nil {
		return nil, err
	}

	canUpdateSlot, err := s.canMakeOperationWithSlot(ctx, updateValidSlot.ID)
	if err != nil {
		return nil, err
	}

	if !canUpdateSlot {
		return nil, err
	}

	slot, err := s.store.Room().UpdateSlot(ctx, *updateValidSlot)
	if err != nil {
		return nil, err
	}

	slot.CreatedAt = updateValidSlot.CreatedAt

	return &bookingpb.UpdateRoomSlotResponse{Slot: pbs.PbRoomSlot(*slot)}, nil
}

func (s Service) RemoveRoomSlot(ctx context.Context, req *bookingpb.RemoveRoomSlotRequest) (*emptypb.Empty, error) {
	slotID, err := uuid.Parse(req.GetUuid())
	if err != nil {
		return nil, status.New(codes.InvalidArgument, "invalid slot_id").Err()
	}

	canUpdateSlot, err := s.canMakeOperationWithSlot(ctx, slotID)
	if err != nil {
		return nil, err
	}

	if !canUpdateSlot {
		return nil, err
	}

	if err = s.store.Room().DeleteSlot(ctx, slotID); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func getSlots(slots []*bookingpb.Slot) ([]model.Slot, error) {
	modelSlots := make([]model.Slot, 0, len(slots))

	for _, pbSlot := range slots {
		if pbSlot.GetPlanSlot() == nil {
			return nil, status.New(codes.InvalidArgument, "invalid plan slot").Err()
		}

		modelSlot, err := getModelSlot(pbSlot.GetPlanSlot().GetUnits(),
			pbSlot.GetFactSlot().GetUnits(),
			pbSlot.GetPlanSlot().GetFragments(),
			pbSlot.GetFactSlot().GetFragments(),
			pbSlot.GetTimeline(),
		)
		if err != nil {
			return nil, err
		}

		modelSlots = append(modelSlots, *modelSlot)
	}

	return modelSlots, nil
}

func validUpdateSlotRequest(req *bookingpb.UpdateRoomSlotRequest) (*model.Slot, error) {
	slot, err := getModelSlot(req.GetSlot().GetPlanSlot().GetUnits(),
		req.GetSlot().GetFactSlot().GetUnits(),
		req.GetSlot().GetPlanSlot().GetFragments(),
		req.GetSlot().GetFactSlot().GetFragments(),
		req.GetSlot().GetTimeline(),
	)
	if err != nil {
		return nil, err
	}

	slotID, err := uuid.Parse(req.GetSlot().GetUuid())
	if err != nil {
		return nil, status.New(codes.InvalidArgument, "invalid slot_id").Err()
	}

	slot.ID = slotID

	return slot, nil
}

func getModelSlot(
	unitsPlanSlot, unitsFactSlot uint64,
	fragmentsPlanSlot, fragmentsFactSlot uint32,
	pbTimeline *bookingpb.Timeline,
) (*model.Slot, error) {
	timeline, err := getTimelineModel(pbTimeline)
	if err != nil {
		return nil, status.New(codes.InvalidArgument, err.Error()).Err()
	}

	if !timeline.IsValid() {
		return nil, status.New(codes.InvalidArgument, "invalid timeline").Err()
	}

	return &model.Slot{
		Timeline:   *timeline,
		PlanSlot: model.NewNotificationUnitsAndFragments(unitsPlanSlot, fragmentsPlanSlot),
		FactSlot: model.NewNotificationUnitsAndFragments(unitsFactSlot, fragmentsFactSlot),
	}, nil
}

func (s Service) canMakeOperationWithSlot(ctx context.Context, slotID uuid.UUID) (bool, error) {
	slot, err := s.store.Room().FindSlotByID(ctx, slotID)
	if err != nil {
		return false, err
	}

	room, err := s.store.Room().FindByID(ctx, slot.RoomID)
	if err != nil {
		return false, err
	}

	if !room.CanUpdateSlot(*auth.FromContext(ctx)) {
		return false, status.New(codes.PermissionDenied, "permission denied").Err()
	}

	return true, nil
}
