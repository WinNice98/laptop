package laptop

import (
	"evm"
	"pc"
	"fmt"
)

type Laptop struct {
	pc.PC
	Mah       uint64
	Wifi      bool
	Bluetooth bool
}

func NewLaptop(name, cpu string, ram uint64, gpu string, vram uint64, ethernet bool, mah uint64, wifi, bluetooth bool) *Laptop {
	a := evm.EVM{name, cpu, ram}
	b := pc.PC{EVM: a, GPU: gpu, VRAM: vram, Ethernet: ethernet}
	return &Laptop{PC: b, Mah: mah, Wifi: wifi, Bluetooth: bluetooth}
}

func showbool(b bool) string {
	if b {
		return "true"
	}
	return "false"
}

func (l *Laptop) ShowSpecs() {
	fmt.Printf("%s\n%s\n%d\n%s\n%d\nEthernet = %s\n%d\nWifi = %s\nBluetooth = %s\n",
		l.Name, l.CPU, l.RAM, l.GPU, l.VRAM, showbool(l.Ethernet), l.Mah, showbool(l.Wifi), showbool(l.Bluetooth))
}

