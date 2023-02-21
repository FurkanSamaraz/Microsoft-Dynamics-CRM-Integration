package crm_services

import (
	"fmt"
	dto "main/dto"
	crm_repository "main/repository/Knowledge_Articles"
	crm_structures "main/structure/Knowledge_Articles"
)

type Service_KnowledgeArticles interface {
	KnowledgeArticlesCreate(model crm_structures.Create_KnowledgeArticles) (*dto.Crm_DTO, error)
	KnowledgeArticlesGet(model crm_structures.Get_KnowledgeArticles) (*dto.Crm_DTO, error)
	KnowledgeArticlesUpdate(model crm_structures.Update_KnowledgeArticles) (*dto.Crm_DTO, error)
}
type KnowledgeArticlesCRMService struct {
	Repo crm_repository.KnowledgerArticlesRepositoryInterface
}

func (t KnowledgeArticlesCRMService) KnowledgeArticlesGet(model crm_structures.Get_KnowledgeArticles) (*dto.Crm_DTO, error) {
	var res dto.Crm_DTO

	fmt.Println("| SERVICE    SUCCESS - KnowledgeArticles |")
	result, err := t.Repo.GetKnowledgerArticles(model)

	if err != nil || !result {
		res.Status = false
		return &res, err
	}

	res = dto.Crm_DTO{Status: result}
	return &res, nil
}
func (t KnowledgeArticlesCRMService) KnowledgeArticlesCreate(model crm_structures.Create_KnowledgeArticles) (*dto.Crm_DTO, error) {
	var res dto.Crm_DTO

	fmt.Println("| SERVICE    SUCCESS - KnowledgeArticles |")
	result, err := t.Repo.CreateKnowledgerArticles(model)

	if err != nil || !result {
		res.Status = false
		return &res, err
	}

	res = dto.Crm_DTO{Status: result}
	return &res, nil
}
func (t KnowledgeArticlesCRMService) KnowledgeArticlesUpdate(model crm_structures.Update_KnowledgeArticles) (*dto.Crm_DTO, error) {
	var res dto.Crm_DTO

	fmt.Println("| SERVICE    SUCCESS - KnowledgeArticles |")
	result, err := t.Repo.UpdateKnowledgerArticles(model)

	if err != nil || !result {
		res.Status = false
		return &res, err
	}

	res = dto.Crm_DTO{Status: result}
	return &res, nil
}
