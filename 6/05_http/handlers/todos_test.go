package handlers

import (
	"net/http/httptest"
	"net/http"
	"testing"
	"strings"
)

func TestGetTodos(t *testing.T) {

	// Инициализируем хендлеры со своими данными
	h := Handler{
		Todos: &[]Todo{
			{Name: "Test", Done: false},
		},
	}

	handler := http.HandlerFunc(h.HandleTodos)

	r, err := http.NewRequest("GET", "/", strings.NewReader(""))
	if err != nil {
		t.Errorf("error: %v", err)
	}

	w := httptest.NewRecorder()

	handler.ServeHTTP(w, r)

	if w.Code != http.StatusOK {
		t.Errorf("Wrong code. Expected %d, got %d", http.StatusOK, w.Code)
	}

	expected := `[{"name":"Test","done":false}]`

	if w.Body.String() != expected {
		t.Errorf(`expected %s, got %s`, expected, w.Body.String())
	}
}