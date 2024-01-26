package monitorHandlers

import (
	"github.com/cardthp/ecommerce-shop/config"
	"github.com/cardthp/ecommerce-shop/modules/monitor"
	"github.com/gofiber/fiber/v2"
)

//handler เป็นส่วนรับ api request/http ที่เข้ามา > รับค่าจาก fiber และ return error เท่านั้น
//ส่วนนี้จะส่งไปใช้ต่อที่ module

type IMonitorHandler interface {
	HealthCheck(c *fiber.Ctx) error //รับค่าจาก fiber และ return error เท่านั้น
}

type monitorHandler struct {
	cfg config.IConfig
}

func MonitorHandler(cfg config.IConfig) IMonitorHandler {
	return &monitorHandler{
		cfg: cfg,
	}
}

func (h *monitorHandler) HealthCheck(c *fiber.Ctx) error {
	res := &monitor.Monitor{
		Name:    h.cfg.App().Name(),
		Version: h.cfg.App().Version(),
	}
	return c.Status(fiber.StatusOK).JSON(res)
}
