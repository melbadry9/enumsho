package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"./apishodan"
)

func main() {
	domain := flag.String("d", "", "[+] Domain to find subdomains")
	shodanKey := flag.String("s", "", "[+] Shodan api key")
	verbose := flag.Bool("v", false, "[+] Show all output")
	flag.Parse()

	if *domain == "" || *shodanKey == "" {
		fmt.Printf("Usage %s -d target.com -s MYShodaNKey", os.Args[0])
		os.Exit(1)
	}

	apiKey := apishodan.New(*shodanKey)

	domainSearch := *domain

	subdomain, err := apiKey.GetSubdomain(domainSearch)
	if err != nil {
		log.Panicln(err)
	}

	if *verbose == true {

		info, err := apiKey.InfoAccount()

		if err != nil {
			log.Panicln(err)
		}

		fmt.Printf(
			"Credits: %d\nScan Credits: %d \n\n",
			info.QueryCredits, info.ScanCredits)

		for _, v := range subdomain.Data {
			d := v.SubD + subdomain.Domain
			fmt.Printf("Domain: %s\nIP/DNS :%s\nLast Scan made by shodan:%s\n", d, v.Value, v.LastSeen)
		}

	} else {
		for _, v := range subdomain.SubDomains {
			fmt.Printf("%s.%s\n", v, domainSearch)
		}
	}
}
