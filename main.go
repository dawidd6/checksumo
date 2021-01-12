package main

// #cgo pkg-config: gio-2.0
// #include "resources.h"
import "C"

import (
	"os"

	"github.com/dawidd6/checksumo/pkg/controller"
	"github.com/dawidd6/checksumo/pkg/model"
	"github.com/dawidd6/checksumo/pkg/view"
)

func main() {
	m := model.New()
	v := view.New()
	c := controller.New(v, m)

	os.Exit(c.View.Application.Run(os.Args))
}
