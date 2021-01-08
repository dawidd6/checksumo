package main

import (
	"os"

	"github.com/dawidd6/checksumo/pkg/controller"
	"github.com/dawidd6/checksumo/pkg/model"
	"github.com/dawidd6/checksumo/pkg/view"
)

var (
	AppID   = "com.github.dawidd6.checksumo"
	Version string
)

func main() {
	m := model.New()
	v := view.New(AppID, Version)
	c := controller.New(v, m)

	os.Exit(c.View.Application.Run(os.Args))
}
