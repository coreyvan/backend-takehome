package app

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"net/http"
)

type HTTP struct {
	host string
	port int
	log  *zap.Logger
	db   *gorm.DB
	g    *gin.Engine
}

func NewHTTP(log *zap.Logger, port int, db *gorm.DB) *HTTP {
	if db == nil {
		panic("db was nil")
	}

	return &HTTP{
		port: port,
		log:  log,
		g:    gin.Default(),
		db:   db,
	}
}

func (h *HTTP) Listen() error {
	h.routes()

	if err := h.migrate(); err != nil {
		return fmt.Errorf("migrating models: %w", err)
	}

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

func (h *HTTP) migrate() error {
	if err := h.db.AutoMigrate(&Location{}); err != nil {
		return fmt.Errorf("migrating locations: %w", err)
	}

	if err := h.db.AutoMigrate(&Event{}); err != nil {
		return fmt.Errorf("migrating locations: %w", err)
	}

	if err := h.db.AutoMigrate(&Waybill{}); err != nil {
		return fmt.Errorf("migrating locations: %w", err)
	}

	if err := h.db.AutoMigrate(&Equipment{}); err != nil {
		return fmt.Errorf("migrating locations: %w", err)
	}

	return nil
}

func (h *HTTP) Equipment() gin.HandlerFunc {
	return func(c *gin.Context) {
		var equipment []Equipment
		result := h.db.Find(&equipment)
		if result.Error != nil {
			c.JSON(http.StatusInternalServerError, "oops")
			return
		}
		c.JSON(http.StatusOK, equipment)
	}
}

func (h *HTTP) Events() gin.HandlerFunc {
	return func(c *gin.Context) {
		var events []Event
		result := h.db.Find(&events)
		if result.Error != nil {
			c.JSON(http.StatusInternalServerError, "oops")
			return
		}
		c.JSON(http.StatusOK, events)
	}
}

func (h *HTTP) Locations() gin.HandlerFunc {
	return func(c *gin.Context) {
		var locations []Location
		result := h.db.Find(&locations)
		if result.Error != nil {
			c.JSON(http.StatusInternalServerError, "oops")
			return
		}
		c.JSON(http.StatusOK, locations)
	}
}

func (h *HTTP) Waybills() gin.HandlerFunc {
	return func(c *gin.Context) {
		var waybills []Waybill
		result := h.db.Find(&waybills)
		if result.Error != nil {
			c.JSON(http.StatusInternalServerError, "oops")
			return
		}
		c.JSON(http.StatusOK, waybills)
	}
}

func (h *HTTP) WaybillsByID() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		if id == "" {
			c.JSON(400, "id not present")
			return
		}

		var waybill Waybill
		result := h.db.First(&waybill, id)
		if result.Error != nil {
			if errors.Is(result.Error, gorm.ErrRecordNotFound) {
				c.JSON(http.StatusNotFound, "Waybill not found")
				return
			}
			c.JSON(http.StatusInternalServerError, "Internal Server Error")
			return
		}

		c.JSON(http.StatusOK, waybill)
	}
}

func (h *HTTP) WaybillEquipment() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		if id == "" {
			c.JSON(400, "id not present")
			return
		}

		c.JSON(http.StatusOK, "OK")
	}
}

func (h *HTTP) WaybillEvents() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		if id == "" {
			c.JSON(400, "id not present")
			return
		}

		c.JSON(http.StatusOK, "OK")
	}
}

func (h *HTTP) WaybillLocations() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		if id == "" {
			c.JSON(400, "id not present")
			return
		}

		c.JSON(http.StatusOK, "OK")
	}
}