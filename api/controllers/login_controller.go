package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/AlaaRadwan94/awesomeProject1/api/auth"
	"github.com/AlaaRadwan94/awesomeProject1/api/models"
	"github.com/AlaaRadwan94/awesomeProject1/api/responses"
	"github.com/AlaaRadwan94/awesomeProject1/api/utils/formaterror"
	"golang.org/x/crypto/bcrypt"
)

//this api will be used to login for users
func (server *Server) Login(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	//user := models.User{}
	userF := models.UserFactory{}
	user := userF.Create()
	err = json.Unmarshal(body, &user)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	user.Prepare()
	err = user.Validate("login")
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	newUser := user.(*models.User)
	token, err := server.SignIn(newUser.Email, newUser.Password)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusUnprocessableEntity, formattedError)
		return
	}
	responses.JSON(w, http.StatusOK, token)
}

//this api will be used to login for companies
func (server *Server) CompanyLogin(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	userF := models.CompanyFactory{}
	user := userF.Create()
	err = json.Unmarshal(body, &user)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	user.Prepare()
	err = user.Validate("login")
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	newUser := user.(*models.User)
	token, err := server.SignIn(newUser.Email, newUser.Password)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusUnprocessableEntity, formattedError)
		return
	}
	responses.JSON(w, http.StatusOK, token)
}

//this function will be used in login api to check the email and the password of the request and create the token
func (server *Server) SignIn(email, password string) (string, error) {

	var err error

	user := models.User{}
	err = server.DB.Debug().Model(models.User{}).Where("email = ?", email).Take(&user).Error
	if err != nil {
		return "", err
	}
	err = models.VerifyPassword(user.Password, password)
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", err
	}
	return auth.CreateToken(user.ID)
}
