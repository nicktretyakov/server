package room

import (
	"context"
	"time"

	"github.com/google/uuid"
	bookingpb "be/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"

	"be/internal/model"
	"be/pkg/auth"
	"be/pkg/server/pbs"
)

func (s Service) AddRoomEquipments(ctx context.Context,
	req *bookingpb.AddRoomEquipmentsRequest,
) (*bookingpb.AddRoomEquipmentsResponse, error) {
	equipmentsFromPB, err := getEquipments(req.GetEquipments())
	if err != nil {
		return nil, err
	}

	if len(equipmentsFromPB) > 0 {
		equipments, err := s.store.Room().AddEquipments(ctx, equipmentsFromPB)
		if err != nil {
			return nil, err
		}

		equipmentIDs := make([]string, 0, len(equipments))

		for _, equipment := range equipments {
			equipmentIDs = append(equipmentIDs, equipment.ID.String())
		}

		return &bookingpb.AddRoomEquipmentsResponse{Uuids: equipmentIDs}, nil
	}

	return &bookingpb.AddRoomEquipmentsResponse{}, nil
}

func (s Service) UpdateRoomEquipment(
	ctx context.Context,
	req *bookingpb.UpdateRoomEquipmentRequest,
) (*bookingpb.UpdateRoomEquipmentResponse, error) {
	updateValidEquipment, err := validateUpdateEquipmentRequest(req)
	if err != nil {
		return nil, err
	}

	canUpdateEquipment, err := s.canMakeOperationsWithEquipment(ctx, updateValidEquipment.ID)
	if err != nil {
		return nil, err
	}

	if !canUpdateEquipment {
		return nil, err
	}

	equipment, err := s.store.Room().UpdateEquipment(ctx, *updateValidEquipment)
	if err != nil {
		return nil, err
	}

	equipment.CreatedAt = updateValidEquipment.CreatedAt

	return &bookingpb.UpdateRoomEquipmentResponse{Equipment: pbs.PbRoomEquipment(*equipment)}, nil
}

func (s Service) RemoveRoomEquipment(ctx context.Context, req *bookingpb.RemoveRoomEquipmentRequest) (*emptypb.Empty, error) {
	equipmentID, err := uuid.Parse(req.GetUuid())
	if err != nil {
		return nil, status.New(codes.InvalidArgument, "invalid equipment_id").Err()
	}

	canUpdateEquipment, err := s.canMakeOperationsWithEquipment(ctx, equipmentID)
	if err != nil {
		return nil, err
	}

	if !canUpdateEquipment {
		return nil, err
	}

	if err = s.store.Room().DeleteEquipment(ctx, equipmentID); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func getEquipments(equipments []*bookingpb.Equipment) ([]model.Equipment, error) {
	modelEquipments := make([]model.Equipment, 0, len(equipments))

	for _, pbEquipment := range equipments {
		modelEquipment, err := getModelEquipment(pbEquipment.GetTitle(),
			pbEquipment.GetDescription(),
			pbEquipment.GetPlanValue(),
			pbEquipment.GetFactValue(),
			pbEquipment.GetTimeline())
		if err != nil {
			return nil, err
		}

		modelEquipments = append(modelEquipments, *modelEquipment)
	}

	return modelEquipments, nil
}

func getTimelineModel(timeline *bookingpb.Timeline) (*model.Timeline, error) {
	start, err := time.Parse(pbs.TimeLayout, timeline.GetStart())
	if err != nil {
		return nil, err
	}

	end, err := time.Parse(pbs.TimeLayout, timeline.GetEnd())
	if err != nil {
		return nil, err
	}

	return &model.Timeline{
		StartAt: start,
		EndAt:   end,
	}, nil
}

func validateUpdateEquipmentRequest(req *bookingpb.UpdateRoomEquipmentRequest) (*model.Equipment, error) {
	equipment, err := getModelEquipment(req.GetEquipment().GetTitle(),
		req.GetEquipment().GetDescription(),
		req.GetEquipment().GetPlanValue(),
		req.GetEquipment().GetFactValue(),
		req.GetEquipment().GetTimeline())
	if err != nil {
		return nil, err
	}

	equipmentID, err := uuid.Parse(req.GetEquipment().GetUuid())
	if err != nil {
		return nil, status.New(codes.InvalidArgument, "invalid equipment_id").Err()
	}

	equipment.ID = equipmentID

	return equipment, nil
}

func getModelEquipment(title, description string, planValue, factValue float32, pbTimeline *bookingpb.Timeline) (*model.Equipment, error) {
	timeline, err := getTimelineModel(pbTimeline)
	if err != nil {
		return nil, status.New(codes.InvalidArgument, err.Error()).Err()
	}

	if !timeline.IsValid() {
		return nil, status.New(codes.InvalidArgument, "invalid timeline").Err()
	}

	return &model.Equipment{
		Title:       title,
		Description: description,
		PlanValue:   planValue,
		FactValue:   factValue,
		Timeline:    *timeline,
	}, nil
}

func (s Service) canMakeOperationsWithEquipment(ctx context.Context, equipmentID uuid.UUID) (bool, error) {
	met, err := s.store.Room().FindEquipmentByID(ctx, equipmentID)
	if err != nil {
		return false, err
	}

	room, err := s.store.Room().FindByID(ctx, met.RoomID)
	if err != nil {
		return false, err
	}

	if !room.CanUpdateEquipment(*auth.FromContext(ctx)) {
		return false, status.New(codes.PermissionDenied, "permission denied").Err()
	}

	return true, nil
}
