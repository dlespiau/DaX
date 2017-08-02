package main

import (
	"fmt"
	"os"

	//dax "github.com/dlespiau/dax/lib"
	"github.com/dlespiau/dax/midi"

	"github.com/urfave/cli"
)

type dax struct {
	seq *midi.Sequencer
}

func (d *dax) Init() {
	d.seq, _ = midi.NewSequencer("DaX", midi.OpenDuplex)
	d.seq.CreateControllerPort("Controller")
}

func (d *dax) listMidiDevices() {
	for _, device := range d.seq.GetDevices() {
		fmt.Printf("%d: %s\n", device.Client, device.Name)
		for _, port := range device.Ports {
			fmt.Printf("    %d: %s\n", port.Port, port.Name)
		}
	}
}

func getDax(ctx *cli.Context) *dax {
	return ctx.App.Metadata["dax"].(*dax)
}

func listMidiDevices(ctx *cli.Context) error {
	dax := getDax(ctx)
	dax.listMidiDevices()
	return nil
}

func main() {
	dax := &dax{}
	dax.Init()

	app := cli.NewApp()
	app.Name = "DaX"
	app.Usage = "Not sure what this is"
	app.Commands = []cli.Command{
		{
			Name:   "list-midi-devices",
			Usage:  "List the MIDI devices on the system",
			Action: listMidiDevices,
		},
	}
	app.HideVersion = true
	app.Metadata = map[string]interface{}{
		"dax": dax,
	}
	app.Run(os.Args)
}
