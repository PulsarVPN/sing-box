//go:build !generate

package main

import "github.com/pulsarvpn/sing-box/log"

func main() {
	if err := mainCommand.Execute(); err != nil {
		log.Fatal(err)
	}
}
