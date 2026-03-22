package format_test

import (
	"testing"

	"github.com/sonynurdianto/test_driven_development/chapter02/format"
)

func TestReuslt(t *testing.T) {
	res := format.Result("", 100_000_000.99)

	expectedRes := "100.000.000.99"

	if res == "" {
		t.Fatalf("expected res is not empty, instead got %q ", "")
	}

	if res != expectedRes {
		t.Fatalf("expceted Res is %s instead got %s", expectedRes, res)
	}
}
