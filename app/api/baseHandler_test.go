package api

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestNewBaseHandler(t *testing.T) {
	db := &gorm.DB{} // Replace with your actual DB instance

	handler := NewBaseHandler(db)

	assert.NotNil(t, handler)
	assert.Equal(t, db, handler.DB)
}
