// This project is created as the alternative to https://github.com/DATA-DOG/godog and is inspirited by it.
// There are a few differences between both solutions:
//
// - GoBDD uses the built-in testing framework
//
// - GoBDD is run as standard test case (not a separate program)
//
// - you can use every Go native feature like build tags, pprof and so on
//
// - the context in every test case contains all the required information to run (values passed from previous steps).
// More information can be found in the readme file https://github.com/go-bdd/gobdd/blob/master/README.md
package gobdd
