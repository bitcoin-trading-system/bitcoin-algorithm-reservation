package models

import (
	"gorm.io/gorm"
)

type Reservation struct {
	ID                  int    `json:"id" gorm:"primaryKey"`
	AlgorithmType       string `json:"algorithm_type"`
	AlgorithmTypeDetail string `json:"algorithm_type_detail"`
	TimeConditionRegex  string `json:"time_condition_regex"`
	IsActive            bool   `json:"is_active"`
}

func NewReservation(algorithmType, algorithmTypeDetail, timeConditionRegex string, isActive bool) *Reservation {
	return &Reservation{
		AlgorithmType:       algorithmType,
		AlgorithmTypeDetail: algorithmTypeDetail,
		TimeConditionRegex:  timeConditionRegex,
		IsActive:            isActive,
	}
}

func (r *Reservation) Create(tx *gorm.DB) error {
	return tx.Create(r).Error
}

func FindReservationByID(tx *gorm.DB, id int) (Reservation, error) {
	var reservation Reservation
	err := tx.First(&reservation, id).Error
	return reservation, err
}

func FindAllReservations(tx *gorm.DB) ([]Reservation, error) {
	var reservations []Reservation
	err := tx.Find(&reservations).Error
	return reservations, err
}

func FindReservationsAndConditions(tx *gorm.DB, algorithmType, algorithmTypeDetail, timeConditionRegex *string, isActive *bool) ([]Reservation, error) {
	whereCondition := map[string]interface{}{}
	if algorithmType != nil {
		whereCondition["algorithm_type"] = *algorithmType
	}
	if algorithmTypeDetail != nil {
		whereCondition["algorithm_type_detail"] = *algorithmTypeDetail
	}
	if timeConditionRegex != nil {
		whereCondition["time_condition_regex"] = *timeConditionRegex
	}
	if isActive != nil {
		whereCondition["is_active"] = *isActive
	}

	var reservations []Reservation
	err := tx.Where(whereCondition).Find(&reservations).Error

	return reservations, err
}

func (r *Reservation) Update(tx *gorm.DB) error {
	return tx.Save(r).Error
}

func (r *Reservation) Delete(tx *gorm.DB) error {
	return tx.Delete(r).Error
}
