// Package docs GENERATED BY SWAG; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "Joscha Henningsen",
            "url": "https://joschas.page"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/course": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "This endpoint returns all information about a course, e.g. its title, description, semester, contacts and more.",
                "consumes": [
                    "application/json",
                    "text/xml"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "course"
                ],
                "summary": "export course information.",
                "parameters": [
                    {
                        "type": "string",
                        "name": "courseID",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "DE or EN, optional",
                        "name": "language",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Course",
                        "schema": {
                            "$ref": "#/definitions/routes.Course"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/routes.ApiError"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/routes.ApiError"
                        }
                    }
                }
            }
        },
        "/course/events": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "This endpoint returns a list containing all events of a course.",
                "consumes": [
                    "application/json",
                    "text/xml"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "course"
                ],
                "summary": "list events for course.",
                "parameters": [
                    {
                        "type": "string",
                        "name": "courseID",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Events",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/routes.SingleEvent"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/routes.ApiError"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/routes.ApiError"
                        }
                    }
                }
            }
        },
        "/course/students": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "This endpoint returns a list containing all students enrolled in a course.",
                "consumes": [
                    "application/json",
                    "text/xml"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "course"
                ],
                "summary": "list enrolled students for course.",
                "parameters": [
                    {
                        "type": "string",
                        "name": "courseID",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "DE or EN, optional",
                        "name": "language",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Students",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/routes.Person"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/routes.ApiError"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/routes.ApiError"
                        }
                    }
                }
            }
        },
        "/organization": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "This endpoint returns all information about an organization and its sub-organization.",
                "consumes": [
                    "application/json",
                    "text/xml"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "organization"
                ],
                "summary": "export an organization.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "DE or EN, optional",
                        "name": "language",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "name": "orgUnitID",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Organization",
                        "schema": {
                            "$ref": "#/definitions/routes.OrgUnit"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/routes.ApiError"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/routes.ApiError"
                        }
                    }
                }
            }
        },
        "/organization/courses": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "This endpoint returns all courses that belong to the specified organization and all courses of sub-orgs. If a course belongs to a subOrg, orgUnitID is the ID of the subOrg.",
                "consumes": [
                    "application/json",
                    "text/xml"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "organization"
                ],
                "summary": "Export an organizations courses.",
                "parameters": [
                    {
                        "type": "integer",
                        "default": 99,
                        "description": "how deep to include children (only relevant if includeChildren is true).",
                        "name": "depth",
                        "in": "query"
                    },
                    {
                        "type": "boolean",
                        "default": true,
                        "description": "whether to include sub-organizations courses.",
                        "name": "includeChildren",
                        "in": "query"
                    },
                    {
                        "enum": [
                            "DE",
                            "EN"
                        ],
                        "type": "string",
                        "name": "language",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "name": "orgUnitID",
                        "in": "query"
                    },
                    {
                        "enum": [
                            "W",
                            "S"
                        ],
                        "type": "string",
                        "name": "teachingTerm",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Courses",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/routes.Course"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/routes.ApiError"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/routes.ApiError"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "routes.Address": {
            "type": "object",
            "properties": {
                "contactName": {
                    "type": "string",
                    "example": "Immobilien (ZA 4)"
                },
                "country": {
                    "type": "string",
                    "example": "Deutschland"
                },
                "floor": {
                    "type": "string",
                    "example": "1.Obergeschoß"
                },
                "locality": {
                    "type": "string",
                    "example": "Garching b. München"
                },
                "pCode": {
                    "type": "string",
                    "example": "85748"
                },
                "region": {
                    "type": "string",
                    "example": "Bayern"
                },
                "roomAdditionalData": {
                    "type": "string",
                    "example": "101, Hörsaal 1, \"Interims I\""
                },
                "roomCode": {
                    "type": "string",
                    "example": "5620.01.101"
                },
                "roomID": {
                    "type": "integer",
                    "example": 62015
                },
                "street": {
                    "type": "string",
                    "example": "Boltzmannstr.    5"
                }
            }
        },
        "routes.ApiError": {
            "type": "object",
            "properties": {
                "msg": {
                    "type": "string",
                    "example": "CampusOnline returned a status code != 200."
                },
                "status": {
                    "type": "integer",
                    "example": 420
                }
            }
        },
        "routes.ContactData": {
            "type": "object",
            "properties": {
                "adr": {
                    "type": "object",
                    "properties": {
                        "country": {
                            "type": "string"
                        },
                        "extadr": {
                            "type": "string"
                        },
                        "locality": {
                            "type": "string"
                        },
                        "pcode": {
                            "type": "string"
                        },
                        "street": {
                            "type": "string"
                        },
                        "text": {
                            "type": "string"
                        }
                    }
                },
                "contact_name": {
                    "type": "object",
                    "properties": {
                        "chardata": {
                            "type": "string"
                        },
                        "text": {
                            "type": "string"
                        }
                    }
                },
                "email": {
                    "type": "string"
                },
                "telephone": {
                    "type": "object",
                    "properties": {
                        "teltype": {
                            "type": "string"
                        },
                        "text": {
                            "type": "string"
                        }
                    }
                },
                "text": {
                    "type": "string"
                },
                "web_link": {
                    "$ref": "#/definitions/routes.WebLink"
                }
            }
        },
        "routes.Course": {
            "type": "object",
            "properties": {
                "admission_info": {
                    "type": "object",
                    "properties": {
                        "admission_description": {
                            "type": "object",
                            "properties": {
                                "user_defined": {
                                    "type": "string"
                                },
                                "web_link": {
                                    "$ref": "#/definitions/routes.WebLink"
                                }
                            }
                        }
                    }
                },
                "contacts": {
                    "type": "object",
                    "properties": {
                        "person": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/routes.Person"
                            }
                        }
                    }
                },
                "course_code": {
                    "type": "string"
                },
                "course_description": {
                    "type": "string"
                },
                "course_id": {
                    "type": "string"
                },
                "course_name": {
                    "type": "object",
                    "properties": {
                        "text": {
                            "type": "string"
                        }
                    }
                },
                "credits": {
                    "type": "object",
                    "properties": {
                        "hours_per_week": {
                            "type": "string"
                        },
                        "text": {
                            "type": "string"
                        }
                    }
                },
                "exam": {
                    "type": "object",
                    "properties": {
                        "info_block": {
                            "$ref": "#/definitions/routes.InfoBlock"
                        }
                    }
                },
                "form_of_assessment": {
                    "type": "string"
                },
                "form_of_teaching": {
                    "type": "string"
                },
                "instruction_language": {
                    "type": "object",
                    "properties": {
                        "teaching_lang": {
                            "type": "string"
                        }
                    }
                },
                "language": {
                    "type": "string"
                },
                "learning_objectives": {
                    "type": "string"
                },
                "level": {
                    "type": "object",
                    "properties": {
                        "web_link": {
                            "$ref": "#/definitions/routes.WebLink"
                        }
                    }
                },
                "syllabus": {
                    "type": "string"
                },
                "teaching_activity": {
                    "type": "object",
                    "properties": {
                        "info_block": {
                            "$ref": "#/definitions/routes.InfoBlock"
                        },
                        "teaching_activity_id": {
                            "type": "string"
                        },
                        "teaching_activity_name": {
                            "type": "object",
                            "properties": {
                                "text": {
                                    "type": "string"
                                }
                            }
                        }
                    }
                },
                "teaching_term": {
                    "type": "string"
                },
                "type_id": {
                    "type": "string"
                },
                "type_name": {
                    "type": "string"
                }
            }
        },
        "routes.InfoBlock": {
            "type": "object",
            "properties": {
                "picture": {
                    "$ref": "#/definitions/routes.Picture"
                },
                "sub_block": {
                    "type": "object",
                    "properties": {
                        "sub_block": {
                            "type": "string"
                        },
                        "text": {
                            "type": "string"
                        },
                        "user_defined": {
                            "type": "string"
                        }
                    }
                },
                "text": {
                    "type": "string"
                },
                "web_link": {
                    "$ref": "#/definitions/routes.WebLink"
                }
            }
        },
        "routes.OrgUnit": {
            "type": "object",
            "properties": {
                "contacts": {
                    "type": "object",
                    "properties": {
                        "contactData": {
                            "$ref": "#/definitions/routes.ContactData"
                        },
                        "text": {
                            "description": "e.g. \"Lehrstuhl\", \"Arbeitsgruppe\", \"Ausschuss\", ...",
                            "type": "string"
                        }
                    }
                },
                "infoBlock": {
                    "$ref": "#/definitions/routes.InfoBlock"
                },
                "orgUnitCode": {
                    "type": "string"
                },
                "orgUnitID": {
                    "type": "string"
                },
                "orgUnitKind": {
                    "type": "object",
                    "properties": {
                        "subBlock": {
                            "type": "array",
                            "items": {
                                "type": "object",
                                "properties": {
                                    "text": {
                                        "description": "e.g. \"Lehrstuhl\", \"Arbeitsgruppe\", \"Ausschuss\", ...",
                                        "type": "string"
                                    },
                                    "userDefined": {
                                        "description": "always \"name\"",
                                        "type": "string"
                                    }
                                }
                            }
                        },
                        "text": {
                            "type": "string"
                        }
                    }
                },
                "orgUnitName": {
                    "type": "object",
                    "properties": {
                        "chardata": {
                            "type": "string"
                        },
                        "text": {
                            "type": "string"
                        }
                    }
                },
                "orgUnits": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/routes.OrgUnit"
                    }
                },
                "text": {
                    "type": "string"
                }
            }
        },
        "routes.Person": {
            "type": "object",
            "properties": {
                "contact_data": {
                    "$ref": "#/definitions/routes.ContactData"
                },
                "ident": {
                    "type": "string"
                },
                "info_block": {
                    "$ref": "#/definitions/routes.InfoBlock"
                },
                "name": {
                    "type": "object",
                    "properties": {
                        "family": {
                            "type": "string"
                        },
                        "given": {
                            "type": "string"
                        },
                        "text": {
                            "type": "string"
                        }
                    }
                },
                "person_id": {
                    "type": "string"
                },
                "role": {
                    "type": "object",
                    "properties": {
                        "chardata": {
                            "type": "string"
                        },
                        "role_id": {
                            "type": "string"
                        },
                        "text": {
                            "type": "string"
                        }
                    }
                },
                "text": {
                    "type": "string"
                }
            }
        },
        "routes.Picture": {
            "type": "object",
            "properties": {
                "text": {
                    "type": "string"
                },
                "web_link": {
                    "type": "object",
                    "properties": {
                        "href": {
                            "type": "string"
                        },
                        "text": {
                            "type": "string"
                        },
                        "user_defined": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "routes.SingleEvent": {
            "type": "object",
            "properties": {
                "address": {
                    "$ref": "#/definitions/routes.Address"
                },
                "comment": {
                    "type": "string",
                    "example": "Videoübertragung aus MW 0001"
                },
                "dtend": {
                    "type": "string",
                    "example": "2020-12-02T14:00:00Z"
                },
                "dtstamp": {
                    "type": "string",
                    "example": "2020-03-20T10:15:50Z"
                },
                "dtstart": {
                    "type": "string",
                    "example": "2020-12-02T13:00:00Z"
                },
                "duration": {
                    "type": "string",
                    "example": "PT01H00M"
                },
                "groupID": {
                    "type": "integer",
                    "example": 850243
                },
                "groupName": {
                    "type": "string",
                    "example": "Standardgruppe"
                },
                "recurringID": {
                    "type": "string",
                    "example": "431418"
                },
                "singleEventID": {
                    "type": "integer",
                    "example": 888021023
                },
                "singleEventTypeID": {
                    "type": "string",
                    "example": "A"
                },
                "singleEventTypeName": {
                    "type": "string",
                    "example": "Abhaltung"
                },
                "status": {
                    "type": "string",
                    "example": "gelöscht"
                },
                "statusID": {
                    "type": "string",
                    "example": "FG"
                }
            }
        },
        "routes.WebLink": {
            "type": "object",
            "properties": {
                "href": {
                    "type": "string"
                },
                "user_defined": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "ApiKeyAuth": {
            "type": "apiKey",
            "name": "x-api-key",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "/api/v1",
	Schemes:          []string{"http", "https"},
	Title:            "CAMPUSOnline Webservice proxy",
	Description:      "This is the proxy server for CAMPUSOnline Webservices, enabling a nicely documented and uniform rest access to CAMPUSOnline resources.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
