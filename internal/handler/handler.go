package handler

import (
	"encoding/json"
	"errors"
	"github.com/tclutin/yandex-calc-service/pkg/calc"
	"github.com/tclutin/yandex-calc-service/pkg/response"
	"log/slog"
	"net/http"
)

type CalcRequest struct {
	Expression string `json:"expression"`
}

type CalcResponse struct {
	Result float64 `json:"result"`
}

type Handler struct {
	logger     *slog.Logger
	calculator *calc.Calc
}

func New(logger *slog.Logger, calc *calc.Calc) *Handler {
	return &Handler{
		logger:     logger,
		calculator: calc,
	}
}

func (h *Handler) Init() *http.ServeMux {
	router := http.NewServeMux()
	router.HandleFunc("/api/v1/calculate", h.Calculate)
	return router
}

func (h *Handler) Calculate(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	defer r.Body.Close()
	w.Header().Set("Content-Type", "application/json")

	var dto CalcRequest
	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		response.SendError(w, http.StatusBadRequest, "Invalid JSON: "+err.Error())
		return
	}

	if dto.Expression == "" {
		response.SendError(w, http.StatusBadRequest, "Missing 'expression' field")
		return
	}

	h.logger.Info("request expression", slog.String("expression", dto.Expression))

	result, err := h.calculator.Calc(dto.Expression)
	if err != nil {
		if errors.Is(err, calc.ErrValidationError) || errors.Is(err, calc.ErrDivisionByZero) {
			response.SendError(w, http.StatusUnprocessableEntity, err.Error())
			return
		}

		response.SendError(w, http.StatusInternalServerError, "Internal server error")
		return
	}

	h.logger.Info("request expression", slog.Float64("result", result))

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(CalcResponse{Result: result})
}
