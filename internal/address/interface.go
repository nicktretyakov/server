package address

import (
	"context"

	"github.com/google/uuid"

	"be/internal/datastore/address"
	"be/internal/model"
)

type IAddressService interface {
	AddAttachment(ctx context.Context, user model.User, attachment model.Attachment,
		typeAddress address.TypeAddress) (*model.Attachment, error)
	RemoveAttachment(ctx context.Context, user model.User, attachmentID uuid.UUID,
		typeAddress address.TypeAddress) error
	RenameAttachment(ctx context.Context, user model.User, attachmentID uuid.UUID, newFileName string,
		typeAddress address.TypeAddress) (*model.Attachment, error)
	AddLink(ctx context.Context, user model.User, link *model.Link, addressID uuid.UUID,
		typeAddress address.TypeAddress) error
	UpdateLink(ctx context.Context, user model.User, link *model.Link, addressID uuid.UUID,
		typeAddress address.TypeAddress) error
	RemoveLink(ctx context.Context, user model.User, link uuid.UUID, addressID uuid.UUID,
		typeAddress address.TypeAddress) error
	AddParticipant(ctx context.Context, user model.User, role string, portalCodeParticipant uint64,
		addressID uuid.UUID, typeAddress address.TypeAddress) (*model.User, string, error)
	UpdateParticipant(ctx context.Context, user model.User, role string, portalCodeParticipantBase, portalCodeParticipantChange uint64,
		addressID uuid.UUID, typeAddress address.TypeAddress) (*model.User, string, error)
	RemoveParticipant(ctx context.Context, user model.User, portalCodeParticipant uint64,
		addressID uuid.UUID, typeAddress address.TypeAddress) error
}
