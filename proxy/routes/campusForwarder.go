package routes

import (
	"encoding/xml"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sonh/qs"
	"log"
	"net/http"
)

const baseURL = "https://campus.tum.de"

var routeMaps = map[string]string{
	"/api/v1/course":               "/tumonlinej/ws/webservice_v1.0/cdm/course/xml",
	"/api/v1/course/students":      "/tumonlinej/ws/webservice_v1.0/cdm/course/students/xml",
	"/api/v1/course/events":        "/tumonlinej/ws/webservice_v1.0/rdm/course/events/xml",
	"/api/v1/organization":         "/tumonlinej/ws/webservice_v1.0/cdm/organization/xml",
	"/api/v1/organization/courses": "/tumonlinej/ws/webservice_v1.0/cdm/organization/courses/xml",
	"/api/v1/organization/persons": "/tumonlinej/ws/webservice_v1.0/cdm/organization/persons/xml",
	"/api/v1/person/":              "/tumonlinej/ws/webservice_v1.0/cdm/person/xml",
	"/api/v1/person/courses":       "/tumonlinej/ws/webservice_v1.0/cdm/person/courses/xml",
}

func campusForward(c *gin.Context, req any) *http.Response {
	encoder := getEncoder()
	values, err := encoder.Values(req)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	values.Set("token", c.MustGet("token").(string))
	u := baseURL + routeMaps[c.Request.URL.Path] + "?" + values.Encode()
	log.Println(u)
	resp, err := http.Get(u)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return resp
}

func getEncoder() *qs.Encoder {
	return qs.NewEncoder(qs.WithTagAlias("form"))
}

type Error struct {
	XMLName xml.Name `xml:"Error" json:"-"`
	Text    string   `xml:",chardata" json:"-"`
	Message string   `xml:"Message"`
	Status  int      `xml:"-" json:"status"`
}

type ApiError struct {
	Status int    `json:"status" example:"420"`
	Msg    string `json:"msg" example:"CampusOnline returned a status code != 200."`
}
