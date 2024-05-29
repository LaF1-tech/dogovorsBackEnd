package controller

import (
	"context"
	"dogovorsBackEnd/internal/use/models"
	"dogovorsBackEnd/internal/use/modules/applications/dto"
	"dogovorsBackEnd/internal/use/modules/applications/entities"
	contracts "dogovorsBackEnd/internal/use/modules/contracts/entities"
)

func (c *controller) acceptApplication(ctx context.Context, user models.User, request dto.PatchApplicationRequestDTO) error {
	application, err := c.repository.GetApplicationByID(ctx, request.ApplicationID)
	if err != nil {
		return err
	}

	student := entities.Student{
		SpecializationID:           application.SpecializationID,
		EducationalEstablishmentID: application.EducationalEstablishmentID,
		StudentName:                application.Name,
		StudentMiddleName:          application.MiddleName,
		StudentLastName:            application.LastName,
		StudentPhoneNumber:         application.PhoneNumber,
	}

	studentID, err := c.repository.CreateStudent(ctx, student)
	if err != nil {
		return err
	}

	for id, data := range application.Types {
		contract := contracts.Contract{
			StudentID:      studentID,
			EmployeeID:     user.UserID,
			TemplateId:     id,
			ExecutionDate:  application.ExecutionDate,
			ExpirationDate: application.ExpirationDate,
		}

		contractID, err := c.repository.CreateContract(ctx, contract)
		if err != nil {
			return err
		}

		contractDR := entities.ContractDataRegistry{
			TemplateID: id,
			ContractID: contractID,
			Data:       data,
		}

		_, err = c.repository.CreateContractDataRegistryEntry(ctx, contractDR)
		if err != nil {
			return err
		}
	}

	return nil
}
