package main

//import _ "github.com/denisenkom/go-mssqldb"
//import "database/sql"
//import "log"
//import "fmt"
//import "flag"
//
//import (
//	"image"
//	"image/jpeg"
//	"os"
//)
//
//
//
//
//var debug = flag.Bool("debug", false, "enable debugging")
//var port *int = flag.Int("port", 1433, "the database port")
//var server = flag.String("server", "10.76.0.214", "the database server")
//var user = flag.String("user", "sa", "the database user")
//var password = flag.String("password", "mari123!", "the database password")
//var database = flag.String("database", "testdb", "the database")
//
//func main() {
//	flag.Parse() // parse the command line args
//
//	if *debug {
//		fmt.Printf(" password:%s\n", *password)
//		fmt.Printf(" port:%d\n", *port)
//		fmt.Printf(" server:%s\n", *server)
//		fmt.Printf(" user:%s\n", *user)
//	}
//
//	// 1. ket noi db
//	dbString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s", *server, *user, *password, *port, *database)
//	if *debug {
//		fmt.Printf(" dbString:%s\n", dbString)
//	}
//
//	db, err := sql.Open("mssql", dbString)
//	if err != nil {
//		log.Fatal("Open dbection failed:", err.Error())
//	}
//	defer db.Close()
//
//	// 2. read image. this can be replace with reading bytes data from uploaded file
//	data, count := readImage()
//	fmt.Printf("read %d bytes: %q\n", count, data[10])
//
//	// 3. insert into db. with test schema as below
//	/*
//	-- drop table Person
//	CREATE TABLE Person
//	(
//	ID int IDENTITY(1,1) PRIMARY KEY,
//	Avatar VARBINARY(8000),
//	FullName varchar(255) default ''
//	);
//	*/
//	stmt, err := db.Prepare("INSERT INTO Person (FullName, Avatar) VALUES(?,?)")
//	checkErr(err)
//
//	// since big file, then just got 8000 bytes as max_size of VARBINARY Column in db.
//	// should choose blog data to store the big fize instead of VARBINARY
//	result, err := stmt.Exec("Nghia Ngo", data[8000])
//	checkErr(err)
//
//	// 4. print out the result.
//	fmt.Printf("%d", result)
//	db.Close()
//}

//func checkErr(err error) {
//	if err != nil {
//		panic(err)
//	}
//}

//func register() {
//	// damn important or else At(), Bounds() functions will
//	// caused memory pointer error!!
//	image.RegisterFormat("jpeg", "jpeg", jpeg.Decode, jpeg.DecodeConfig)
//}
//
//func readImage() ([]byte, int) {
//	register()
//
//	imgfile, err := os.Open("./img.jpg")
//
//	fmt.Println(imgfile.Name())
//
//	if err != nil {
//		fmt.Println("img.jpg file not found!")
//		os.Exit(1)
//	}
//	defer imgfile.Close()
//
//	fi, err := imgfile.Stat()
//	if err != nil {
//		// Could not obtain stat, handle error
//	}
//
//	data := make([]byte, fi.Size())
//	count, err := imgfile.Read(data)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	return data, count
//}