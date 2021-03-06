package dics

import (
	{{.ModulePluralLowercase}} "{{.PackageName}}/{{.ModulePluralLowercase}}"
	models "{{.PackageName}}/{{.ModulePluralLowercase}}/models"
	validations "{{.PackageName}}/{{.ModulePluralLowercase}}/validations"
	"github.com/sarulabs/dingo/v4"
)

var {{.Module}} = []dingo.Def{
	{
		Name:  "module:{{.ModuleLowercase}}:model",
		Build: (*models.{{.Module}})(nil),
        Params: dingo.Params{
			"Model": dingo.Service("bima:model"),
		},
	},
	{
		Name:  "module:{{.ModuleLowercase}}:validation",
		Build: (*validations.{{.Module}})(nil),
	},
	{
		Name:  "module:{{.ModuleLowercase}}",
		Build: (*{{.ModulePluralLowercase}}.Module)(nil),
		Params: dingo.Params{
			"Module":    dingo.Service("bima:module"),
			"Validator": dingo.Service("module:{{.ModuleLowercase}}:validation"),
		},
	},
	{
		Name:  "module:{{.ModuleLowercase}}:server",
		Build: (*{{.ModulePluralLowercase}}.Server)(nil),
		Params: dingo.Params{
			"Server": dingo.Service("bima:server"),
			"Module": dingo.Service("module:{{.ModuleLowercase}}"),
		},
	},
}
