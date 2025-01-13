package usecase

import (
	"reflect"
	"testing"

	"github.com/bitcoin-trading-system/bitcoin-algorithm-reservation/config"
	"github.com/bitcoin-trading-system/bitcoin-algorithm-reservation/models"
	"github.com/bitcoin-trading-system/bitcoin-algorithm-reservation/utils"
)

func TestReservationUseCase_CreateReservation(t *testing.T) {
	type fields struct {
		Config config.Config
	}
	type args struct {
		algorithmType       string
		algorithmTypeDetail string
		timeConditionRegex  string
		isActive            *bool
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "CreateReservation",
			fields: fields{
				Config: UseCaseTestConfig,
			},
			args: args{
				algorithmType:       utils.GenerateRandomString(10),
				algorithmTypeDetail: utils.GenerateRandomString(10),
				timeConditionRegex:  utils.GenerateRandomString(10),
				isActive:            boolPointer(utils.GenerateRandomBool()),
			},
			wantErr: false,
		},
		{
			name: "CreateReservation_RequiredAlgorithmType",
			fields: fields{
				Config: UseCaseTestConfig,
			},
			args: args{
				algorithmType:       "",
				algorithmTypeDetail: utils.GenerateRandomString(10),
				timeConditionRegex:  utils.GenerateRandomString(10),
				isActive:            boolPointer(utils.GenerateRandomBool()),
			},
			wantErr: true,
		},
		{
			name: "CreateReservation_RequiredAlgorithmTypeDetail",
			fields: fields{
				Config: UseCaseTestConfig,
			},
			args: args{
				algorithmType:       utils.GenerateRandomString(10),
				algorithmTypeDetail: "",
				timeConditionRegex:  utils.GenerateRandomString(10),
				isActive:            boolPointer(utils.GenerateRandomBool()),
			},
			wantErr: true,
		},
		{
			name: "CreateReservation_RequiredTimeConditionRegex",
			fields: fields{
				Config: UseCaseTestConfig,
			},
			args: args{
				algorithmType:       utils.GenerateRandomString(10),
				algorithmTypeDetail: utils.GenerateRandomString(10),
				timeConditionRegex:  "",
				isActive:            boolPointer(utils.GenerateRandomBool()),
			},
			wantErr: true,
		},
		{
			name: "CreateReservation_RequiredIsActive",
			fields: fields{
				Config: UseCaseTestConfig,
			},
			args: args{
				algorithmType:       utils.GenerateRandomString(10),
				algorithmTypeDetail: utils.GenerateRandomString(10),
				timeConditionRegex:  utils.GenerateRandomString(10),
				isActive:            nil,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ru := NewReservationUseCase(tt.fields.Config)
			if err := ru.CreateReservation(tt.args.algorithmType, tt.args.algorithmTypeDetail, tt.args.timeConditionRegex, tt.args.isActive); err != nil {
				if !tt.wantErr {
					t.Errorf("ReservationUseCase.CreateReservation() error = %v, wantErr %v", err, tt.wantErr)
				}
			}
		})
	}
}

func TestReservationUseCase_FindReservationByID(t *testing.T) {
	// テスト用のデータを作成
	reservation := models.NewReservation(utils.GenerateRandomString(10), utils.GenerateRandomString(10), utils.GenerateRandomString(10), utils.GenerateRandomBool())
	if err := reservation.Create(models.GetterDB()); err != nil {
		panic(err)
	}

	type fields struct {
		Config config.Config
	}
	type args struct {
		id int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    models.Reservation
		wantErr bool
	}{
		{
			name: "FindReservationByID",
			fields: fields{
				Config: UseCaseTestConfig,
			},
			args: args{
				id: reservation.ID,
			},
			want:    *reservation,
			wantErr: false,
		},
		{
			name: "FindReservationByID_RequiredID",
			fields: fields{
				Config: UseCaseTestConfig,
			},
			args: args{
				id: 0,
			},
			want:    models.Reservation{},
			wantErr: true,
		},
		{
			name: "FindReservationByID_NotFound",
			fields: fields{
				Config: UseCaseTestConfig,
			},
			args: args{
				id: 999999,
			},
			want:    models.Reservation{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ru := NewReservationUseCase(tt.fields.Config)
			got, err := ru.FindReservationByID(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("ReservationUseCase.FindReservationByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReservationUseCase.FindReservationByID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReservationUseCase_FindReservations(t *testing.T) {
	// テスト用のデータを作成
	algorithmType := utils.GenerateRandomString(10)
	algorithmTypeDetail := utils.GenerateRandomString(10)
	timeConditionRegex := utils.GenerateRandomString(10)
	isActive := utils.GenerateRandomBool()

	reservation := models.NewReservation(algorithmType, algorithmTypeDetail, timeConditionRegex, isActive)
	if err := reservation.Create(models.GetterDB()); err != nil {
		panic(err)
	}

	type fields struct {
		Config config.Config
	}
	type args struct {
		algorithmType       *string
		algorithmTypeDetail *string
		timeConditionRegex  *string
		isActive            *bool
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "FindReservations",
			fields: fields{
				Config: UseCaseTestConfig,
			},
			args: args{
				algorithmType:       &algorithmType,
				algorithmTypeDetail: &algorithmTypeDetail,
				timeConditionRegex:  &timeConditionRegex,
				isActive:            &isActive,
			},
			wantErr: false,
		},
		{
			name: "FindReservations_NotFound",
			fields: fields{
				Config: UseCaseTestConfig,
			},
			args: args{
				algorithmType:       stringPointer(utils.GenerateRandomString(10)),
				algorithmTypeDetail: stringPointer(utils.GenerateRandomString(10)),
				timeConditionRegex:  stringPointer(utils.GenerateRandomString(10)),
				isActive:            boolPointer(!isActive),
			},
			wantErr: false,
		},
		{
			name: "FindReservations_AlgorithmType_Nil",
			fields: fields{
				Config: UseCaseTestConfig,
			},
			args: args{
				algorithmType:       nil,
				algorithmTypeDetail: &algorithmTypeDetail,
				timeConditionRegex:  &timeConditionRegex,
				isActive:            &isActive,
			},
			wantErr: false,
		},
		{
			name: "FindReservations_AlgorithmTypeDetail_Nil",
			fields: fields{
				Config: UseCaseTestConfig,
			},
			args: args{
				algorithmType:       &algorithmType,
				algorithmTypeDetail: nil,
				timeConditionRegex:  &timeConditionRegex,
				isActive:            &isActive,
			},
			wantErr: false,
		},
		{
			name: "FindReservations_TimeConditionRegex_Nil",
			fields: fields{
				Config: UseCaseTestConfig,
			},
			args: args{
				algorithmType:       &algorithmType,
				algorithmTypeDetail: &algorithmTypeDetail,
				timeConditionRegex:  nil,
				isActive:            &isActive,
			},
			wantErr: false,
		},
		{
			name: "FindReservations_IsActive_Nil",
			fields: fields{
				Config: UseCaseTestConfig,
			},
			args: args{
				algorithmType:       &algorithmType,
				algorithmTypeDetail: &algorithmTypeDetail,
				timeConditionRegex:  &timeConditionRegex,
				isActive:            nil,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ru := NewReservationUseCase(tt.fields.Config)
			gots, err := ru.FindReservations(tt.args.algorithmType, tt.args.algorithmTypeDetail, tt.args.timeConditionRegex, tt.args.isActive)
			if err != nil {
				if !tt.wantErr {
					t.Errorf("ReservationUseCase.FindReservations() error = %v, wantErr %v", err, tt.wantErr)
				}
				return
			}
			for _, got := range gots {
				if tt.args.algorithmType != nil && got.AlgorithmType != *tt.args.algorithmType {
					t.Errorf("ReservationUseCase.FindReservations() = %v, want %v", got.AlgorithmType, *tt.args.algorithmType)
				}
				if tt.args.algorithmTypeDetail != nil && got.AlgorithmTypeDetail != *tt.args.algorithmTypeDetail {
					t.Errorf("ReservationUseCase.FindReservations() = %v, want %v", got.AlgorithmTypeDetail, *tt.args.algorithmTypeDetail)
				}
				if tt.args.timeConditionRegex != nil && got.TimeConditionRegex != *tt.args.timeConditionRegex {
					t.Errorf("ReservationUseCase.FindReservations() = %v, want %v", got.TimeConditionRegex, *tt.args.timeConditionRegex)
				}
				if tt.args.isActive != nil && got.IsActive != *tt.args.isActive {
					t.Errorf("ReservationUseCase.FindReservations() = %v, want %v", got.IsActive, *tt.args.isActive)
				}
			}
		})
	}
}

func TestReservationUseCase_UpdateReservation(t *testing.T) {
	// テスト用のデータを作成
	reservation := models.NewReservation(utils.GenerateRandomString(10), utils.GenerateRandomString(10), utils.GenerateRandomString(10), utils.GenerateRandomBool())
	if err := reservation.Create(models.GetterDB()); err != nil {
		panic(err)
	}

	type fields struct {
		Config config.Config
	}
	type args struct {
		id                  int
		algorithmType       *string
		algorithmTypeDetail *string
		timeConditionRegex  *string
		isActive            *bool
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "UpdateReservation",
			fields: fields{
				Config: UseCaseTestConfig,
			},
			args: args{
				id:                  reservation.ID,
				algorithmType:       stringPointer(utils.GenerateRandomString(10)),
				algorithmTypeDetail: stringPointer(utils.GenerateRandomString(10)),
				timeConditionRegex:  stringPointer(utils.GenerateRandomString(10)),
				isActive:            boolPointer(utils.GenerateRandomBool()),
			},
			wantErr: false,
		},
		{
			name: "UpdateReservation_RequiredID",
			fields: fields{
				Config: UseCaseTestConfig,
			},
			args: args{
				id:                  0,
				algorithmType:       stringPointer(utils.GenerateRandomString(10)),
				algorithmTypeDetail: stringPointer(utils.GenerateRandomString(10)),
				timeConditionRegex:  stringPointer(utils.GenerateRandomString(10)),
				isActive:            boolPointer(utils.GenerateRandomBool()),
			},
			wantErr: true,
		},
		{
			name: "UpdateReservation_RequiredAlgorithmType",
			fields: fields{
				Config: UseCaseTestConfig,
			},
			args: args{
				id:                  reservation.ID,
				algorithmType:       nil,
				algorithmTypeDetail: stringPointer(utils.GenerateRandomString(10)),
				timeConditionRegex:  stringPointer(utils.GenerateRandomString(10)),
				isActive:            boolPointer(utils.GenerateRandomBool()),
			},
			wantErr: true,
		},
		{
			name: "UpdateReservation_RequiredAlgorithmTypeDetail",
			fields: fields{
				Config: UseCaseTestConfig,
			},
			args: args{
				id:                  reservation.ID,
				algorithmType:       stringPointer(utils.GenerateRandomString(10)),
				algorithmTypeDetail: nil,
				timeConditionRegex:  stringPointer(utils.GenerateRandomString(10)),
				isActive:            boolPointer(utils.GenerateRandomBool()),
			},
			wantErr: true,
		},
		{
			name: "UpdateReservation_RequiredTimeConditionRegex",
			fields: fields{
				Config: UseCaseTestConfig,
			},
			args: args{
				id:                  reservation.ID,
				algorithmType:       stringPointer(utils.GenerateRandomString(10)),
				algorithmTypeDetail: stringPointer(utils.GenerateRandomString(10)),
				timeConditionRegex:  nil,
				isActive:            boolPointer(utils.GenerateRandomBool()),
			},
			wantErr: true,
		},
		{
			name: "UpdateReservation_RequiredIsActive",
			fields: fields{
				Config: UseCaseTestConfig,
			},
			args: args{
				id:                  reservation.ID,
				algorithmType:       stringPointer(utils.GenerateRandomString(10)),
				algorithmTypeDetail: stringPointer(utils.GenerateRandomString(10)),
				timeConditionRegex:  stringPointer(utils.GenerateRandomString(10)),
				isActive:            nil,
			},
			wantErr: true,
		},
		{
			name: "UpdateReservation_NotFound",
			fields: fields{
				Config: UseCaseTestConfig,
			},
			args: args{
				id:                  999999,
				algorithmType:       stringPointer(utils.GenerateRandomString(10)),
				algorithmTypeDetail: stringPointer(utils.GenerateRandomString(10)),
				timeConditionRegex:  stringPointer(utils.GenerateRandomString(10)),
				isActive:            boolPointer(utils.GenerateRandomBool()),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ru := NewReservationUseCase(tt.fields.Config)
			if err := ru.UpdateReservation(tt.args.id, tt.args.algorithmType, tt.args.algorithmTypeDetail, tt.args.timeConditionRegex, tt.args.isActive); err != nil {
				if !tt.wantErr {
					t.Errorf("ReservationUseCase.UpdateReservation() error = %v, wantErr %v", err, tt.wantErr)
				}
				return
			}

			got, err := models.FindReservationByID(models.GetterDB(), tt.args.id)
			if err != nil {
				t.Errorf("ReservationUseCase.UpdateReservation() error = %v", err)
			}

			if tt.args.algorithmType != nil && got.AlgorithmType != *tt.args.algorithmType {
				t.Errorf("ReservationUseCase.UpdateReservation() = %v, want %v", got.AlgorithmType, *tt.args.algorithmType)
			}

			if tt.args.algorithmTypeDetail != nil && got.AlgorithmTypeDetail != *tt.args.algorithmTypeDetail {
				t.Errorf("ReservationUseCase.UpdateReservation() = %v, want %v", got.AlgorithmTypeDetail, *tt.args.algorithmTypeDetail)
			}

			if tt.args.timeConditionRegex != nil && got.TimeConditionRegex != *tt.args.timeConditionRegex {
				t.Errorf("ReservationUseCase.UpdateReservation() = %v, want %v", got.TimeConditionRegex, *tt.args.timeConditionRegex)
			}

			if tt.args.isActive != nil && got.IsActive != *tt.args.isActive {
				t.Errorf("ReservationUseCase.UpdateReservation() = %v, want %v", got.IsActive, *tt.args.isActive)
			}
		})
	}
}

func TestReservationUseCase_DeleteReservation(t *testing.T) {
	type fields struct {
		Config config.Config
	}
	type args struct {
		id int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "DeleteReservation",
			fields: fields{
				Config: UseCaseTestConfig,
			},
			args: args{
				id: func() int {
					reservation := models.NewReservation(utils.GenerateRandomString(10), utils.GenerateRandomString(10), utils.GenerateRandomString(10), utils.GenerateRandomBool())
					if err := reservation.Create(models.GetterDB()); err != nil {
						panic(err)
					}
					return reservation.ID
				}(),
			},
			wantErr: false,
		},
		{
			name: "DeleteReservation_RequiredID",
			fields: fields{
				Config: UseCaseTestConfig,
			},
			args: args{
				id: 0,
			},
			wantErr: true,
		},
		{
			name: "DeleteReservation_NotFound",
			fields: fields{
				Config: UseCaseTestConfig,
			},
			args: args{
				id: 999999,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ru := NewReservationUseCase(tt.fields.Config)
			if err := ru.DeleteReservation(tt.args.id); err != nil {
				if !tt.wantErr {
					t.Errorf("ReservationUseCase.DeleteReservation() error = %v, wantErr %v", err, tt.wantErr)
				}
				return
			}

			_, err := models.FindReservationByID(models.GetterDB(), tt.args.id)
			if err == nil {
				t.Errorf("ReservationUseCase.DeleteReservation() error = %v", err)
			}
		})
	}
}

func stringPointer(s string) *string {
	return &s
}

func boolPointer(b bool) *bool {
	return &b
}
