package auth

import (
	"testing"
	"time"
)

// TestNewUser проверяет корректность создания нового пользователя.
func TestNewUser(t *testing.T) {
	// Определяем имя и пароль для тестового пользователя.
	name := "testuser"
	password := "testpass"

	// Создаем пользователя.
	user := NewUser(name, password)

	// Проверяем, что поля пользователя соответствуют ожидаемым.
	if user.Name != name {
		t.Errorf("expected name %s, got %s", name, user.Name)
	}
	if user.Password != password {
		t.Errorf("expected password %s, got %s", password, user.Password)
	}
	if user.IsAdmin {
		t.Errorf("expected IsAdmin false, got %t", user.IsAdmin)
	}
	if user.AmountOfAgents != 0 {
		t.Errorf("expected AmountOfAgents 0, got %d", user.AmountOfAgents)
	}
	expectedTime := 3 * time.Second
	if user.TimeToCalc != expectedTime {
		t.Errorf("expected TimeToCalc %v, got %v", expectedTime, user.TimeToCalc)
	}
}
