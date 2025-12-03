package subscription

import (
	"database/sql"
	"errors"
	apperrors "subscription/internal/error"
	"subscription/internal/model"
	"time"
)

type SubscriptionRepoSQL struct {
	Db *sql.DB
}

func NewSubscriptionRepoSQL(db *sql.DB) *SubscriptionRepoSQL {
	return &SubscriptionRepoSQL{
		Db: db,
	}
}

func (r *SubscriptionRepoSQL) CreateSubscription(sub *model.Subscription) error {
	query := `
	INSERT INTO subscriptions (user_id, service_name, price, start_date, end_date)
	VALUES($1, $2, $3, $4, $5)`

	_, err := r.Db.Exec(query, sub.UserId, sub.ServiceName, sub.Price, sub.DateStart, sub.DateEnd)
	if err != nil {
		return err
	}

	return nil
}

func (r *SubscriptionRepoSQL) GetServiceNameByUserId(userId string, serviceName string) error {
	query := `
	SELECT service_name FROM subscriptions
	WHERE user_id = $1 AND service_name = $2`

	var service string
	err := r.Db.QueryRow(query, userId, serviceName).Scan(&service)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return apperrors.ErrNoService
		}
		return err
	}

	return nil
}

func (r *SubscriptionRepoSQL) GetStartSubDate(userId string, serviceName string) (time.Time, error) {
	query := `
	SELECT start_date FROM subscriptions
	WHERE user_id = $1 AND service_name = $2`
	
	var date time.Time
	err := r.Db.QueryRow(query, userId, serviceName).Scan(&date)
	if err != nil {
		return time.Time{}, err
	}

	return date, nil
}

func (r *SubscriptionRepoSQL) GetEndSubDate(userId string, serviceName string) (sql.NullTime, error) {
	query := `
	SELECT end_date FROM subscriptions
	WHERE user_id = $1 AND service_name = $2`
	
	var date sql.NullTime
	err := r.Db.QueryRow(query, userId, serviceName).Scan(&date)
	if err != nil {
		return date, err
	}

	return date, nil
}

func (r *SubscriptionRepoSQL) GetCost(sub *model.SubscriptionFind) (int, error) {
	query := `
	SELECT SUM(price) * ((EXTRACT(YEAR FROM $1::DATE) - EXTRACT(YEAR FROM $2::DATE)) * 12 
	+ (EXTRACT(MONTH FROM $1) - EXTRACT(MONTH FROM $2)))
	FROM subscriptions
	WHERE user_id = $3 AND service_name = $4`

	var cost int
	err := r.Db.QueryRow(query, sub.DateEnd, sub.DateStart, sub.UserId, sub.ServiceName).Scan(&cost)
	if err != nil {
		return 0, err
	}

	return cost, nil
}
