package main

import (
  "database/sql"
  "fmt"
  "github.com/gin-gonic/gin.v1"
  "github.com/jinzhu/gorm"

  _ "github.com/lib/pq"
)


const (
  host     = "localhost"
  port     = 5432
  // user     = "postgres"
  dbname   = "bbpractice"
)

// func main() {
//   psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
//     "password=%s dbname=%s sslmode=disable",
//     host, port, user, dbname)
//   db, err := sql.Open("postgres", psqlInfo)
//   if err != nil {
//     panic(err)
//   }
//   defer db.Close()
//
//   err = db.Ping()
//   if err != nil {
//     panic(err)
//   }
//
//   fmt.Println("Successfully connected!")
// }

//////////

var db *gorm.DB

func init() {
	//open a db connection
  //Sprintf prints a function into a string and .fmt formats it'
  psqlInfo := fmt.Sprintf("host=%s port=%d "+
    "dbname=%s",
    host, port, dbname)

	var err error
	db, err = gorm.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

  fmt.Println("Successfully connected!")
	//Migrate the schema
	db.AutoMigrate(&todoModel{})
}

//init is a magic function in Go that runs once is first run and does some stuff needed before the project itself runs https://medium.com/golangspec/init-functions-in-go-eac191b3860a

func main() {

	router := gin.Default()

	v1 := router.Group("/api/v1/todos")
	{
		v1.POST("/", createTodo)
		v1.GET("/", fetchAllTodo)
		v1.GET("/:id", fetchSingleTodo)
		v1.PUT("/:id", updateTodo)
		v1.DELETE("/:id", deleteTodo)
	}
	router.Run()

}

type (
	// todoModel describes a todoModel type
	todoModel struct {
		gorm.Model
		Title     string `json:"title"`
		Completed int    `json:"completed"`
	}

	// transformedTodo represents a formatted todo
	transformedTodo struct {
		ID        uint   `json:"id"`
		Title     string `json:"title"`
		Completed bool   `json:"completed"`
	}
)

// createTodo add a new todo
func createTodo(c *gin.Context) {
	completed, _ := strconv.Atoi(c.PostForm("completed"))
	todo := todoModel{Title: c.PostForm("title"), Completed: completed}
	db.Save(&todo)
	c.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "message": "Todo item created successfully!", "resourceId": todo.ID})
}

// fetchAllTodo fetch all todos
func fetchAllTodo(c *gin.Context) {
	var todos []todoModel
	var _todos []transformedTodo

	db.Find(&todos)

	if len(todos) <= 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No todo found!"})
		return
	}

	//transforms the todos for building a good response
	for _, item := range todos {
		completed := false
		if item.Completed == 1 {
			completed = true
		} else {
			completed = false
		}
		_todos = append(_todos, transformedTodo{ID: item.ID, Title: item.Title, Completed: completed})
	}
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": _todos})
}

// fetchSingleTodo fetch a single todo
func fetchSingleTodo(c *gin.Context) {
	var todo todoModel
	todoID := c.Param("id")

	db.First(&todo, todoID)

	if todo.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No todo found!"})
		return
	}

	completed := false
	if todo.Completed == 1 {
		completed = true
	} else {
		completed = false
	}

	_todo := transformedTodo{ID: todo.ID, Title: todo.Title, Completed: completed}
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": _todo})
}

// updateTodo update a todo
func updateTodo(c *gin.Context) {
	var todo todoModel
	todoID := c.Param("id")

	db.First(&todo, todoID)

	if todo.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No todo found!"})
		return
	}

	db.Model(&todo).Update("title", c.PostForm("title"))
	completed, _ := strconv.Atoi(c.PostForm("completed"))
	db.Model(&todo).Update("completed", completed)
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Todo updated successfully!"})
}

// deleteTodo remove a todo
func deleteTodo(c *gin.Context) {
	var todo todoModel
	todoID := c.Param("id")

	db.First(&todo, todoID)

	if todo.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No todo found!"})
		return
	}

	db.Delete(&todo)
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Todo deleted successfully!"})
}
