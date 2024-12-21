package handler

import (
	"bytes"
	"github.com/tclutin/yandex-calc-service/pkg/calc"
	"github.com/tclutin/yandex-calc-service/pkg/logger"
	"net/http"
	"net/http/httptest"
	"testing"
)

// TODO: желательно замокать калькулятор
func TestHandler_Success(t *testing.T) {
	data := `{"expression": "2+2"}`

	h := New(logger.New(), calc.New())

	req := httptest.NewRequest("POST", "/api/v1/calculate", bytes.NewBufferString(data))
	req.Header.Set("Content-Type", "application/json")

	recorder := httptest.NewRecorder()

	h.Init().ServeHTTP(recorder, req)

	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("expected status 200, got %v", status)
	}
}

func TestHandler_MethodNotAllowed(t *testing.T) {
	data := `{"expression": "2+2"}`

	h := New(logger.New(), calc.New())

	req := httptest.NewRequest("GET", "/api/v1/calculate", bytes.NewBufferString(data))
	req.Header.Set("Content-Type", "application/json")

	recorder := httptest.NewRecorder()

	h.Init().ServeHTTP(recorder, req)

	if status := recorder.Code; status != http.StatusMethodNotAllowed {
		t.Errorf("expected status 405, got %v", status)
	}
}

func TestHandler_ValidJsonBody(t *testing.T) {
	data := `{"expression":}`

	h := New(logger.New(), calc.New())

	req := httptest.NewRequest("POST", "/api/v1/calculate", bytes.NewBufferString(data))
	req.Header.Set("Content-Type", "application/json")

	recorder := httptest.NewRecorder()

	h.Init().ServeHTTP(recorder, req)

	if status := recorder.Code; status != http.StatusBadRequest {
		t.Errorf("expected status 400, got %v", status)
	}
}

func TestHandler_RequiredJsonField(t *testing.T) {
	data := `{"expressin": "2+2"}`

	h := New(logger.New(), calc.New())

	req := httptest.NewRequest("POST", "/api/v1/calculate", bytes.NewBufferString(data))
	req.Header.Set("Content-Type", "application/json")

	recorder := httptest.NewRecorder()

	h.Init().ServeHTTP(recorder, req)

	if status := recorder.Code; status != http.StatusBadRequest {
		t.Errorf("expected status 400, got %v", status)
	}
}
