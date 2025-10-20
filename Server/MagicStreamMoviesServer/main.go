package main
import (
	"fmt"
	controllers "github.com/MemoMeto35/Server/MagicStreamMoviesServer/controllers"
	"github.com/gin-gonic/gin"
)
func main(){
	router := gin.Default()
	router.GET("/hello", func(c *gin.Context){
		c.String(200, "Hello, MovieStream!")
	})
	router.GET("/movies", controllers.GetMovies())
	if err:=router.Run(":8080"); err!=nil{
		fmt.Println("Failed to start server:", err)
	}
}