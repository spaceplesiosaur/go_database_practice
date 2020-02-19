//first I did brew install go.  This gave me access to the Go command line tools.
//In Go, all projects live in the same place (workspace).  This is by default a folder called 'Go' in your home directory, (this is what Go expects) and if you type go env GOPATH in your terminal you will see it.

//It's what Go expects, but you have to make one anyway.  So make a folder named go.

//Inside that folder I made SRC, which is where all of my source code goes - it's where I imported all of my go get's, which are like npm installs.  I think src is kind of like a package.json for all of my go projects.  My project folder is in SRC as well, and I am able to import the things I've installed there direction with just the file name.  Apparently all of your project folders go in here.

//putting the package directly in source with your project makes it so that it has EXACTLY the right packages and versions for what you're doing

//gofmt -w main.go will format stuff for you

package main

//Every go project starts with a package, and if you want it to work it starts with 'main'.  You can't just run arbitrary files, so you have to tell it which file will be the package that it runs.  This is sort of like your root.

//GO will not let you import anything you aren't using
import (
	// "database/sql"
	"fmt"
	//format - functions related to input and output
	"net/http"
	//this is nice for errors
	"strconv"
  //we need this for ParseFloat
	"github.com/gin-gonic/gin"
	//This is like express, it's a library for handling http requests and responses
	"github.com/jinzhu/gorm"
	//gorm is an ORM, and object relationship manager.  It is kind of like knex. This builds your database
	_ "github.com/lib/pq"
	//this is our import to use postgres
)

//underscore is in place of the alias

const (
	host = "localhost"
	port = 5432
	//postgres is running on 5432
	// user     = "postgres"
	dbname = "bbpractice"
)

//Here we are just defining what we want our host to be - for this db, we don't need a user or pw

var db *gorm.DB

// gorm is like knex.  this is wrapping our database in gorm and saying that when we call db we mean the gorm wrapped database. LOOK UP gorm.DB

func init() {
	//open a db connection
	//Sprintf prints a function into a string and .fmt formats it'
	psqlInfo := fmt.Sprintf("host=%s port=%d "+
		"dbname=%s sslmode=disable",
		host, port, dbname)
	//This is a connection string, a special way to connect to the pg db that the underlying postgres library uses.
	var err error
	//error is likely built into Go.

	db, err = gorm.Open("postgres", psqlInfo)

	//the way this is written is like destructuring in JS.  gorm.open takes two arguments, the name of the server, and the connection string.  It returns two values, the db connection object (db is of a type defined in gorm.DB) and the error.

	if err != nil {
		panic(err)
		//I'm guessing this is like throw in js.
	}

	fmt.Println("Successfully connected!")
	//fmt is a library for printing things nicely.  It is printing this line for us.
	db.AutoMigrate(&songModel{})
	//We now have a db called db of the type defined in gorm.DB.  It has methods on it, like AutoMigrate.
	//Migrate the schema according to the songModel table below
	//makes sure db table exists and set it up that way
	//the & makes sure this passes the memory address of the function so you're manipulating the orig, not the copy
}

//init is a magic function in Go that runs once is first run and does some stuff needed before the project itself runs https://medium.com/golangspec/init-functions-in-go-eac191b3860a

//main is the function that wraps all of our funtionality. It needs to be in every file.  It takes no arguments.  It is magic and runs when we run the program.

//variable := blah blah is a quick way of writing var variable = blah blah

func main() {

	router := gin.Default()
	//Default is a method on gin that return the default type of router it likes to use
	//now that router is the name for this router we called out of gin, we can run methods on it.  Group allows us to make a bunch of endpoints off of this one path.

	songs := router.Group("/api/v1/songs")
	{
		songs.POST("/", addSong)
		songs.GET("/", fetchAllSongs)
		songs.GET("/:id", fetchSong)
		// v1.PUT("/:id", changeSong)
		// v1.DELETE("/:id", removeSong)
	}
	router.Run()
	//this is the end of our main() function, and is telling the router to run and listen for incoming connections and requests.
}

type (
	// todoModel describes a todoModel type
	//gorm.Model is a structure inside gorm for a table
	songModel struct {
		gorm.Model
		Title         string  `json:"title"`
		SpotifyId     string  `json:"spotifyid"`
		URL           string  `json:"url"`
		Delay         float64 `json:"delay"`
		AvBarDuration float64 `json:"avbarduration"`
		Duration      float64 `json:"duration"`
		Tempo         float64 `json:"tempo"`
		TimeSignature uint    `json: "timesignature"`
	}

	// transformedTodo represents a formatted todo, how you show it in json
	transformedSong struct {
		ID            uint    `json:"id"`
		Title         string  `json:"title"`
		SpotifyId     string  `json:"spotifyid"`
		URL           string  `json:"url"`
		Delay         float64 `json:"delay"`
		AvBarDuration float64 `json:"avbarduration"`
		Duration      float64 `json:"duration"`
		Tempo         float64 `json:"tempo"`
		TimeSignature uint    `json: "timesignature"`
	}
)

//addSong add a new song

//Context is the most important part of gin. It allows us to pass variables between middleware, manage the flow, validate the JSON of a request and render a JSON response for example. https://godoc.org/github.com/gin-gonic/gin

func addSong(context *gin.Context) {

  //so...PostForm returns a string!  We said we wanted floats, so we have to parse the values from the form into floats.  The thing is, floats return a value and an error, so we have to do some error handling.

  delay, err := strconv.ParseFloat(context.PostForm("delay"), 64)
  if err != nil {
    context.JSON(http.StatusUnprocessableEntity, gin.H{"status": http.StatusUnprocessableEntity, "message": "delay should be a float!", "Err": err})
		return
	}
  avbarduration, err := strconv.ParseFloat(context.PostForm("avbarduration"), 64)
  if err != nil {
    context.JSON(http.StatusUnprocessableEntity, gin.H{"status": http.StatusUnprocessableEntity, "message": "avbarduration should be a float!"})
		return
	}
  duration, err := strconv.ParseFloat(context.PostForm("duration"), 64)
  if err != nil {
    context.JSON(http.StatusUnprocessableEntity, gin.H{"status": http.StatusUnprocessableEntity, "message": "duration should be a float!"})
		return
	}
  tempo, err := strconv.ParseFloat(context.PostForm("tempo"), 64)
  if err != nil {
    context.JSON(http.StatusUnprocessableEntity, gin.H{"status": http.StatusUnprocessableEntity, "message": "tempo should be a float!"})
		return
	}
  timesignature, err := strconv.ParseUint(context.PostForm("timesignature"), 10, 32)
  if err != nil {
    context.JSON(http.StatusUnprocessableEntity, gin.H{"status": http.StatusUnprocessableEntity, "message": "timesignature should be an integer!"})
		return
	}

  //Kit just knows that ParseUint expects a 32, I definitely don't see it in the docs.


  //https://golang.org/src/net/http/status.go for Go status codes
  //notice Delay and the others are now assigned to variables defined above.

	song := songModel{
    Title: context.PostForm("title"),
    SpotifyId: context.PostForm("spotifyid"),
    URL: context.PostForm("url"),
    Delay: delay,
    AvBarDuration: avbarduration,
    Duration: duration,
    Tempo: tempo,
    TimeSignature: uint(timesignature),
  }

  //The parse functions return the widest type (float64, int64, and uint64), but if the size argument specifies a narrower width the result can be converted to that narrower type without data loss:

	db.Save(&song)
  //this is a gorm method that saves the song to the database
	context.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "message": "Song created successfully!", "resourceId": song.ID})
}

// fetchAllSongs fetch all todos

//context is both a request and a response

func fetchAllSongs(context *gin.Context) {
	//what is a gin context?  It's probably like the request, response objects you pass in as arguments in express.
	var songs []songModel
	//this is saying that var songs is equal to an array that follows the structure of songModel

	db.Find(&songs)
	//db find, and for everything everything you find, put it into the songs array.
	// SELECT * FROM users;

	if len(songs) <= 0 {
		context.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No songs found!"})
		return
	}
	// gin.H is a shortcut for map[string]interface{},
	//A map in Go is like a hash, or object.  This is how Go makes objects - define it as a map, define the key as a string, and define the value as an interface.   The conext.json method takes a status, and then something it can convert to a string representation of json.  In this case, we're defining an object as we feed it in.

	//the function len applied to an array (in this case songs) is like array.length in js.  This is saying that if there's nothing in the array, return the error

	//c.json is referring to the context in the argument above - c was assigned to gin.context.  I renamed it 'context'

	context.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": songs})
}

// fetchSong fetch a single song

func fetchSong(context *gin.Context) {
	var song songModel
	songID := context.Param("id")
	//again, context is like the request/response objects in express - context.Param is like the request param.

	db.First(&song, songID)
	//First is a function that means find the first item that fits the arguments - a song with the passed in song id.

	// SELECT * FROM users WHERE id = 10;

	if song.ID == 0 {
		context.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No song found!"})
		return
	}

	_song := transformedSong{ID: song.ID, Title: song.Title, SpotifyId: song.SpotifyId, URL: song.URL, Delay: song.Delay, AvBarDuration: song.AvBarDuration, Duration: song.Duration, Tempo: song.Tempo, TimeSignature: song.TimeSignature}

	context.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": _song})
}

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
