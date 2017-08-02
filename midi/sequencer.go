package midi

// #cgo pkg-config: alsa
// #include <alsa/asoundlib.h>
import "C"

import (
	"errors"
	"unsafe"
)

type Device struct {
	Name   string
	Client int
	Ports  []Port
}

type Port struct {
	Name string
	Port int
}

type Sequencer struct {
	handle *C.snd_seq_t
}

type OpenDirection int

const (
	OpenOutput OpenDirection = 1
	OpenInput                = 2
	OpenDuplex               = 3
)

func NewSequencer(name string, direction OpenDirection) (*Sequencer, error) {
	var handle *C.snd_seq_t
	var ret C.int

	cdefault := C.CString("default")
	ret = C.snd_seq_open(&handle, cdefault, C.int(direction), 0)
	C.free(unsafe.Pointer(cdefault))
	if ret < 0 {
		return nil, errors.New("alsa: couldn't open sequencer")
	}

	cname := C.CString(name)
	C.snd_seq_set_client_name(handle, cname)
	C.free(unsafe.Pointer(cname))

	seq := new(Sequencer)
	seq.handle = handle

	return seq, nil
}

func (seq *Sequencer) Close() {
	C.snd_seq_close(seq.handle)
	seq.handle = nil
}

func (seq *Sequencer) CreateControllerPort(name string) Port {
	var port C.int

	cname := C.CString(name)
	port = C.snd_seq_create_simple_port(seq.handle, cname,
		C.SND_SEQ_PORT_CAP_WRITE|C.SND_SEQ_PORT_CAP_SUBS_WRITE,
		C.SND_SEQ_PORT_TYPE_APPLICATION)
	C.free(unsafe.Pointer(cname))

	return Port{Port: int(port)}
}

func (seq *Sequencer) getNbClients() int {
	var system_info *C.snd_seq_system_info_t

	C.snd_seq_system_info_malloc(&system_info)
	C.snd_seq_system_info(seq.handle, system_info)
	n_clients := int(C.snd_seq_system_info_get_cur_clients(system_info))
	C.snd_seq_system_info_free(system_info)

	return n_clients
}

func (seq *Sequencer) queryNextClient(info *C.snd_seq_client_info_t) bool {
	return C.snd_seq_query_next_client(seq.handle, info) == 0
}

func getNumPorts(info *C.snd_seq_client_info_t) int {
	return int(C.snd_seq_client_info_get_num_ports(info))
}

func (seq *Sequencer) GetDevices() []Device {
	var client_info *C.snd_seq_client_info_t
	var port_info *C.snd_seq_port_info_t

	n_clients := seq.getNbClients()
	devices := make([]Device, n_clients, n_clients)

	C.snd_seq_client_info_malloc(&client_info)
	C.snd_seq_port_info_malloc(&port_info)

	d := 0
	C.snd_seq_client_info_set_client(client_info, -1)
	for ; seq.queryNextClient(client_info) && d < n_clients; d++ {
		// client number
		cclient := C.snd_seq_client_info_get_client(client_info)
		devices[d].Client = int(cclient)

		// name
		cname := C.snd_seq_client_info_get_name(client_info)
		devices[d].Name = C.GoString(cname)

		// list of ports
		C.snd_seq_port_info_set_client(port_info, cclient)

		n_ports := getNumPorts(client_info)
		devices[d].Ports = make([]Port, n_ports, n_ports)

		p := 0
		C.snd_seq_port_info_set_port(port_info, -1)
		for ; C.snd_seq_query_next_port(seq.handle, port_info) >= 0; p++ {
			cname := C.snd_seq_port_info_get_name(port_info)
			devices[d].Ports[p].Name = C.GoString(cname)

			port := C.snd_seq_port_info_get_port(port_info)
			devices[d].Ports[p].Port = int(port)
		}
	}

	C.snd_seq_client_info_free(client_info)
	C.snd_seq_port_info_free(port_info)

	return devices
}
