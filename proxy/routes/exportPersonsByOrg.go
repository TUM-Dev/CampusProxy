package routes

import (
	"encoding/xml"
	"github.com/gin-gonic/gin"
	"net/http"
)

type exportPersonsByOrgRequest struct {
	OrgUnitID       string `json:"orgUnitID" form:"orgUnitID" xml:"orgUnitID" validate:"required"`
	Language        string `json:"language" form:"language" xml:"language" validate:"optional" enums:"DE,EN"`
	IncludeChildren *bool  `json:"includeChildren" form:"includeChildren" xml:"includeChildren" default:"true" validate:"optional"` // whether to include sub-organizations persons.
	Depth           *int   `json:"depth" form:"depth" xml:"depth" default:"99" validate:"optional"`                                 // how deep to include children (only relevant if includeChildren is true).
}

// ExportPersonsByOrg returns all persons of an organization.
//
// @Summary export persons of an organization.
// @Description This endpoint returns all persons that work at an organization.
// @Failure 400 {object} string "Bad Request"
// @Failure 401 {object} ApiError "Unauthorized"
// @Failure 404 {object} ApiError "Not Found"
// @Success 200 {object} []Person "Persons"
// @Accept json,xml
// @Param SearchRequest query exportPersonsByOrgRequest true "request"
// @Produce json
// @Router /organization/persons [get]
// @Tags organization
// @Security ApiKeyAuth
func ExportPersonsByOrg(c *gin.Context) {
	var req exportPersonsByOrgRequest
	err := c.Bind(&req)
	if err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	// fill defaults:
	if req.Depth == nil {
		req.Depth = new(int)
		*req.Depth = 99
	}
	if req.IncludeChildren == nil {
		req.IncludeChildren = new(bool)
		*req.IncludeChildren = true
	}

	resp := campusForward(c, req)
	var cdm exportPersonsByOrgResp
	err = xml.NewDecoder(resp.Body).Decode(&cdm)
	if err != nil {
		c.Set("ApiError", resp.StatusCode)
		return
	}
	c.JSON(200, cdm.OrgUnit.getPersons(*req.IncludeChildren, *req.Depth))
}

type exportPersonsByOrgResp struct {
	XMLName                   xml.Name       `xml:"CDM"`
	Text                      string         `xml:",chardata"`
	Xsi                       string         `xml:"xsi,attr"`
	NoNamespaceSchemaLocation string         `xml:"noNamespaceSchemaLocation,attr"`
	Properties                Properties     `xml:"properties"`
	OrgUnit                   OrgUnitWPerson `xml:"orgUnit"`
}

func (u OrgUnitWPerson) getPersons(includeSubOrgs bool, depth int) []Person {
	var persons []Person
	persons = append(persons, u.Persons...)
	if includeSubOrgs && depth > 0 {
		for _, o := range u.OrgUnits {
			persons = append(persons, o.getPersons(includeSubOrgs, depth+1)...)
		}
	}
	return persons
}

type OrgUnitWPerson struct {
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
	InfoBlock InfoBlock        `xml:"infoBlock" json:"infoBlock"`
	OrgUnits  []OrgUnitWPerson `xml:"orgUnit" json:"orgUnits,omitempty"`
	Persons   []Person         `xml:"person"`
}
