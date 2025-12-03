package service

import "subscription/internal/repository"

type Service struct {
	SubscriptionRepo repository.SubscriptionRepo
}

func NewService(subscriptionRepo repository.SubscriptionRepo) *Service {
	return &Service{
		SubscriptionRepo: subscriptionRepo,
	}
}
