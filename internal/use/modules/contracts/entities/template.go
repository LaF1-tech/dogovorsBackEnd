package entities

type TemplateVarType string

const (
	TemplateVarTypeString   TemplateVarType = "string"
	TemplateVarTypeInteger  TemplateVarType = "integer"
	TemplateVarTypeDate     TemplateVarType = "date"
	TemplateVarTypeTime     TemplateVarType = "time"
	TemplateVarTypeDateTime TemplateVarType = "datetime"
	TemplateVarTypeBoolean  TemplateVarType = "boolean"
)

type AggregatedTemplate struct {
	ContractID    int
	TemplateName  string
	TemplateText  string
	NecessaryData map[string]TemplateVarType
	Data          map[string]any
}
