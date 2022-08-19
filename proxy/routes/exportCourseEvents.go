package routes

import (
	"encoding/xml"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

type exportCourseEventsRequest struct {
	CourseID string `json:"courseID" form:"courseID" xml:"courseID" `
}

// ExportCourseEvents returns a list of all events in the given course.
//
// @Summary list events for course.
// @Description This endpoint returns a list containing all events of a course.
// @Failure 400 {object} string "Bad Request"
// @Failure 401 {object} Error "Unauthorized"
// @Failure 404 {object} Error "Not Found"
// @Success 200 {object} []SingleEvent "Events"
// @Accept json,xml
// @Param SearchRequest query exportCourseEventsRequest true "request"
// @Produce json
// @Router /course/events [get]
// @Tags course
// @Security ApiKeyAuth
func ExportCourseEvents(c *gin.Context) {
	var req exportCourseEventsRequest
	err := c.Bind(&req)
	if err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	resp := campusForward(c, req)
	var cdm courseEventsResp
	err = xml.NewDecoder(resp.Body).Decode(&cdm)
	if err != nil {
		c.Set("ApiError", resp.StatusCode)
		return
	}
	xmlEvents := cdm.Resource.Description.ResourceGroup.Description.Resource
	jsonEvents := make([]SingleEvent, len(xmlEvents))
	for i := range xmlEvents {
		jsonEvents[i] = xmlEvents[i].toJsonEvent()
	}
	c.JSON(200, jsonEvents)
}

// types:

type xmlSingleEvent struct {
	Text        string `xml:",chardata"`
	TypeID      string `xml:"typeID,attr"`
	Description struct {
		Text      string `xml:",chardata"`
		Attribute []struct {
			Text          string `xml:",chardata"`
			AttrID        string `xml:"attrID,attr"`
			AttrDataType  string `xml:"attrDataType,attr"`
			CharacterData string `xml:"characterData"`
		} `xml:"attribute"`
	} `xml:"description"`
}

func getOrNilStr(m *map[string]attrCharDataWrapper, key string) *string {
	if s, ok := (*m)[key]; ok {
		return &s.text
	}
	return nil
}

func getOrNilInt(m *map[string]attrCharDataWrapper, key string) *int {
	s := getOrNilStr(m, key)
	if s == nil {
		return nil
	}
	i, err := strconv.Atoi(*s)
	if err != nil {
		return nil
	}
	return &i
}

func getOrNilDt(m *map[string]attrCharDataWrapper, key string) *time.Time {
	s := getOrNilStr(m, key)
	if s == nil {
		return nil
	}
	t, err := time.ParseInLocation("20060102T150405", *s, time.Local)
	if err != nil {
		return nil
	}
	return &t
}

func getOrNilCharData(m *map[string]attrCharDataWrapper, key string) *string {
	if s, ok := (*m)[key]; ok {
		return &s.charData
	}
	return nil
}

type attrCharDataWrapper struct {
	text     string
	charData string
}

func (e xmlSingleEvent) toJsonEvent() SingleEvent {
	a := make(map[string]attrCharDataWrapper)
	for _, attr := range e.Description.Attribute {
		a[attr.AttrID] = attrCharDataWrapper{attr.Text, attr.CharacterData}
	}
	return SingleEvent{
		SingleEventID:       getOrNilInt(&a, "singleEventID"),
		SingleEventTypeID:   getOrNilStr(&a, "singleEventTypeID"),
		SingleEventTypeName: getOrNilStr(&a, "singleEventTypeName"),
		GroupID:             getOrNilInt(&a, "groupID"),
		GroupName:           getOrNilStr(&a, "groupName"),
		Dtstamp:             getOrNilDt(&a, "dtstamp"),
		Dtstart:             getOrNilDt(&a, "dtstart"),
		Dtend:               getOrNilDt(&a, "dtend"),
		Duration:            getOrNilStr(&a, "duration"),
		Comment:             getOrNilCharData(&a, "comment"),
		StatusID:            getOrNilStr(&a, "statusID"),
		Status:              getOrNilStr(&a, "status"),
		RecurringID:         getOrNilStr(&a, "recurringID"),
		Address: &Address{
			RoomID:      getOrNilInt(&a, "adr/roomID"),
			RoomCode:    getOrNilStr(&a, "adr/roomCode"),
			RoomName:    getOrNilCharData(&a, "adr/roomAdditionalInfo"),
			ContactName: getOrNilStr(&a, "adr/contactName"),
			Floor:       getOrNilStr(&a, "adr/floor"),
			Street:      getOrNilStr(&a, "adr/street"),
			PCode:       getOrNilStr(&a, "adr/pcode"),
			Locality:    getOrNilStr(&a, "adr/locality"),
			Region:      getOrNilStr(&a, "adr/region"),
			Country:     getOrNilStr(&a, "adr/country"),
		},
	}
}

type courseEventsResp struct {
	XMLName        xml.Name `xml:"RDM"`
	Text           string   `xml:",chardata"`
	Cor            string   `xml:"cor,attr"`
	Xsi            string   `xml:"xsi,attr"`
	SchemaLocation string   `xml:"schemaLocation,attr"`
	Resource       struct {
		Text        string `xml:",chardata"`
		ID          string `xml:"id,attr"`
		TypeID      string `xml:"typeID,attr"`
		Description struct {
			Text      string `xml:",chardata"`
			Attribute []struct {
				Text         string `xml:",chardata"` // 950494023, Grundlagen: Be...
				AttrID       string `xml:"attrID,attr"`
				AttrDataType string `xml:"attrDataType,attr"`
			} `xml:"attribute"`
			ResourceGroup struct {
				Text        string `xml:",chardata"`
				TypeID      string `xml:"typeID,attr"`
				Description struct {
					Text     string           `xml:",chardata"`
					Resource []xmlSingleEvent `xml:"resource"`
				} `xml:"description"`
			} `xml:"resourceGroup"`
		} `xml:"description"`
	} `xml:"resource"`
}

// target Types:

type SingleEvent struct {
	SingleEventID       *int       `json:"singleEventID" example:"888021023"`
	SingleEventTypeID   *string    `json:"singleEventTypeID" example:"A"`
	SingleEventTypeName *string    `json:"singleEventTypeName" example:"Abhaltung"`
	GroupID             *int       `json:"groupID" example:"850243"`
	GroupName           *string    `json:"groupName" example:"Standardgruppe"`
	Dtstamp             *time.Time `json:"dtstamp" example:"2020-03-20T10:15:50Z"`
	Dtstart             *time.Time `json:"dtstart" example:"2020-12-02T13:00:00Z"`
	Dtend               *time.Time `json:"dtend" example:"2020-12-02T14:00:00Z"`
	Duration            *string    `json:"duration" example:"PT01H00M"`
	Comment             *string    `json:"comment" example:"Videoübertragung aus MW 0001"`
	StatusID            *string    `json:"statusID" example:"FG"`
	Status              *string    `json:"status" example:"gelöscht"`
	RecurringID         *string    `json:"recurringID" example:"431418"`
	Address             *Address   `json:"address"`
}

type Address struct {
	RoomID      *int    `json:"roomID" example:"62015"`
	RoomCode    *string `json:"roomCode" example:"5620.01.101"`
	RoomName    *string `json:"roomAdditionalData" example:"101, Hörsaal 1, \"Interims I\""`
	ContactName *string `json:"contactName" example:"Immobilien (ZA 4)"`
	Floor       *string `json:"floor" example:"1.Obergeschoß"`
	Street      *string `json:"street" example:"Boltzmannstr.    5"`
	PCode       *string `json:"pCode" example:"85748"`
	Locality    *string `json:"locality" example:"Garching b. München"`
	Region      *string `json:"region" example:"Bayern"`
	Country     *string `json:"country" example:"Deutschland"`
}
