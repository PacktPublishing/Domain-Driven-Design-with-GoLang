# GOBDD

[![GoDoc](https://godoc.org/github.com/go-bdd/gobdd?status.svg)](https://godoc.org/github.com/go-bdd/gobdd) [![Coverage Status](https://coveralls.io/repos/github/go-bdd/gobdd/badge.svg?branch=master)](https://coveralls.io/github/go-bdd/gobdd?branch=master)

This is a BDD testing framework. Uses gherkin for the test's syntax. From version 1.0, the API is stable.

## Why did I make the library?

There is godog library for BDD tests in Go. I found this library useful but it runs as an external application which compiles our code. It has several disadvantages:

-   no debugging (breakpoints) in the test. Sometimes it’s useful to go through the whole execution step by step
-   metrics don’t count the test run this way
-   some style checkers recognise tests as dead code
-   it’s impossible to use built-in features like build constraints.
-   no context in steps - so the state have to be stored somewhere else - in my opinion, it makes the maintenance harder

## Quick start

Add the package to your project:

```
go get github.com/go-bdd/gobdd
```

Inside `features` folder create your scenarios. Here is an example:

```gherkin
Feature: math operations
  Scenario: add two digits
    When I add 1 and 2
    Then I the result should equal 3
```

Add a new test `main_test.go`:

```go
func add(t gobdd.StepTest, ctx gobdd.Context, var1, var2 int) {
	res := var1 + var2
	ctx.Set("sumRes", res)
}

func check(t gobdd.StepTest, ctx gobdd.Context, sum int) {
	received, err := ctx.GetInt("sumRes")
	if err != nil {
		t.Error(err)
		return
	}

	if sum != received {
		t.Error(errors.New("the math does not work for you"))
	}
}

func TestScenarios(t *testing.T) {
	suite := NewSuite(t)
	suite.AddStep(`I add (\d+) and (\d+)`, add)
	suite.AddStep(`I the result should equal (\d+)`, check)
	suite.Run()
}
```

and run tests

```
go test ./...
```

More detailed documentation can be found on the docs page: https://go-bdd.github.io/gobdd/. A sample application is available in [a separate repository](https://github.com/go-bdd/sample-app).

# Contributing

All contributions are very much welcome. If you'd like to help with GoBDD development, please see open issues and submit your pull request via GitHub.

# Support

If you didn't find the answer to your question in the documentation, feel free to ask us directly!

Please join us on the `#gobdd-library` channel on the [Gophers slack](https://gophers.slack.com/): You can get [an invite here](https://gophersinvite.herokuapp.com/).
You can find updates about the progress on Twitter: [GoBdd](https://twitter.com/Go_BDD).

You can support my work using [issuehunt](https://issuehunt.io/r/go-bdd) or by [buying me a coffee](https://www.buymeacoffee.com/bklimczak).
