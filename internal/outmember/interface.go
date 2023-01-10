package outmember

import (
	"context"

	"github.com/google/uuid"

	addressStore "be/internal/datastore/address"
	"be/internal/model"
)

type IOutmemberService interface {
	NewInitialOutmember(ctx context.Context, user model.User,
		addressType addressStore.TypeAddress, outmember model.Outmember) error
	NewAcceptanceOutmember(ctx context.Context, user model.User,
		addressType addressStore.TypeAddress, outmember model.Outmember) error
	RegisterFinalOutmember(ctx context.Context, user model.User, finalReportID uuid.UUID, outmember model.Outmember) (*model.Booking, error)
	AcceptFinalOutmember(ctx context.Context, user model.User, finalReportID uuid.UUID, outmember model.Outmember) (*model.Booking, error)
}
