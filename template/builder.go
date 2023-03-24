package template

import (
	_ "embed"
	"os"
	"text/template"

	"github.com/samuelsih/sucms/config"
	"github.com/samuelsih/sucms/pkg"
)

//go:embed one.tmpl
var file string

func Build(c config.Extracted) {
	tmplDatas := Generate(c)

	t, err := template.New("one_tmpl").Parse(file)
	if err != nil {
		pkg.LogFail(err.Error())
	}

	for _, tdata := range tmplDatas {
		f, err := os.OpenFile(tdata.Filename, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0600)
		if err != nil {
			pkg.LogFail(err.Error())
		}

		defer f.Close()

		t.Execute(f, tdata)
	}
}
