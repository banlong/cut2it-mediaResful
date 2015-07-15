package main
import _ "github.com/denisenkom/go-mssqldb"
import (
	"image"
	"image/jpeg"
	"time"
	"strconv"
	"net/http"
	"io/ioutil"

)

const MAX_THUMB_SIZE  = 1000

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


/**
 * GENERATE ID
 * Generate id for description table & title table
 */
func generateId(prev int, padd string) string{
	now := time.Now()
	yearday:= time.Now().YearDay()
	year:= now.Year()
	autoInt := prev + 1;
	retVal := padd + strconv.Itoa(year) + strconv.Itoa(yearday) +  strconv.Itoa(autoInt)
	return retVal
}


func getImageBytes(r *http.Request, fieldname string)([]byte, int){
	//if no file name input
	_,fHeader,err := r.FormFile(fieldname)
	if(err != nil) {
		return nil, 0
	}

	file,err:= fHeader.Open()
	//if cannot open file
	if(err != nil) {
		return nil, -1
	}

	data,_:=ioutil.ReadAll(file)
	count:= len(data)

	//if picture size > max_size
	if count > MAX_THUMB_SIZE{
		return nil, -2
	}

	return data, count
}

func imageValidate(w http.ResponseWriter, flag int) int{
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	//w.WriteHeader(http.StatusBadRequest)

	switch flag {
		case -1:
			http.Error(w, "Cannot open the image file" , 400)
			return -1
		case -2:
			http.Error(w, "Image size is over " + strconv.Itoa(MAX_THUMB_SIZE) + " bytes" , 400)
			return -2
	}

	return 0;
}

