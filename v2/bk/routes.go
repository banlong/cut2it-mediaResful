package main

import "net/http"

//ROUTE STRUCT
type Route struct {
	Name        string				// route name
	Method      string				// http method
	Pattern     string    			// path of resource GET/POST...
	HandlerFunc http.HandlerFunc  	//func to handle the request
}

type Routes []Route

var routes = Routes{
	//HANDLE THE ROOT "/"
	Route{
		"Index",
		"GET",
		"/",
		Index,
	},

	Route{
		"TodoIndex",
		"GET",
		"/todos",
		TodoIndex,
	},

	Route{
		"TodoCreate",
		"POST",
		"/todos",
		TodoCreate,
	},
	Route{
		"TodoShow",
		"GET",
		"/todos/{todoId}",
		TodoShow,
	},

	//ROUTE FOR LANGUAGES
	Route{
		"ShowLanguages",
		"GET",
		"/languages",
		goLanguages,
	},

	Route{
		"GetLanguagesById",
		"GET",
		"/languages/{languageId}",
		goGetLanguageById,
	},

	Route{
		"UpdateLanguage",
		"PUT",
		"/languages",
		goUpdateLanguage,
	},

	Route{
		"InsertLanguage",
		"POST",
		"/languages",
		goInsertLanguage,
	},

	Route{
		"DeleteLanguage",
		"DELETE",
		"/languages/{languageId}",
		goDeleteLanguage,
	},

	Route{
		"Greet",
		"GET",
		"/greet", // great.html
		Greet,
	},

	Route{
		"Greeter",
		"POST",
		"/greeter", // greeter.html
		Greeter,
	},


}
