package entities

import "dogovorsBackEnd/internal/use/entities"

type TemplateVarType string

const (
	TemplateVarTypeString    TemplateVarType = "string"
	TemplateVarTypeInteger   TemplateVarType = "integer"
	TemplateVarTypeDate      TemplateVarType = "date"
	TemplateVarTypeTime      TemplateVarType = "time"
	TemplateVarTypeDateTime  TemplateVarType = "datetime"
	TemplateVarTypeDateRange TemplateVarType = "daterange"
	TemplateVarTypeBoolean   TemplateVarType = "boolean"
)

type Template struct {
	entities.Timed

	TemplateID      int
	TemplateName    string
	TemplateContent string
	TemplateStyles  string
	NecessaryData   map[string]TemplateVarType
}

type TemplatePreview struct {
	TemplateID   int
	TemplateName string
}

type FullTemplate struct {
	ContractID      int
	TemplateName    string
	TemplateContent string
	TemplateStyles  string
	NecessaryData   map[string]TemplateVarType
	Data            map[string]any
}
