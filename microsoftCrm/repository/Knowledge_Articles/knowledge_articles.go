package crm_repository

import (
	"fmt"
	crm_structures "main/structure/Knowledge_Articles"

	"gorm.io/gorm"
)

type KnowledgerArticlesDB struct {
	CrmCollection *gorm.DB
}
type KnowledgerArticlesRepositoryInterface interface {
	GetKnowledgerArticles(model crm_structures.Get_KnowledgeArticles) (bool, error)
	CreateKnowledgerArticles(model crm_structures.Create_KnowledgeArticles) (bool, error)
	UpdateKnowledgerArticles(model crm_structures.Update_KnowledgeArticles) (bool, error)
}

func (t KnowledgerArticlesDB) GetKnowledgerArticles(model crm_structures.Get_KnowledgeArticles) (bool, error) {
	fmt.Println("| REPOSITORY SUCCESS - KnowledgerArticles |")
	return true, nil
}
func (t KnowledgerArticlesDB) CreateKnowledgerArticles(model crm_structures.Create_KnowledgeArticles) (bool, error) {
	fmt.Println("| REPOSITORY SUCCESS - KnowledgerArticles |")
	return true, nil
}
func (t KnowledgerArticlesDB) UpdateKnowledgerArticles(model crm_structures.Update_KnowledgeArticles) (bool, error) {
	fmt.Println("| REPOSITORY SUCCESS - KnowledgerArticles |")
	return true, nil
}
