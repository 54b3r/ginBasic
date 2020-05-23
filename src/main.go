package main

import (
	"fmt"
	"io/ioutil"
	"log"

	jwt "github.com/54b3r/ginBasic/middleware"
	"github.com/gin-gonic/gin"
)

// IndexPage handler curl localhost:8080
func IndexPage(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "You have reached the index page!",
	})
}

// PostHomepage func curl -X POST localhost:8080 -d 'hello'
func PostHomepage(c *gin.Context) {
	body := c.Request.Body
	value, err := ioutil.ReadAll(body)
	if err != nil {
		log.Printf("[ERROR]: Could not read request body %s", err.Error())
	}

	c.JSON(200, gin.H{
		"message": string(value),
	})
}

// QueryStringParams use url query strings
// curl localhost:8080/query?name=james&age=45
func QueryStringParams(c *gin.Context) {
	name := c.Query("name")
	age := c.Query("age")

	c.JSON(200, gin.H{
		"message": "Using QueryString Parameters with Gin",
		"name":    name,
		"age":     age,
	})
}

// QueryPathParams use url query strings
// curl localhost:8080/query/james/45
func QueryPathParams(c *gin.Context) {
	name := c.Param("name")
	age := c.Param("age")

	c.JSON(200, gin.H{
		"message": "Using QueryPath Parameters with Gin",
		"name":    name,
		"age":     age,
	})
}

func JwtSample(c *gin.Context) {
	validToken, err := jwt.GenerateJWT()
	if err != nil {
		fmt.Fprintf(c.Writer, err.Error())
	}

	c.JSON(200, gin.H{
		"access_token": validToken,
	})
}

func main() {
	r := gin.Default()
	r.GET("/", IndexPage)
	r.POST("/", PostHomepage)
	r.GET("/query", QueryStringParams)          // /query?name=james&age=45
	r.GET("/query/:name/:age", QueryPathParams) // /path/james/45
	r.GET("/jwt/sample", JwtSample)

	r.Run()
}
