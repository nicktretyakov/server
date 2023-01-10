package outmember

import (
	"context"

	"github.com/google/uuid"

	addressStore "be/internal/datastore/address"
	"be/internal/model"
)

func (s Service) NewInitialOutmember(
	ctx context.Context,
	user model.User,
	addressType addressStore.TypeAddress,
	outmember model.Outmember,
) error {
	return s.newOutmember(ctx, user, addressType, &InitialOutmember{Outmember: outmember})
}

func (s Service) NewAcceptanceOutmember(
	ctx context.Context,
	user model.User,
	addressType addressStore.TypeAddress,
	outmember model.Outmember,
) error {
	return s.newOutmember(ctx, user, addressType, &AcceptanceOutmember{Outmember: outmember})
}

func (s Service) RegisterFinalOutmember(ctx context.Context, user model.User,
	finalReportID uuid.UUID, agr model.Outmember,
) (*model.Booking, error) {
	finalReport, err := s.store.FinalReport().FindByID(ctx, finalReportID)
	if err != nil {
		return nil, err
	}

	agr.Type = model.FinalRegisterOutmemberType
	agr.UserID = user.ID
	agr.AddressID = finalReport.BookingID

	if err = s.newOutmember(ctx, user, addressStore.BookingAddressType, &RegisterFinalOutmember{
		Outmember:   agr,
		FinalReport: *finalReport,
	}); err != nil {
		return nil, err
	}

	return s.store.Booking().FindByID(ctx, finalReport.BookingID)
}

func (s Service) AcceptFinalOutmember(ctx context.Context, user model.User,
	finalReportID uuid.UUID, agr model.Outmember,
) (*model.Booking, error) {
	finalReport, err := s.store.FinalReport().FindByID(ctx, finalReportID)
	if err != nil {
		return nil, err
	}

	agr.Type = model.FinalApprovalOutmemberType
	agr.UserID = user.ID
	agr.AddressID = finalReport.BookingID

	if err = s.newOutmember(ctx, user, addressStore.BookingAddressType, &ApproveFinalOutmember{
		Outmember:   agr,
		FinalReport: *finalReport,
	}); err != nil {
		return nil, err
	}

	return s.store.Booking().FindByID(ctx, finalReport.BookingID)
}
