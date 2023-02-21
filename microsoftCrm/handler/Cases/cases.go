package crm_handlers

import (
	crm_services "main/service/Cases"

	"github.com/gofiber/fiber/v2"
)

type CasesHandler struct {
	Service crm_services.Service_Cases
}

func (h CasesHandler) Cases(c *fiber.Ctx) error {
	return nil
}

func (h CasesHandler) CasesCreate(c *fiber.Ctx) error {
	return nil
}
func (h CasesHandler) CasesById(c *fiber.Ctx) error {
	return nil
}
func (h CasesHandler) CasesUpdate(c *fiber.Ctx) error {
	return nil
}
func (h CasesHandler) CasesDelete(c *fiber.Ctx) error {
	return nil
}
