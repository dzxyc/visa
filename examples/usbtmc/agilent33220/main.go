// Copyright (c) 2017 The visa developers. All rights reserved.
// Project site: https://github.com/gotmc/visa
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

package main

import (
	"log"
	"time"

	"github.com/gotmc/visa"
	_ "github.com/gotmc/visa/drivers/tcpip"
	_ "github.com/gotmc/visa/drivers/usbtmc"
)

// Can use either a USBTMC or TCP/IP socker to communicate with the function
// generator. Below are two different VISA address strings.
const (
	usbAddress string = "USB0::2391::1031::MY44035349::INSTR"
	tcpAddress string = "TCPIP::10.12.100.242::2055::SOCKET"
)

func main() {

	start := time.Now()
	fgen, err := visa.NewResource(usbAddress)
	if err != nil {
		log.Fatal("Couldn't open the resource for the function generator")
	}
	defer fgen.Close()

	log.Printf("%.2fs to setup VISA resource\n", time.Since(start).Seconds())

	fgen.WriteString("apply:sinusoid 2340, 0.1, 0.0")
	fgen.WriteString("burst:internal:period 0.112")
	fgen.WriteString("burst:ncycles 131")
	fgen.WriteString("burst:state on")
	fgen.WriteString("*idn?")

	start = time.Now()
	var buf []byte
	bytesRead, err := fgen.Read(buf)
	log.Printf("%.2fs to read %d bytes\n", time.Since(start).Seconds(), bytesRead)
	if err != nil {
		log.Printf("Error reading: %s", err)
	}
	// fmt.Printf("Read %d bytes = %s", bytesRead, buf[12:bytesRead])
	// fmt.Printf("Last rune read = %x\n", buf[bytesRead-1:bytesRead])
	// fmt.Printf("Last rune read = %q\n", buf[bytesRead-1:bytesRead])

	log.Println("Made it here!")

	defer fgen.Close()

}