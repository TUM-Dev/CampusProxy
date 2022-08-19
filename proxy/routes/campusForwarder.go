package routes

import (
	"encoding/xml"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sonh/qs"
	"net/http"
)

const baseURL = "https://campus.tum.de"

var routeMaps = map[string]string{
	"/api/v1/course":          "/tumonlinej/ws/webservice_v1.0/cdm/course/xml",
	"/api/v1/course/students": "/tumonlinej/ws/webservice_v1.0/cdm/course/students/xml",
	"/api/v1/course/events":   "/tumonlinej/ws/webservice_v1.0/rdm/course/events/xml",
	"/api/v1/organization":    "/tumonlinej/ws/webservice_v1.0/cdm/organization/xml",
}

func campusForward(c *gin.Context, req any) *http.Response {
	encoder := getEncoder()
	values, err := encoder.Values(req)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	values.Set("token", c.MustGet("token").(string))
	resp, err := http.Get(baseURL + routeMaps[c.Request.URL.Path] + "?" + values.Encode())
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
