package controllers

import (
	"github.com/AlaaRadwan94/awesomeProject1/api/middlewares"
)

func (s *Server) initializeRoutes() {

	// Home Route
	//s.Router.HandleFunc("/", middlewares.SetMiddlewareJSON(s.Home)).Methods("GET")

	// Login Route
	s.Router.HandleFunc("/login", middlewares.SetMiddlewareJSON(s.Login)).Methods("POST")
	s.Router.HandleFunc("/company-login", middlewares.SetMiddlewareJSON(s.CompanyLogin)).Methods("POST")

	//Users routes
	s.Router.HandleFunc("/users", middlewares.SetMiddlewareJSON(s.CreateUser)).Methods("POST")
	s.Router.HandleFunc("/users-promotion/{promotion}", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.UpdatePromotion))).Methods("PUT")
	s.Router.HandleFunc("/users/{id}", middlewares.SetMiddlewareJSON(s.GetUser)).Methods("GET")
	s.Router.HandleFunc("/users/{id}", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.UpdateUser))).Methods("PUT")
	s.Router.HandleFunc("/users/{id}", middlewares.SetMiddlewareAuthentication(s.DeleteUser)).Methods("DELETE")
	//
	////Company routes
	s.Router.HandleFunc("/companies", middlewares.SetMiddlewareJSON(s.CreateCompany)).Methods("POST")
	s.Router.HandleFunc("/companies-promotion/{promotion}", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.UpdateCompanyPromotion))).Methods("PUT")
	s.Router.HandleFunc("/companies/{id}", middlewares.SetMiddlewareJSON(s.GetCompany)).Methods("GET")
	s.Router.HandleFunc("/companies/{id}", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.UpdateCompany))).Methods("PUT")
	s.Router.HandleFunc("/companies/{id}", middlewares.SetMiddlewareAuthentication(s.DeleteCompany)).Methods("DELETE")

}
