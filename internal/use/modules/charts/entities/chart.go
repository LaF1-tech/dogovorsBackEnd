package entities

import "time"

type PeriodChart struct {
	Period        time.Time
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
