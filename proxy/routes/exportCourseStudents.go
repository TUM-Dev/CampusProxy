package routes

import (
	"encoding/xml"
	"github.com/gin-gonic/gin"
	"net/http"
)

type exportCourseStudentsRequest struct {
	CourseID string  `json:"courseID" form:"courseID" xml:"courseID" `
	Language *string `json:"language" form:"language" xml:"language"` // DE or EN, optional
}

// ExportCourseStudents returns a list of all students in the given course.
//
// @Summary list enrolled students for course.
// @Description This endpoint returns a list containing all students enrolled in a course.
// @Failure 400 {object} string "Bad Request"
// @Failure 401 {object} Error "Unauthorized"
// @Failure 404 {object} Error "Not Found"
// @Success 200 {object} []Person "Students"
// @Accept json,xml
// @Param SearchRequest query exportCourseStudentsRequest true "request"
// @Produce json
// @Router /api/v1/course/students [get]
// @Tags course
func ExportCourseStudents(c *gin.Context) {
	var req exportCourseStudentsRequest
	err := c.Bind(&req)
	if err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	resp := campusForward(c, req)
	var cdm CourseStudentResp
	err = xml.NewDecoder(resp.Body).Decode(&cdm)
	if err != nil {
		c.Set("ApiError", resp.StatusCode)
		return
	}
	c.JSON(200, cdm.Person)
}

/**
 * types:
 */

type CourseStudentResp struct {
	XMLName                   xml.Name   `xml:"CDM" json:"XMLName"`
	Text                      XmlText    `xml:",chardata" json:"text,omitempty"`
	Xsi                       string     `xml:"xsi,attr" json:"xsi,omitempty"`
	NoNamespaceSchemaLocation string     `xml:"noNamespaceSchemaLocation,attr" json:"noNamespaceSchemaLocation,omitempty"`
	Language                  string     `xml:"language,attr" json:"language,omitempty"`
	Properties                Properties `xml:"properties" json:"properties"`
	Course                    Course     `xml:"course" json:"course"`
	Person                    []Person   `xml:"person" json:"person,omitempty"`
}
