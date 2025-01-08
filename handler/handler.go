package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/bitcoin-trading-system/bitcoin-algorithm-reservation/config"
	"github.com/bitcoin-trading-system/bitcoin-algorithm-reservation/usecase"
)

type handler struct {
	ReservationUseCase usecase.IReservationUseCase
	Config             config.Config
}

type IHandler interface {
	PostReservation(ctx *gin.Context)
	GetReservationByID(ctx *gin.Context)
	GetReservations(ctx *gin.Context)
	UpdateReservation(ctx *gin.Context)
	DeleteReservation(ctx *gin.Context)
}

func NewHandler(config config.Config) IHandler {
	reservationUseCase := usecase.NewReservationUseCase(config)

	return &handler{
		ReservationUseCase: reservationUseCase,
		Config:             config,
	}
}

func (h *handler) PostReservation(ctx *gin.Context) {
	type PostReservationRequestBody struct {
		AlgorithmType       string `json:"algorithm_type" binding:"required"`
		AlgorithmTypeDetail string `json:"algorithm_type_detail" binding:"required"`
		TimeConditionRegex  string `json:"time_condition_regex" binding:"required"`
		IsActive            *bool  `json:"is_active" binding:"required"`
	}

	parsePostReservationRequestBody := func(ctx *gin.Context) (PostReservationRequestBody, error) {
		var req PostReservationRequestBody
		err := ctx.ShouldBindJSON(&req)
		return req, err
	}

	req, err := parsePostReservationRequestBody(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.ReservationUseCase.CreateReservation(req.AlgorithmType, req.AlgorithmTypeDetail, req.TimeConditionRegex, req.IsActive); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	ctx.JSON(http.StatusCreated, gin.H{"status": "ok"})
}

func (h *handler) GetReservationByID(ctx *gin.Context) {
	rID := ctx.Param("id")
	rIDInt, err := strconv.Atoi(rID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	reservations, err := h.ReservationUseCase.FindReservationByID(rIDInt)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"reservations": reservations})
}

func (h *handler) GetReservations(ctx *gin.Context) {
	type GetReservationsQuery struct {
		AlgorithmType       *string `form:"algorithm_type"`
		AlgorithmTypeDetail *string `form:"algorithm_type_detail"`
		TimeConditionRegex  *string `form:"time_condition_regex"`
		IsActive            *bool   `form:"is_active"`
	}

	parseGetReservationsQuery := func(ctx *gin.Context) (GetReservationsQuery, error) {
		var query GetReservationsQuery
		err := ctx.ShouldBindQuery(&query)
		return query, err
	}

	query, err := parseGetReservationsQuery(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	reservations, err := h.ReservationUseCase.FindReservations(query.AlgorithmType, query.AlgorithmTypeDetail, query.TimeConditionRegex, query.IsActive)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, reservations)
}

func (h *handler) UpdateReservation(ctx *gin.Context) {
	type UpdateReservationRequestBody struct {
		AlgorithmType       *string `json:"algorithm_type"`
		AlgorithmTypeDetail *string `json:"algorithm_type_detail"`
		TimeConditionRegex  *string `json:"time_condition_regex"`
		IsActive            *bool   `json:"is_active"`
	}

	parseUpdateReservationRequestBody := func(ctx *gin.Context) (UpdateReservationRequestBody, error) {
		var req UpdateReservationRequestBody
		err := ctx.ShouldBindJSON(&req)
		return req, err
	}

	rID := ctx.Param("id")
	rIDInt, err := strconv.Atoi(rID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	req, err := parseUpdateReservationRequestBody(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.ReservationUseCase.UpdateReservation(rIDInt, req.AlgorithmType, req.AlgorithmTypeDetail, req.TimeConditionRegex, req.IsActive); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "ok"})
}

func (h *handler) DeleteReservation(ctx *gin.Context) {
	rID := ctx.Param("id")
	rIDInt, err := strconv.Atoi(rID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.ReservationUseCase.DeleteReservation(rIDInt); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "ok"})
}
