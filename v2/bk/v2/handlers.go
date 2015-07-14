package main

import (
	"encoding/json"				//JSON endoding
	"fmt"						//printf
	"io"						//exit, file
	"io/ioutil"					//buffer
	"net/http"					//http handling
	"strconv"					//string conversion
	"github.com/gorilla/mux"	//implements a request router and dispatcher.
	"html/template"				//data driven templates for generating textual output such as HTML
	"bufio"						//buffer io, wrap io.Reader/Writer
	"os"						//platform independent inteface to os functionality
	//"image"					//Package image implements a basic 2-D image library. The fundamental interface is
								//called Image. An Image contains colors, which are described in the image/color package.
	//"bytes"					//Package bytes implements functions for the manipulation of byte slices.

	"mime/multipart"
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

//UPLOAD FILE TO A POSITION ON SERVER
func goUploadImage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method)
	r.ParseMultipartForm(32 << 20)
	file, handler, err := r.FormFile("uploadfile")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	fmt.Fprintf(w, "%v", handler.Header)
	f, err := os.OpenFile("./"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}

	io.Copy(f, file)
	defer f.Close()
	return
}

func goUploadImageToSQLServer(w http.ResponseWriter, r *http.Request){

	//Get file name
	fmt.Println("method:", r.Method)
	r.ParseMultipartForm(32 << 20)
	_, handler, _ := r.FormFile("uploadfile")

	if(handler.Filename == ""){
		fmt.Println("Empty file name")
		return
	}

	fmt.Println("File name: "+ handler.Filename )
	dat,_ := ioutil.ReadFile(handler.Filename)
	fmt.Print(string(dat)) //nothing to print here

	f, _ := os.Open("/tmp/dat")
	b1 := make([]byte, 5)
	n1, _ := f.Read(b1)
	fmt.Printf("%d bytes: %s\n", n1, string(b1)) //no bytes upto here

	o2, _ := f.Seek(6, 0)
	b2 := make([]byte, 2)
	n2, _ := f.Read(b2)
	fmt.Printf("%d bytes @ %d: %s\n", n2, o2, string(b2))
	o3, _ := f.Seek(6, 0)
	b3 := make([]byte, 2)
	n3, _ := io.ReadAtLeast(f, b3, 2)
	fmt.Printf("%d bytes @ %d: %s\n", n3, o3, string(b3))
	_, err = f.Seek(0, 0)
	r4 := bufio.NewReader(f)
	b4, _ := r4.Peek(5)
	fmt.Printf("5 bytes: %s\n", string(b4))
	f.Close()

	//--------------------
//	fmt.Println("Gere")
//	//open file
//
//	f, err := os.OpenFile("./"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//	defer f.Close()
//
//
//	//convert file into bytes
//	//io.Copy(f, file)
//	fileInfo, _ := f.Stat()
//	var size int64 = fileInfo.Size()
//	fmt.Printf("Size - %d \n", size)
//	imgByte := make([]byte, size)
//
//	buffer := bufio.NewReader(f)
//	_, err = buffer.Read(imgByte)
//
//	//convert byte to image
//	//img, _, _ := image.(bytes.NewReader(imgByte))
//
//	sqlStatement := "INSERT INTO profile(id, picture)VALUES(1, ?)"
//	fmt.Println(sqlStatement)
//
//	st, err := db.Prepare(sqlStatement)
//	if err != nil{
//		fmt.Printf("Incorrect format -  ")
//		fmt.Print( err );
//		os.Exit(-1)
//	}
//
//	//get the number of affected rows
//	st.Exec(imgByte)

}

func temp(w http.ResponseWriter, r *http.Request){

	const _24K = (1 << 20) * 24
	//var status int
	if err = r.ParseMultipartForm(_24K); nil != err {
		//status = http.StatusInternalServerError
		return
	}
	for _, fheaders := range r.MultipartForm.File {
		for _, hdr := range fheaders {
			// open uploaded
			var infile multipart.File
			if infile, err = hdr.Open(); nil != err {
				//status = http.StatusInternalServerError
				fmt.Println(err)
				return
			}
			// open destination
			var outfile *os.File
			if outfile, err = os.Create("./uploaded/" + hdr.Filename); nil != err {
				//status = http.StatusInternalServerError
				fmt.Println(err)
				return
			}
			// 32K buffer copy
			var written int64
			if written, err = io.Copy(outfile, infile); nil != err {
				//status = http.StatusInternalServerError
				fmt.Println(err)
				return
			}
			//fmt.Println([]byte("uploaded file:" + hdr.Filename + ";length:" + strconv.Itoa(int(written))))
			fmt.Println("uploaded file:" + hdr.Filename + ";length:" + strconv.Itoa(int(written)))
		}
	}
}