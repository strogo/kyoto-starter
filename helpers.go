package main

import (
	"html/template"

	"github.com/yuriizinets/kyoto"
)

func mktemplate(page string) *template.Template {
	t := template.New(page)
	t = t.Funcs(kyoto.TFuncMap())
	t, _ = t.ParseGlob("*.html")
	t, _ = t.ParseGlob("uikit/twui/*.html")
	return t
}
