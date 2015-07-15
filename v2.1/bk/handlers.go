package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
	"github.com/gorilla/mux"
	"html/template"

)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome CUT2IT GO API!\n")
}

func TodoIndex(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(todos); err != nil {
		panic(err)
	}
}

func TodoShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var todoId int
	var err error
	if todoId, err = strconv.Atoi(vars["todoId"]); err != nil {
		panic(err)
	}
	todo := RepoFindTodo(todoId)
	if todo.Id > 0 {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(todo); err != nil {
			panic(err)
		}
		return
	}

	// If we didn't find it, 404
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusNotFound)
	if err := json.NewEncoder(w).Encode(jsonErr{Code: http.StatusNotFound, Text: "Not Found"}); err != nil {
		panic(err)
	}

}

func TodoCreate(w http.ResponseWriter, r *http.Request) {
	var todo Todo
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}
	if err := json.Unmarshal(body, &todo); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}

	t := RepoCreateTodo(todo)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(t); err != nil {
		panic(err)
	}
}



/*
 * GET ALL LANGUAGE
 * Handle request of all langguages
 */
func goLanguages(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	var langs Languages = getLanguages();
	if err := json.NewEncoder(w).Encode(langs); err != nil {
		panic(err)
	}
}

/**
 * GET A LANGUAGE BY ITS ID
 */
func goGetLanguageById(w http.ResponseWriter, r *http.Request) {
	//parse array of input from request message
	vars := mux.Vars(r)

	//get id input
	var langId string
	langId = vars["languageId"]

	//execute get method
	lang := getLanguageById(langId)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	if lang.Id != "" {
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(lang); err != nil {
			panic(err)
		}
		return
	}

	// If we didn't find it, 404
	w.WriteHeader(http.StatusNotFound)
	if err := json.NewEncoder(w).Encode(jsonErr{Code: http.StatusNotFound, Text: "Not Found"}); err != nil {
		panic(err)
	}

}

func goUpdateLanguage(w http.ResponseWriter, r *http.Request) {
	//get infor from form
	langId := r.FormValue("langId")
	langTitle := r.FormValue("langTitle")
	lang := Language{Id: langId, Title:langTitle}

	//execute get method
	updateLanguage(lang)
	return
}

func goInsertLanguage(w http.ResponseWriter, r *http.Request){
	//get infor from form
	langId := r.FormValue("langId")
	langTitle := r.FormValue("langTitle")
	lang := Language{Id: langId, Title:langTitle}

	//execute get method
	insertLanguage(lang)
	return
}

func goDeleteLanguage(w http.ResponseWriter, r *http.Request){
	//parse array of input from request message
	vars := mux.Vars(r)

	//get id input
	var langId string
	langId = vars["languageId"]

	//execute get method
	ret := deleteLanguageById(langId)
	if ret == -1 {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusNotFound)

	}else{
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
	}
	return
}

func Greet(w http.ResponseWriter, r *http.Request) {

	t, _ := template.ParseFiles("html/greet.html")
	t.Execute(w, nil)
}

func Greeter(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	t, _ := template.ParseFiles("html/greeter.html")
	err := t.Execute(w, username)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}


