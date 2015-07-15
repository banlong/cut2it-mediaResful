package main

import (
	"encoding/json"				//JSON endoding
	"fmt"						//printf
	"net/http"					//http handling
	"github.com/gorilla/mux"	//implements a request router and dispatcher.
	"strconv"					//string conversion
	"time"						//time format

	//"io/ioutil"					//file, buffer handling
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
	mediaCol := getAllMedias();
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

	var media Media;
	//get media info from form
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

	//add the first title
	var tempTitle MediaTitle
	tempTitle.TitleId = generateId(autoInt,"TT")
	tempTitle.LanguageId = r.FormValue("titleLangId")
	tempTitle.TitleShort = r.FormValue("titleShort")
	tempTitle.TitleExtended = r.FormValue("titleExtended")
	tempTitle.MediaId = r.FormValue("mediaId")
	tempTitle.Region = r.FormValue("titleRegion")
	media.TitleCol = append(media.TitleCol, tempTitle)

	//add the first description
	var tempDescription MediaDescription
	tempDescription.DescriptionId = generateId(autoInt, "DE")
	tempDescription.MediaId = r.FormValue("mediaId")
	tempDescription.DescriptionShort = r.FormValue("descriptionShort")
	tempDescription.DescriptionExpanded= r.FormValue("descriptionExpanded")
	tempDescription.LanguageId = r.FormValue("descriptionLangId")
	tempDescription.Region = r.FormValue("descriptionRegion")
	media.DescriptionCol = append(media.DescriptionCol, tempDescription)

	//add Thumbnail
	media.ThumbnailCol.MediaId = r.FormValue("mediaId")
	media.ThumbnailCol.ThumbId = generateId(autoInt, "TN")
	media.ThumbnailCol.TargetDemographic =  r.FormValue("targetDemographic")

	var databyte int
	media.ThumbnailCol.LowQualityImage,databyte = getImageBytes(r, "lowQualityThumb")
	if(imageValidate(w, databyte) < 0){
		return
	}

	media.ThumbnailCol.MediumQualityImage,_ = getImageBytes(r, "medQualityThumb")
	if(imageValidate(w, databyte) < 0){
		return
	}

	media.ThumbnailCol.HighQualityImage,_ = getImageBytes(r, "higQualityThumb")
	if(imageValidate(w, databyte) < 0){
		return
	}
	//execute get method
	insertMedia(media)
	insertTitle(media)
	insertDescription(media)
	insertThumb(media)
	autoInt++
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	return
}

