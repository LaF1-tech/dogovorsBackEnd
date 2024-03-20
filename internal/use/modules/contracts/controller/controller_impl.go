package controller

import (
	"context"
	"github.com/speedata/bagme/document"
	"github.com/speedata/boxesandglue/backend/bag"

	"dogovorsBackEnd/internal/use/models"
	"dogovorsBackEnd/internal/use/modules/contracts/dto"
	"dogovorsBackEnd/internal/use/modules/contracts/entities"
	tEntities "dogovorsBackEnd/internal/use/modules/templates/entities"
	"dogovorsBackEnd/internal/use/utils/uslices"
)

func (c *controller) GetAllContracts(ctx context.Context, user models.User) (dto.ContractsViewResponseDTO, error) {
	if err := user.AssertPermission(models.PermissionAdmin); err != nil {
		return dto.ContractsViewResponseDTO{}, err
	}

	applications, err := c.repository.GetAllContracts(ctx)
	if err != nil {
		return dto.ContractsViewResponseDTO{}, err
	}

	return dto.ContractsViewResponseDTO{
		List: uslices.MapFunc(applications, func(item entities.AggregatedContract) dto.ContractAggregatedResponseDTO {
			return dto.ContractAggregatedResponseDTO{
				ContractID:        item.ContractID,
				StudentName:       item.StudentName,
				StudentLastName:   item.StudentLastName,
				EmployeeFirstName: item.EmployeeFirstName,
				EmployeeLastName:  item.EmployeeLastName,
				ApplicationType:   item.ApplicationType,
				ExecutionDate:     item.ExecutionDate,
				ExpirationDate:    item.ExpirationDate,
			}
		}),
	}, nil
}

func (c *controller) GetContractByID(ctx context.Context, user models.User, id int) (dto.ContractItemResponseDTO, error) {
	if err := user.AssertPermission(models.PermissionAdmin); err != nil {
		return dto.ContractItemResponseDTO{}, err
	}
	contract, err := c.repository.GetContractByID(ctx, id)
	if err != nil {
		return dto.ContractItemResponseDTO{}, err
	}
	return dto.ContractItemResponseDTO{
		ContractID:      contract.ContractID,
		StudentID:       contract.StudentID,
		EmployeeID:      contract.EmployeeID,
		ApplicationType: contract.ApplicationType,
		ExecutionDate:   contract.ExecutionDate,
		ExpirationDate:  contract.ExpirationDate,
	}, nil
}

func (c *controller) PatchContractByID(ctx context.Context, user models.User, request dto.PatchContractRequestDTO) error {
	if err := user.AssertPermission(models.PermissionAdmin); err != nil {
		return err
	}
	contract := entities.Contract{
		ContractID:      request.ContractID,
		StudentID:       request.StudentID,
		EmployeeID:      request.EmployeeID,
		ApplicationType: request.ApplicationType,
		ExecutionDate:   request.ExecutionDate,
		ExpirationDate:  request.ExpirationDate,
	}
	err := c.repository.PatchContractByID(ctx, contract)
	if err != nil {
		return err
	}
	return nil
}

func (c *controller) generateTemplateHTML(ctx context.Context, id int) (tEntities.FullTemplate, string, error) {
	template, err := c.repository.GetAggregatedTemplateByContractID(ctx, id)
	if err != nil {
		return tEntities.FullTemplate{}, "", nil
	}

	content, err := c.repository.GenerateTemplateContent(ctx, template)
	return template, content, err
}

func (c *controller) GenerateContractPDF(ctx context.Context, user models.User, id int) (string, error) {
	if err := user.AssertPermission(models.PermissionAdmin); err != nil {
		return "", err
	}

	template, html, err := c.generateTemplateHTML(ctx, id)
	if err != nil {
		return "", err
	}
	d, err := document.New("out.pdf")
	if err != nil {
		return "", err
	}
	if d.AddCSS(template.TemplateStyles); err != nil {
		return "", err
	}
	if err = d.Frontend.LoadIncludedFonts(); err != nil {
		return "", err
	}
	wd := bag.MustSp("280pt")
	colText := bag.MustSp("140pt")
	rowText := bag.MustSp("23cm")
	if err = d.OutputAt(html, wd, colText, rowText); err != nil {
		return "", err
	}

	return html, d.Finish()
}