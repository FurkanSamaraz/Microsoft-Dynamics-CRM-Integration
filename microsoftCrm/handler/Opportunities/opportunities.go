package crm_handlers

import (
	crm_services "main/service/Opportunities"

	"github.com/gofiber/fiber/v2"
)

type OpportunitiesHandler struct {
	Service crm_services.Service_Opportunities
}

func (h OpportunitiesHandler) Opportunities(c *fiber.Ctx) error {
	return nil
}

func (h OpportunitiesHandler) OpportunitiesCreate(c *fiber.Ctx) error {
	return nil
}
func (h OpportunitiesHandler) OpportunitiesById(c *fiber.Ctx) error {
	return nil
}
func (h OpportunitiesHandler) OpportunitiesUpdate(c *fiber.Ctx) error {
	return nil
}
func (h OpportunitiesHandler) OpportunitiesDelete(c *fiber.Ctx) error {
	return nil
}
