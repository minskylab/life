package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"gopkg.in/yaml.v3"
)

func loadValidFiles(seedPath string) []string {
	validFiles := []string{}
	err := filepath.Walk(seedPath, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}

		extension := filepath.Ext(info.Name())

		if extension != ".yaml" && extension != ".yml" {
			return nil
		}

		validFiles = append(validFiles, path)

		return nil
	})
	if err != nil {
		log.Fatal(err)
	}

	return validFiles
}

func getFilesContents(files ...string) map[string][]byte {
	var wg sync.WaitGroup
	var m sync.Mutex

	filesLength := len(files)
	contents := make(map[string][]byte, filesLength)
	wg.Add(filesLength)

	for _, file := range files {
		go func(file string) {
			content, err := ioutil.ReadFile(file)

			if err != nil {
				log.Fatal(err)
			}

			m.Lock()
			contents[file] = content
			m.Unlock()
			wg.Done()
		}(file)
	}

	wg.Wait()

	return contents
}

func partToEnts(data []byte, parts [][]byte) [][]byte {
	if parts == nil {
		parts = [][]byte{}
	}

	cutoff := bytes.Index(data, []byte("---"))

	if cutoff >= 0 {
		parts = append(parts, data[0:cutoff])
		return partToEnts(data[cutoff+3:], parts)
	}

	parts = append(parts, data)
	return parts
}

// EntityProp ...
type EntityProp struct {
	Name      string   `yaml:"name"`
	Type      string   `yaml:"type"`
	Required  bool     `yaml:"required"`
	Unique    bool     `yaml:"unique"`
	Immutable bool     `yaml:"immutable"`
	Many      bool     `yaml:"many"`
	Ref       string   `yaml:"ref"`
	Values    []string `yaml:"values"` // useful for ent
}

// RawEntity ...
type RawEntity struct {
	Name       string                 `yaml:"name"`
	Type       string                 `yaml:"type"`
	Attributes map[string]interface{} `yaml:"attributes"`
	Relations  map[string]interface{} `yaml:"relations"`
	Values     []string               `yaml:"values"`
}

// Alias ...
type Alias struct {
	Type      string `yaml:"type"`
	Required  bool   `yaml:"required"`
	Unique    bool   `yaml:"unique"`
	Immutable bool   `yaml:"immutable"`
	Many      bool   `yaml:"many"`
	Ref       string `yaml:"ref"`
}

// Entity ...
type Entity struct {
	Name       string
	Attributes []EntityProp
	Relations  []EntityProp
	Values     []string
}

func parseProp(name string, prop interface{}, alias map[string]Alias) EntityProp {
	switch prop.(type) {
	case string:
		fAttr := EntityProp{}
		sAttr := prop.(string)
		// TODO Here
		//

		fAttr.Name = name

		al, exists := alias[strings.TrimRight(sAttr, "[]! \n")]
		if exists {
			fAttr.Type = strings.Trim(al.Type, "[]! \n")
			fAttr.Required = al.Required
			fAttr.Many = al.Many
			fAttr.Ref = al.Ref
			fAttr.Immutable = al.Immutable
			fAttr.Unique = al.Unique
		} else {
			fAttr.Type = strings.Trim(sAttr, "[]! \n")

			if strings.HasSuffix(sAttr, "!") {
				fAttr.Required = true
				sAttr = strings.TrimRight(sAttr, "! \n")
			}

			if strings.HasPrefix(sAttr, "[") && strings.HasSuffix(sAttr, "]") {
				fAttr.Many = true
			}
		}

		return fAttr
	case map[string]interface{}:
		fAttr := EntityProp{}
		sAttr := prop.(map[string]interface{})

		fAttr.Name = name

		if attrType, ok := sAttr["type"].(string); ok {
			fAttr.Type = strings.Trim(attrType, "[]! \n")

			if strings.HasSuffix(attrType, "!") {
				fAttr.Required = true
				attrType = strings.TrimRight(attrType, "! \n")
			}

			if strings.HasPrefix(attrType, "[") && strings.HasSuffix(attrType, "]") {
				fAttr.Many = true
			}
		}

		if isReq, ok := sAttr["required"].(bool); ok {
			fAttr.Required = isReq
		}

		if isMany, ok := sAttr["many"].(bool); ok {
			fAttr.Many = isMany
		}

		if relation, ok := sAttr["ref"].(string); ok {
			fAttr.Ref = relation
		}

		return fAttr
	default:
		log.Fatal("invalid attribute type")
	}

	return EntityProp{}
}

type kind string

// EnumKind ...
const EnumKind kind = "enum"

// EntityKind ...
const EntityKind kind = "ent"

// AliasKind ...
const AliasKind kind = "alias"

func entKind(ent RawEntity) kind {
	if (len(ent.Attributes) == 0 && len(ent.Relations) == 0) && len(ent.Values) != 0 {
		return EnumKind
	}

	if ent.Name == "" {
		return AliasKind
	}

	return EntityKind
}

func main() {
	seedPath := "seed"
	files := loadValidFiles(seedPath)
	contents := getFilesContents(files...)

	entities := map[string][]RawEntity{}
	alias := map[string]Alias{}

	for f, data := range contents {
		parts := partToEnts(data, nil)
		var ents []RawEntity

		for _, p := range parts {
			var e RawEntity
			if err := yaml.Unmarshal(p, &e); err != nil {
				log.Fatal(err)
			}

			if e.Name == "" { // isn't an entity
				newAlias := map[string]Alias{}
				if err := yaml.Unmarshal(p, &newAlias); err != nil {
					log.Fatal(err)
				}

				for k, v := range newAlias {
					alias[k] = v
				}

				continue
			}

			ents = append(ents, e)
		}
		entities[f] = ents

	}

	finalEntities := make([]Entity, 0)

	enums := map[string][]string{}

	for file, ents := range entities {
		for _, ent := range ents {
			fmt.Printf("from: %s | [%s] T: %s, A: %d, R: %d\n", file, entKind(ent), ent.Name, len(ent.Attributes), len(ent.Relations))

			if len(ent.Values) != 0 {
				enums[ent.Name] = ent.Values
			}

			attributes := make([]EntityProp, 0)
			relations := make([]EntityProp, 0)

			for name, attr := range ent.Attributes {
				p := parseProp(name, attr, alias)

				values, isEnum := enums[p.Type]

				if isEnum {
					p.Values = values
					continue
				}

				attributes = append(attributes, p)
			}

			for name, rel := range ent.Relations {
				relations = append(relations, parseProp(name, rel, alias))
			}

			finalEntities = append(finalEntities, Entity{
				Name:       ent.Name,
				Attributes: attributes,
				Relations:  relations,
				Values:     ent.Values,
			})
		}
	}

	// spew.Dump(finalEntities)

	if err := executeEntGenerator("ent/schema", finalEntities); err != nil {
		log.Fatal(err)
	}

	if err := executeGraphQLGenerator("graphql/schema", finalEntities); err != nil {
		log.Fatal(err)
	}
}
