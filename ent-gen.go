package main

import (
	"bytes"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path"
	"strings"
	"text/template"

	"github.com/pkg/errors"
)

func generateEntSchema(entity Entity, tmpl *template.Template, out io.Writer) error {
	return tmpl.Execute(out, entity)
}

func generateEntSchemaDir(where string, entities []Entity, tmpl *template.Template) error {
	buff := bytes.NewBuffer([]byte{})

	_ = os.MkdirAll(where, os.ModeDir)

	for _, ent := range entities {
		if err := generateEntSchema(ent, tmpl, buff); err != nil {
			return errors.WithStack(err)
		}

		filename := path.Join(where, strings.ToLower(ent.Name))

		if err := ioutil.WriteFile(filename, buff.Bytes(), 0644); err != nil {
			return errors.WithStack(err)
		}

		buff.Reset()
	}

	return nil
}

func executeEntGenerator(where string, entities []Entity) error {
	tmpl, err := template.New("ent.go").ParseFiles("ent.go.tmpl")
	if err != nil {
		log.Fatal(err)
	}

	return generateEntSchemaDir(where, entities, tmpl)
}
