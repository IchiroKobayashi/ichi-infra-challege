package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/ichi-infra-challenge/docker-local/api/src/model"
)

func main() {
	// Start HTTP server
	r := gin.Default()
	r.GET("/search", searchExample)
	r.GET("/create", createData)

	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}

//あいまい文字列検索
func searchExample(c *gin.Context) {
	searchTerm := c.Query("search")
	if searchTerm == "" {
		c.JSON(http.StatusInternalServerError, "search not specified")
		return
	}

	resp, err := model.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, "Something went wrong")
		return
	}
	c.IndentedJSON(http.StatusOK, resp)
}

func createData(c *gin.Context) {

	name := c.Query("search")
	resp, err := model.Create(name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "Something went wrong")
		return
	}
	c.IndentedJSON(http.StatusOK, resp)
}
