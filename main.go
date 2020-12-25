package main

import (
	"github.com/dawidd6/checksumo/pkg/controller"
	"github.com/dawidd6/checksumo/pkg/model"

	"github.com/dawidd6/checksumo/pkg/view"
)

const (
	appID = "com.github.dawidd6.checksumo"
)

func main() {
	m := model.New()
	v := view.New(appID)
	c := controller.New(v, m)
	c.Run()
}
