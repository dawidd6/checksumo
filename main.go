package main

import (
	"os"

	"github.com/dawidd6/checksumo/pkg/controller"
	"github.com/dawidd6/checksumo/pkg/model"
	"github.com/dawidd6/checksumo/pkg/view"
)

// #cgo pkg-config: gio-2.0
// #include "resources.h"
import "C"

const AppID = "com.github.dawidd6.checksumo"

func main() {
	m := model.New()
	v := view.New(AppID)
	c := controller.New(v, m)

	os.Exit(c.View.Application.Run(os.Args))
}
