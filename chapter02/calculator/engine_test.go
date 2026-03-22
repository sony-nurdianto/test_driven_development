package calculator_test

import (
	"log"
	"os"
	"testing"

	"github.com/sonynurdianto/test_driven_development/chapter02/calculator"
)

func setup() {
	log.Println("setting up")
}

func teardown() {
	log.Println("Tearing down")
}

func TestMain(m *testing.M) {
	setup()

	e := m.Run()

	teardown()

	os.Exit(e)
}

func TestAdd(t *testing.T) {
	defer func() {
		log.Println("Deffered Tearing Down")
	}()

	// arange
	var e calculator.Engine

	// act
	x, y := 2.5, 3.5
	expect := 6.0

	got := e.Add(x, y)

	// assert
	if got != expect {
		t.Errorf("Add(%2.f + %2.f) incorrect, got %2.f, want: %2.f", 2.5, 3.5, got, expect)
	}
}

func TestAdd_ExampleSubtest(t *testing.T) {
	var e calculator.Engine

	actAssert := func(x, y, want float64) {
		got := e.Add(x, y)

		if got != want {
			t.Errorf("Add(%2.f + %2.f) incorrect, got %2.f, want: %2.f", x, y, got, want)
		}
	}

	t.Run("positive input", func(t *testing.T) {
		x, y := 2.5, 3.5
		want := 6.0
		actAssert(x, y, want)
	})

	t.Run("negative input", func(t *testing.T) {
		x, y := -2.5, -3.5
		want := -6.0
		actAssert(x, y, want)
	})
}

func BenchmarkAdd(b *testing.B) {
	e := calculator.Engine{}
	for b.Loop() {
		e.Add(2, 3)
	}
}
