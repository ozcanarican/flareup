package main

import (
	"flag"
	"fmt"
	"log"
	"strings"

	"github.com/ozcanarican/flareup/internal/cloudflare"
	"github.com/ozcanarican/flareup/internal/ip"
)

func main() {
	var ipadres string
	argFull := flag.String("d", "", "Sub domain to redirect")
	argIp := flag.String("i", "", "IP address to redirect. Leave empty for current public ip")
	argForce := flag.Bool("f", false, "Forcefuly redirect")
	argRemove := flag.Bool("r", false, "Remove existing record")

	flag.Parse()

	if *argFull == "" {
		log.Fatalln("Requires a sub domain name.\nUsage: flareup -d <sub.domain.name>")
	}

	if *argIp != "" {
		ipadres = *argIp
	} else {
		ipadres = ip.PublicIP()
	}
	parts := strings.Split(*argFull, ".")
	domain := strings.Replace(*argFull, fmt.Sprintf("%v.", parts[0]), "", 1)
	fmt.Println(domain)
	fmt.Printf("Domain: %v, IP Address: %v, Forceful: %v", *argFull, ipadres, *argForce)

	if *argRemove {
		cloudflare.RemoveRecord(*argFull, domain)
	} else {
		cloudflare.AddRecord(*argFull, domain, ipadres, *argForce)
	}
}
