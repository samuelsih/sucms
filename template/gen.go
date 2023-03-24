package template

import (
	"strings"

	"github.com/samuelsih/sucms/config"
	"github.com/samuelsih/sucms/pkg"
)

type TmplData struct {
	Name  string
	Filename string
	Table []string
}

func Generate(c config.Extracted) []TmplData {
	var lists []TmplData

	for table, schemas := range c.SchemasDefined {
		var tmplData TmplData

		name := strings.Split(table, "|")
		tmplData.Name = name[0] 
		tmplData.Filename = name[1]

		for columnName, columnType := range schemas {
			switch columnType {
			case "id":
				tmplData.Table = append(tmplData.Table, id(columnName))
			case "string":
				tmplData.Table = append(tmplData.Table, str(columnName))
			case "integer":
				tmplData.Table = append(tmplData.Table, integer(columnName))
			case "text":
				tmplData.Table = append(tmplData.Table, text(columnName))
			case "boolean":
				tmplData.Table = append(tmplData.Table, boolean(columnName))
			case "uuid":
				tmplData.Table = append(tmplData.Table, uuid(columnName))
			case "date":
				tmplData.Table = append(tmplData.Table, date(columnName))
			case "bigInteger":
				tmplData.Table = append(tmplData.Table, bigInteger(columnName))
			case "bigIncrements":
				tmplData.Table = append(tmplData.Table, bigIncrements(columnName))
			default:
				pkg.LogFail("Unknown type " + columnType)
			}
		} 

		lists = append(lists, tmplData)
	}

	return lists
}