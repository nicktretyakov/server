package server

import (
	"net"

	grpcMiddleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpcRecovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"github.com/rs/zerolog"
	bookingpb "be/proto"
	//grpcLogger "gitlab.tn.ru/superapp/golang/grpc-zerolog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	outmemberService "be/internal/outmember"
	"be/internal/datastore"
	"be/internal/filestorage"
	addressService "be/internal/address"
	"be/internal/notecreator"
	roomService "be/internal/room"
	"be/internal/profile"
	bookingService "be/internal/booking"
	"be/pkg/auth"
	"be/pkg/notesender"
	"be/pkg/server/outmember"
	authServer "be/pkg/server/auth"
	"be/pkg/server/dictionary"
	"be/pkg/server/employee"
	"be/pkg/server/finalreport"
	"be/pkg/server/address"
	"be/pkg/server/note"
	"be/pkg/server/room"
	"be/pkg/server/booking"
	"be/pkg/server/report"
	"be/pkg/server/stage"
)

type Server struct {
	logger zerolog.Logger
	srv    *grpc.Server
}

func (s *Server) Shutdown() {
	s.srv.GracefulStop()
}

type Opts struct {
	AuthService         auth.IAuth
	Store               datastore.IDatastore
	Logger              zerolog.Logger
	BookingService      bookingService.IBookingService
	RoomService      roomService.IRoomService
	OutmemberService    outmemberService.IOutmemberService
	AddressService addressService.IAddressService
	ProfileAPI          profile.IProfile
	FileLoader          filestorage.IFileLoader
	NoteCreator *notecreator.NoteCreator
	NoteSender  *notesender.NoteSender
}

//nolint:funlen
func New(cfg Config, opts Opts) *Server {
	logger := opts.Logger
	authService := opts.AuthService
	store := opts.Store

	panicRecoverHandler := func(p interface{}) (err error) {
		logger.Error().Msgf("panic triggered: %v", p)
		return status.Errorf(codes.Unknown, "panic triggered: %v", p)
	}

	optsRecovery := []grpcRecovery.Option{
		grpcRecovery.WithRecoveryHandler(panicRecoverHandler),
	}

	srv := grpc.NewServer(
		grpcMiddleware.WithUnaryServerChain(
		//	grpcLogger.NewUnaryServerInterceptor(logger),
			grpcRecovery.UnaryServerInterceptor(optsRecovery...),
			auth.Interceptor(authService),
		),
	)

	bookingpb.RegisterAuthServer(srv, authServer.NewService(authService))

	bookingServer := booking.NewService(
		cfg.Booking(),
		store,
		opts.BookingService,
		opts.NoteCreator,
		booking.WithLogger(logger),
		booking.WithFileLoader(opts.FileLoader),
	)

	addressServer := address.NewService(
		opts.AddressService,
		store,
		opts.NoteCreator,
		address.WithLogger(logger),
		address.WithFileLoader(opts.FileLoader),
	)

	roomServer := room.NewService(store, opts.RoomService, opts.NoteCreator, room.WithLogger(logger))

	bookingpb.RegisterBookingServiceServer(srv, bookingServer)
	bookingpb.RegisterRoomServiceServer(srv, roomServer)
	bookingpb.RegisterAddressServiceServer(srv, addressServer)
	bookingpb.RegisterDictionaryServiceServer(srv, dictionary.NewService(store, logger))
	bookingpb.RegisterNoteServer(srv, note.NewService(store, note.WithLogger(logger)))
	bookingpb.RegisterOutmemberServiceServer(srv, outmember.NewService(outmember.Opts{
		Store:            store,
		BookingService:   opts.BookingService,
		Logger:           opts.Logger,
		Notificator:      opts.NoteCreator,
		OutmemberService: opts.OutmemberService,
	}))
	bookingpb.RegisterFinalReportServiceServer(srv, finalreport.NewService(finalreport.Opts{
		Store:          store,
		BookingService: opts.BookingService,
		Logger:         opts.Logger,
		Notificator:    opts.NoteCreator,
	}))

	reportServOpts := report.Opts{
		Store:          store,
		Logger:         opts.Logger,
		BookingService: opts.BookingService,
		Notificator:    opts.NoteCreator,
	}
	bookingpb.RegisterReportServiceServer(srv, report.NewService(reportServOpts))

	bookingpb.RegisterEmployeeServiceServer(
		srv, employee.New(employee.Opts{EmployeeRep: opts.ProfileAPI, Logger: opts.Logger}),
	)

	stageServOpts := stage.Opts{Store: store, Logger: opts.Logger, BookingService: opts.BookingService}

	bookingpb.RegisterStageServiceServer(srv, stage.NewService(stageServOpts))

	if opts.NoteCreator != nil {
		opts.NoteCreator.Run()
	}

	if opts.NoteSender != nil {
		opts.NoteSender.Run()
	}

	return &Server{
		logger: logger,
		srv:    srv,
	}
}

func Run(s *Server, ip, port string) error {
	address := net.JoinHostPort(ip, port)

	lis, err := net.Listen("tcp", address)
	if err != nil {
		return err
	}

	return RunListener(s, lis)
}

func RunListener(s *Server, lis net.Listener) error {
	go func() {
		s.logger.Info().Msgf("Start GRPC on %s", lis.Addr().String())

		if err := s.srv.Serve(lis); err != nil {
			s.logger.Fatal().Err(err).Msg("failed to serve")
		}
	}()

	return nil
}
