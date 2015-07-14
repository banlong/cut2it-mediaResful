package main

import (
	"encoding/json"				//JSON endoding
	"fmt"						//printf
	"net/http"					//http handling
	"github.com/gorilla/mux"	//implements a request router and dispatcher.
	"strconv"					//string conversion
	"time"						//time format
	"io/ioutil"					//file, buffer handling

	//"io"						//exit, file
	//"html/template"			//data driven templates for generating textual output such as HTML
	//"bufio"					//buffer io, wrap io.Reader/Writer
	//"os"						//platform independent inteface to os functionality


)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome CUT2IT GO API!\n")
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

/**
 * UPDATE A LANGUAGE
 */
func goUpdateLanguage(w http.ResponseWriter, r *http.Request) {
	//get infor from form
	langId := r.FormValue("langId")
	langTitle := r.FormValue("langTitle")
	lang := Language{Id: langId, Title:langTitle}

	//execute get method
	updateLanguage(lang)
	return
}

/**
 * ADD NEW LANGUAGE
 */
func goInsertLanguage(w http.ResponseWriter, r *http.Request){
	//get infor from form
	langId := r.FormValue("langId")
	langTitle := r.FormValue("langTitle")
	lang := Language{Id: langId, Title:langTitle}

	//execute get method
	insertLanguage(lang)
	return
}

/**
 * DELETE A LANGUAGE
 * Return 200 OK - Success
 *        404 Not Found - invalid id or id not found
 */
func goDeleteLanguage(w http.ResponseWriter, r *http.Request){
	//parse array of input from request message
	vars := mux.Vars(r)

	//get id input
	var langId string
	langId = vars["languageId"]

	//execute get method
	ret := deleteLanguageById(langId)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	switch  {
		case ret <= 0:
			w.WriteHeader(http.StatusNotFound)
		default:
			w.WriteHeader(http.StatusOK)
	}
	return
}

/**
 * GET ALL MEDIAS IN DB
 *
 */
func goGetAllMedia(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	mediaCol := getAllMedia();
	if err := json.NewEncoder(w).Encode(mediaCol); err != nil {
		panic(err)
	}
}

/**
 * GET MEDIA BY TITLE
 */
func goGetMediaByTitle(w http.ResponseWriter, r *http.Request){

	//get title from the form
	vars := mux.Vars(r)
	mediaTitle := vars["mediaTitle"]
	fmt.Println("Handler: " + mediaTitle)
	//execute get method
	medias := getMediaByTitle(mediaTitle)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	if(medias != nil){
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(medias); err != nil {
			panic(err)
		}
	} else{
		// If we didn't find it, 404
		w.WriteHeader(http.StatusNotFound)
		if err := json.NewEncoder(w).Encode(jsonErr{Code: http.StatusNotFound, Text: "Not Found"}); err != nil {
			panic(err)
		}
	}
}

/**
 * ADD A MEDIA
 */
func goAddMedia(w http.ResponseWriter, r *http.Request){
	//get info from form
	var media Media;
	var relInfo RelatedData
	media.MediaId = r.FormValue("mediaId")
	media.MediaHead = r.FormValue("mediaHead")
	media.Codec = r.FormValue("codec")
	media.Container = r.FormValue("container")
	media.UniqueViews, _ =  strconv.Atoi(r.FormValue("uniqueViews"))
	media.TargetDemographic = r.FormValue("targetDemographic")
	media.SizeInBytes, _ = strconv.Atoi(r.FormValue("sizeInBytes"))
	media.UploadedDate= time.Now()
	width:= r.FormValue("widthInPixels")
	media.WidthInPixels, _ = strconv.Atoi(width)

	height:= r.FormValue("heighInPixels")
	media.HeightInPixels, _ = strconv.Atoi(height)

	media.Orientation = r.FormValue("orientation")

	bStr := r.FormValue("definitionHD")
	if(bStr == "true"){
		media.DefinitionHD = true
	}else{
		media.DefinitionHD = false
	}

	media.RegionOfOrgin = r.FormValue("regionOfOrgin")
	media.ModifiedDate =time.Now()

	media.Duration = r.FormValue("duration")
	media.Title = r.FormValue("titleShort")
	media.ETitle = r.FormValue("titleExtended")
	media.Description = r.FormValue("descriptionShort")
	media.EDescription= r.FormValue("descriptionExpanded")

	relInfo.LanguageId = r.FormValue("langId")

	//execute get method
	insertMedia(media, relInfo)
	return
}

func goAddPicture(w http.ResponseWriter, r *http.Request){
	//get bytes from imagefile
	person:= r.FormValue("person")
	_,fHeader,err := r.FormFile("image")
	file, _:= fHeader.Open()
	data,err:=ioutil.ReadAll(file)
	count:= len(data)

	//fmt.Println(fHeader.Filename)
	//fmt.Println(person)
	fmt.Printf("read %d bytes", count)

	//insert image into database
	stmt, err := db.Prepare("INSERT INTO Person (FullName, Avatar) VALUES(?,?)")
	checkErr(err)

	// since big file, then just got 8000 bytes as max_size of VARBINARY Column in db.
	// should choose blog data to store the big fize instead of VARBINARY
	result, err := stmt.Exec(person, data)
	checkErr(err)

	// 4. print out the result.
	fmt.Printf("%d", result)
}