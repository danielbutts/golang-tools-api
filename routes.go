package toolexchange

import "net/http"

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		Index,
	},
	Route{
		"ToolIndex",
		"GET",
		"/tools",
		ToolIndex,
	},
	Route{
		"TodoShow",
		"GET",
		"/tools/{toolId}",
		ToolShow,
	},
}
