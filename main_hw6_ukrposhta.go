package main

import (
	"github.com/kyrylyuk-andriy/golang-homework/hw6_ukrposhta"
)

func main() {
	giftbox := hw6_ukrposhta.Package{
		Item:            "box",
		SenderName:      "Bob",
		SenderAddress:   "Kyiv",
		ReceiverName:    "Robert",
		ReceiverAddress: "Lviv",
	}

	giftenvelope := hw6_ukrposhta.Package{
		Item:            "envelope",
		SenderName:      "Peter",
		SenderAddress:   "Ternopil",
		ReceiverName:    "Dan",
		ReceiverAddress: "Rivne",
	}

	giftparcel := hw6_ukrposhta.Package{
		Item:            "parcel",
		SenderName:      "Peter",
		SenderAddress:   "Ternopil",
		ReceiverName:    "Dan",
		ReceiverAddress: "Rivne",
	}

	hw6_ukrposhta.SendPackage(giftbox)
	hw6_ukrposhta.SendPackage(giftenvelope)
	hw6_ukrposhta.SendPackage(giftparcel)
}
