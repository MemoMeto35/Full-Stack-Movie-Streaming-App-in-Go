package controllers
import(
	"context"
	"time"
	"github.com/gin-gonic/gin"
)
func GetMovies() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		var movies []
	}
}