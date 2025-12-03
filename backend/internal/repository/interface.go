package repository

import (
	"database/sql"
	"subscription/internal/model"
	"time"
)

type SubscriptionRepo interface {
	CreateSubscription(sub *model.Subscription) error
	GetCost(sub *model.SubscriptionFind) (int, error)
	GetServiceNameByUserId(userId string, serviceName string) error
	GetStartSubDate(userId string, serviceName string) (time.Time, error)
	GetEndSubDate(userId string, serviceName string) (sql.NullTime, error)
}
