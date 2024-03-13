package validator

type TemplateValidator interface {
}

type validator struct {
}

func New() TemplateValidator {
	return &validator{}
}
