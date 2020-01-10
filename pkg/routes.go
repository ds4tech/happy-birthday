package birthday

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	for _, route := range routes {

		var handler http.Handler
		handler = route.HandlerFunc
		handler = Logger(handler, route.Name)
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
		//Handler(route.HandlerFunc)
	}
	return router
}

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		Index,
	},
	Route{
		"Hello",
		"GET",
		"/hello",
		Hello,
	},
	Route{
		"HelloSomebody",
		"GET",
		"/hello/{somebody}",
		HelloSomebody,
	},
	Route{
		"SaveSmbsName",
		"PUT",
		"/hello/{somebody}",
		SaveSmbsName,
	},

	Route{
		"Count",
		"GET",
		"/count",
		Counter,
	},
	Route{
		"Parser",
		"GET",
		"/parse",
		Parser,
	},
}