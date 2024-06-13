package entities

type PeriodChart struct {
	Period        string
	ContractCount int
}

type EducationalEstablishmentCountChart struct {
	EducationalEstablishmentName string
	ContractCount                int
}

type SpecializationNameCountChart struct {
	SpecializationName string
	ContractCount      int
}

type TemplateNameCountChart struct {
	TemplateName  string
	ContractCount int
}
