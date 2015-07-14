package main
import (
	"time"
	"image"
)


//Media structure
type Media struct {
	MediaId   			string			`json:"mediaId"`
	MediaHead   		string  		`json:"mediaHead"`
	Codec       		string  		`json:"codec"`
	Container			string			`json:"container"`
	UniqueViews 		int  			`json:"uniqueViews"`
	TargetDemographic 	string			`json:"targetDemographic"`
	SizeInBytes			int				`json:"sizeInBytes"`
	UploadedDate      	time.Time		`json:"uploadedDate"`
	WidthInPixels     	int 			`json:"widthInPixels"`
	HeightInPixels    	int				`json:"heightInPixels"`
	Orientation       	string			`json:"orientation"`
	DefinitionHD      	bool			`json:"definitionHD"`
	RegionOfOrgin     	string			`json:"regionOfOrgin"`
	ModifiedDate      	time.Time		`json:"modifiedDate"`
	Duration          	string			`json:"duration"`
	Title				string			`json:"title"`
	ETitle				string			`json:"etitle"`
	TitleRegion			string			`json:"titleRegion"`
	Description			string			`json:"description"`
	EDescription		string			`json:"edescription"`
	LowImage			image.Image		`json:"lQuaImage"`
	MedImage			image.Image		`json:"mQuaImage"`
	HiImage				image.Image		`json:"hQuaImage"`
}

//Array of media
type Medias []Media


type Language struct{
	Id 		string		`json:"id"`
	Title     string    `json:"title"`
}

type Languages []Language

type RelatedData struct{
	LanguageId		string		`json:"langId"`

}