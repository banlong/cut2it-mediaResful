package main
import (
	"fmt"								//format library
	"database/sql"						//first import to work with database
	_"github.com/denisenkom/go-mssqldb" //using mssql driver
	"log"								//log
	"os"
)

	var db *sql.DB
	var err error

func init() {
	fmt.Println("_Database connection inititalized.")
	db, err = sql.Open("mssql", "server=127.0.0.1;port=1433;user id=coredbadmin;password=ilovecut2it;database=media")
	if err != nil {
		log.Fatal(err)
		fmt.Println("Error!.")
		os.Exit(1)
	}

}



//func main2() {
//
//	//abstract the database using MS SQL Driver
//	db, err:= sql.Open("mssql", "server=127.0.0.1;port=1433;user id=nghiepnds;password=sunghiep;database=media")
//
//	if err != nil {
//		log.Fatal(err)
//		fmt.Println("Error!.")
//		os.Exit(1)
//	}
//
//	fmt.Println("Before Query")
//	displayLanguage(db)
//
//	//EXECUTE A STORE PROCEDURE, NO RETURN SET
//	pst, err := db.Prepare("EXEC Insert_Languages 'CU', 'Cuban'")
//	pst.Query()
//	if err != nil {
//		fmt.Printf("Call procedure failed, ")
//		fmt.Print(err)
//	}
//	pst.Close()
//
//	//must close pst before open another one
//	fmt.Println("After Query")
//	displayLanguage(db)
//
//	defer db.Close()
//}


//DISPLAY LANGUAGE IN DATABASE
//func displayLanguage(db *sql.DB) {
//	st, err := db.Prepare("SELECT * FROM Languages")
//	if err != nil{
//		fmt.Printf("Statement = select * from language, ")
//		fmt.Print( err );
//		os.Exit(1)
//	}
//
//	rows, err := st.Query()
//	if err != nil {
//		fmt.Printf("Query = select * from language, ")
//		fmt.Print( err )
//		os.Exit(1)
//	}
//
//	i := 0
//	var title string
//	var id string
//	for rows.Next() {
//		i++
//		err = rows.Scan(&id, &title )
//		fmt.Printf("%s - %s \n", id, title )
//	}
//	fmt.Println("Total: " , i )
//
//	//close the result set so that we can reopen it later
//	st.Close()
//	rows.Close()
//	return
//}

//GET ALL LANGUAGE
func getLanguages() Languages {
	st, err := db.Prepare("SELECT * FROM Languages")
	if err != nil{
		fmt.Printf("Statement = select * from language, ")
		fmt.Print( err );
		os.Exit(1)
	}

	rows, err := st.Query()
	if err != nil {
		fmt.Printf("Query = select * from language, ")
		fmt.Print( err )
		os.Exit(1)
	}

	var langs Languages

	i := 0
	var title string
	var id string
	for rows.Next() {
		i++
		err = rows.Scan(&id, &title )
		fmt.Printf("%s - %s \n", id, title )

		langs = append(langs, Language{Id: id, Title: title})
	}
	fmt.Println("Total: " , i )

	//close the result set so that we can reopen it later
	st.Close()
	rows.Close()
	//defer db.Close()

	return langs
}

//GET LANGUAGE BY ID
func getLanguageById(langId string) Language {
	sqlStatement := "EXEC getLanguageById '" + langId + "'"
	//fmt.Println(sqlStatement)

	st, err := db.Prepare(sqlStatement)
	if err != nil{
		fmt.Printf("Incorrect format -  ")
		fmt.Print( err );
		os.Exit(1)
	}

	rows, err := st.Query()
	if err != nil {
		fmt.Printf("Exec SQL failed - ")
		fmt.Print( err )
		os.Exit(1)
	}

	var lang Language
	for rows.Next() {
		err = rows.Scan(&lang.Id, &lang.Title)
	}
	return lang
}

func updateLanguage(lang Language){

	sqlStatement := "EXEC Change_Languages '" + lang.Id + "', '" +  lang.Title + "'"
	//fmt.Println(sqlStatement)

	st, err := db.Prepare(sqlStatement)
	if err != nil{
		fmt.Printf("Incorrect format -  ")
		fmt.Print( err );
		os.Exit(1)
	}

	st.Query()
	return
}


func insertLanguage(lang Language){
	sqlStatement := "EXEC Insert_Languages '" + lang.Id + "', '" +  lang.Title + "'"
	//fmt.Println(sqlStatement)

	st, err := db.Prepare(sqlStatement)
	if err != nil{
		fmt.Printf("Incorrect format -  ")
		fmt.Print( err );
		os.Exit(1)
	}

	st.Query()
	return
}

func deleteLanguageById(langId string)int64{

	sqlStatement := "EXEC Delete_Language '" + langId + "'"
	//fmt.Println(sqlStatement)

	st, err := db.Prepare(sqlStatement)
	if err != nil{
		fmt.Printf("Incorrect format -  ")
		fmt.Print( err );
		os.Exit(-1)
	}


	//this give rows effected = 1 always
	result, er := st.Exec()
	retVal, er := result.RowsAffected()

	if er != nil{
		fmt.Print( err );
		return -1
	}

	//this always give Row Effected = 0
//	var retVal int
//	result, err:= st.Query()
//	err = result.Scan(&retVal)


	fmt.Printf("Row Effected: %d \n", retVal  )
	return retVal

	//why retVal always = 1?
	//SP return @@count
	//How to get this value?
}