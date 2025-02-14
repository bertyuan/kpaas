// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT

package docs

import (
	"bytes"
	"encoding/json"
	"strings"

	"github.com/alecthomas/template"
	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{.Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "Support",
            "url": "http://github.com/kpaas-io/kpaas/issues",
            "email": "support@kpaas.io"
        },
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/v1/deploy/wizard/checks": {
            "get": {
                "description": "Get the result of the check node",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "checking"
                ],
                "summary": "Get the result of check node",
                "operationId": "GetCheckNodeListResult",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/api.GetCheckingResultResponse"
                        }
                    }
                }
            },
            "post": {
                "description": "Check if the node meets the pre-deployment requirements",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "checking"
                ],
                "summary": "check node list",
                "operationId": "CheckNodeList",
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/api.SuccessfulOption"
                        }
                    }
                }
            }
        },
        "/api/v1/deploy/wizard/clusters": {
            "post": {
                "description": "Store new cluster information",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "cluster"
                ],
                "summary": "Set Cluster Information",
                "operationId": "SetCluster",
                "parameters": [
                    {
                        "description": "RequiredFields: shortName, name, kubeAPIServerConnectType",
                        "name": "cluster",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/api.Cluster"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/api.SuccessfulOption"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/h.AppErr"
                        }
                    },
                    "409": {
                        "description": "Conflict",
                        "schema": {
                            "$ref": "#/definitions/h.AppErr"
                        }
                    }
                }
            }
        },
        "/api/v1/deploy/wizard/deploys": {
            "get": {
                "description": "Get the result of the deployment",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "deploy"
                ],
                "summary": "Get the result of deployment",
                "operationId": "GetDeploymentReport",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/api.GetDeploymentReportResponse"
                        }
                    }
                }
            },
            "post": {
                "description": "Launch deployment",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "deploy"
                ],
                "summary": "Launch deployment",
                "operationId": "LaunchDeployment",
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/api.SuccessfulOption"
                        }
                    }
                }
            }
        },
        "/api/v1/deploy/wizard/logs/{id}": {
            "get": {
                "description": "Download the deployment log details to check the cause of the error",
                "produces": [
                    "text/plain"
                ],
                "tags": [
                    "log"
                ],
                "summary": "Download the log detail",
                "operationId": "DownloadLog",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Log ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Log File Content",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/h.AppErr"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/h.AppErr"
                        }
                    }
                }
            }
        },
        "/api/v1/deploy/wizard/nodes": {
            "post": {
                "description": "Add deployment candidate to node list",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "node"
                ],
                "summary": "Add Node Information",
                "operationId": "AddNode",
                "parameters": [
                    {
                        "description": "Node information",
                        "name": "node",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/api.NodeData"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/api.NodeData"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/h.AppErr"
                        }
                    },
                    "409": {
                        "description": "Conflict",
                        "schema": {
                            "$ref": "#/definitions/h.AppErr"
                        }
                    }
                }
            }
        },
        "/api/v1/deploy/wizard/nodes/{ip}": {
            "put": {
                "description": "Update a node information which in deployment candidate node list",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "node"
                ],
                "summary": "Update Node Information",
                "operationId": "UpdateNode",
                "parameters": [
                    {
                        "description": "Node information",
                        "name": "node",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/api.NodeData"
                        }
                    },
                    {
                        "type": "integer",
                        "description": "Node IP Address",
                        "name": "ip",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/api.NodeData"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/h.AppErr"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/h.AppErr"
                        }
                    },
                    "409": {
                        "description": "Conflict",
                        "schema": {
                            "$ref": "#/definitions/h.AppErr"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete a node from deployment candidate node list",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "node"
                ],
                "summary": "Delete a node",
                "operationId": "DeleteNode",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Node IP Address",
                        "name": "ip",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {},
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/h.AppErr"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/h.AppErr"
                        }
                    },
                    "409": {
                        "description": "Conflict",
                        "schema": {
                            "$ref": "#/definitions/h.AppErr"
                        }
                    }
                }
            }
        },
        "/api/v1/deploy/wizard/progresses": {
            "get": {
                "description": "Get all data, include current progress, cluster and node data. deploying progress or error.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "wizard"
                ],
                "summary": "Get all of current deploy wizard data",
                "operationId": "GetWizardProgress",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/api.GetWizardResponse"
                        }
                    }
                }
            },
            "delete": {
                "description": "Clear all data, include current progress, cluster and node data. deploying progress or error.",
                "tags": [
                    "wizard"
                ],
                "summary": "Clear all of current deploy wizard data",
                "operationId": "ClearWizard",
                "responses": {
                    "204": {}
                }
            }
        },
        "/api/v1/ssh/tests": {
            "post": {
                "description": "Try to connection a node using ssh",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "ssh"
                ],
                "summary": "Test a node connection",
                "operationId": "TestSSH",
                "parameters": [
                    {
                        "description": "Node information",
                        "name": "node",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/api.ConnectionData"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/api.SuccessfulOption"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/h.AppErr"
                        }
                    },
                    "409": {
                        "description": "Conflict",
                        "schema": {
                            "$ref": "#/definitions/h.AppErr"
                        }
                    }
                }
            }
        },
        "/api/v1/ssh_certificates": {
            "get": {
                "description": "Get SSH login certificate keys list",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "ssh_certificate"
                ],
                "summary": "Get SSH login keys list",
                "operationId": "GetSSHCertificate",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/api.GetSSHCertificateListResponse"
                        }
                    }
                }
            },
            "post": {
                "description": "Add SSH login private key",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "ssh_certificate"
                ],
                "summary": "Add SSH login private key",
                "operationId": "AddSSHCertificate",
                "parameters": [
                    {
                        "description": "Certificate information",
                        "name": "certificate",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/api.SSHCertificate"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/api.SuccessfulOption"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "api.Annotation": {
            "type": "object",
            "required": [
                "key",
                "value"
            ],
            "properties": {
                "key": {
                    "type": "string"
                },
                "value": {
                    "type": "string"
                }
            }
        },
        "api.CheckingItem": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "object",
                    "$ref": "#/definitions/api.Error"
                },
                "point": {
                    "type": "string"
                },
                "result": {
                    "description": "Checking Result",
                    "type": "string",
                    "enum": [
                        "checking",
                        "passed",
                        "failed"
                    ]
                }
            }
        },
        "api.CheckingResultResponseData": {
            "type": "object",
            "properties": {
                "items": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/api.CheckingItem"
                    }
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "api.Cluster": {
            "type": "object",
            "required": [
                "kubeAPIServerConnectType",
                "name",
                "shortName"
            ],
            "properties": {
                "annotations": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/api.Annotation"
                    }
                },
                "kubeAPIServerConnectType": {
                    "description": "kube-apiserver connect type",
                    "type": "string",
                    "enum": [
                        "firstMasterIP",
                        "keepalived",
                        "loadbalancer"
                    ]
                },
                "labels": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/api.Label"
                    }
                },
                "loadbalancerIP": {
                    "description": "kube-apiserver loadbalancer ip when kubeAPIServerConnectType is loadbalancer required",
                    "type": "string",
                    "maxLength": 15
                },
                "loadbalancerPort": {
                    "description": "kube-apiserver loadbalancer port when kubeAPIServerConnectType is loadbalancer required",
                    "type": "integer",
                    "maximum": 65535,
                    "minimum": 1
                },
                "name": {
                    "type": "string"
                },
                "netInterfaceName": {
                    "description": "keepalived listen net interface name",
                    "type": "string",
                    "maxLength": 30
                },
                "nodePortMaximum": {
                    "type": "integer",
                    "default": 32767,
                    "maximum": 65535
                },
                "nodePortMinimum": {
                    "type": "integer",
                    "default": 30000,
                    "minimum": 1
                },
                "shortName": {
                    "type": "string",
                    "maxLength": 20,
                    "minLength": 1
                },
                "vip": {
                    "description": "keepalived listen virtual ip",
                    "type": "string",
                    "maxLength": 15
                }
            }
        },
        "api.ConnectionData": {
            "type": "object",
            "required": [
                "ip",
                "port",
                "username"
            ],
            "properties": {
                "authorizationType": {
                    "description": "type of authorization",
                    "type": "string",
                    "enum": [
                        "password",
                        "privateKey"
                    ]
                },
                "ip": {
                    "description": "node ip",
                    "type": "string",
                    "maxLength": 15,
                    "minLength": 1
                },
                "password": {
                    "description": "login password",
                    "type": "string"
                },
                "port": {
                    "description": "ssh port",
                    "type": "integer",
                    "default": 22,
                    "maximum": 65535,
                    "minimum": 1
                },
                "privateKeyName": {
                    "description": "the private key name of login",
                    "type": "string"
                },
                "username": {
                    "description": "ssh username",
                    "type": "string",
                    "maxLength": 128
                }
            }
        },
        "api.DeploymentNode": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "object",
                    "$ref": "#/definitions/api.Error"
                },
                "result": {
                    "description": "Checking Result",
                    "type": "string",
                    "enum": [
                        "pending",
                        "deploying",
                        "completed",
                        "failed",
                        "aborted"
                    ]
                }
            }
        },
        "api.DeploymentResponseData": {
            "type": "object",
            "properties": {
                "nodes": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/api.DeploymentNode"
                    }
                },
                "role": {
                    "type": "string"
                }
            }
        },
        "api.Error": {
            "type": "object",
            "properties": {
                "detail": {
                    "description": "Why is it wrong, what is the judgment condition?",
                    "type": "string"
                },
                "fixMethods": {
                    "description": "How to improve to meet the conditions",
                    "type": "string"
                },
                "logId": {
                    "description": "ID used to get the log file",
                    "type": "integer"
                },
                "reason": {
                    "description": "Reason of Error message",
                    "type": "string"
                }
            }
        },
        "api.GetCheckingResultResponse": {
            "type": "object",
            "properties": {
                "nodes": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/api.CheckingResultResponseData"
                    }
                },
                "result": {
                    "description": "Overall inspection status",
                    "type": "string",
                    "enum": [
                        "notRunning",
                        "checking",
                        "passed",
                        "failed"
                    ]
                }
            }
        },
        "api.GetDeploymentReportResponse": {
            "type": "object",
            "properties": {
                "deployClusterStatus": {
                    "description": "The cluster deployment status",
                    "type": "string",
                    "enum": [
                        "notRunning",
                        "running",
                        "successful",
                        "failed",
                        "workedButHaveError"
                    ]
                },
                "nodes": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/api.DeploymentResponseData"
                    }
                }
            }
        },
        "api.GetSSHCertificateListResponse": {
            "type": "object",
            "properties": {
                "names": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                }
            }
        },
        "api.GetWizardResponse": {
            "type": "object",
            "properties": {
                "checkResult": {
                    "description": "Nodes check result",
                    "type": "string",
                    "enum": [
                        "notRunning",
                        "checking",
                        "passed",
                        "failed"
                    ]
                },
                "checkingData": {
                    "description": "Check result",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/api.CheckingResultResponseData"
                    }
                },
                "cluster": {
                    "description": "Cluster Information",
                    "type": "object",
                    "$ref": "#/definitions/api.Cluster"
                },
                "deployClusterStatus": {
                    "description": "Cluster deployment status",
                    "type": "string",
                    "enum": [
                        "notRunning",
                        "running",
                        "successful",
                        "failed",
                        "workedButHaveError"
                    ]
                },
                "deploymentData": {
                    "description": "Deployment result",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/api.DeploymentResponseData"
                    }
                },
                "mode": {
                    "description": "Wizard mode, normal or advanced",
                    "type": "string"
                },
                "nodes": {
                    "description": "Nodes Information",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/api.NodeData"
                    }
                },
                "progress": {
                    "description": "Wizard progress",
                    "type": "string"
                }
            }
        },
        "api.Label": {
            "type": "object",
            "required": [
                "key",
                "value"
            ],
            "properties": {
                "key": {
                    "type": "string"
                },
                "value": {
                    "type": "string"
                }
            }
        },
        "api.NodeData": {
            "type": "object",
            "required": [
                "ip",
                "name",
                "port",
                "username"
            ],
            "properties": {
                "authorizationType": {
                    "description": "type of authorization",
                    "type": "string",
                    "enum": [
                        "password",
                        "privateKey"
                    ]
                },
                "description": {
                    "description": "node description",
                    "type": "string"
                },
                "ip": {
                    "description": "node ip",
                    "type": "string",
                    "maxLength": 15,
                    "minLength": 1
                },
                "labels": {
                    "description": "Node labels",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/api.Label"
                    }
                },
                "name": {
                    "description": "node name",
                    "type": "string",
                    "maxLength": 64,
                    "minLength": 1
                },
                "password": {
                    "description": "login password",
                    "type": "string"
                },
                "port": {
                    "description": "ssh port",
                    "type": "integer",
                    "default": 22,
                    "maximum": 65535,
                    "minimum": 1
                },
                "privateKeyName": {
                    "description": "the private key name of login",
                    "type": "string"
                },
                "taint": {
                    "description": "Node taints",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/api.Taint"
                    }
                },
                "username": {
                    "description": "ssh username",
                    "type": "string",
                    "maxLength": 128
                }
            }
        },
        "api.SSHCertificate": {
            "type": "object",
            "required": [
                "content",
                "name"
            ],
            "properties": {
                "content": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "api.SuccessfulOption": {
            "type": "object",
            "properties": {
                "success": {
                    "type": "boolean"
                }
            }
        },
        "api.Taint": {
            "type": "object",
            "required": [
                "key",
                "value"
            ],
            "properties": {
                "effect": {
                    "type": "string",
                    "enum": [
                        "NoSchedule",
                        "NoExecute",
                        "PreferNoSchedule"
                    ]
                },
                "key": {
                    "type": "string"
                },
                "value": {
                    "type": "string"
                }
            }
        },
        "h.AppErr": {
            "type": "object",
            "properties": {
                "msg": {
                    "type": "string"
                },
                "payload": {
                    "type": "object"
                }
            }
        }
    }
}`

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Schemes     []string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = swaggerInfo{
	Version:     "0.1",
	Host:        "localhost:8080",
	BasePath:    "/api",
	Schemes:     []string{},
	Title:       "kpaasRestfulApi",
	Description: "KPaaS RESTful API service for frontend and using Deploy service API to deployment kubernetes cluster.",
}

type s struct{}

func (s *s) ReadDoc() string {
	sInfo := SwaggerInfo
	sInfo.Description = strings.Replace(sInfo.Description, "\n", "\\n", -1)

	t, err := template.New("swagger_info").Funcs(template.FuncMap{
		"marshal": func(v interface{}) string {
			a, _ := json.Marshal(v)
			return string(a)
		},
	}).Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, sInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register(swag.Name, &s{})
}
