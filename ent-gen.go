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

type genTypeExtension string

const graphqlGen genTypeExtension = ".graphql"
const entGen genTypeExtension = ".go"

func generateEntSchema(entity Entity, tmpl *template.Template, out io.Writer) error {
	return tmpl.Execute(out, entity)
}

func generateSchemaDir(where string, entities []Entity, tmpl *template.Template, ext genTypeExtension) error {
	buff := bytes.NewBuffer([]byte{})

	_ = os.MkdirAll(where, os.ModePerm)

	for _, ent := range entities {
		if entityKind(ent) == EnumKind {

		}

		if ext == entGen && (len(ent.Values) != 0 || (len(ent.Relations) == 0 && len(ent.Attributes) == 0)) { // scalar
			continue
		}

		if err := generateEntSchema(ent, tmpl, buff); err != nil {
			return errors.WithStack(err)
		}

		filename := path.Join(where, strings.ToLower(ent.Name)+string(ext))

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

	return generateSchemaDir(where, entities, tmpl, entGen)
}

func executeGraphQLGenerator(where string, entities []Entity) error {
	tmpl, err := template.ParseFiles("gql.graphql.tmpl")
	if err != nil {
		log.Fatal(err)
	}

	return generateSchemaDir(where, entities, tmpl, graphqlGen)
}
