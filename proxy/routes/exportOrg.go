package routes

import (
	"encoding/xml"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type exportOrganizationRequest struct {
	OrgUnitID string  `json:"orgUnitID" form:"orgUnitID" xml:"orgUnitID" `
	Language  *string `json:"language" form:"language" xml:"language"` // DE or EN, optional
}

// ExportOrganization returns the information of an organization and its children.
//
// @Summary export an organization.
// @Description This endpoint returns all information about an organization and its sub-organization.
// @Failure 400 {object} string "Bad Request"
// @Failure 401 {object} Error "Unauthorized"
// @Failure 404 {object} Error "Not Found"
// @Success 200 {object} OrgUnit "Organization"
// @Accept json,xml
// @Param SearchRequest query exportOrganizationRequest true "request"
// @Produce json
// @Router /organization [get]
// @Tags organization
// @Security ApiKeyAuth
func ExportOrganization(c *gin.Context) {
	fmt.Println("ExportOrg")
	var req exportOrganizationRequest
	err := c.Bind(&req)
	if err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	resp := campusForward(c, req)
	var cdm OrgRes
	err = xml.NewDecoder(resp.Body).Decode(&cdm)
	if err != nil {
		c.Set("ApiError", resp.StatusCode)
		return
	}
	c.JSON(200, cdm.OrgUnit)
}

// Types:

type OrgUnit struct {
	Text        XmlText `xml:",chardata" json:"text,omitempty"`
	OrgUnitID   string  `xml:"orgUnitID" json:"orgUnitID,omitempty"`
	OrgUnitName struct {
		Chardata string  `xml:",chardata" json:"chardata,omitempty"`
		Text     XmlText `xml:"text" json:"text,omitempty"`
	} `xml:"orgUnitName" json:"orgUnitName"`
	OrgUnitCode string `xml:"orgUnitCode" json:"orgUnitCode,omitempty"`
	OrgUnitKind struct {
		Text     XmlText `xml:",chardata" json:"text,omitempty"`
		SubBlock []struct {
			Text        XmlText `xml:",chardata" json:"text,omitempty"`               // e.g. "Lehrstuhl", "Arbeitsgruppe", "Ausschuss", ...
			UserDefined string  `xml:"userDefined,attr" json:"userDefined,omitempty"` // always "name"
		} `xml:"subBlock" json:"subBlock,omitempty"`
	} `xml:"orgUnitKind" json:"orgUnitKind"`
	Contacts struct {
		Text        XmlText     `xml:",chardata" json:"text,omitempty"`
		ContactData ContactData `xml:"contactData" json:"contactData"`
	} `xml:"contacts" json:"contacts"`
	InfoBlock InfoBlock `xml:"infoBlock" json:"infoBlock"`
	OrgUnits  []OrgUnit `xml:"orgUnit" json:"orgUnits,omitempty"`
}

type OrgRes struct {
	XMLName                   xml.Name   `xml:"CDM"`
	Text                      XmlText    `xml:",chardata"`
	Xsi                       string     `xml:"xsi,attr"`
	NoNamespaceSchemaLocation string     `xml:"noNamespaceSchemaLocation,attr"`
	Properties                Properties `xml:"properties"`
	OrgUnit                   OrgUnit    `xml:"orgUnit"`
}
