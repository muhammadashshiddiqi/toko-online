package app

import "toko-online/app/controllers"

func (server *Server) initializeRoutes()  {
	server.Router.HandleFunc("/", controllers.Home).Methods("GET")
}