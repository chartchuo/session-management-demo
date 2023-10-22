package token

import (
	"log"
	"os"
	"testing"
	"time"
)

type mockTime struct {
	mockNow time.Time
}

func newMockTime() *mockTime {
	return &mockTime{mockNow: time.Now()}
}
func (m *mockTime) now() time.Time {
	return m.mockNow
}
func (m *mockTime) forward(d time.Duration) {
	m.mockNow = m.mockNow.Add(d)
}

var mock = newMockTime()

func TestMain(m *testing.M) {
	// replease time now function

	now = mock.now
	exitVal := m.Run()

	os.Exit(exitVal)
}

func TestMock(t *testing.T) {
	var start = now()

	mock.forward(refreshExp)
	var end = now()
	if !end.After(start) {
		log.Printf("start: %v\n", start)
		log.Printf("end: %v\n", end)
		t.Errorf("fastforward error")
	}
}
