/*
Copyright © 2022 Leonardo Biffi <leonardobiffi@outlook.com>
*/
package template

import (
	"bytes"
	"text/template"
)

var TemplateDockerfile = `FROM {{ .Image }}`

type Dockerfile struct {
	Image string
}

func File(dockerfile Dockerfile) (file string, err error) {
	tmpl, err := template.New("Dockerfile").Parse(TemplateDockerfile)
	if err != nil {
		return
	}

	var content bytes.Buffer
	if err = tmpl.Execute(&content, dockerfile); err != nil {
		return
	}

	return content.String(), nil
}
