package report

import (
	"context"
	"strings"
	"time"
	"unicode/utf8"

	bookingpb "be/proto"

	"github.com/google/uuid"
	"github.com/pkg/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"be/internal/lib"
	"be/internal/model"
	"be/pkg/auth"
	"be/pkg/server/pbs"
	"be/pkg/server/stage"
	"be/pkg/server/validators"
)

func (s Service) SendBookingReport(
	ctx context.Context,
	req *bookingpb.SendBookingReportRequest,
) (*bookingpb.SendBookingReportResponse, error) {
	reportID, err := pbs.ParseID(req)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "invalid report_id")
	}

	var user *model.User

	if user = auth.FromContext(ctx); user == nil {
		return nil, status.Error(codes.Unauthenticated, "auth required")
	}

	if len(strings.TrimSpace(req.GetEvents())) == 0 {
		return nil, status.New(codes.InvalidArgument, "events is required field").Err()
	}

	if utf8.RuneCountInString(strings.TrimSpace(req.GetEvents())) < stage.MinTitleLength {
		return nil, status.Newf(codes.InvalidArgument, "invalid events (min. length %d characters)", stage.MinTitleLength).Err()
	}

	slot, err := validators.Notification(req.GetSlot())
	if err != nil {
		return nil, status.Newf(codes.InvalidArgument, "slot invalid: %s", err.Error()).Err()
	}

	endAt, err := validators.Time(req.GetEndDate())
	if err != nil {
		return nil, status.Newf(codes.InvalidArgument, "time invalid: %s", err.Error()).Err()
	}

	rep, err := s.bookingService.SendReport(ctx, *user, model.ReportBooking{
		ID:      reportID,
		Slot:    lib.Notification(slot),
		EndAt:   lib.Time(endAt),
		Events:  lib.String(req.GetEvents()),
		Reasons: lib.String(req.GetReasons()),
		Comment: lib.String(req.GetComment()),
	})
	if err != nil {
		return nil, err
	}

	book, err := s.store.Booking().FindByID(ctx, rep.BookingID)
	if err != nil {
		return nil, err
	}

	go s.notificator.CreateBookingNote(user, model.SentReportNotify, []*model.Booking{book}, rep.Period.Time())

	return &bookingpb.SendBookingReportResponse{
		Report: pbs.BookingReport(*rep),
	}, nil
}

func (s Service) SendRoomReport(ctx context.Context, req *bookingpb.SendRoomReportRequest) (*bookingpb.ReportRoomResponse, error) {
	valRequest, err := validateSendReportRequest(req)
	if err != nil {
		return nil, err
	}

	user := auth.FromContext(ctx)

	room, err := s.store.Room().FindByID(ctx, valRequest.RoomID)
	if err != nil {
		return nil, err
	}

	if !room.CanSendReport(*user) {
		return nil, status.New(codes.PermissionDenied, "permission denied").Err()
	}

	slots, err := s.store.Room().FindSlotsByID(ctx, valRequest.SlotsID)
	if err != nil {
		return nil, err
	}

	equipments, err := s.store.Room().FindEquipmentsByID(ctx, valRequest.EquipmentsID)
	if err != nil {
		return nil, err
	}

	releases, err := s.store.Room().FindReleasesByID(ctx, valRequest.ReleasesID)
	if err != nil {
		return nil, err
	}

	report := model.ReportRoom{
		RoomID:     valRequest.RoomID,
		Timeline:   valRequest.Timeline,
		Comment:    &valRequest.Comment,
		Slots:      slots,
		Equipments: equipments,
		Releases:   releases,
	}

	if err := checkTimelineConsistency(report); err != nil {
		return nil, status.Newf(codes.InvalidArgument, "invalid timeline: %s", err.Error()).Err()
	}

	updReport, err := s.store.Report().CreateReportRoom(ctx, report)
	if err != nil {
		return nil, err
	}

	return &bookingpb.ReportRoomResponse{
		RoomReport: pbs.RoomReport(*updReport),
	}, nil
}

type validateRequest struct {
	RoomID       uuid.UUID
	Timeline     model.Timeline
	Comment      string
	SlotsID      []uuid.UUID
	EquipmentsID []uuid.UUID
	ReleasesID   []uuid.UUID
}

func validateSendReportRequest(req *bookingpb.SendRoomReportRequest) (*validateRequest, error) {
	roomID, err := uuid.Parse(req.GetRoomID())
	if err != nil {
		return nil, status.New(codes.InvalidArgument, "invalid room id").Err()
	}

	start, err := time.Parse(pbs.TimeLayout, req.GetTimeline().GetStart())
	if err != nil {
		return nil, status.New(codes.InvalidArgument, "invalid timeline start").Err()
	}

	end, err := time.Parse(pbs.TimeLayout, req.GetTimeline().GetEnd())
	if err != nil {
		return nil, status.New(codes.InvalidArgument, "invalid timeline end").Err()
	}

	slotIDs, err := lib.ParseUUIDStrings(req.GetSlotsID())
	if err != nil {
		return nil, status.New(codes.InvalidArgument, "invalid slot id").Err()
	}

	equipmentIDs, err := lib.ParseUUIDStrings(req.GetEquipmentsID())
	if err != nil {
		return nil, status.New(codes.InvalidArgument, "invalid equipment id").Err()
	}

	releaseIDs, err := lib.ParseUUIDStrings(req.GetReleasesID())
	if err != nil {
		return nil, status.New(codes.InvalidArgument, "invalid release id").Err()
	}

	return &validateRequest{
		RoomID:       roomID,
		Comment:      req.Comment,
		SlotsID:      slotIDs,
		EquipmentsID: equipmentIDs,
		ReleasesID:   releaseIDs,
		Timeline: model.Timeline{
			StartAt: start,
			EndAt:   end,
		},
	}, nil
}

func checkTimelineConsistency(report model.ReportRoom) error {
	if err := checkSlotsTimeline(report.Slots, report.Timeline); err != nil {
		return err
	}

	if err := checkEquipmentsTimeline(report.Equipments, report.Timeline); err != nil {
		return err
	}

	if err := checkReleasesTimeline(report.Releases, report.Timeline); err != nil {
		return err
	}

	return nil
}

func checkSlotsTimeline(slots []model.Slot, timeline model.Timeline) error {
	for _, slot := range slots {
		if slot.Timeline.EndAt.Before(timeline.StartAt) || slot.Timeline.StartAt.After(timeline.EndAt) {
			return errors.New("invalid slot timeline")
		}
	}

	return nil
}

func checkEquipmentsTimeline(equipments []model.Equipment, timeline model.Timeline) error {
	for _, equipment := range equipments {
		if equipment.Timeline.EndAt.Before(timeline.StartAt) || equipment.Timeline.StartAt.After(timeline.EndAt) {
			return errors.New("invalid equipment timeline")
		}
	}

	return nil
}

func checkReleasesTimeline(releases []model.Release, timeline model.Timeline) error {
	for _, release := range releases {
		if release.Date.Before(timeline.StartAt) || release.Date.After(timeline.EndAt) {
			return errors.New("invalid release date")
		}
	}

	return nil
}

func (s Service) GetRoomReport(ctx context.Context, req *bookingpb.GetRoomReportRequest) (*bookingpb.ReportRoomResponse, error) {
	reportID, err := uuid.Parse(req.GetReportID())
	if err != nil {
		return nil, status.New(codes.InvalidArgument, "invalid report id").Err()
	}

	report, err := s.store.Report().FindRoomReportByID(ctx, reportID)
	if err != nil {
		return nil, err
	}

	return &bookingpb.ReportRoomResponse{
		RoomReport: pbs.RoomReport(*report),
	}, nil
}
