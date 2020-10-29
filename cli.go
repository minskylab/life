package main

import (
	"path"

	"github.com/pkg/errors"
	"github.com/urfave/cli/v2"
)

func buildCLI() *cli.App {
	app := &cli.App{
		Name:        "society",
		Description: "A utility to scaffold ent + gqlgen projects. Create entities and make a society.",
		Commands: []*cli.Command{
			{
				Name:        "init",
				Description: "Bootstrap a simple society example.",
				Action: func(c *cli.Context) error {
					return nil
				},
			},
			{
				Name:        "generate",
				Description: "Generate ent schema and gqlgen graphql types definitions",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:    "seed",
						Aliases: []string{"s"},
						Usage:   "set the seed source filepath",
						Value:   "seed",
					},
					&cli.StringFlag{
						Name:    "output",
						Aliases: []string{"o"},
						Usage:   "output ent/schema and graphql/schema filepath",
						Value:   ".",
					},
				},
				Action: func(c *cli.Context) error {
					seedPath := c.String("seed")
					outEntSchema := path.Join(c.String("output"), "ent/schema")
					outGraphqlSchema := path.Join(c.String("output"), "graphql/schema")

					entities, err := entitiesFromSeed(seedPath)
					if err != nil {
						return errors.WithStack(err)
					}

					if err := executeEntGenerator(outEntSchema, entities); err != nil {
						return errors.WithStack(err)
					}

					if err := executeGraphQLGenerator(outGraphqlSchema, entities); err != nil {
						return errors.WithStack(err)
					}

					return nil
				},
			},
		},
	}
	return app
}
