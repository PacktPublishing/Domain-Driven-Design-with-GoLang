package chapter8

import (
	"testing"

	"github.com/go-bdd/gobdd"
)

func add(t gobdd.StepTest, ctx gobdd.Context, first, second int) {
	res := first + second
	ctx.Set("result", res)
}

func check(t gobdd.StepTest, ctx gobdd.Context, sum int) {
	received, err := ctx.GetInt("result")
	if err != nil {
		t.Fatal(err)
		return
	}
	if sum != received {
		t.Fatalf("expected %d but got %d", sum, received)
	}
}

func TestScenarios(t *testing.T) {
	suite := gobdd.NewSuite(t)
	suite.AddStep(`I add (\d+) and (\d+)`, add)
	suite.AddStep(`the result should equal (\d+)`, check)
	suite.Run()
}
