package app

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"net/http"
	"time"
)

type HTTP struct {
	port string
	log  *zap.Logger
	db   *gorm.DB
	g    *gin.Engine
}

func NewHTTP(log *zap.Logger, port string, db *gorm.DB) *HTTP {
	if db == nil {
		panic("db was nil")
	}

	gin.SetMode(gin.ReleaseMode)

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

	return h.g.Run(fmt.Sprintf(":%s", h.port))
}

func (h *HTTP) routes() {
	h.g.GET("/equipment", h.Equipment())
	h.g.GET("/events", h.Events())
	h.g.GET("/locations", h.Locations())
	h.g.GET("/waybills", h.Waybills())
	h.g.GET("/waybills/:id", h.WaybillsByID())
	h.g.GET("/waybills/:id/equipment", h.WaybillEquipment())
	h.g.GET("/waybills/:id/events", h.WaybillEvents())
	h.g.GET("/waybills/:id/locations", h.WaybillLocations())
	h.g.GET("/waybills/:id/route", h.WaybillRoute())
	h.g.GET("/waybills/:id/parties", h.WaybillParties())
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
			h.log.Sugar().Errorf("finding all equipment: %v", result.Error)
			c.JSON(http.StatusInternalServerError, "Internal server error")
			return
		}
		c.JSON(http.StatusOK, equipment)
	}
}

func (h *HTTP) Events() gin.HandlerFunc {
	return func(c *gin.Context) {
		where := h.db.Model(&Event{})
		after := c.Query("after")
		if after != "" {
			t, err := time.Parse(time.RFC3339, after)
			if err != nil {
				h.log.Sugar().Errorf("parsing query param after: %v", err)
				c.JSON(http.StatusBadRequest, "could not parse query param after")
				return
			}
			where = where.Where("events.posting_date > ?", t)
		}

		var events []Event
		result := where.Find(&events)
		if result.Error != nil {
			h.log.Sugar().Errorf("finding events: %v", result.Error)
			c.JSON(http.StatusInternalServerError, "Internal Server Error")
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
			h.log.Sugar().Errorf("finding all locations: %v", result.Error)
			c.JSON(http.StatusInternalServerError, "Internal Server Error")
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
			h.log.Sugar().Errorf("finding all waybills: %v", result.Error)
			c.JSON(http.StatusInternalServerError, "Internal Server Error")
			return
		}
		c.JSON(http.StatusOK, waybills)
	}
}

func (h *HTTP) WaybillsByID() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		if id == "" {
			c.JSON(http.StatusBadRequest, "id not present")
			return
		}

		var waybill Waybill
		result := h.db.First(&waybill, id)
		if result.Error != nil {
			if errors.Is(result.Error, gorm.ErrRecordNotFound) {
				c.JSON(http.StatusNotFound, "Waybill not found")
				return
			}
			h.log.Sugar().Errorf("finding waybill by id: %v", result.Error)
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
			c.JSON(http.StatusBadRequest, "id not present")
			return
		}

		var equipment []Equipment
		h.db.Raw("SELECT equipment.* FROM waybills JOIN equipment on waybills.equipment_id = equipment.equipment_id WHERE waybills.id = ?", id).Scan(&equipment)

		c.JSON(http.StatusOK, equipment)
	}
}

func (h *HTTP) WaybillEvents() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		if id == "" {
			c.JSON(http.StatusBadRequest, "id not present")
			return
		}

		where := h.db.Where("waybill_id = ?", id)
		after := c.Query("after")
		if after != "" {
			t, err := time.Parse(time.RFC3339, after)
			if err != nil {
				h.log.Sugar().Errorf("parsing query param after: %v", err)
				c.JSON(http.StatusBadRequest, "could not parse query param after")
				return
			}
			where = where.Where("waybill_id = ? AND posting_date > ?", id, t)
		}

		var events []Event
		where.Find(&events)

		c.JSON(http.StatusOK, events)
	}
}

func (h *HTTP) WaybillLocations() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		if id == "" {
			c.JSON(http.StatusBadRequest, "id not present")
			return
		}

		var waybill Waybill
		result := h.db.First(&waybill, id)
		if result.Error != nil {
			if errors.Is(result.Error, gorm.ErrRecordNotFound) {
				c.JSON(http.StatusNotFound, "Waybill not found")
				return
			}
			h.log.Sugar().Errorf("finding waybill by id: %v", result.Error)
			c.JSON(http.StatusInternalServerError, "Internal Server Error")
			return
		}

		var locations []Location
		h.db.Raw("SELECT * FROM locations where id IN (?,?)", waybill.OriginID, waybill.DestinationID).Scan(&locations)

		c.JSON(http.StatusOK, locations)
	}
}

func (h *HTTP) WaybillRoute() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		if id == "" {
			c.JSON(http.StatusBadRequest, "id not present")
			return
		}

		var waybill Waybill
		result := h.db.First(&waybill, id)
		if result.Error != nil {
			if errors.Is(result.Error, gorm.ErrRecordNotFound) {
				c.JSON(http.StatusNotFound, "Waybill not found")
				return
			}
			h.log.Sugar().Errorf("finding waybill by id: %v", result.Error)
			c.JSON(http.StatusInternalServerError, "Internal Server Error")
			return
		}

		var route []RoutePart
		if err := json.Unmarshal([]byte(waybill.Routes), &route); err != nil {
			h.log.Sugar().Errorf("unmarshaling routes: %v", err)
			c.JSON(http.StatusInternalServerError, "Internal Server Error")
			return
		}

		c.JSON(http.StatusOK, route)
	}
}

func (h *HTTP) WaybillParties() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		if id == "" {
			c.JSON(http.StatusBadRequest, "id not present")
			return
		}

		var waybill Waybill
		result := h.db.First(&waybill, id)
		if result.Error != nil {
			if errors.Is(result.Error, gorm.ErrRecordNotFound) {
				c.JSON(http.StatusNotFound, "Waybill not found")
				return
			}
			h.log.Sugar().Errorf("finding waybill by id: %v", result.Error)
			c.JSON(http.StatusInternalServerError, "Internal Server Error")
			return
		}

		var parties []Party
		if err := json.Unmarshal([]byte(waybill.Parties), &parties); err != nil {
			h.log.Sugar().Errorf("unmarshaling parties: %v", err)
			c.JSON(http.StatusInternalServerError, "Internal Server Error")
			return
		}

		c.JSON(http.StatusOK, parties)
	}
}
