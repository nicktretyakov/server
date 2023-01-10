package handlers

import (
	"github.com/gorilla/mux"
	//"github.com/rs/zerolog"

	"be/internal/datastore"
	//"be/internal/filestorage"
	"be/internal/profile"
	"be/pkg/auth"
	"be/pkg/middleware"
)

// Bind binds http-handlers to router
// @title TNPromo mini-app API
// @version 0.0.1
// @description This is a tn-promo mini app server.
// @contact.name API
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @schemes https http
// @name Authorization
// @BasePath /api.
func Bind(authApp auth.IAuth,
	userStore datastore.IUserStore,
	profileAPI profile.IProfile,
	/*fileStorage filestorage.IFileStorage,*/
	r *mux.Router, /*logger zerolog.Logger,*/
) error {
	a := &api{
		//logger:      logger,
		auth:        authApp,
		userStore:   userStore,
		profileAPI:  profileAPI,
		//fileStorage: fileStorage,
	}

	apiRoot := r.PathPrefix("/api/v1").Subrouter()

	authRqRouter := apiRoot.PathPrefix("").Subrouter()
	authRqRouter.Use(middleware.AuthMiddleware(authApp))

	authRqRouter.HandleFunc("/avatar/{portal_code}", a.GetEmployeeAvatar).Methods("GET")
	authRqRouter.HandleFunc("/attachment", a.UploadTemporaryFile).Methods("POST")

	return nil
}

type api struct {
	//logger      zerolog.Logger
	auth        auth.IAuth
	userStore   datastore.IUserStore
	profileAPI  profile.IProfile
	//fileStorage filestorage.IFileStorage
}
