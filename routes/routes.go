package routes

import (
	"curso3/controllers"
	middleware "curso3/middlewares"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func HandleRequest() {
	r := mux.NewRouter()
	r.Use(middleware.ContentTypeMiddleware)
	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/api/personalidades", controllers.TodasPersonalidades).Methods("Get")
	r.HandleFunc("/api/personalidades/{id}", controllers.RetornaPersonalidade).Methods("Get")
	r.HandleFunc("/api/personalidades", controllers.Criar).Methods("Post")
	r.HandleFunc("/api/personalidades", controllers.Deletar).Methods("Delete")
	r.HandleFunc("/api/personalidades/{id}", controllers.Editar).Methods("Put")
	log.Fatal(http.ListenAndServe(":8001", r))

}
