package main

import (
	"github.com/dawidd6/ghashverifier/pkg/controller"
	"github.com/dawidd6/ghashverifier/pkg/model"

	"github.com/dawidd6/ghashverifier/pkg/view"
)

const (
	appID = "com.github.dawidd6.ghashverifier"
)

func main() {
	m := model.New()
	v := view.New(appID)
	c := controller.New(v, m)
	c.Run()
}
