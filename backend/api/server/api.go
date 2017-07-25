package server

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/handlers"
	"github.com/op/go-logging"
	"net/http"
	"os"
)

const (
	//GET = "GET"
	GET = "GET"
	//POST = "POST"
	POST = "POST"
	//PUT = "PUT"
	PUT = "PUT"
	//DELETE = "DELETE"
	DELETE = "DELETE"
)

//API defines an API
type API struct{}

var log = logging.MustGetLogger("api")

//Start starts the api on the given port
func (api *API) Start(port int) {
	portString := fmt.Sprintf(":%d", port)

	log.Infof("Listening to port %d", port)
	http.ListenAndServe(portString, handlers.LoggingHandler(os.Stdout, http.DefaultServeMux))
}

//AddController adds a GetHandler to a specific path
func (api *API) AddController(controller Controller, path string) {
	http.HandleFunc("/api"+path, api.controller(controller))

	log.Infof("Added controller for endpoint %s", path)
}

func (api *API) controller(controller Controller) http.HandlerFunc {
	return func(rw http.ResponseWriter, request *http.Request) {
		var data interface{}
		var code int

		request.ParseForm()
		values := request.Form

		switch request.Method {
		case GET:
			code, data = controller.Get(values)
		case POST:
			code, data = controller.Post(values)
		case PUT:
			code, data = controller.Put(values)
		case DELETE:
			code, data = controller.Delete(values)
		default:
			rw.WriteHeader(405)
			return
		}

		content, err := json.Marshal(data)
		if err != nil {
			rw.WriteHeader(500)
		}

		rw.Header().Add("Access-Control-Allow-Origin", "*")
		rw.WriteHeader(code)
		rw.Write(content)
	}
}
