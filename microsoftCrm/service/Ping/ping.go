package crm_services

import (
	"fmt"
	dto "main/dto"
	crm_repository "main/repository/Ping"
	crm_structures "main/structure/Ping"
)

type Service_Ping interface {
	PingGet(model crm_structures.GetPing) (*dto.Crm_DTO, error)
}
type PingCRMService struct {
	Repo crm_repository.PingRepositoryInterface
}

func (t PingCRMService) PingGet(model crm_structures.GetPing) (*dto.Crm_DTO, error) {
	var res dto.Crm_DTO

	fmt.Println("| SERVICE    SUCCESS - Ping |")
	result, err := t.Repo.GetPing(model)

	if err != nil || !result {
		res.Status = false
		return &res, err
	}

	res = dto.Crm_DTO{Status: result}
	return &res, nil
}
