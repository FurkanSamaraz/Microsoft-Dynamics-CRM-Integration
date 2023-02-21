package crm_services

import (
	"fmt"
	dto "main/dto"
	crm_repository "main/repository/Opportunities"
	crm_structures "main/structure/Opportunities"
)

type Service_Opportunities interface {
	OpportunitiesCreate(model crm_structures.CreateOpportunities) (*dto.Crm_DTO, error)
	OpportunitiesGet(model crm_structures.GetOpportunities) (*dto.Crm_DTO, error)
	OpportunitiesUpdate(model crm_structures.UpdateOpportunities) (*dto.Crm_DTO, error)
}
type OpportunitiesCRMService struct {
	Repo crm_repository.OpportunitiesRepositoryInterface
}

func (t OpportunitiesCRMService) OpportunitiesGet(model crm_structures.GetOpportunities) (*dto.Crm_DTO, error) {
	var res dto.Crm_DTO

	fmt.Println("| SERVICE    SUCCESS - Opportunities |")
	result, err := t.Repo.GetOpportunities(model)

	if err != nil || !result {
		res.Status = false
		return &res, err
	}

	res = dto.Crm_DTO{Status: result}
	return &res, nil
}
func (t OpportunitiesCRMService) OpportunitiesCreate(model crm_structures.CreateOpportunities) (*dto.Crm_DTO, error) {
	var res dto.Crm_DTO

	fmt.Println("| SERVICE    SUCCESS - Opportunities |")
	result, err := t.Repo.CreateOpportunities(model)

	if err != nil || !result {
		res.Status = false
		return &res, err
	}

	res = dto.Crm_DTO{Status: result}
	return &res, nil
}
func (t OpportunitiesCRMService) OpportunitiesUpdate(model crm_structures.UpdateOpportunities) (*dto.Crm_DTO, error) {
	var res dto.Crm_DTO

	fmt.Println("| SERVICE    SUCCESS - Opportunities |")
	result, err := t.Repo.UpdateOpportunities(model)

	if err != nil || !result {
		res.Status = false
		return &res, err
	}

	res = dto.Crm_DTO{Status: result}
	return &res, nil
}
