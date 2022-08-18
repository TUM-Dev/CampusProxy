package proxy

import (
	"CampusOnlineWebservices/proxy/routes"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Proxy struct {
}

// Start starts the CAMPUSOnline Webservice proxy server
// @title CAMPUSOnline Webservice proxy
// @version 1.0
// @description This is the proxy server for CAMPUSOnline Webservices, enabling a nicely documented and uniform rest access to CAMPUSOnline resources.
// @host localhost:8081
// @BasePath /api/v1
// @schemes http https
func (p *Proxy) Start() {
	r := gin.New()
	api := r.Group("/api/v1")
	api.Use(func(c *gin.Context) {
		type r struct {
			Token string `json:"token" form:"token" header:"x-api-key"`
		}
		var req r
		_ = c.Bind(&req)
		if req.Token != "" {
			c.Set("token", req.Token)
			return
		}
		_ = c.BindHeader(&req)
		c.Set("token", req.Token)
	})
	api.Use(func(c *gin.Context) {
		c.Next()
		if _, exists := c.Get("ApiError"); exists {
			c.AbortWithStatusJSON(c.MustGet("ApiError").(int), ApiError{
				Status: c.MustGet("ApiError").(int),
				Msg:    "CampusOnline returned a status code != 200.",
			})
		}
		if len(c.Errors) > 0 {
			c.Writer.Write([]byte(c.Errors[0].Error()))
		}
	})
	api.GET("/course", routes.ExportCourse)
	api.GET("/course/students", routes.ExportCourseStudents)
	fmt.Println(http.ListenAndServe(":8020", r))
}

type ApiError struct {
	Status int
	Msg    string
}
