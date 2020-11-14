package controllers

import (
	"net/http"

	"github.com/AlaaRadwan94/awesomeProject1/api/responses"
)

func (server *Server) Home(w http.ResponseWriter, r *http.Request) {
	responses.JSON(w, http.StatusOK, "Welcome To Note Enterprise API")

}
