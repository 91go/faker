package faker

import (
	"reflect"
)

type Generator struct {
	Locale_ string
	Pkg     string
}

func NewGenerator(localVar ...string) Generator {
	newGenerator := Generator{}
	if len(localVar) > 0 {
		newGenerator.Locale_ = localVar[0]
	} else {
		newGenerator.Locale_ = "zh_CN"
	}
	newGenerator.Pkg = reflect.TypeOf(newGenerator).PkgPath()
	return newGenerator
}

func (g *Generator) SetLanguage(localVar string) {
	g.Locale_ = localVar
}

func (g *Generator) Generics(intended string) string {
	return ""
}
