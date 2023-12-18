package api

import (
	"github.com/gin-gonic/gin"
	"github.com/onasunnymorning/chocolatelab/pkg/batch"
)

type BatchCreateRequest struct {
	Name string `json:"name"`
}

// CreateBatch godoc
// @Summary Create a new batch
// @Description Create a new batch
// @Tags batches
// @Accept json
// @Produce json
// @Param batch body BatchCreateRequest true "Batch"
// @Success 201 {object} batch.Batch
// @Failure 400
// @Failure 500
// @Router /batches [post]
func (h *BaseHandler) CreateBatch(ctx *gin.Context) {
	var req BatchCreateRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	batch := batch.NewBatch(req.Name)
	result := h.DB.WithContext(ctx).Create(&batch)
	if result.Error != nil {
		ctx.JSON(500, gin.H{"error": result.Error.Error()})
		return
	}
	ctx.JSON(201, batch)
}
