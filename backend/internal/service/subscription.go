package service

import (
	"database/sql"
	"errors"
	"fmt"
	"subscription/internal/dto"
	apperrors "subscription/internal/error"
	"subscription/internal/model"
	"time"

	"go.uber.org/zap"
)

func (s *Service) CreateSubscription(req *dto.CreateSubscriptionReq) error {
	createSubLogger := s.Logger.With(zap.String("Function", "CreateSubscription"))
	if req.Price < 0 {
		createSubLogger.Warn("Error price negative", zap.Int("Price", req.Price))
		return apperrors.ErrNegativePrice
	}
	t1, err := s.timeParse(req.StartDate)
	if err != nil {
		return apperrors.ErrIncorrectData
	}

	var EndDate sql.NullString
	if req.EndDate != "" {
		t2, err := s.timeParse(req.EndDate)
		if err != nil {
			createSubLogger.Warn("Error parse time", zap.String("End date", req.EndDate))
			return apperrors.ErrIncorrectData
		}
		if t2.Unix() <= t1.Unix() {
			createSubLogger.Warn("Error end date less than start date",
				zap.String("Start date", req.StartDate),
				zap.String("End date", req.EndDate))
			return apperrors.ErrInvalidDataPeriod
		}
		EndDate = sql.NullString{
			Valid:  true,
			String: t2.Format("2006-01-02"),
		}
	} else {
		EndDate = sql.NullString{
			Valid: false,
		}
	}

	subModel := model.Subscription{
		UserId:      req.UserId,
		ServiceName: req.ServiceName,
		Price:       req.Price,
		DateStart:   t1.Format("2006-01-02"),
		DateEnd:     EndDate,
	}

	err = s.SubscriptionRepo.CreateSubscription(&subModel)
	if err != nil {
		createSubLogger.Error("Cannot create subscription",
			zap.String("Error", err.Error()))
		return err
	}

	return nil
}

func (s *Service) GetCost(req *dto.GetCostReq) (int, error) {
	getCostLogger := s.Logger.With(zap.String("Function", "GetCost"))
	err := s.SubscriptionRepo.GetServiceNameByUserId(req.UserId, req.ServiceName)
	if err != nil {
		if errors.Is(err, apperrors.ErrNoService) {
			getCostLogger.Warn("Service not found", zap.String("Error", "Not found"))
		}
		getCostLogger.Error("Cannot get service name", zap.String("Error", err.Error()))
		return 0, err
	}
	t1, err := s.timeParse(req.StartDate)
	if err != nil {
		getCostLogger.Warn("Error parse start date", zap.String("Start date", req.StartDate))
		return 0, apperrors.ErrIncorrectData
	}
	t2, err := s.timeParse(req.EndDate)
	if err != nil {
		getCostLogger.Warn("Error parse end date", zap.String("End date", req.EndDate))
		return 0, apperrors.ErrIncorrectData
	}

	if t1.Unix() >= t2.Unix() {
		getCostLogger.Warn("Error end date less than start date",
			zap.String("Start date", req.StartDate),
			zap.String("End date", req.EndDate))
		return 0, apperrors.ErrInvalidDataPeriod
	}

	date, err := s.SubscriptionRepo.GetStartSubDate(req.UserId, req.ServiceName)
	if err != nil {
		getCostLogger.Error("Cannot take start date", zap.String("Error", err.Error()))
		return 0, err
	}

	if t1.Unix() < date.Unix() {
		getCostLogger.Error("User dont have sub at this time", zap.Time("Start date", date))
		return 0, apperrors.ErrUserWithoutSub
	}

	dateEnd, err := s.SubscriptionRepo.GetEndSubDate(req.UserId, req.ServiceName)
	if err != nil {
		getCostLogger.Error("Cannot take end date", zap.String("Error", err.Error()))
		return 0, err
	}
	if dateEnd.Valid {
		if dateEnd.Time.Unix() < t2.Unix() {
			t2 = dateEnd.Time
		}
	}

	subFind := model.SubscriptionFind{
		UserId:      req.UserId,
		ServiceName: req.ServiceName,
		DateStart:   t1.Format("2006-01-02"),
		DateEnd:     t2.Format("2006-01-02"),
	}

	cost, err := s.SubscriptionRepo.GetCost(&subFind)
	if err != nil {
		getCostLogger.Error("Cannot take cost of user subs", zap.String("Error", err.Error()))
		return 0, err
	}

	return cost, nil
}

func (s *Service) timeParse(date string) (time.Time, error) {
	const layout = "02-01-2006"
	t, err := time.Parse(layout, fmt.Sprintf("01-%s", date))

	return t, err
}
