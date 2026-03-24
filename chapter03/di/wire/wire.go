//go:build wireinject

package main

import "github.com/google/wire"

func InitCalc() *Calculator {
	wire.Build(wire.NewSet(NewEngine, wire.Bind(new(Adder), new(*Engine)), NewCalculator))
	return nil
}
