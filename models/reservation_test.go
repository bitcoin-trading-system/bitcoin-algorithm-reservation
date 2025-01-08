package models

import (
	"reflect"
	"testing"

	"github.com/bitcoin-trading-system/bitcoin-algorithm-reservation/utils"
)

func TestReservation_Create(t *testing.T) {
	type fields struct {
		AlgorithmType       string
		AlgorithmTypeDetail string
		TimeConditionRegex  string
		IsActive            bool
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "Create",
			fields: fields{
				AlgorithmType:       utils.GenerateRandomString(10),
				AlgorithmTypeDetail: utils.GenerateRandomString(10),
				TimeConditionRegex:  utils.GenerateRandomString(10),
				IsActive:            true,
			},
		},
		{
			name: "Create",
			fields: fields{
				AlgorithmType:       "",
				AlgorithmTypeDetail: "",
				TimeConditionRegex:  "",
				IsActive:            false,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewReservation(tt.fields.AlgorithmType, tt.fields.AlgorithmTypeDetail, tt.fields.TimeConditionRegex, tt.fields.IsActive)
			if err := r.Create(GetterDB()); (err != nil) != tt.wantErr {
				t.Errorf("Reservation.Create() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestFindReservationByID(t *testing.T) {
	tests := []struct {
		name      string
		want      *Reservation
		preCreate bool
		wantErr   bool
	}{
		{
			name:      "FindReservationByID",
			want:      NewReservation(utils.GenerateRandomString(10), utils.GenerateRandomString(10), utils.GenerateRandomString(10), true),
			preCreate: true,
			wantErr:   false,
		},
		{
			name:      "FindReservationByID_NotFound",
			want:      NewReservation(utils.GenerateRandomString(10), utils.GenerateRandomString(10), utils.GenerateRandomString(10), true),
			preCreate: false,
			wantErr:   true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.preCreate {
				tt.want.Create(GetterDB())
			}

			got, err := FindReservationByID(GetterDB(), tt.want.ID)
			if err != nil {
				if !tt.wantErr {
					t.Errorf("FindReservationByID() error = %v, wantErr %v", err, tt.wantErr)
				}
				return
			}

			if !reflect.DeepEqual(got, *tt.want) {
				t.Errorf("FindReservationByID() = %v, want %v", got, *tt.want)
			}
		})
	}
}

func TestFindAllReservations(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "FindAllReservations",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewReservation(utils.GenerateRandomString(10), utils.GenerateRandomString(10), utils.GenerateRandomString(10), true)
			if err := r.Create(GetterDB()); err != nil {
				t.Errorf("Reservation.Create() error = %v", err)
			}
			got, err := FindAllReservations(GetterDB())
			if err != nil {
				t.Errorf("FindAllReservations() error = %v", err)
				return
			}

			if len(got) == 0 {
				t.Errorf("FindAllReservations() = %v, want not empty", got)
			}
		})

	}
}

func TestFindReservationsAndConditions(t *testing.T) {
	type condition struct {
		algorithmType       *string
		algorithmTypeDetail *string
		timeConditionRegex  *string
		isActive            *bool
	}

	tests := []struct {
		name      string
		condition condition
	}{
		{
			name: "FindReservationsAndConditions",
			condition: condition{
				algorithmType:       stringPointer(utils.GenerateRandomString(10)),
				algorithmTypeDetail: stringPointer(utils.GenerateRandomString(10)),
				timeConditionRegex:  stringPointer(utils.GenerateRandomString(10)),
				isActive:            boolPointer(true),
			},
		},
		{
			name: "FindReservationsAndConditions",
			condition: condition{
				algorithmType:       nil,
				algorithmTypeDetail: nil,
				timeConditionRegex:  nil,
				isActive:            nil,
			},
		},
		{
			name: "FindReservationsAndConditions",
			condition: condition{
				algorithmType:       stringPointer(utils.GenerateRandomString(10)),
				algorithmTypeDetail: nil,
				timeConditionRegex:  nil,
				isActive:            nil,
			},
		},
		{
			name: "FindReservationsAndConditions",
			condition: condition{
				algorithmType:       nil,
				algorithmTypeDetail: stringPointer(utils.GenerateRandomString(10)),
				timeConditionRegex:  nil,
				isActive:            nil,
			},
		},
		{
			name: "FindReservationsAndConditions",
			condition: condition{
				algorithmType:       nil,
				algorithmTypeDetail: nil,
				timeConditionRegex:  stringPointer(utils.GenerateRandomString(10)),
				isActive:            nil,
			},
		},
		{
			name: "FindReservationsAndConditions",
			condition: condition{
				algorithmType:       nil,
				algorithmTypeDetail: nil,
				timeConditionRegex:  nil,
				isActive:            boolPointer(true),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			algorithmType := utils.GenerateRandomString(10)
			algorithmTypeDetail := utils.GenerateRandomString(10)
			timeConditionRegex := utils.GenerateRandomString(10)
			isActive := true

			if tt.condition.algorithmType != nil {
				algorithmType = *tt.condition.algorithmType
			}
			if tt.condition.algorithmTypeDetail != nil {
				algorithmTypeDetail = *tt.condition.algorithmTypeDetail
			}
			if tt.condition.timeConditionRegex != nil {
				timeConditionRegex = *tt.condition.timeConditionRegex
			}
			if tt.condition.isActive != nil {
				isActive = *tt.condition.isActive
			}

			r := NewReservation(algorithmType, algorithmTypeDetail, timeConditionRegex, isActive)
			if err := r.Create(GetterDB()); err != nil {
				t.Errorf("Reservation.Create() error = %v", err)
			}

			got, err := FindReservationsAndConditions(GetterDB(), tt.condition.algorithmType, tt.condition.algorithmTypeDetail, tt.condition.timeConditionRegex, tt.condition.isActive)
			if err != nil {
				t.Errorf("FindReservationsAndConditions() error = %v", err)
				return
			}

			for _, g := range got {
				if tt.condition.algorithmType != nil {
					if g.AlgorithmType != *tt.condition.algorithmType {
						t.Errorf("FindReservationsAndConditions() = %v, want %v", g.AlgorithmType, *tt.condition.algorithmType)
					}
				}

				if tt.condition.algorithmTypeDetail != nil {
					if g.AlgorithmTypeDetail != *tt.condition.algorithmTypeDetail {
						t.Errorf("FindReservationsAndConditions() = %v, want %v", g.AlgorithmTypeDetail, *tt.condition.algorithmTypeDetail)
					}
				}

				if tt.condition.timeConditionRegex != nil {
					if g.TimeConditionRegex != *tt.condition.timeConditionRegex {
						t.Errorf("FindReservationsAndConditions() = %v, want %v", g.TimeConditionRegex, *tt.condition.timeConditionRegex)
					}
				}

				if tt.condition.isActive != nil {
					if g.IsActive != *tt.condition.isActive {
						t.Errorf("FindReservationsAndConditions() = %v, want %v", g.IsActive, *tt.condition.isActive)
					}
				}
			}

			if tt.condition.algorithmType == nil && tt.condition.algorithmTypeDetail == nil && tt.condition.timeConditionRegex == nil && tt.condition.isActive == nil {
				allReservations, err := FindAllReservations(GetterDB())
				if err != nil {
					t.Errorf("FindAllReservations() error = %v", err)
					return
				}
				if len(got) != len(allReservations) {
					t.Errorf("FindReservationsAndConditions() = %v, want %v", got, allReservations)
					return
				}
			}
		})
	}
}

func TestReservation_Update(t *testing.T) {
	type fields struct {
		AlgorithmType       string
		AlgorithmTypeDetail string
		TimeConditionRegex  string
		IsActive            bool
	}
	tests := []struct {
		name         string
		fields       fields
		updateFields fields
		wantErr      bool
	}{
		{
			name: "Update",
			fields: fields{
				AlgorithmType:       utils.GenerateRandomString(10),
				AlgorithmTypeDetail: utils.GenerateRandomString(10),
				TimeConditionRegex:  utils.GenerateRandomString(10),
				IsActive:            true,
			},
			updateFields: fields{
				AlgorithmType:       utils.GenerateRandomString(10),
				AlgorithmTypeDetail: utils.GenerateRandomString(10),
				TimeConditionRegex:  utils.GenerateRandomString(10),
				IsActive:            false,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewReservation(tt.fields.AlgorithmType, tt.fields.AlgorithmTypeDetail, tt.fields.TimeConditionRegex, tt.fields.IsActive)
			if err := r.Create(GetterDB()); err != nil {
				t.Errorf("Reservation.Create() error = %v", err)
			}

			r.AlgorithmType = tt.updateFields.AlgorithmType
			r.AlgorithmTypeDetail = tt.updateFields.AlgorithmTypeDetail
			r.TimeConditionRegex = tt.updateFields.TimeConditionRegex
			r.IsActive = tt.updateFields.IsActive

			if err := r.Update(GetterDB()); (err != nil) != tt.wantErr {
				t.Errorf("Reservation.Update() error = %v, wantErr %v", err, tt.wantErr)
			}

			got, err := FindReservationByID(GetterDB(), r.ID)
			if err != nil {
				t.Errorf("FindReservationByID() error = %v", err)
				return
			}

			if !reflect.DeepEqual(got, *r) {
				t.Errorf("Reservation.Update() = %v, want %v", got, *r)
			}
		})
	}
}

func TestReservation_Delete(t *testing.T) {
	type fields struct {
		ID                  int
		AlgorithmType       string
		AlgorithmTypeDetail string
		TimeConditionRegex  string
		IsActive            bool
	}
	tests := []struct {
		name      string
		preInsert bool
		fields    fields
		wantErr   bool
	}{
		{
			name: "Delete",
			preInsert: true,
			fields: fields{
				AlgorithmType:       utils.GenerateRandomString(10),
				AlgorithmTypeDetail: utils.GenerateRandomString(10),
				TimeConditionRegex:  utils.GenerateRandomString(10),
				IsActive:            true,
			},
			wantErr: false,
		},
		{
			name: "Delete_NotFound",
			preInsert: false,
			fields: fields{
				AlgorithmType:       utils.GenerateRandomString(10),
				AlgorithmTypeDetail: utils.GenerateRandomString(10),
				TimeConditionRegex:  utils.GenerateRandomString(10),
				IsActive:            true,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewReservation(tt.fields.AlgorithmType, tt.fields.AlgorithmTypeDetail, tt.fields.TimeConditionRegex, tt.fields.IsActive)
			if tt.preInsert {
				if err := r.Create(GetterDB()); err != nil {
					t.Errorf("Reservation.Create() error = %v", err)
				}
			}

			if err := r.Delete(GetterDB()); (err != nil) != tt.wantErr {
				t.Errorf("Reservation.Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func boolPointer(b bool) *bool {
	return &b
}

func stringPointer(s string) *string {
	return &s
}
