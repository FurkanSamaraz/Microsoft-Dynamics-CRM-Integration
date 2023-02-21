package crm_handlers

import (
	crm_services "main/service/Ping"

	"github.com/gofiber/fiber/v2"
)

type PingHandler struct {
	Service crm_services.Service_Ping
}

func (h PingHandler) Ping(c *fiber.Ctx) error {
	return nil
}
