package controllers

import (
	"context"
	"database/sql"
	"net/http"
	"strings"

	"github.com/jobutterfly/olives/utils"
)

func (h *Handler) GetUser(w http.ResponseWriter, r *http.Request) {
	v, err := utils.GetPathValues(strings.Split(r.URL.Path, "/"))
	if err != nil {
		utils.NewError(w, http.StatusInternalServerError, err.Error())
		return
	}

	user, err := h.q.GetUser(context.Background(), int32(v.Id))
	if err != nil {
		if err == sql.ErrNoRows {
			utils.NewError(w, http.StatusNotFound, err.Error())
			return
		}

		utils.NewError(w, http.StatusInternalServerError, "error when getting user")
		return
	}

	utils.NewResponse(w, http.StatusOK, user)
	return
}
