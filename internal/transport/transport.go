package transport

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
)

type HTTP struct {
	host string
	port int
	log  *zap.Logger
	g    *gin.Engine
}

func NewHTTP(log *zap.Logger, port int) *HTTP {
	return &HTTP{
		port: port,
		log:  log,
		g:    gin.Default(),
	}
}

func (h *HTTP) Listen() error {
	h.routes()
	return h.g.Run(fmt.Sprintf(":%d", h.port))
}

func (h *HTTP) routes() {
	h.g.GET("/equipment", h.Equipment())
	h.g.GET("/events", h.Events())
	h.g.GET("/locations", h.Locations())
	h.g.GET("/waybills", h.Waybills())
	h.g.GET("/waybills/:id", h.WaybillsByID())
	h.g.GET("/waybills/:id/equipment", h.WaybillEquipment())
	h.g.GET("/waybills/events", h.WaybillEvents())
	h.g.GET("/waybills/locations", h.WaybillLocations())
}

func (h *HTTP) Equipment() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, "OK")
	}
}

func (h *HTTP) Events() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, "OK")
	}
}

func (h *HTTP) Locations() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, "OK")
	}
}

func (h *HTTP) Waybills() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, "OK")
	}
}

func (h *HTTP) WaybillsByID() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		if id == "" {
			c.JSON(400, "id not present")
		}

		c.JSON(http.StatusOK, "OK")
	}
}

func (h *HTTP) WaybillEquipment() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		if id == "" {
			c.JSON(400, "id not present")
		}

		c.JSON(http.StatusOK, "OK")
	}
}

func (h *HTTP) WaybillEvents() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		if id == "" {
			c.JSON(400, "id not present")
		}

		c.JSON(http.StatusOK, "OK")
	}
}

func (h *HTTP) WaybillLocations() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		if id == "" {
			c.JSON(400, "id not present")
		}

		c.JSON(http.StatusOK, "OK")
	}
}
