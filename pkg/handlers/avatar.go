package handlers

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/pkg/errors"

	"be/internal/datastore/base"
	"be/internal/model"
	"be/pkg/response"
)

// GetEmployeeAvatar
// @Summary Return employees avatar
// @Router /v1/avatar/{portal_code} [get].
func (a *api) GetEmployeeAvatar(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	portalCode, err := strconv.ParseUint(mux.Vars(r)["portal_code"], 0, 0)
	if err != nil {
		response.WriteJSON(w, http.StatusBadRequest, NewAPIError(err))
		return
	}

	user, err := a.userStore.FindUserByPortalCode(ctx, portalCode)
	if err != nil && !errors.Is(err, base.ErrNotFound) {
		response.WriteJSON(w, http.StatusNotFound, NewAPIError(err))
		return
	}

	var employee *model.Employee

	if user != nil {
		employee = &user.Employee
	} else if _employee, err := a.profileAPI.FindEmployeeByPortalCode(ctx, portalCode); err == nil {
		_e := _employee.Cast()
		employee = &_e
	}

	if employee == nil {
		response.WriteJSON(w, http.StatusNotFound, NewAPIError(errors.New("not found")))
		return
	}

	w.Header().Set("Content-Type", "image/jpeg")
	w.Header().Set("Cache-control", "max-age=5")

	if _, err = w.Write([]byte(employee.GetAvatar())); err != nil {
		response.WriteJSON(w, http.StatusInternalServerError, NewAPIError(err))
		return
	}
}
