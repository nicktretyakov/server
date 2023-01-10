package note

import (
	"context"

	bookingpb "be/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"be/internal/datastore/address"
	"be/internal/datastore/sorting"
	"be/internal/lib"
	"be/internal/model"
	"be/pkg/auth"
	"be/pkg/server/pbs"
)

func (s *Service) GetNoteSettings(ctx context.Context,
	_ *bookingpb.GetNoteSettingsRequest,
) (*bookingpb.GetNoteSettingsResponse, error) {
	user := auth.FromContext(ctx)

	settings, err := s.store.User().GetNoteSettingByUser(ctx, &user.ID)
	if err != nil {
		return nil, err
	}

	return &bookingpb.GetNoteSettingsResponse{
		IsEmailOn: settings.EmailOn,
		IsLifeOn:  settings.LifeOn,
	}, nil
}

func (s *Service) SetNoteSettings(ctx context.Context,
	in *bookingpb.SetNoteSettingsRequest,
) (*bookingpb.SetNoteSettingsResponse, error) {
	user := auth.FromContext(ctx)

	if in.GetType() == bookingpb.SetNoteSettingsRequest_EMAIL {
		if err := s.store.User().SetEmailNoteSetting(ctx, &user.ID, in.GetIsOn()); err != nil {
			return nil, err
		}
	} else {
		if err := s.store.User().SetLifeNoteSetting(ctx, &user.ID, in.GetIsOn()); err != nil {
			return nil, err
		}
	}

	return &bookingpb.SetNoteSettingsResponse{}, nil
}

func (s *Service) GetNotes(ctx context.Context, in *bookingpb.GetNotesRequest) (*bookingpb.GetNotesResponse, error) {
	user := auth.FromContext(ctx)

	status := model.Read

	if in.GetType() == bookingpb.STATUS_NOTE_NOT_READ {
		status = model.NotRead
	}

	notes, err := s.store.Address().SystemNoteList(
		ctx,
		user.ID,
		status,
		in.GetLimit(),
		in.GetOffset(),
		withSorting(in.GetSorting()),
	)
	if err != nil {
		return nil, err
	}

	return &bookingpb.GetNotesResponse{
		Notes: pbs.PbSystemsNotes(notes),
	}, nil
}

func (s *Service) GetNotesCount(
	ctx context.Context,
	_ *bookingpb.GetNotesCountRequest,
) (*bookingpb.GetNotesCountResponse, error) {
	user := auth.FromContext(ctx)

	count, err := s.store.Address().GetSystemNotesCount(ctx, user.ID)
	if err != nil {
		return nil, err
	}

	return &bookingpb.GetNotesCountResponse{Count: int64(count)}, nil
}

func (s *Service) ReadNotes(
	ctx context.Context,
	in *bookingpb.ReadNotesRequest,
) (*bookingpb.ReadNotesResponse, error) {
	noteIds, err := lib.ParseUUIDStrings(in.GetIds())
	if err != nil {
		return nil, status.New(codes.InvalidArgument, err.Error()).Err()
	}

	if err = s.store.Address().UpdateSystemNotes(ctx, noteIds); err != nil {
		return nil, err
	}

	return &bookingpb.ReadNotesResponse{}, nil
}

func withSorting(pbSorting *bookingpb.NoteSorting) (s sorting.Sorting) {
	switch pbSorting.GetType() {
	case bookingpb.NoteSorting_BY_CREATED_AT:
		s = address.NewSortingByCreatedAt(pbSorting.GetAsc())
	case bookingpb.NoteSorting_BY_READ_AT:
		s = address.NewSortingByReadAt(pbSorting.GetAsc())
	default:
		s = address.NewSortingByReadAt(false)
	}

	return s
}
