package crm_services

import (
	"fmt"
	dto "main/dto"
	crm_repository "main/repository/Cases"
	crm_structures "main/structure/Cases"
)

type Service_Cases interface {
	CasesCreate(model crm_structures.CreateCases) (*dto.Crm_DTO, error)
	CasesGet(model crm_structures.GetCases) (*dto.Crm_DTO, error)
}
type CasesCRMService struct {
	Repo crm_repository.CasesRepositoryInterface
}

func (t CasesCRMService) CasesGet(model crm_structures.GetCases) (*dto.Crm_DTO, error) {
	var res dto.Crm_DTO

	fmt.Println("| SERVICE    SUCCESS - Cases |")
	result, err := t.Repo.GetCases(model)

	if err != nil || !result {
		res.Status = false
		return &res, err
	}

	res = dto.Crm_DTO{Status: result}
	return &res, nil
}
func (t CasesCRMService) CasesCreate(model crm_structures.CreateCases) (*dto.Crm_DTO, error) {
	var res dto.Crm_DTO

	fmt.Println("| SERVICE    SUCCESS - Cases |")
	result, err := t.Repo.CreateCases(model)

	if err != nil || !result {
		res.Status = false
		return &res, err
	}

	res = dto.Crm_DTO{Status: result}
	return &res, nil
}
