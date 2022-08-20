package routes

import (
	"encoding/xml"
	"github.com/gin-gonic/gin"
	"net/http"
)

type exportPersonRequest struct {
	PersonID string `json:"personID" form:"personID" xml:"personID" validate:"required"`
	Language string `json:"language" form:"language" xml:"language" validate:"optional" enums:"DE,EN"`
}

// ExportPerson returns all info about a person.
//
// @Summary export person.
// @Description This endpoint returns all info about a person.
// @Failure 400 {object} string "Bad Request"
// @Failure 401 {object} ApiError "Unauthorized"
// @Failure 404 {object} ApiError "Not Found"
// @Success 200 {object} Person "Person"
// @Accept json,xml
// @Param SearchRequest query exportPersonRequest true "request"
// @Produce json
// @Router /person [get]
// @Tags person
// @Security ApiKeyAuth
func ExportPerson(c *gin.Context) {
	var req exportPersonRequest
	err := c.Bind(&req)
	if err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	resp := campusForward(c, req)
	var cdm exportPersonResp
	err = xml.NewDecoder(resp.Body).Decode(&cdm)
	if err != nil {
		c.Set("ApiError", resp.StatusCode)
		return
	}
	c.JSON(200, cdm.Person)
}

type exportPersonResp struct {
	XMLName                   xml.Name   `xml:"CDM"`
	Text                      string     `xml:",chardata"`
	Xsi                       string     `xml:"xsi,attr"`
	NoNamespaceSchemaLocation string     `xml:"noNamespaceSchemaLocation,attr"`
	Language                  string     `xml:"language,attr"`
	Properties                Properties `xml:"properties"`
	Person                    Person     `xml:"person"`
}
