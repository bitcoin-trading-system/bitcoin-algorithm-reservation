package usecase

import (
	"errors"

	"github.com/bitcoin-trading-system/bitcoin-algorithm-reservation/config"
	"github.com/bitcoin-trading-system/bitcoin-algorithm-reservation/models"
)

type ReservationUseCase struct {
	Config config.Config
}

type IReservationUseCase interface {
	CreateReservation(algorithmType, algorithmTypeDetail, timeConditionRegex string, isActive *bool) error
	FindReservationByID(id int) (models.Reservation, error)
	FindReservations(algorithmType, algorithmTypeDetail, timeConditionRegex *string, isActive *bool) ([]models.Reservation, error)
	UpdateReservation(id int, algorithmType, algorithmTypeDetail, timeConditionRegex *string, isActive *bool) error
	DeleteReservation(id int) error
}

func NewReservationUseCase(config config.Config) IReservationUseCase {
	return &ReservationUseCase{
		Config: config,
	}
}

func (ru *ReservationUseCase) CreateReservation(algorithmType, algorithmTypeDetail, timeConditionRegex string, isActive *bool) error {
	if algorithmType == "" {
		return errors.New("algorithm type is required")
	}
	if algorithmTypeDetail == "" {
		return errors.New("algorithm type detail is required")
	}
	if timeConditionRegex == "" {
		return errors.New("time condition regex is required")
	}
	if isActive == nil {
		return errors.New("is active is required")
	}

	reservation := models.NewReservation(algorithmType, algorithmTypeDetail, timeConditionRegex, *isActive)
	return reservation.Create(models.GetterDB())
}

func (ru *ReservationUseCase) FindReservationByID(id int) (models.Reservation, error) {
	if id == 0 {
		return models.Reservation{}, errors.New("id is required")
	}

	return models.FindReservationByID(models.GetterDB(), id)
}

func (ru *ReservationUseCase) FindReservations(algorithmType, algorithmTypeDetail, timeConditionRegex *string, isActive *bool) ([]models.Reservation, error) {
	return models.FindReservationsAndConditions(models.GetterDB(), algorithmType, algorithmTypeDetail, timeConditionRegex, isActive)
}

func (ru *ReservationUseCase) UpdateReservation(id int, algorithmType, algorithmTypeDetail, timeConditionRegex *string, isActive *bool) error {
	if id == 0 {
		return errors.New("id is required")
	}
	if algorithmType == nil {
		return errors.New("algorithm type is required")
	}
	if algorithmTypeDetail == nil {
		return errors.New("algorithm type detail is required")
	}
	if timeConditionRegex == nil {
		return errors.New("time condition regex is required")
	}
	if isActive == nil {
		return errors.New("is active is required")
	}

	r, err := models.FindReservationByID(models.GetterDB(), id)
	if err != nil {
		return err
	}

	r.AlgorithmType = *algorithmType
	r.AlgorithmTypeDetail = *algorithmTypeDetail
	r.TimeConditionRegex = *timeConditionRegex
	r.IsActive = *isActive

	return r.Update(models.GetterDB())
}

func (ru *ReservationUseCase) DeleteReservation(id int) error {
	if id == 0 {
		return errors.New("id is required")
	}

	r, err := models.FindReservationByID(models.GetterDB(), id)
	if err != nil {
		return err
	}

	return r.Delete(models.GetterDB())
}
