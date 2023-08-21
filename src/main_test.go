package main

import (
	"os"
	"testing"

	"github.com/deltachat/deltachat-rpc-client-go/deltachat"
)

var acfactory *deltachat.AcFactory

func TestMain(m *testing.M) {
	acfactory = &deltachat.AcFactory{Debug: os.Getenv("TEST_DEBUG") == "1"}
	acfactory.TearUp()
	defer acfactory.TearDown()
	m.Run()
}
