package main
import (
	"fmt"								//format library
	"database/sql"						//first import to work with database
	_"github.com/denisenkom/go-mssqldb" //using mssql driver
	"log"								//log
	"os"
	"strconv"
	"time"
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
 * GET ALL MEDIA
 */
func getAllMedia() Medias{
	sqlStr:="EXEC getAllMedias"
	st, err := db.Prepare(sqlStr)
	if err != nil{
		fmt.Print( err );
		os.Exit(1)
	}

	rows, err := st.Query()
	if err != nil {
		fmt.Print( err )
		os.Exit(1)
	}

	var mediaCol Medias
	var tempMedia Media
	i := 0
	for rows.Next() {
		err = rows.Scan(&tempMedia.MediaId,
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
			&tempMedia.Duration,
			&tempMedia.Title,
			&tempMedia.ETitle,
			&tempMedia.Description,
			&tempMedia.EDescription,
			&tempMedia.LowImage,
			&tempMedia.MedImage,
			&tempMedia.HiImage	)

		if err != nil{
			fmt.Println(err)
		}
		i++
		mediaCol = append(mediaCol, tempMedia)
		fmt.Println(tempMedia)
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
	sqlStatement := "EXEC getMediabyTitle '%" + title + "%'"
	fmt.Println(sqlStatement)

	st, _ := db.Prepare(sqlStatement)
	rows,_ := st.Query()
	var mediaCol Medias
	var tempMedia Media
	i := 0
	for rows.Next() {
		err = rows.Scan(&tempMedia.MediaId,
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
						&tempMedia.Duration,
						&tempMedia.Title,
						&tempMedia.ETitle,
						&tempMedia.Description,
						&tempMedia.EDescription,
						&tempMedia.LowImage,
						&tempMedia.MedImage,
						&tempMedia.HiImage	)

		if err != nil{
			fmt.Println(err)
		}
		i++
		mediaCol = append(mediaCol, tempMedia)
		fmt.Println(tempMedia)
	}

	if(i== 0){
		return nil
	}
	return mediaCol
}

func getMediaByCondition(medTitle string, uView int) Medias{

	sqlStr:= "SELECT Media.*," +
	"Media_Title.titleShort,"+
	"Media_Title.titleExtended,"+
	"Media_Description.descriptionShort,"+
	"Media_Description.descriptionExpanded,"+
	"Thumbnail.lowQualityImage,"+
	"Thumbnail.mediumQualityImage,"+
	"Thumbnail.highQualityImage "+
	"FROM "+
	"Media INNER JOIN "+
	"Media_Title ON Media.mediaId = Media_Title.mediaId INNER JOIN "+
	"Media_Description ON Media.mediaId = Media_Description.mediaId INNER JOIN "+
	"Thumbnail ON Media.mediaId = Thumbnail.mediaId "+
	"WHERE " +
	"(Media_Title.titleShort LIKE '%?'"

	if(uView >= 0){
		sqlStr = sqlStr + " AND uniqueViews = " + strconv.Itoa(uView)

	}

	fmt.Println(sqlStr)

	st, _ := db.Prepare(sqlStr)
	rows,_ := st.Query(medTitle)
	var mediaCol Medias
	var tempMedia Media
	i := 0
	for rows.Next() {
		err = rows.Scan(&tempMedia.MediaId,
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
			&tempMedia.Duration,
			&tempMedia.Title,
			&tempMedia.ETitle,
			&tempMedia.Description,
			&tempMedia.EDescription,
			&tempMedia.LowImage,
			&tempMedia.MedImage,
			&tempMedia.HiImage	)

		if err != nil{
			fmt.Println(err)
		}
		i++
		mediaCol = append(mediaCol, tempMedia)
		fmt.Println(tempMedia)
	}

	if(i== 0){
		return nil
	}
	return mediaCol
}

func insertMedia(media Media, relInfo RelatedData){
	const longForm = "Jan 2, 2006 at 3:04pm (MST)"
	const shortForm = "2006-Jan-02"

//	sqlStatement := "EXEC Insert_Media "+
//	"'" + media.MediaId + "'," +
//	"'" + media.MediaHead + "'," +
//	"'" + media.Codec + "'," +
//	"'" + media.Container + "'," +
//	"" + strconv.Itoa(media.UniqueViews) + "," +
//	"'" + media.TargetDemographic + "'," +
//	"" + strconv.Itoa(media.SizeInBytes) + "," +
//	"'" + media.UploadedDate.Format(longForm) + "'," +
//	"" + strconv.Itoa(media.WidthInPixels) + "," +
//	"" +strconv.Itoa( media.HeightInPixels) + "," +
//	"'" + media.Orientation + "'," +
//	"" + strconv.FormatBool(media.DefinitionHD) + "," +
//	"'" + media.RegionOfOrgin + "'," +
//	"'" + media.ModifiedDate.Format(longForm) + "'," +
//	"'" + media.Duration + "'," +
//	"'" + generateId(0, "TT") + "'," +
//	"'" + relInfo.LanguageId + "'," +                     //emp
//	"'" + media.TitleRegion + "'," +                   //emp
//	"'" + media.Title + "'," +
//	"'" + media.ETitle + "'," +
//	"'" + generateId(0, "DE") + "'," +
//	"'" + media.Description + "'," +
//	"'" + media.EDescription + "'"

	sqlStatement := "EXEC	Insert_Media " +
	"@mediaId = '" +media.MediaId + "'," +
	"@mediaHead = '" + media.MediaHead + "'," +
	"@codec = '" + media.Codec + "'," +
	"@container = '" + media.Container + "'," +
	"@uniqueViews = " + strconv.Itoa(media.UniqueViews) + "," +
	"@targetDemographic = '" + media.TargetDemographic + "'," +
	"@sizeInBytes = " +  strconv.Itoa(media.SizeInBytes) + "," +
	"@uploadedDate = '" +  media.UploadedDate.Format(shortForm) + "'," +
	"@widthInPixels = " + strconv.Itoa(media.WidthInPixels) + "," +
	"@heightInPixels = " + strconv.Itoa( media.HeightInPixels) + "," +
	"@orientation = '" +media.Orientation + "'," +
	"@definitionHD = " + strconv.FormatBool(media.DefinitionHD) + "," +
	"@regionOfOrgin = '" + media.RegionOfOrgin + "'," +
	"@modifiedDate = '" +  media.ModifiedDate.Format(shortForm) + "'," +
	"@duration = '" +  media.Duration + "'," +
	"@titleId = '" + generateId(autoInt, "TT") + "'," +
	"@languageId = '" + relInfo.LanguageId + "'," +
	"@region = '" + media.TitleRegion + "'," +
	"@titleShort = '" + media.Title + "'," +
	"@titleExtended = '" + media.ETitle + "'," +
	"@descriptionId = '" + generateId(autoInt, "DE") + "'," +
	"@descriptionShort = '" + media.Description + "'," +
	"@descriptionExpanded = '"+ media.EDescription + "'"


	//fmt.Println(sqlStatement)
	st, err := db.Prepare(sqlStatement)
	if err != nil{
		fmt.Printf("Incorrect format -  ")
		fmt.Print( err );
		os.Exit(1)
	}

	rst, er :=st.Exec();
	if er != nil{
		fmt.Print(er)
		os.Exit(1)
	}

	fmt.Println(rst.LastInsertId())
	autoInt++
	return

}

/**
 * CONVERT String to Datetime
 */
func str2Date(dateStr string, dateForm string)time.Time{
	//get infor from form
	const longForm = "Jan 2, 2006 at 3:04pm (MST)"
	const shortForm = "2006-Jan-02"
	var datetimeVal time.Time
	var err error
	if(dateForm == "L"){
		datetimeVal,err = time.Parse(longForm, dateStr)
	}else{
		datetimeVal,err = time.Parse(shortForm, dateStr)
	}

	if(err != nil){
		return time.Time{}
	}
	return datetimeVal
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