package main

import (
	"os/exec"
	"strconv"
	"testing"
	"time"

	"github.com/bitcoin-trading-system/bitcoin-algorithm-reservation/config"
	"github.com/bitcoin-trading-system/bitcoin-algorithm-reservation/models"
	"github.com/bitcoin-trading-system/bitcoin-algorithm-reservation/utils"
)

func TestMainFunction(t *testing.T) {
	// main関数をゴルーチンで実行
	go main()

	config := config.NewConfig("toml/local.toml", "env/.env.local")
	models.Init(config)

	type args struct {
		curl []string
	}

	tests := []struct {
		name string
		args args
	}{
		{
			name: "health",
			args: args{
				curl: []string{"curl", "-w", "%{http_code}", "-o", "/dev/null", "-s", "http://localhost:8003/health"},
			},
		},
		{
			name: "post_reservation",
			args: args{
				curl: []string{"curl", "-w", "%{http_code}", "-o", "/dev/null", "-s", "-X", "POST", "-H", "Content-Type: application/json", "-d", `{"algorithm_type": "type1", "algorithm_type_detail": "detail1", "time_condition_regex": "regex1", "is_active": true}`, "http://localhost:8003/reservations"},
			},
		},
		{
			name: "get_reservation_by_id",
			args: args{
				curl: func() []string {
					r := models.NewReservation(utils.GenerateRandomString(10), utils.GenerateRandomString(10), utils.GenerateRandomString(10), utils.GenerateRandomBool())
					if err := r.Create(models.GetterDB()); err != nil {
						panic(err)
					}

					reservations, err := models.FindAllReservations(models.GetterDB())
					if err != nil {
						panic(err)
					}
					return []string{"curl", "-w", "%{http_code}", "-o", "/dev/null", "-s", "http://localhost:8003/reservations/" + strconv.Itoa(reservations[0].ID)}
				}(),
			},
		},
		{
			name: "get_reservations",
			args: args{
				curl: []string{"curl", "-w", "%{http_code}", "-o", "/dev/null", "-s", "http://localhost:8003/reservations?is_active=true"},
			},
		},
		{
			name: "update_reservation",
			args: args{
				curl: func() []string {
					reservations, err := models.FindAllReservations(models.GetterDB())
					if err != nil {
						panic(err)
					}
					return []string{"curl", "-w", "%{http_code}", "-o", "/dev/null", "-s", "-X", "PUT", "-H", "Content-Type: application/json", "-d", `{"algorithm_type": "type1", "algorithm_type_detail": "detail1", "time_condition_regex": "regex1", "is_active": true}`, "http://localhost:8003/reservations/" + strconv.Itoa(reservations[0].ID)}
				}(),
			},
		},
		{
			name: "delete_reservation",
			args: args{
				curl: func() []string {
					reservations, err := models.FindAllReservations(models.GetterDB())
					if err != nil {
						panic(err)
					}
					return []string{"curl", "-w", "%{http_code}", "-o", "/dev/null", "-s", "-X", "DELETE", "http://localhost:8003/reservations/" + strconv.Itoa(reservations[0].ID)}
				}(),
			},
		},
	}

	// サーバーが起動するのを待つ
	time.Sleep(1 * time.Second)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// curlコマンドを実行
			cmd := exec.Command(tt.args.curl[0], tt.args.curl[1:]...)
			output, err := cmd.CombinedOutput()
			if err != nil {
				t.Fatalf("Failed to execute curl command: %v", err)
			}

			// ステータスコードを確認
			statusCode := string(output[len(output)-3:]) // 最後の3文字がステータスコード
			if statusCode != "200" && statusCode != "201" {
				t.Fatalf("Server returned non-200/201 status code: %s", statusCode)
			}
		})
	}
}
