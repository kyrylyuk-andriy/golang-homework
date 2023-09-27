package hw6_delivery

import (
	"fmt"
)

type Delivery interface {
	GetReceiver() string
	GetSender() string
	Send()
}

type Package struct {
	Item            string
	SenderName      string
	SenderAddress   string
	ReceiverName    string
	ReceiverAddress string
}

func (p Package) GetReceiver() string {
	return p.ReceiverName
}

func (p Package) GetSender() string {
	return p.SenderName
}

func (p Package) Send() {
	if p.Item == "box" {
		fmt.Printf("box was sent via Terminal1 from %s to $%s\n", p.SenderAddress, p.ReceiverAddress)
	} else if p.Item == "envelope" {
		fmt.Printf("envelope was sent via Terminal2 from %s to $%s\n", p.SenderAddress, p.ReceiverAddress)
	} else {
		fmt.Printf("Ukpposhta was not able to detect what is the item and return to the %s\n", p.SenderName)
	}
}

func SendPackage(office Delivery) {
	office.Send()
}
