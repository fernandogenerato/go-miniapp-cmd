package app

import (
	"fmt"
	"github.com/urfave/cli"
	"log"
	"net"
)

// Build new cmdline to be executed
func Build() *cli.App {
	app := cli.NewApp()
	app.Name = "Command line App"
	app.Usage = "Search IPs and names of servers from internet"

	flags := []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Value: "github.com/fernandogenerato",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:   "ip",
			Usage:  "Search IPs of address in internet",
			Flags:  flags,
			Action: findIPs,
		},
		{
			Name:   "servers",
			Usage:  "Search name of servers in internet",
			Flags:  flags,
			Action: searchServer,
		},
	}
	return app
}
func searchServer(c *cli.Context) {
	host := c.String("host")
	servers, err := net.LookupNS(host)
	if err != nil {
		log.Fatal(err)
	}

	for _, s := range servers {
		fmt.Println(s.Host)
	}

}
func findIPs(c *cli.Context) {
	host := c.String("host")

	ips, err := net.LookupIP(host)
	if err != nil {
		log.Fatal(err)
	}

	for _, ip := range ips {
		fmt.Println(ip)
	}

}
