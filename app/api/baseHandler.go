package api

import "gorm.io/gorm"

// BaseHandler is a struct that contains the database connection where all handlers can access it
type BaseHandler struct {
	DB *gorm.DB
}

// NewBaseHandler is a constructor for BaseHandler
func NewBaseHandler(db *gorm.DB) *BaseHandler {
	return &BaseHandler{
		DB: db,
	}
}
