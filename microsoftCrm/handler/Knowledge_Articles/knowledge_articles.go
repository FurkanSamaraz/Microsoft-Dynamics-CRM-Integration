package crm_handlers

import (
	crm_services "main/service/Knowledge_Articles"

	"github.com/gofiber/fiber/v2"
)

type Knowledge_ArticlesHandler struct {
	Service crm_services.Service_KnowledgeArticles
}

func (h Knowledge_ArticlesHandler) KnowledgeArticles(c *fiber.Ctx) error {
	return nil
}

func (h Knowledge_ArticlesHandler) KnowledgeArticlesCreate(c *fiber.Ctx) error {
	return nil
}
func (h Knowledge_ArticlesHandler) KnowledgeArticlesById(c *fiber.Ctx) error {
	return nil
}
func (h Knowledge_ArticlesHandler) KnowledgeArticlesUpdate(c *fiber.Ctx) error {
	return nil
}
func (h Knowledge_ArticlesHandler) KnowledgeArticlesDelete(c *fiber.Ctx) error {
	return nil
}
