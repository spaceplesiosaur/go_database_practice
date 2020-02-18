package main


//putting the package directly in source with your project makes it so that it has EXACTLY the right packages and versions for what you're doing

//GO will not let you import anything you aren't using
import (
  // "database/sql"
  "fmt"
  "net/http"
	// "strconv"
  "github.com/gin-gonic/gin"
  "github.com/jinzhu/gorm"

  _ "github.com/lib/pq"
)
//underscore is in place of the alias

const (
  host     = "localhost"
  port     = 5432
  // user     = "postgres"
  dbname   = "bbpractice"
)


var db *gorm.DB
// gorm is like knex

func init() {
	//open a db connection
  //Sprintf prints a function into a string and .fmt formats it'
  psqlInfo := fmt.Sprintf("host=%s port=%d "+
    "dbname=%s sslmode=disable" ,
    host, port, dbname)

	var err error
	db, err = gorm.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

  fmt.Println("Successfully connected!")
	//Migrate the schema
	db.AutoMigrate(&songModel{})
  //makes sure db table exists and set it up that way
  //the & makes sure this passes the memory address of the function so you're manipulating the orig, not the copy
}

//init is a magic function in Go that runs once is first run and does some stuff needed before the project itself runs https://medium.com/golangspec/init-functions-in-go-eac191b3860a

func main() {

	router := gin.Default()

	v1 := router.Group("/api/v1/songs")
	{
		// v1.POST("/", addSong)
		v1.GET("/", fetchAllSongs)
		// v1.GET("/:id", fetchSong)
		// v1.PUT("/:id", changeSong)
		// v1.DELETE("/:id", removeSong)
	}
	router.Run()

}

type (
	// todoModel describes a todoModel type
	songModel struct {
		gorm.Model
		Title          string `json:"title"`
		SpotifyId      string `json:"spotifyid"`
    URL            string `json:"url"`
    Delay          float64  `json:"delay"`
    AvBarDuration  float64 `json:"avbarduration"`
    Duration       float64 `json:"duration"`
    Tempo          float64 `json:"tempo"`
    TimeSignature  uint `json: "timesignature"`

	}

	// transformedTodo represents a formatted todo
	transformedSong struct {
		ID        uint   `json:"id"`
    Title          string `json:"title"`
		SpotifyId      string `json:"spotifyid"`
    URL            string `json:"url"`
    Delay          float64  `json:"delay"`
    AvBarDuration  float64 `json:"avbarduration"`
    Duration       float64 `json:"duration"`
    Tempo          float64 `json:"tempo"`
    TimeSignature  uint `json: "timesignature"`
	}
)

// addSong add a new todo
// func addSong(c *gin.Context) {
	// song := songModel{Title: c.PostForm("title"), SpotifyId: c.PostForm("spotifyid"), URL: c.PostForm("url"), Delay: c.PostForm("delay"), AvBarDuration: c.PostForm("avbarduration"), Duration: c.PostForm("duration"), Tempo: c.PostForm("tempo"), TimeSignature: c.PostForm("timesignature")}
	// db.Save(&song)
	// c.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "message": "Song created successfully!", "resourceId": song.ID})
// }

// fetchAllSongs fetch all todos

//context is both a request and a response

func fetchAllSongs(c *gin.Context) {
	var songs []songModel
	// var _songs []transformedSong

	db.Find(&songs)

	if len(songs) <= 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No songs found!"})
		return
	}


	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": songs})
}

// fetchSong fetch a single todo
// func fetchSong(c *gin.Context) {
// 	var song songModel
// 	songID := c.Param("id")
//
// 	db.First(&song, songID)
//
// 	if song.ID == 0 {
// 		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No song found!"})
// 		return
// 	}

	// completed := false
	// if todo.Completed == 1 {
	// 	completed = true
	// } else {
	// 	completed = false
	// }

// 	_song := transformedSong{ID: song.ID, Title: song.Title}
// 	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": _song})
// }

// changeSong update a todo
// func changeSong(c *gin.Context) {
// 	var song songModel
// 	songID := c.Param("id")
//
// 	db.First(&song, songID)
//
// 	if song.ID == 0 {
// 		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No song found!"})
// 		return
// 	}
//
// 	db.Model(&song).Update("title", c.PostForm("title"))
// 	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Song updated successfully!"})
// }

// removeSong remove a todo
// func removeSong(c *gin.Context) {
// 	var song songModel
// 	songID := c.Param("id")
//
// 	db.First(&song, songID)
//
// 	if song.ID == 0 {
// 		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No song found!"})
// 		return
// 	}
//
// 	db.Delete(&song)
// 	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Song deleted successfully!"})
// }
