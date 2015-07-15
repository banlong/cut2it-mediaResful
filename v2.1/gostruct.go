package main
import (
	"time"

)

type MediaDescription struct{
	DescriptionId       string			`json:"descriptionId"`
	LanguageId          string  		`json:"languageId"`
	Region              string  		`json:"region"`
	DescriptionShort    string  		`json:"descriptionShort"`
	DescriptionExpanded string  		`json:"descriptionExpanded"`
	MediaId             string  		`json:"mediaId"`
}

type MediaTitle struct {
	TitleId       string			`json:"TitleId"`
	LanguageId    string			`json:"LanguageId"`
	Region        string			`json:"Region"`
	TitleShort    string			`json:"TitleShort"`
	TitleExtended string			`json:"TitleExtended"`
	MediaId       string			`json:"MediaId"`
}

type Thumbnail struct{
	ThumbId            string		`json:"thumbId"`
	LowQualityImage    []byte		`json:"lowQualityImage"`
	MediumQualityImage []byte		`json:"mediumQualityImage"`
	HighQualityImage   []byte		`json:"highQualityImage"`
	TargetDemographic  string		`json:"targetDemographic"`
	MediaId            string		`json:"mediaId"`
}

//Media
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
	TitleCol			[]MediaTitle
	DescriptionCol		[]MediaDescription
	ThumbnailCol		Thumbnail
}

//Array of media
type Medias []Media


type Language struct{
	Id 		string		`json:"id"`
	Title     string    `json:"title"`
}

type Languages []Language

