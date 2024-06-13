package controller

import (
	"context"
	"fmt"
	"github.com/SebastiaanKlippert/go-wkhtmltopdf"
	"strings"
	"time"

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
				TemplateName:      item.TemplateName,
				StudentLastName:   item.StudentLastName,
				StudentName:       item.StudentName,
				StudentMiddleName: item.StudentMiddleName,
				ContractStatus:    item.ContractStatus,
				EmployeeFirstName: item.EmployeeFirstName,
				EmployeeLastName:  item.EmployeeLastName,
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
		ContractID:     contract.ContractID,
		StudentID:      contract.StudentID,
		EmployeeID:     contract.EmployeeID,
		TemplateID:     contract.TemplateId,
		ExecutionDate:  contract.ExecutionDate,
		ExpirationDate: contract.ExpirationDate,
	}, nil
}

func (c *controller) PatchContractByID(ctx context.Context, user models.User, request dto.PatchContractRequestDTO) error {
	if err := user.AssertPermission(models.PermissionAdmin); err != nil {
		return err
	}
	contract := entities.ContractExecutionControl{
		ContractID:     request.ContractID,
		ContractStatus: request.ContractStatus,
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

	filename := fmt.Sprintf("./contracts/[%d] %s - %s [%s].pdf", template.ContractID, template.TemplateName, template.Data["student_name"], time.Now().Format(time.DateTime))
	filename = strings.ReplaceAll(filename, ":", "-")

	pdfg, err := wkhtmltopdf.NewPDFGenerator()
	if err != nil {
		return "", err
	}

	page := wkhtmltopdf.NewPageReader(strings.NewReader(html))
	page.EnableLocalFileAccess.Set(true)

	pdfg.AddPage(page)
	pdfg.Dpi.Set(300)
	pdfg.Orientation.Set(wkhtmltopdf.OrientationPortrait)
	pdfg.PageSize.Set(wkhtmltopdf.PageSizeA4)

	if err := pdfg.Create(); err != nil {
		return "", err
	}

	if err := pdfg.WriteFile(filename); err != nil {
		return "", err
	}

	return filename, nil
}
