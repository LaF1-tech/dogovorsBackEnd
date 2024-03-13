package controller

import "context"

func (c *controller) generateTemplateHTML(ctx context.Context, id int) (string, error) {
	fullTemplateData, err := c.repository.GetFullTemplateDataByContractID(ctx, id)
	if err != nil {
		return "", err
	}
	return fullTemplateData.TemplateText, err
}
