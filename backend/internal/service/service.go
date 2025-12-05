package service

import (
	"subscription/internal/repository"

	"go.uber.org/zap"
)

type Service struct {
	SubscriptionRepo repository.SubscriptionRepo
	Logger           *zap.Logger
}

func NewService(subscriptionRepo repository.SubscriptionRepo, logger *zap.Logger) *Service {
	return &Service{
		SubscriptionRepo: subscriptionRepo,
		Logger: logger,
	}
}
