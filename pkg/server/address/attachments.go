package address

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

func (s Service) AddAttachment(ctx context.Context, req *bookingpb.CreateAttachmentRequest) (*bookingpb.CreateAttachmentResponse, error) {
	addressID, err := uuid.Parse(req.GetAddressID())
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "invalid address_id")
	}

	file, err := s.fileLoader.DownloadFile(ctx, req.GetAttachmentURL())
	if err != nil {
		return nil, err
	}

	user := auth.FromContext(ctx)

	attachment := model.Attachment{
		AddressID: addressID,
		Filename:       req.FileName,
		Mime:           req.GetMime(),
		Size:           file.Size(),
	}

	attachment.SetSource(file.Source())

	createdAttachment, err := s.addressService.AddAttachment(ctx, *user, attachment, getTypeAddress(req.GetAddressType()))
	if err != nil {
		return nil, addressErrorToStatus(err)
	}

	return &bookingpb.CreateAttachmentResponse{Attachment: pbs.Attachment(*createdAttachment)}, nil
}

func (s Service) RemoveAttachment(ctx context.Context, request *bookingpb.RemoveAttachmentRequest) (*emptypb.Empty, error) {
	attachmentID, err := uuid.Parse(request.GetAttachmentId())
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "invalid attachment_id")
	}

	user := auth.FromContext(ctx)
	if err = s.addressService.RemoveAttachment(ctx, *user, attachmentID, getTypeAddress(request.GetAddressType())); err != nil {
		return nil, addressErrorToStatus(err)
	}

	return &emptypb.Empty{}, err
}

func (s Service) RenameAttachment(ctx context.Context,
	request *bookingpb.RenameAttachmentRequest,
) (*bookingpb.CreateAttachmentResponse, error) {
	attachmentID, err := uuid.Parse(request.GetAttachmentId())
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "invalid attachment_id")
	}

	user := auth.FromContext(ctx)

	updatedAttachment, err := s.addressService.
		RenameAttachment(
			ctx,
			*user,
			attachmentID,
			request.GetFileName(),
			getTypeAddress(request.GetAddressType()),
		)
	if err != nil {
		return nil, addressErrorToStatus(err)
	}

	return &bookingpb.CreateAttachmentResponse{Attachment: pbs.Attachment(*updatedAttachment)}, nil
}
