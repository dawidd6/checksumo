package main

import (
	"github.com/dawidd6/checksumo/pkg/controller"
	"github.com/dawidd6/checksumo/pkg/model"

	"github.com/dawidd6/checksumo/pkg/view"
)

var (
	appID   = "com.github.dawidd6.checksumo"
	version string
)

func main() {
	m := model.New()
	v := view.New(appID, version)
	c := controller.New(v, m)
	c.Run()
}
