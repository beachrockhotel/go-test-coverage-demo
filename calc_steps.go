package main

import (
	"fmt"
	"github.com/cucumber/godog"
)

var a, b, result int

func givenFirstNumber(arg1 int) error {
	a = arg1
	return nil
}

func andSecondNumber(arg1 int) error {
	b = arg1
	return nil
}

func whenIAddThem() error {
	result = Add(a, b)
	return nil
}

func thenTheResultShouldBe(arg1 int) error {
	if result != arg1 {
		return fmt.Errorf("expected %d but got %d", arg1, result)
	}
	return nil
}

func InitializeScenario(ctx *godog.ScenarioContext) {
	ctx.Step(`^the first number is (\d+)$`, givenFirstNumber)
	ctx.Step(`^the second number is (\d+)$`, andSecondNumber)
	ctx.Step(`^I add the numbers$`, whenIAddThem)
	ctx.Step(`^the result should be (\d+)$`, thenTheResultShouldBe)
}
