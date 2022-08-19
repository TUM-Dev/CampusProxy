package routes

import (
	"encoding/xml"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type exportCoursesByOrgRequest struct {
	OrgUnitID       string  `json:"orgUnitID" form:"orgUnitID" xml:"orgUnitID"`
	TeachingTerm    *string `json:"teachingTerm" form:"teachingTerm" xml:"teachingTerm" enums:"W,S" validate:"optional"`
	Language        *string `json:"language" form:"language" xml:"language" enums:"DE,EN" validate:"optional"`
	IncludeChildren *bool   `json:"includeChildren" form:"includeChildren" xml:"includeChildren" default:"true" validate:"optional"` // whether to include sub-organizations courses.
	Depth           *int    `json:"depth" form:"depth" xml:"depth" default:"99" validate:"optional"`                                 // how deep to include children (only relevant if includeChildren is true).
}

// ExportCoursesByOrg returns the all courses that belong to an organization.
//
// @Summary Export an organizations courses.
// @Description This endpoint returns all courses that belong to the specified organization and all courses of sub-orgs. If a course belongs to a subOrg, orgUnitID is the ID of the subOrg.
// @Failure 400 {object} string "Bad Request"
// @Failure 401 {object} ApiError "Unauthorized"
// @Failure 404 {object} ApiError "Not Found"
// @Success 200 {object} []Course "Courses"
// @Accept json,xml
// @Param SearchRequest query exportCoursesByOrgRequest true "request"
// @Produce json
// @Router /organization/courses [get]
// @Tags organization
// @Security ApiKeyAuth
func ExportCoursesByOrg(c *gin.Context) {
	fmt.Println("ExportCoursesByOrg")
	var req exportCoursesByOrgRequest
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
	var orgResp coursesByOrgResp
	err = xml.NewDecoder(resp.Body).Decode(&orgResp)
	if err != nil {
		c.Set("ApiError", resp.StatusCode)
		return
	}
	courses := make([]Course, 0)
	for _, o := range orgResp.OrgUnit {
		courses = append(courses, o.getCourses(*req.IncludeChildren, *req.Depth)...)
	}
	c.JSON(200, courses)
}

type orgUnitCourses struct {
	Text        string `xml:",chardata"`
	OrgUnitID   string `xml:"orgUnitID"` // 15426
	OrgUnitName struct {
		Chardata string `xml:",chardata"`
		Text     string `xml:"text"` // Informatik 1 - Lehrstuhl ...
	} `xml:"orgUnitName"`
	OrgUnitCode string `xml:"orgUnitCode"` // TUINI01
	OrgUnitKind struct {
		Text     string `xml:",chardata"`
		SubBlock struct {
			Text        string `xml:",chardata"` // Lehrstuhl
			UserDefined string `xml:"userDefined,attr"`
		} `xml:"subBlock"`
	} `xml:"orgUnitKind"`
	Contacts  string `xml:"contacts"`
	InfoBlock struct {
		Text    string `xml:",chardata"`
		WebLink struct {
			Text        string `xml:",chardata"`
			UserDefined string `xml:"userDefined,attr"`
			Href        string `xml:"href"` // https://campus.tum.de/tum...
		} `xml:"webLink"`
	} `xml:"infoBlock"`
	Course  []Course         `xml:"course"`
	OrgUnit []orgUnitCourses `xml:"orgUnit"`
}

type coursesByOrgResp struct {
	XMLName                   xml.Name         `xml:"CDM"`
	Text                      string           `xml:",chardata"`
	Xsi                       string           `xml:"xsi,attr"`
	NoNamespaceSchemaLocation string           `xml:"noNamespaceSchemaLocation,attr"`
	Properties                Properties       `xml:"properties"`
	OrgUnit                   []orgUnitCourses `xml:"orgUnit"`
}

func (u orgUnitCourses) getCourses(includeChildren bool, depth int) []Course {
	var courses []Course
	courses = append(courses, u.Course...)
	if includeChildren && depth > 0 {
		for _, o := range u.OrgUnit {
			courses = append(courses, o.getCourses(includeChildren, depth-1)...)
		}
	}
	return courses
}
