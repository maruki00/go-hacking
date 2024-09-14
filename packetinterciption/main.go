package main

// go get
import (
	"fmt"
	"log"

	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
)

var (
	iface    = "wlan0"
	devFound = false
	snaplen  = int32(1600)
	timeout  = pcap.BlockForever
	filter   = "tcp and port 80"
	promisc  = false
)

var keywords = []string{
	"username",
	"password",
	"uname",
	"pass",
	"urname",
	"uemail",
}

func main() {

	devs, err := pcap.FindAllDevs()
	if err != nil {
		panic("hello world")
	}

	for _, dev := range devs {
		if dev.Name == iface {
			devFound = true
		}
		if !devFound {
			log.Fatalf("device %s not found", dev.Name)
		}
		handel, err := pcap.OpenLive(iface, snaplen, promisc, timeout)
		if err != nil {
			log.Fatal(err.Error())
		}

		defer handel.Close()
		err = handel.SetBPFFilter(filter)
		if err != nil {
			log.Fatal(err.Error())
		}

		src := gopacket.NewPacketSource(handel, handel.LinkType())

		for pkt := range src.Packets() {
			applayer := pkt.ApplicationLayer()
			if applayer == nil {
				continue
			}

			payload := applayer.Payload()
			fmt.Println(string(payload))
			// for _, s := range keywords {
			// 	index := strings.Index(string(payload), s)
			// 	if index != -1 {
			// 		fmt.Println(string(payload[index:]))
			// 	}
			// }

		}

	}
}
