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
			c.AbortWithStatusJSON(c.MustGet("ApiError").(int), routes.ApiError{
				Status: c.MustGet("ApiError").(int),
				Msg:    "CampusOnline returned a status code != 200.",
			})
		}
		if len(c.Errors) > 0 {
			c.Writer.Write([]byte(c.Errors[0].Error()))
		}
	})

	course := api.Group("/course")
	{
		course.GET("", routes.ExportCourse)
		course.GET("/students", routes.ExportCourseStudents)
		course.GET("/events", routes.ExportCourseEvents)
	}

	org := api.Group("/organization")
	{
		org.GET("", routes.ExportOrganization)
		org.GET("/courses", routes.ExportCoursesByOrg)
		org.GET("/persons", routes.ExportPersonsByOrg)
	}

	persons := api.Group("/person")
	{
		persons.GET("", routes.ExportPerson)
		persons.GET("/courses", routes.ExportCoursesByPerson)
	}
	fmt.Println("Started proxy server.")
	fmt.Println(http.ListenAndServe(":8020", r))
}
