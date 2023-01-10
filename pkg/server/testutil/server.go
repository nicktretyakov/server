package testutil

import (
	"context"
	"net"
	//"testing"

	//"github.com/golang/mock/gomock"
	//"github.com/rs/zerolog"
	//"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	//"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/test/bufconn"

	//"be/internal/lib"
	//"be/internal/model"
	//"be/pkg/server/server"
	//"be/test/mocks"
)

const (
	bufferSize = 1024 * 1024
)

// type MockGRPC struct {
// 	BookingStoreMock        *mocks.MockIBookingStore
// 	RoomStoreMock        *mocks.MockIRoomStore
// 	ReportStoreMock         *mocks.MockIReportStore
// 	UserStoreMock           *mocks.MockIUserStore
// 	StageStoreMock          *mocks.MockIStageStore
// 	ClientContextFn         func(ctx context.Context) context.Context
// 	ClientContextWithUserFn func(ctx context.Context) (context.Context, *model.User)
// 	Session                 model.Session
// 	User                    model.User
// 	BookingServiceMock      *mocks.MockIBookingService
// 	RoomServiceMock      *mocks.MockIRoomService
// 	AddressServiceMock *mocks.MockIAddressService
// 	FileLoaderMock          *mocks.MockIFileLoader

// 	authMock *mocks.MockIAuth
// 	dsMock   *mocks.MockIDatastore

// 	ctrl *gomock.Controller
// }

// func (m *MockGRPC) Finish() {
// 	m.ctrl.Finish()
// }

// func NewMockGRPCServer(ctx context.Context, t *testing.T) (*grpc.ClientConn, *MockGRPC) {
// 	t.Helper()

// 	m := createMockGRPC(t)

// 	user := &model.User{ID: lib.UUID(), Role: model.Admin}
// 	sess := model.Session{
// 		ID:     lib.UUID(),
// 		UserID: user.ID,
// 	}

// 	m.authMock.EXPECT().ParseToken(gomock.Any(), "jwt_token").Return(user, &sess, nil)
// 	m.ClientContextFn = func(ctx context.Context) context.Context {
// 		return metadata.NewOutgoingContext(ctx, metadata.Pairs("authorization", "bearer jwt_token"))
// 	}
// 	m.ClientContextWithUserFn = func(ctx context.Context) (context.Context, *model.User) {
// 		return metadata.NewOutgoingContext(ctx, metadata.Pairs("authorization", "bearer jwt_token")), user
// 	}
// 	m.Session = sess
// 	m.User = *user

// 	s := server.New(server.Config{Secret: "example"}, server.Opts{
// 		AuthService:         m.authMock,
// 		Store:               m.dsMock,
// 		Logger:              zerolog.Logger{},
// 		BookingService:      m.BookingServiceMock,
// 		RoomService:      m.RoomServiceMock,
// 		AddressService: m.AddressServiceMock,
// 		FileLoader:          m.FileLoaderMock,
// 	})
// 	lis := bufconn.Listen(bufferSize)
// 	require.NoError(t, server.RunListener(s, lis))

// 	conn, err := Dial(ctx, lis)
// 	require.NoError(t, err)

// 	return conn, m
// }

// nolint: staticcheck
func Dial(ctx context.Context, lis *bufconn.Listener) (*grpc.ClientConn, error) {
	dial := func(context.Context, string) (net.Conn, error) { return lis.Dial() }

	return grpc.DialContext(ctx, "localhost:0", grpc.WithContextDialer(dial), grpc.WithInsecure())
}

// func createMockGRPC(t *testing.T) *MockGRPC {
// 	t.Helper()

// 	ctrl := gomock.NewController(t)

// 	dsMock := mocks.NewMockIDatastore(ctrl)
// 	bookingMock := mocks.NewMockIBookingStore(ctrl)
// 	roomMock := mocks.NewMockIRoomStore(ctrl)
// 	userMock := mocks.NewMockIUserStore(ctrl)
// 	stageStoreMock := mocks.NewMockIStageStore(ctrl)
// 	authMock := mocks.NewMockIAuth(ctrl)
// 	bookingServiceMock := mocks.NewMockIBookingService(ctrl)
// 	roomServiceMock := mocks.NewMockIRoomService(ctrl)
// 	addressServiceMock := mocks.NewMockIAddressService(ctrl)
// 	fileLoaderMock := mocks.NewMockIFileLoader(ctrl)
// 	reportMock := mocks.NewMockIReportStore(ctrl)

// 	dsMock.EXPECT().Booking().Return(bookingMock).AnyTimes()
// 	dsMock.EXPECT().Room().Return(roomMock).AnyTimes()
// 	dsMock.EXPECT().Report().Return(reportMock).AnyTimes()
// 	dsMock.EXPECT().User().Return(userMock).AnyTimes()
// 	dsMock.EXPECT().Stage().Return(stageStoreMock).AnyTimes()

// 	return &MockGRPC{
// 		BookingStoreMock:        bookingMock,
// 		RoomStoreMock:        roomMock,
// 		UserStoreMock:           userMock,
// 		StageStoreMock:          stageStoreMock,
// 		authMock:                authMock,
// 		dsMock:                  dsMock,
// 		BookingServiceMock:      bookingServiceMock,
// 		RoomServiceMock:      roomServiceMock,
// 		AddressServiceMock: addressServiceMock,
// 		FileLoaderMock:          fileLoaderMock,
// 		ReportStoreMock:         reportMock,
// 		ctrl:                    ctrl,
// 	}
// }
