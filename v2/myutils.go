package main
import _ "github.com/denisenkom/go-mssqldb"
//import "database/sql"
import "log"
import "fmt"
//import "flag"
import (
	"image"
	"image/jpeg"
	"os"
)
//ROUTE FOR date
//Route{
//"InsertDate",
//"POST",
//"/dates/add",
//goInsertDate,
//},
//func goInsertDate(w http.ResponseWriter, r *http.Request){
//	//get infor from form
//	const longForm = "Jan 2, 2006 at 3:04pm (MST)"
//	const shortForm = "2006-Jan-02"
//
//	Id := r.FormValue("id")
//	Dates := r.FormValue("date")
//	datetimeVal, _ := time.Parse(longForm, Dates)
//
//	//execute get method
//	sqlStatement := "INSERT INTO Profile (id, dateData)VALUES (?,?)"
//	fmt.Println(sqlStatement)
//
//	st, _ := db.Prepare(sqlStatement)
//	st.Query(Id, datetimeVal)
//	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
//	w.WriteHeader(http.StatusOK)
//}



func getImageBytes(filename string) ([]byte, int) {
	imageRegister()


	imgfile, err := os.Open(filename)

	fmt.Println(imgfile.Name())


	if err != nil {
		fmt.Println("img.jpg file not found!")
		os.Exit(1)
	}
	defer imgfile.Close()

	fi, err := imgfile.Stat()
	if err != nil {
		// Could not obtain stat, handle error
	}

	data := make([]byte, fi.Size())
	count, err := imgfile.Read(data)
	if err != nil {
		log.Fatal(err)
	}

	return data, count
}

func imageRegister() {
	// damn important or else At(), Bounds() functions will
	// caused memory pointer error!!
	image.RegisterFormat("jpeg", "jpeg", jpeg.Decode, jpeg.DecodeConfig)
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}