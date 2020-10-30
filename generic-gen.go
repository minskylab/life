package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path"
	"strings"
	"text/template"

	"github.com/pkg/errors"
)

type genTypeExtension string

const graphqlGen genTypeExtension = ".graphqls"
const entGen genTypeExtension = ".go"

func generateTemplate(entity Entity, tmpl *template.Template, out io.Writer) error {
	return tmpl.Execute(out, entity)
}

func generateSchemaDir(where string, entities []Entity, tmpl *template.Template, ext genTypeExtension) error {
	buff := bytes.NewBuffer([]byte{})

	_ = os.MkdirAll(where, os.ModePerm)

	scalarsPath := path.Join(where, "scalars"+string(ext))
	var scalarsFile *os.File

	for _, ent := range entities {
		kind := entityKind(ent)

		if kind == ScalarKind {
			if ext == graphqlGen {
				if scalarsFile == nil {
					if fileExists(scalarsPath) {
						_ = os.Remove(scalarsPath)
					}
					scalarsFile, _ = os.OpenFile(scalarsPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
					defer scalarsFile.Close()
				}

				if _, err := scalarsFile.WriteString(fmt.Sprintf("scalar %s\n", ent.Name)); err != nil {
					return errors.WithStack(err)
				}
			}
			continue
		}

		if ext == entGen && kind == EnumKind {
			continue
		}

		if err := generateTemplate(ent, tmpl, buff); err != nil {
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
		return errors.WithStack(err)
	}

	return generateSchemaDir(where, entities, tmpl, entGen)
}

func executeGraphQLGenerator(where string, entities []Entity) error {
	tmpl, err := template.ParseFiles("gql.graphqls.tmpl")
	if err != nil {
		return errors.WithStack(err)
	}

	if err := generateSchemaDir(where, entities, tmpl, graphqlGen); err != nil {
		return errors.WithStack(err)
	}

	return nil
}
