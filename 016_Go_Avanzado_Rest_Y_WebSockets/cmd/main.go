package main

import (
	"github.com/jsovalles/rest-ws/internal/bootstrap"
	"go.uber.org/fx"
)

func main() {
	fx.New(bootstrap.Module).Run()
}
