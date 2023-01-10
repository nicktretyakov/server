package handlers

import (
	"net/http"
	"time"

	//"be/pkg/response"
)

const tmpLinkFileLifeTime = time.Hour * 24

type TemporaryFileUploadResponse struct {
	Link string `json:"link"`
}

// UploadTemporaryFile uploads temporary file
// @Router /v1/attachment [post].
func (a *api) UploadTemporaryFile(w http.ResponseWriter, r *http.Request) {
	//ctx := r.Context()

	// file, fileHeaders, err := r.FormFile("file")
	// if err != nil {
	// 	response.WriteJSON(w, http.StatusInternalServerError, NewAPIError(err))
	// 	return
	// }

	// filename := fileHeaders.Filename
	// mime := fileHeaders.Header.Get("content-type")

	// key, err := a.fileStorage.AddTemporaryFile(ctx, file, filename, mime)
	// if err != nil {
	// 	response.WriteJSON(w, http.StatusInternalServerError, NewAPIError(err))
	// 	return
	// }

	// signedLink, err := a.fileStorage.Link(key, tmpLinkFileLifeTime)
	// if err != nil {
	// 	response.WriteJSON(w, http.StatusInternalServerError, NewAPIError(err))
	// 	return
	// }

	// response.WriteJSON(w, http.StatusOK, TemporaryFileUploadResponse{
	// 	Link: signedLink,
	// })
}
