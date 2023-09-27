package main

import (
	"github.com/kyrylyuk-andriy/golang-homework/hw6_delivery"
)

func main() {
	giftbox := hw6_delivery.Package{
		Item:            "box",
		SenderName:      "Bob",
		SenderAddress:   "Kyiv",
		ReceiverName:    "Robert",
		ReceiverAddress: "Lviv",
	}

	giftenvelope := hw6_delivery.Package{
		Item:            "envelope",
		SenderName:      "Peter",
		SenderAddress:   "Ternopil",
		ReceiverName:    "Dan",
		ReceiverAddress: "Rivne",
	}

	giftparcel := hw6_delivery.Package{
		Item:            "parcel",
		SenderName:      "Peter",
		SenderAddress:   "Ternopil",
		ReceiverName:    "Dan",
		ReceiverAddress: "Rivne",
	}

	hw6_delivery.SendPackage(giftbox)
	hw6_delivery.SendPackage(giftenvelope)
	hw6_delivery.SendPackage(giftparcel)
}
