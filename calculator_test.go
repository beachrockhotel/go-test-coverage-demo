package main

import (
	"os"
	"testing"

	"github.com/cucumber/godog"
)

func TestAdd(t *testing.T) {
	if Add(2, 3) != 5 {
		t.Error("Expected 5")
	}
}

func TestMain(m *testing.M) {
	opts := godog.Options{
		Format: "pretty",
		Paths:  []string{"features"},
	}

	status := godog.TestSuite{
		Name:                "calculator",
		ScenarioInitializer: InitializeScenario,
		Options:             &opts,
	}.Run()

	if st := m.Run(); st > status {
		status = st
	}
	os.Exit(status)
}
