package main
import (
	"fmt"								//format library
	"database/sql"						//first import to work with database
	_"github.com/denisenkom/go-mssqldb" //using mssql driver
	"log"								//log
	"os"
	//"strconv"


)

	var db *sql.DB
	var err error
	var autoInt int = 0
func init() {
	fmt.Println("_Database connection inititalized.")
	db, err = sql.Open("mssql", "server=127.0.0.1;port=1433;user id=coredbadmin;password=ilovecut2it;database=media")
	if err != nil {
		log.Fatal(err)
		fmt.Println("Error!.")
		os.Exit(1)
	}

}

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

//CHANGE LANGUAGE INFO
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

//CREATE NEW LANGUAGE
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

//DELETE LANGUAGE
func deleteLanguageById(langId string)int64{

	sqlStatement := "EXEC Delete_Language '" + langId + "'"
	//fmt.Println(sqlStatement)

	st, err := db.Prepare(sqlStatement)
	if err != nil{
		fmt.Printf("Incorrect format -  ")
		fmt.Print( err );
		os.Exit(-1)
	}

	//get the number of affected rows
	result, er := st.Exec()
	retVal, er := result.RowsAffected()

	if er != nil{
		fmt.Print( err );
		return -1
	}

	//fmt.Printf("Row Effected: %d \n", retVal  )
	return retVal
}

/**
 * GET ALL MEDIA IN DATABASE
 */
func getAllMedias() Medias{
	sqlStr:="SELECT * FROM media"
	st, err := db.Prepare(sqlStr)
	checkErr(err);

	rows, err := st.Query()
	checkErr(err);


	var mediaCol Medias
	var tempMedia Media
	//var count int
	i := 0
	for rows.Next() {
		err = rows.Scan(
			&tempMedia.MediaId,
			&tempMedia.MediaHead,
			&tempMedia.Codec,
			&tempMedia.Container,
			&tempMedia.UniqueViews,
			&tempMedia.TargetDemographic,
			&tempMedia.SizeInBytes,
			&tempMedia.UploadedDate,
			&tempMedia.WidthInPixels,
			&tempMedia.HeightInPixels,
			&tempMedia.Orientation,
			&tempMedia.DefinitionHD,
			&tempMedia.RegionOfOrgin,
			&tempMedia.ModifiedDate,
			&tempMedia.Duration)
		tempMedia.TitleCol, _ = getTitleByMediaId(tempMedia.MediaId)
		tempMedia.DescriptionCol, _=  getDescriptionByMediaId(tempMedia.MediaId)
		tempMedia.ThumbnailCol = getThumbnailByMediaId(tempMedia.MediaId)
		checkErr(err)
		i++
		mediaCol = append(mediaCol, tempMedia)

	}
	fmt.Println("Total: " , i )

	//close the result set so that we can reopen it later
	st.Close()
	rows.Close()
	return mediaCol
}

/**
 * GET MEDIA BY TITLE
 */
func getMediaByTitle(title string) Medias{
	sqlStatement := "EXEC getMediabyTitle '" + title + "'"
	st, _ := db.Prepare(sqlStatement)
	rows,_ := st.Query()
	var mediaCol Medias
	var tempMedia Media
	i := 0
	for rows.Next() {
		//for media info
		err = rows.Scan(
						&tempMedia.MediaId,
						&tempMedia.MediaHead,
						&tempMedia.Codec,
						&tempMedia.Container,
						&tempMedia.UniqueViews,
						&tempMedia.TargetDemographic,
						&tempMedia.SizeInBytes,
						&tempMedia.UploadedDate,
						&tempMedia.WidthInPixels,
						&tempMedia.HeightInPixels,
						&tempMedia.Orientation,
						&tempMedia.DefinitionHD,
						&tempMedia.RegionOfOrgin,
						&tempMedia.ModifiedDate,
						&tempMedia.Duration)
		tempMedia.TitleCol, _ = getTitleByMediaId(tempMedia.MediaId)
		tempMedia.DescriptionCol, _=  getDescriptionByMediaId(tempMedia.MediaId)
		tempMedia.ThumbnailCol = getThumbnailByMediaId(tempMedia.MediaId)
		checkErr(err)
		i++
		mediaCol = append(mediaCol, tempMedia)

	}
	fmt.Println("#Media: " , i )

	//close the result set so that we can reopen it later
	st.Close()
	rows.Close()
	return mediaCol

}

func insertMedia(media Media){
	sqlStatement := "INSERT INTO Media(mediaId,mediaHead, codec,container, uniqueViews,targetDemographic,"+
					"sizeInBytes,uploadedDate,widthInPixels,heightInPixels,	orientation,definitionHD,"+
					"regionOfOrgin,	modifiedDate, duration)	VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)"
	//fmt.Println(sqlStatement)
	st, err := db.Prepare(sqlStatement)
	checkErr(err)

	rst, er :=st.Exec(
		media.MediaId,
		media.MediaHead,
		media.Codec,
		media.Container,
		media.UniqueViews,
		media.TargetDemographic,
		media.SizeInBytes,
		media.UploadedDate,
		media.WidthInPixels,
		media.HeightInPixels,
		media.Orientation,
		media.DefinitionHD,
		media.RegionOfOrgin,
		media.ModifiedDate,
		media.Duration)
	checkErr(er)

	count, er := rst.RowsAffected()
	fmt.Printf("rows affected: %d \n", count)
}

func insertTitle(media Media){
	sqlStatement := "INSERT INTO Media_Title (titleId, languageId, region, titleShort, titleExtended, mediaId) "+
					"VALUES (?, ?, ?, ?, ?, ?)"
	//fmt.Println(sqlStatement)
	st, err := db.Prepare(sqlStatement)
	checkErr(err)

	_, er :=st.Exec(
		media.TitleCol[0].TitleId,
		media.TitleCol[0].LanguageId,
		media.TitleCol[0].Region,
		media.TitleCol[0].TitleShort,
		media.TitleCol[0].TitleExtended,
		media.TitleCol[0].MediaId)
	checkErr(er)

}

func insertDescription(media Media){
	sqlStatement := "INSERT INTO Media_Description (descriptionId, languageId, region, descriptionShort, "+
					"descriptionExpanded, mediaId) VALUES (?, ? ,? ,? ,? ,? )"
	//fmt.Println(sqlStatement)
	st, err := db.Prepare(sqlStatement)
	checkErr(err)

	_, er :=st.Exec(
		media.DescriptionCol[0].DescriptionId,
		media.DescriptionCol[0].LanguageId,
		media.DescriptionCol[0].Region,
		media.DescriptionCol[0].DescriptionShort,
		media.DescriptionCol[0].DescriptionExpanded,
		media.DescriptionCol[0].MediaId)
	checkErr(er)
}

func insertThumb(media Media){
	sqlStatement := "INSERT INTO Thumbnail (thumbId, lowQualityImage, mediumQualityImage, highQualityImage, "+
	"targetDemographic, mediaId) VALUES (?,? ,? ,? ,? ,? )"
	//fmt.Println(sqlStatement)
	st, err := db.Prepare(sqlStatement)
	checkErr(err)


	_, er :=st.Exec(
		media.ThumbnailCol.ThumbId,
		media.ThumbnailCol.LowQualityImage,
		media.ThumbnailCol.MediumQualityImage,
		media.ThumbnailCol.HighQualityImage,
		media.ThumbnailCol.TargetDemographic,
		media.ThumbnailCol.MediaId)
	checkErr(er)
}

func getTitleByMediaId(mediaId string) ([]MediaTitle, int){
	sqlStr:="SELECT * FROM Media_Title WHERE mediaId = '" + mediaId +"'"
	st, err := db.Prepare(sqlStr)
	checkErr(err);

	rows, err := st.Query()
	checkErr(err);
	var titles []MediaTitle
	var tempTitle MediaTitle
	i := 0
	for rows.Next() {
		err = rows.Scan(
			&tempTitle.TitleId,
			&tempTitle.LanguageId,
			&tempTitle.Region,
			&tempTitle.TitleShort,
			&tempTitle.TitleExtended,
			&tempTitle.MediaId)
		i++
		titles = append(titles, tempTitle)
	}
	return titles, i
}

func getDescriptionByMediaId(mediaId string) ([]MediaDescription, int){
	sqlStr:="SELECT * FROM Media_Description WHERE mediaId = '" + mediaId +"'"
	st, err := db.Prepare(sqlStr)
	checkErr(err);

	rows, err := st.Query()
	checkErr(err);
	var descriptionCol []MediaDescription
	var tempDescription MediaDescription
	i := 0
	for rows.Next() {
		err = rows.Scan(
			&tempDescription.DescriptionId,
			&tempDescription.LanguageId,
			&tempDescription.Region,
			&tempDescription.DescriptionShort,
			&tempDescription.DescriptionExpanded,
			&tempDescription.MediaId)
		i++
		descriptionCol = append(descriptionCol, tempDescription)
	}
	return descriptionCol, i
}

func getThumbnailByMediaId(mediaId string) Thumbnail{
	sqlStr:="SELECT * FROM Thumbnail WHERE mediaId = '" + mediaId +"'"
	st, err := db.Prepare(sqlStr)
	checkErr(err);

	row := st.QueryRow()
	var thumb Thumbnail
	err = row.Scan(
			&thumb.ThumbId,
			&thumb.LowQualityImage,
			&thumb.MediumQualityImage,
			&thumb.HighQualityImage,
			&thumb.TargetDemographic,
			&thumb.MediaId)
	return thumb
}