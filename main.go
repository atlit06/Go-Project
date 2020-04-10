package main

import (
	"net/http"

	"github.com/atlit06/webservice/controllers"
)

func main() {
	controllers.RegisterControllers()
	http.ListenAndServe(":3000", nil)
}
