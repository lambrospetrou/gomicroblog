package main

import (
	"flag"
	"fmt"
	"github.com/lambrospetrou/gomicroblog/gen"
	"github.com/lambrospetrou/gomicroblog/view"
	"log"
	"path/filepath"
	"runtime"
)

func main() {
	fmt.Println("Go Microblog service started!")

	// use all the available cores
	runtime.GOMAXPROCS(runtime.NumCPU())

	var dir_site = flag.String("site", "", "specify a directory that contains the site to generate")
	flag.Parse()

	log.Println("site:", *dir_site)
	if len(*dir_site) > 0 {
		ViewBuilder := view.NewBuilder(filepath.Join(*dir_site, "_layouts"))
		err := gen.GenerateSite(*dir_site, ViewBuilder, filepath.Join(*dir_site, "config.json"))
		if err != nil {
			panic(err)
		}
		return
	} else {
		log.Fatalln("Site source directory not given")
		return
	}
}
