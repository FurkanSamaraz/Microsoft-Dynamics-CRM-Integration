package crm_repository

import (
	"fmt"
	crm_structures "main/structure/Opportunities"

	"gorm.io/gorm"
)

type OpportunitiesDB struct {
	CrmCollection *gorm.DB
}
type OpportunitiesRepositoryInterface interface {
	GetOpportunities(model crm_structures.GetOpportunities) (bool, error)
	CreateOpportunities(model crm_structures.CreateOpportunities) (bool, error)
	UpdateOpportunities(model crm_structures.UpdateOpportunities) (bool, error)
}

func (t OpportunitiesDB) GetOpportunities(model crm_structures.GetOpportunities) (bool, error) {
	fmt.Println("| REPOSITORY SUCCESS - Opportunities |")
	return true, nil
}
func (t OpportunitiesDB) CreateOpportunities(model crm_structures.CreateOpportunities) (bool, error) {
	fmt.Println("| REPOSITORY SUCCESS - Opportunities |")
	return true, nil
}
func (t OpportunitiesDB) UpdateOpportunities(model crm_structures.UpdateOpportunities) (bool, error) {
	fmt.Println("| REPOSITORY SUCCESS - Opportunities |")
	return true, nil
}
