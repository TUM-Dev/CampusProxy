package routes

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

type exportCourseRequest struct {
	CourseID string  `json:"courseID" form:"courseID" xml:"courseID" `
	Language *string `json:"language" form:"language" xml:"language"` // DE or EN, optional
}

// ExportCourse returns the course information for the given course ID.
//
// @Summary export course information.
// @Description This endpoint returns all information about a course, e.g. its title, description, semester, contacts and more.
// @Failure 400 {object} string "Bad Request"
// @Failure 401 {object} ApiError "Unauthorized"
// @Failure 404 {object} ApiError "Not Found"
// @Success 200 {object} Course "Course"
// @Accept json,xml
// @Param SearchRequest query exportCourseRequest true "request"
// @Produce json
// @Router /course [get]
// @Tags course
// @Security ApiKeyAuth
func ExportCourse(c *gin.Context) {
	fmt.Println("ExportCourse")
	var req exportCourseRequest
	err := c.Bind(&req)
	if err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	resp := campusForward(c, req)
	var cdm CDM
	err = xml.NewDecoder(resp.Body).Decode(&cdm)
	if err != nil {
		c.Set("ApiError", resp.StatusCode)
		return
	}
	c.JSON(200, cdm.Course)
}

/**
 * types:
**/

// XmlText is the text body of all xml nodes. Since most nodes just contain children, the text is mostly empty or whitespace.
// This struct therefore trims all whitespace from the text when encoding to json.
type XmlText string

func (t XmlText) MarshalJSON() ([]byte, error) {
	return json.Marshal(strings.Trim(string(t), " \n\t"))
}

type Picture struct {
	Text    XmlText `xml:",chardata" json:"text,omitempty"`
	WebLink *struct {
		Text        XmlText `xml:",chardata" json:"text,omitempty"`
		UserDefined string  `xml:"userDefined,attr" json:"user_defined,omitempty"`
		Href        string  `xml:"href" json:"href,omitempty"`
	} `xml:"webLink" json:"web_link"`
}

type WebLink struct {
	Text        XmlText `xml:",chardata" json:"-" swaggerignore:"true"`
	UserDefined string  `xml:"userDefined,attr" json:"user_defined,omitempty"`
	Href        string  `xml:"href" json:"href,omitempty"`
}

type InfoBlock struct {
	Text     XmlText  `xml:",chardata" json:"text,omitempty"`
	WebLink  *WebLink `xml:"webLink" json:"web_link"`
	SubBlock *struct {
		Text        XmlText `xml:",chardata" json:"text,omitempty"`
		UserDefined string  `xml:"userDefined,attr" json:"user_defined,omitempty"`
		SubBlock    string  `xml:"subBlock" json:"sub_block,omitempty"`
	} `xml:"subBlock" json:"sub_block"`
	Picture *Picture `xml:"picture" json:"picture"`
}

type ContactData struct {
	Text        XmlText `xml:",chardata" json:"text,omitempty"`
	ContactName *struct {
		Chardata string  `xml:",chardata" json:"chardata,omitempty"`
		Text     XmlText `xml:"text" json:"text,omitempty"`
	} `xml:"contactName" json:"contact_name"`
	Adr *struct {
		Text     XmlText `xml:",chardata" json:"text,omitempty"`
		Extadr   string  `xml:"extadr" json:"extadr,omitempty"`
		Street   string  `xml:"street" json:"street,omitempty"`
		Locality string  `xml:"locality" json:"locality,omitempty"`
		Pcode    string  `xml:"pcode" json:"pcode,omitempty"`
		Country  string  `xml:"country" json:"country,omitempty"`
	} `xml:"adr" json:"adr"`
	Email     string `xml:"email" json:"email,omitempty"`
	Telephone *struct {
		Text    XmlText `xml:",chardata" json:"text,omitempty"`
		Teltype string  `xml:"teltype,attr" json:"teltype,omitempty"`
	} `xml:"telephone" json:"telephone"`
	WebLink *WebLink `xml:"webLink" json:"web_link"`
}

type Person struct {
	Text     XmlText `xml:",chardata" json:"text,omitempty"`
	Ident    string  `xml:"ident,attr"`
	PersonID string  `xml:"personID" json:"person_id,omitempty"`
	Name     struct {
		Text   XmlText `xml:",chardata" json:"text,omitempty"`
		Given  string  `xml:"given" json:"given,omitempty"`
		Family string  `xml:"family" json:"family,omitempty"`
	} `xml:"name" json:"name"`
	Role *struct {
		Chardata string  `xml:",chardata" json:"chardata,omitempty"`
		RoleID   string  `xml:"roleID,attr" json:"role_id,omitempty"`
		Text     XmlText `xml:"text" json:"text,omitempty"`
	} `xml:"role" json:"role"`
	ContactData ContactData `xml:"contactData" json:"contact_data"`
	InfoBlock   InfoBlock   `xml:"infoBlock" json:"info_block"`
}

type Course struct {
	Text       XmlText `xml:",chardata" json:"-" swaggerignore:"true"`
	Language   string  `xml:"language,attr" json:"language,omitempty"`
	TypeID     string  `xml:"typeID,attr" json:"type_id,omitempty"`
	TypeName   string  `xml:"typeName,attr" json:"type_name,omitempty"`
	CourseID   string  `xml:"courseID" json:"course_id,omitempty"`
	CourseName struct {
		Chardata string  `xml:",chardata" json:"-" swaggerignore:"true"`
		Text     XmlText `xml:"text" json:"text,omitempty"`
	} `xml:"courseName" json:"course_name"`
	CourseCode        string `xml:"courseCode" json:"course_code,omitempty"`
	CourseDescription string `xml:"courseDescription" json:"course_description,omitempty"`
	Level             struct {
		Text    XmlText `xml:",chardata" json:"-" swaggerignore:"true"`
		WebLink WebLink `xml:"webLink" json:"web_link"`
	} `xml:"level" json:"level"`
	TeachingTerm string `xml:"teachingTerm" json:"teaching_term,omitempty"`
	Credits      struct {
		Text         XmlText `xml:",chardata" json:"text,omitempty"`
		HoursPerWeek string  `xml:"hoursPerWeek,attr" json:"hours_per_week,omitempty"`
	} `xml:"credits" json:"credits"`
	LearningObjectives string `xml:"learningObjectives" json:"learning_objectives,omitempty"`
	AdmissionInfo      struct {
		Text                 XmlText `xml:",chardata" json:"-" swaggerignore:"true"`
		AdmissionDescription struct {
			Text        XmlText `xml:",chardata" json:"-" swaggerignore:"true"`
			UserDefined string  `xml:"userDefined,attr" json:"user_defined,omitempty"`
			WebLink     WebLink `xml:"webLink" json:"web_link"`
		} `xml:"admissionDescription" json:"admission_description"`
	} `xml:"admissionInfo" json:"admission_info"`
	FormOfTeaching      string `xml:"formOfTeaching" json:"form_of_teaching,omitempty"`
	FormOfAssessment    string `xml:"formOfAssessment" json:"form_of_assessment,omitempty"`
	InstructionLanguage struct {
		Text         XmlText `xml:",chardata" json:"-" swaggerignore:"true"`
		TeachingLang string  `xml:"teachingLang,attr" json:"teaching_lang,omitempty"`
	} `xml:"instructionLanguage" json:"instruction_language"`
	Syllabus XmlText `xml:"syllabus" json:"syllabus,omitempty"`
	Exam     struct {
		Text      XmlText   `xml:",chardata" json:"-" swaggerignore:"true"`
		InfoBlock InfoBlock `xml:"infoBlock" json:"info_block"`
	} `xml:"exam" json:"exam"`
	TeachingActivity struct {
		Text                 XmlText `xml:",chardata" json:"-" swaggerignore:"true"`
		TeachingActivityID   string  `xml:"teachingActivityID" json:"teaching_activity_id,omitempty"`
		TeachingActivityName struct {
			Chardata string  `xml:",chardata" json:"-" swaggerignore:"true"`
			Text     XmlText `xml:"text" json:"text,omitempty"`
		} `xml:"teachingActivityName" json:"teaching_activity_name"`
		InfoBlock InfoBlock `xml:"infoBlock" json:"info_block"`
	} `xml:"teachingActivity" json:"teaching_activity"`
	Contacts struct {
		Text    XmlText  `xml:",chardata" json:"-" swaggerignore:"true"`
		Persons []Person `xml:"person" json:"person,omitempty"`
	} `xml:"contacts" json:"contacts"`
}

type Properties struct {
	Text       XmlText `xml:",chardata" json:"text,omitempty"`
	Datasource string  `xml:"datasource" json:"datasource,omitempty"`
	Datetime   struct {
		Text XmlText `xml:",chardata" json:"text,omitempty"`
		Date string  `xml:"date,attr" json:"date,omitempty"`
		Time string  `xml:"time,attr" json:"time,omitempty"`
	} `xml:"datetime" json:"datetime"`
}

type CDM struct {
	XMLName                   xml.Name   `xml:"CDM" json:"xml_name"`
	Text                      XmlText    `xml:",chardata" json:"text,omitempty"`
	Xsi                       string     `xml:"xsi,attr" json:"xsi,omitempty"`
	NoNamespaceSchemaLocation string     `xml:"noNamespaceSchemaLocation,attr" json:"no_namespace_schema_location,omitempty"`
	Language                  string     `xml:"language,attr" json:"language,omitempty"`
	Properties                Properties `xml:"properties" json:"properties"`
	Course                    Course     `xml:"course" json:"course"`
}
