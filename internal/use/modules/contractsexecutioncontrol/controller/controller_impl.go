package controller

import (
	"context"
	"dogovorsBackEnd/internal/use/models"
	"dogovorsBackEnd/internal/use/modules/contractsexecutioncontrol/dto"
	"dogovorsBackEnd/internal/use/modules/contractsexecutioncontrol/entities"
	"dogovorsBackEnd/internal/use/utils/uslices"
)

func (c *controller) GetContractsExecutionControlByContractID(ctx context.Context, user models.User, id int) (dto.AggregatedContractsExecutionControlResponseDTO, error) {
	if err := user.AssertPermission(models.PermissionAdmin); err != nil {
		return dto.AggregatedContractsExecutionControlResponseDTO{}, err
	}

	cecs, err := c.repository.GetContractsExecutionControlByContractID(ctx, id)
	if err != nil {
		return dto.AggregatedContractsExecutionControlResponseDTO{}, err
	}

	return dto.AggregatedContractsExecutionControlResponseDTO{
		List: uslices.MapFunc(cecs, func(item entities.AggregatedContractExecutionControl) dto.AggregatedContractExecutionControlResponseDTO {
			return dto.AggregatedContractExecutionControlResponseDTO{
				Id:              item.Id,
				ContractId:      item.ContractId,
				TemplateName:    item.TemplateName,
				StudentLastName: item.StudentLastName,
				StudentName:     item.StudentName,
				ContractStatus:  item.ContractStatus,
				ControlDate:     item.ControlDate,
			}
		}),
	}, nil
}

func (c *controller) CreateContractExecutionControl(ctx context.Context, user models.User, request dto.CreateContractExecutionControlRequestDTO) (dto.ContractExecutionControlResponseDTO, error) {
	if err := user.AssertPermission(models.PermissionAdmin); err != nil {
		return dto.ContractExecutionControlResponseDTO{}, err
	}

	if err := c.validator.CreateContractExecutionControlRequest(ctx, request); err != nil {
		return dto.ContractExecutionControlResponseDTO{}, err
	}

	cec := entities.ContractExecutionControl{
		ContractId:     request.ContractId,
		ControlDate:    request.ControlDate,
		ContractStatus: request.ControlStatus,
	}

	id, err := c.repository.CreateContractExecutionControl(ctx, cec)
	if err != nil {
		return dto.ContractExecutionControlResponseDTO{}, err
	}

	cec.Id = id

	return dto.ContractExecutionControlResponseDTO{
		Id:             cec.Id,
		ContractId:     cec.ContractId,
		ContractStatus: cec.ContractStatus,
		ControlDate:    cec.ControlDate,
	}, nil
}

func (c *controller) PatchContractExecutionControl(ctx context.Context, user models.User, request dto.PatchContractExecutionControlRequestDTO) error {
	if err := user.AssertPermission(models.PermissionAdmin); err != nil {
		return err
	}

	if err := c.validator.PatchContractExecutionControlRequest(ctx, request); err != nil {
		return err
	}

	cec := entities.ContractExecutionControl{
		Id:             request.Id,
		ControlDate:    request.ControlDate,
		ContractStatus: request.ControlStatus,
	}

	err := c.repository.PatchContractExecutionControl(ctx, cec)
	if err != nil {
		return err
	}
	return nil
}
