package main

import (
	"fathil/gin-file-upload-ca/request"
	"fathil/gin-file-upload-ca/usecase"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// Note: Add this line inside the server.go after the gin instance created
	// Set a lower memory limit for multipart forms (default is 32 MiB)
	router.MaxMultipartMemory = 8 << 20 // 8 MiB

	router.POST("file", func(c *gin.Context) {
		var request request.Request
		err := c.ShouldBind(&request)
		if err != nil {
			c.JSON(400, err.Error())
			return
		}

		err = usecase.Store(&request)
		if err != nil {
			c.JSON(400, err.Error())
			return
		}

		c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", request.File.Filename))
	})

	router.Run(":8080")
}
