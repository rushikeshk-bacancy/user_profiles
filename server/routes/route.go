package routes

import (
	handlers "server/handlers"

	"github.com/gin-gonic/gin"
)

func Init(r *gin.Engine) {
	r.POST("/add", handlers.AddUser())
	r.POST("/update", handlers.UpdateUser())
	r.GET("/view", handlers.GetUsers())
	r.POST("/getUser", handlers.GetOneUser())
	r.POST("/delete", handlers.RemoveUser())
	r.POST("/upload", handlers.UploadUserPicture())
}
