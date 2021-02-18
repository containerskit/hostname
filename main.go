package main

import (
	"io/ioutil"
	"log"
	"os"
	"syscall"
)

const usage = `
hostname HOSTNAMEFILE

Set hostname from the file provided. This tool can be used to set an instance
hostname from the metadata file.
`

func main() {
	log.SetOutput(os.Stdout)
	log.SetPrefix("")
	log.SetFlags(0)

	if len(os.Args) < 2 {
		log.Fatalln(usage)
	}

	file := os.Args[1]

	hostname, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatalf("error: unable to read hostname file %s\n", file)
	}

	if err := syscall.Sethostname(hostname); err != nil {
		log.Fatalf("error: unable to set hostname: %v\n", err)
	}

	log.Printf("hostname: %s\n", string(hostname))
}
