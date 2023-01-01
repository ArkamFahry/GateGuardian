package env

import "testing"

func TestInitRequiredEnv(t *testing.T) {
	got := InitRequiredEnv()

	if got == nil {
		t.Errorf("Error initializing required envs")
	}
}
