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

func generateSchemaDir(where string, entities []Entity, tmpl *template.Template, ext string, passEnums bool) error {
	buff := bytes.NewBuffer([]byte{})

	_ = os.MkdirAll(where, os.ModePerm)

	for _, ent := range entities {
		if passEnums && len(ent.Values) != 0 || (len(ent.Relations) == 0 && len(ent.Attributes) == 0) { // scalar
			continue
		}

		if err := generateEntSchema(ent, tmpl, buff); err != nil {
			return errors.WithStack(err)
		}

		filename := path.Join(where, strings.ToLower(ent.Name)+ext)

		if err := ioutil.WriteFile(filename, buff.Bytes(), 0644); err != nil {
			return errors.WithStack(err)
		}

		buff.Reset()
	}

	return nil
}

func executeEntGenerator(where string, entities []Entity) error {
	tmpl, err := template.ParseFiles("ent.go.tmpl")
	if err != nil {
		log.Fatal(err)
	}

	return generateSchemaDir(where, entities, tmpl, ".go", true)
}

func executeGraphQLGenerator(where string, entities []Entity) error {
	tmpl, err := template.ParseFiles("gql.graphql.tmpl")
	if err != nil {
		log.Fatal(err)
	}

	return generateSchemaDir(where, entities, tmpl, ".graphql", false)
}
