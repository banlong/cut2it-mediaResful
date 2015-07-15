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

	/*************************ROUTE FOR LANGUAGE*********************************************/
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

    /*************************ROUTE FOR MEDIA*********************************************/
	Route{
		"ShowAllMedia",
		"GET",
		"/medias",
		goGetAllMedia,
	},

	Route{
		"ShowMediaByTitle",
		"GET",
		"/medias/{mediaTitle}",
		goGetMediaByTitle,
	},

	Route{
		"InsertAMedia",
		"POST",
		"/medias/insert",
		goAddMedia,
	},


}
