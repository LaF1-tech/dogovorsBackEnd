package dto

import "dogovorsBackEnd/internal/use/modules/templates/entities"

type TemplateResponseDTO struct {
	TemplateID      int                                 `json:"template_id,omitempty"`
	TemplateName    string                              `json:"template_name,omitempty"`
	TemplateContent string                              `json:"template_content,omitempty""`
	TemplateStyles  string                              `json:"template_styles,omitempty"`
	NecessaryData   map[string]entities.TemplateVarType `json:"necessary_data,omitempty"`
}

type TemplatePreviewItemResponseDTO struct {
	TemplateID   int    `json:"template_id,omitempty"`
	TemplateName string `json:"template_name,omitempty"`
}

type TemplatePreviewResponseDTO struct {
	List []TemplatePreviewItemResponseDTO `json:"list,omitempty"`
}
