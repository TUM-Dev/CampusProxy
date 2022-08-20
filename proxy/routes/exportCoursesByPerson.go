package routes

import (
	"encoding/xml"
	"github.com/gin-gonic/gin"
	"net/http"
)

type exportCoursesByPersonRequest struct {
	PersonID     string `json:"personID" form:"personID" xml:"personID" validate:"required"`
	TeachingTerm string `json:"teachingTerm" form:"teachingTerm" xml:"teachingTerm" validate:"optional" enums:"W,S"`
	Language     string `json:"language" form:"language" xml:"language" validate:"optional" enums:"DE,EN"`
}

// ExportCoursesByPerson returns the courses a person is teaching/participating in (e.g. organising).
//
// @Summary export person's courses.
// @Description This endpoint returns all courses a person participates in (e.g. by teaching it, or being another part of the team). This endpoint does **not** return courses a student is enrolled in.
// @Failure 400 {object} string "Bad Request"
// @Failure 401 {object} ApiError "Unauthorized"
// @Failure 404 {object} ApiError "Not Found"
// @Success 200 {object} []Course "Courses"
// @Accept json,xml
// @Param SearchRequest query exportCoursesByPersonRequest true "request"
// @Produce json
// @Router /person/courses [get]
// @Tags person
// @Security ApiKeyAuth
func ExportCoursesByPerson(c *gin.Context) {
	var req exportCoursesByPersonRequest
	err := c.Bind(&req)
	if err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	resp := campusForward(c, req)
	var cdm coursesByPersonResp
	err = xml.NewDecoder(resp.Body).Decode(&cdm)
	if err != nil {
		c.Set("ApiError", resp.StatusCode)
		return
	}
	c.JSON(200, cdm.Course)
}

type coursesByPersonResp struct {
	XMLName                   xml.Name   `xml:"CDM"`
	Text                      string     `xml:",chardata"`
	Xsi                       string     `xml:"xsi,attr"`
	NoNamespaceSchemaLocation string     `xml:"noNamespaceSchemaLocation,attr"`
	Language                  string     `xml:"language,attr"`
	Properties                Properties `xml:"properties"`
	Course                    []Course   `xml:"course"`
}
