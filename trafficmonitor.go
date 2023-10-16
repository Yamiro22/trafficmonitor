package main

import (
	"fmt"
	"log"

	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
)

func main() {
	// Open the main network device for packet capture
	handle, err := pcap.OpenLive("wlo1", 1600, true, pcap.BlockForever)
	if err != nil {
		log.Fatal(err)
	}
	defer handle.Close()

	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())

	fmt.Println("Capturing packets... Press Ctrl+C to stop.")
	for packet := range packetSource.Packets() {
		// Print packet information
		fmt.Println(packet)
	}
}
