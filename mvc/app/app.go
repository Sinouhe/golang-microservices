package app

import (
	"net/http"
	"golang-microservices/mvc/controllers"
)

func StartApp() {
	http.HandleFunc("/users", controllers.GetUsers)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err);
	}
}
