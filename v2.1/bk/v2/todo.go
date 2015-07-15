package main

import "time"

type Todo struct {
	Id        int       `json:"id"`
	Name      string    `json:"name"`
	Completed bool      `json:"completed"`
	Due       time.Time `json:"due"`
}


//TODO: this model class
type Todos []Todo

type Togo struct {
	Id        string       `json:"id"`
	Title     string    `json:"title"`
}

type Togos []Togo

//Media structure
type Media struct {
	mediaId   			string			`json:"mediaId"`
	mediaHead   		string  		`json:"mediaHead"`
	codec       		string  		`json:"codec"`
	container			string			`json:"container"`
	uniqueViews 		int  			`json:"uniqueViews"`
	targetDemographic 	string			`json:"targetDemographic"`
	sizeInBytes			int				`json:"sizeInBytes"`
	uploadedDate      	time.Time 		`json:"uploadedDate"`
	widthInPixels     	int 			`json:"widthInPixels"`
	heightInPixels    	int				`json:"heightInPixels"`
	orientation       	string			`json:"orientation"`
	definitionHD      	bool			`json:"definitionHD"`
	regionOfOrgin     	string			`json:"regionOfOrgin"`
	modifiedDate      	time.Time		`json:"modifiedDate"`
	duration          	time.Duration	`json:"duration"`
}

//Array of media
type Medias []Media


type Language struct{
	Id 		string		`json:"id"`
	Title     string    `json:"title"`
}

type Languages []Language