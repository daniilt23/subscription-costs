package handler

import (
	"errors"
	"net/http"
	"subscription/internal/dto"
	apperrors "subscription/internal/error"

	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateSubscription(c *gin.Context) {
	var req dto.CreateSubscriptionReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = h.Service.CreateSubscription(&req)
	if err != nil {
		if errors.Is(err, apperrors.ErrIncorrectData) {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": apperrors.ErrIncorrectData.Error(),
			})
			return
		}
		if errors.Is(err, apperrors.ErrNegativePrice) {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": apperrors.ErrNegativePrice.Error(),
			})
			return
		}
		if errors.Is(err, apperrors.ErrInvalidDataPeriod) {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": apperrors.ErrInvalidDataPeriod.Error(),
			})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, dto.CreateSubscriptionResp{
		Message: "Successfully create subscription",
	})
}

func (h *Handler) GetCost(c *gin.Context) {
	var req dto.GetCostReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	cost, err := h.Service.GetCost(&req)
	if err != nil {
		if errors.Is(err, apperrors.ErrNoService) {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": apperrors.ErrNoService.Error(),
			})
			return
		}
		if errors.Is(err, apperrors.ErrInvalidDataPeriod) {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": apperrors.ErrInvalidDataPeriod.Error(),
			})
			return
		}
		if errors.Is(err, apperrors.ErrUserWithoutSub) {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": apperrors.ErrUserWithoutSub.Error(),
			})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, dto.GetCostResp{
		Total: cost,
	})
}
