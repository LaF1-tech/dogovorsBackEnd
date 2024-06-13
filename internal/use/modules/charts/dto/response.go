package dto

type PeriodChartResponseDTO struct {
	Period        string `json:"period,omitempty"`
	ContractCount int    `json:"contract_count,omitempty"`
}

type EducationalEstablishmentChartResponseDTO struct {
	EducationalEstablishmentName string `json:"educational_establishment_name,omitempty"`
	ContractCount                int    `json:"contract_count,omitempty"`
}

type SpecializationChartResponseDTO struct {
	SpecializationName string `json:"specialization_name,omitempty"`
	ContractCount      int    `json:"contract_count,omitempty"`
}

type TemplateChartResponseDTO struct {
	TemplateName  string `json:"template_name,omitempty"`
	ContractCount int    `json:"contract_count,omitempty"`
}

type PeriodsChartResponseDTO struct {
	List []PeriodChartResponseDTO `json:"list,omitempty"`
}

type EducationalEstablishmentsChartResponseDTO struct {
	List []EducationalEstablishmentChartResponseDTO `json:"list,omitempty"`
}

type SpecializationsChartResponseDTO struct {
	List []SpecializationChartResponseDTO `json:"list,omitempty"`
}

type TemplatesChartResponseDTO struct {
	List []TemplateChartResponseDTO `json:"list,omitempty"`
}
