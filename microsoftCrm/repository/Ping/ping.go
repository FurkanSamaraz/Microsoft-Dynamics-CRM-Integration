package crm_repository

import (
	"fmt"
	crm_structures "main/structure/Ping"

	"gorm.io/gorm"
)

type PingDB struct {
	CrmCollection *gorm.DB
}
type PingRepositoryInterface interface {
	GetPing(model crm_structures.GetPing) (bool, error)
}

func (t PingDB) GetPing(model crm_structures.GetPing) (bool, error) {
	fmt.Println("| REPOSITORY SUCCESS - Ping |")
	return true, nil
}
