package template

import (
	"bytes"
	"html/template"
)

type Template struct {
	Policy string
}

func Concat(policy, configFile string) []byte {
	tpl := template.Must(template.ParseFiles(configFile))

	p := Template{Policy: policy}

	output := bytes.Buffer{}

	if err := tpl.Execute(&output, p); err != nil {
		panic(err)
	}

	return output.Bytes()
}
