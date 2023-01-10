package room

import (
	"context"
	"time"
	"unicode/utf8"

	"github.com/google/uuid"
	bookingpb "be/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"be/internal/model"
	"be/pkg/auth"
)

const (
	minTitleLen       = 3
	minDescriptionLen = 3
)

func (s Service) CreateInitialRoom(ctx context.Context,
	req *bookingpb.CreateInitialRoomRequest,
) (*bookingpb.CreateInitialRoomResponse, error) {
	author := auth.FromContext(ctx)

	newRoom, equipmentIDs, slotIDs, err := validateCreateDraftRequest(req)
	if err != nil {
		return nil, err
	}

	newRoom.Author = author
	ownerPortalCode, assigneePortalCode := uint64(req.GetPortalCodeOwner()), uint64(req.GetPortalCodeAssignee())

	createdRoom, err := s.roomService.Create(ctx, *newRoom, ownerPortalCode, assigneePortalCode, equipmentIDs, slotIDs)
	if err != nil {
		return nil, err
	}

	pbRoom, err := s.roomFromDB(ctx, createdRoom.ID)
	if err != nil {
		return nil, err
	}

	return &bookingpb.CreateInitialRoomResponse{
		Room: pbRoom,
	}, nil
}

type iCreateInitialRoomRequest interface {
	GetTitle() string
	GetDescription() string
	GetTargetAudience() string
	GetPortalCodeOwner() uint32
	GetPortalCodeAssignee() uint32
	GetCreationDate() string
	GetSlots() []string
	GetEquipments() []string
}

//nolint:gocognit,gocyclo,cyclop
func validateCreateDraftRequest(request iCreateInitialRoomRequest) (*model.Room, []uuid.UUID, []uuid.UUID, error) {
	if len(request.GetTitle()) < minTitleLen {
		return nil, nil, nil, status.New(codes.InvalidArgument, "invalid title").Err()
	}

	if utf8.RuneCountInString(request.GetDescription()) < minDescriptionLen {
		return nil, nil, nil, status.New(codes.InvalidArgument, "invalid description").Err()
	}

	if request.GetPortalCodeOwner() == 0 {
		return nil, nil, nil, status.New(codes.InvalidArgument, "invalid portal code owner").Err()
	}

	if request.GetPortalCodeAssignee() == 0 {
		return nil, nil, nil, status.New(codes.InvalidArgument, "invalid portal code assignee").Err()
	}

	room := model.Room{
		Title:          request.GetTitle(),
		Description:    request.GetDescription(),
		Status:         model.InitialAddressStatus,
		TargetAudience: request.GetTargetAudience(),
	}

	if request.GetCreationDate() != "" {
		creationDate, err := time.Parse(time.RFC3339, request.GetCreationDate())
		if err != nil {
			return nil, nil, nil, status.Error(codes.InvalidArgument, "invalid creation date format")
		}

		room.CreationDate = creationDate
	} else {
		return nil, nil, nil, status.Error(codes.InvalidArgument, "invalid creation date")
	}

	var (
		equipmentIDs, slotIDs []uuid.UUID
		err                  error
	)

	if request.GetSlots() != nil && len(request.GetSlots()) > 0 {
		slotIDs, err = getIDs(request.GetSlots())
		if err != nil {
			return nil, nil, nil, status.New(codes.InvalidArgument, "invalid slot id").Err()
		}
	}

	if request.GetEquipments() != nil && len(request.GetEquipments()) > 0 {
		equipmentIDs, err = getIDs(request.GetEquipments())
		if err != nil {
			return nil, nil, nil, status.New(codes.InvalidArgument, "invalid equipment id").Err()
		}
	}

	return &room, equipmentIDs, slotIDs, nil
}

func getIDs(ids []string) ([]uuid.UUID, error) {
	UUIDs := make([]uuid.UUID, 0, len(ids))

	for _, id := range ids {
		UUID, err := uuid.Parse(id)
		if err != nil {
			return nil, err
		}

		UUIDs = append(UUIDs, UUID)
	}

	return UUIDs, nil
}
