package utils

import (
	"reflect"

	"github.com/gotk3/gotk3/gtk"
)

func BindWidgets(v interface{}, uiResourcePath string) {
	// Load UI from resources
	builder, err := gtk.BuilderNewFromResource(uiResourcePath)
	if err != nil {
		panic(err)
	}

	// Get widgets from UI definition
	vStruct := reflect.ValueOf(v).Elem()
	for i := 0; i < vStruct.NumField(); i++ {
		field := vStruct.Field(i)
		structField := vStruct.Type().Field(i)
		widget := structField.Tag.Get("gtk")

		if widget == "" {
			continue
		}

		obj, err := builder.GetObject(widget)
		if err != nil {
			panic(err)
		}

		field.Set(reflect.ValueOf(obj).Convert(field.Type()))
	}
}
