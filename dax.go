package main

import (
	"fmt"
	"os"

	//dax "github.com/dlespiau/dax/lib"
	"github.com/dlespiau/dax/lib/flag"
	"github.com/dlespiau/dax/lib/midi"
)

type Dax struct {
	options struct {
		list_midi_devices bool
	}
}

var app Dax

func init() {
	flag.BoolVar(&app.options.list_midi_devices,
		[]string{"-list-midi-devices"}, false,
		"List the MIDI devices connected to the system")
	flag.Parse()
}

func list_midi_devices() {
	seq, _ := midi.NewSequencer("DaX", midi.OpenDuplex)

	for _, device := range seq.GetDevices() {
		fmt.Printf("%d: %s\n", device.Client, device.Name)
		for _, port := range device.Ports {
			fmt.Printf("    %d: %s\n", port.Port, port.Name)
		}
	}
}

func main() {
	quit_early := false

	if app.options.list_midi_devices {
		list_midi_devices()
		quit_early = true
	}

	if quit_early {
		os.Exit(0)
	}
}
