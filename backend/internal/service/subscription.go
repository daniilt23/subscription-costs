package service

import (
	"database/sql"
	"fmt"
	"subscription/internal/dto"
	apperrors "subscription/internal/error"
	"subscription/internal/model"
	"time"
)

func (s *Service) CreateSubscription(req *dto.CreateSubscriptionReq) error {
	if req.Price < 0 {
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
			return apperrors.ErrIncorrectData
		}
		if t2.Unix() <= t1.Unix() {
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
		return err
	}

	return nil
}

func (s *Service) GetCost(req *dto.GetCostReq) (int, error) {
	err := s.SubscriptionRepo.GetServiceNameByUserId(req.UserId, req.ServiceName)
	if err != nil {
		return 0, err
	}
	t1, err := s.timeParse(req.StartDate)
	if err != nil {
		return 0, apperrors.ErrIncorrectData
	}
	t2, err := s.timeParse(req.EndDate)
	if err != nil {
		return 0, apperrors.ErrIncorrectData
	}

	if t1.Unix() >= t2.Unix() {
		return 0, apperrors.ErrInvalidDataPeriod
	}

	date, err := s.SubscriptionRepo.GetStartSubDate(req.UserId, req.ServiceName)
	if err != nil {
		return 0, err
	}

	if t1.Unix() < date.Unix() {
		return 0, apperrors.ErrUserWithoutSub
	}

	dateEnd, err := s.SubscriptionRepo.GetEndSubDate(req.UserId, req.ServiceName)
	if err != nil {
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
		return 0, err
	}

	return cost, nil
}

func (s *Service) timeParse(date string) (time.Time, error) {
	const layout = "02-01-2006"
	t, err := time.Parse(layout, fmt.Sprintf("01-%s", date))

	return t, err
}
