package proxy

import (
	"CampusOnlineWebservices/proxy/routes"
	"fmt"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"

	_ "CampusOnlineWebservices/docs"
)

type Proxy struct {
}

// Start starts the CAMPUSOnline Webservice proxy server
func (p *Proxy) Start() {
	r := gin.New()
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

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
	api.GET("/course/events", routes.ExportCourseEvents)

	api.GET("/organization", routes.ExportOrganization)
	fmt.Println(http.ListenAndServe(":8020", r))
}

type ApiError struct {
	Status int
	Msg    string
}
