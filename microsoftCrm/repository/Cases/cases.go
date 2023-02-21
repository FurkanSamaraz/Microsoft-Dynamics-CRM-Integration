package crm_repository

import (
	"fmt"
	crm_structures "main/structure/Cases"

	"gorm.io/gorm"
)

type CasesDB struct {
	CrmCollection *gorm.DB
}
type CasesRepositoryInterface interface {
	GetCases(model crm_structures.GetCases) (bool, error)
	CreateCases(model crm_structures.CreateCases) (bool, error)
}

func (t CasesDB) GetCases(model crm_structures.GetCases) (bool, error) {
	fmt.Println("| REPOSITORY SUCCESS - Cases |")
	return true, nil
}
func (t CasesDB) CreateCases(model crm_structures.CreateCases) (bool, error) {
	fmt.Println("| REPOSITORY SUCCESS - Cases |")
	return true, nil
}
