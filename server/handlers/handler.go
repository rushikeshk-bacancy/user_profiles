package handlers

import (
	"net/http"

	services "server/services"

	"github.com/gin-gonic/gin"
)

// Chnage this path according to your file structure
var storagePath = "/home/rushikeshk/Workspace/projects/user_profiles/server/App_File_Storage/"

// AddUser :  Handler for adding the user into the database
func AddUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestBody := services.NewUser()
		c.Bind(&requestBody)
		services.AddNewUser(requestBody)
		c.JSON(http.StatusOK, services.GetUsersList())
	}
}

// UpdateUser : Handler for updating the user in the database
func UpdateUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestBody := services.NewUser()
		c.Bind(&requestBody)
		services.UpdateUserService(requestBody)
		c.JSON(http.StatusOK, services.GetUsersList())
	}
}

// RemoveUser : Handler for removing the user in the database
func RemoveUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestBody := services.NewUser()
		c.Bind(&requestBody)
		services.RemoveUserService(requestBody)
		c.JSON(http.StatusOK, services.GetUsersList())
	}
}

// GetUsers : Handler for viewing all the users in the database
func GetUsers() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, services.GetUsersList())
	}
}

// GetOneUser : Handler for viewing one of the user in the database
func GetOneUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestBody := services.NewUser()
		c.Bind(&requestBody)
		// services.GetSpecifiedUser(requestBody)
		c.JSON(http.StatusOK, services.GetSpecifiedUser(requestBody))
	}
}

// UploadUserPicture : Handler for uploading the file on static server
func UploadUserPicture() gin.HandlerFunc {
	return func(c *gin.Context) {
		file, err := c.FormFile("file")
		// The file cannot be received.
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"message": "No file is received",
			})
			return
		}
		// Retrieve file information
		// extension := filepath.Ext(file.Filename)

		// The file is received, so let's save it
		if err := c.SaveUploadedFile(file, storagePath+file.Filename); err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"message": "Unable to save the file",
			})
			return
		}

		// File saved successfully. Return proper result
		c.JSON(http.StatusOK, gin.H{"message": "Your file has been successfully uploaded."})
	}
}
