package main

import (
	"bytes"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	findStr := []byte(`    name="SrKomodo.Software.shadowfoxUpdater"
`)

	os.Remove("manifest.xml")

	manifest, err := ioutil.ReadFile("_manifest.xml")
	if err != nil {
		if !os.IsNotExist(err) {
			panic(err)
		}
	}

	err = ioutil.WriteFile(
		"manifest.xml",
		bytes.Replace(
			manifest,
			findStr,
			append(
				findStr,
				[]byte("    version=\""+strings.TrimPrefix(os.Args[1], "v")+".0\"\n")...,
			),
			-1,
		),
		0644,
	)
	if err != nil {
		panic(err)
	}
}
