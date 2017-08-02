package main

import (
	"fmt"
	"os"

	"github.com/dlespiau/dax"
	"github.com/urfave/cli"
)

type examples struct {
	list []*Example
}

func (e *examples) findExampleByID(id string) (*Example, error) {
	for _, example := range e.list {
		if example.ID() == id {
			return example, nil
		}
	}
	return nil, fmt.Errorf("no example widh id '%s'", id)
}

func getExamples(ctx *cli.Context) *examples {
	return ctx.App.Metadata["examples"].(*examples)
}

func list(ctx *cli.Context) error {
	examples := getExamples(ctx)
	for _, example := range examples.list {
		fmt.Printf("%s: %s\n", example.ID(), example.Description)
	}
	return nil
}

func run(ctx *cli.Context) error {
	if ctx.NArg() != 1 {
		return fmt.Errorf("run: need the name of the example to run")
	}
	name := ctx.Args().First()

	examples := getExamples(ctx)
	example, err := examples.findExampleByID(name)
	if err != nil {
		return err
	}

	app := dax.NewApplication(example.Name)
	window := app.CreateWindow(app.Name+" Example", 800, 600)
	window.SetScene(example.Scene)
	app.Run()

	return nil
}

var daxExamples = &examples{
	list: []*Example{
		&gfxPolylineExample,
		&gfxScenegraphExample,
		&winsysEventsExample,
	},
}

func main() {
	app := cli.NewApp()
	app.Name = "dax-examples"
	app.Usage = "Show off what DaX can do"
	app.Commands = []cli.Command{
		{
			Name:      "run",
			Usage:     "Run an example",
			ArgsUsage: "<example name>",
			Action:    run,
		},
		{
			Name:   "list",
			Usage:  "Print the list of examples",
			Action: list,
		},
	}
	app.HideVersion = true
	app.Metadata = map[string]interface{}{
		"examples": daxExamples,
	}
	app.Run(os.Args)
}
