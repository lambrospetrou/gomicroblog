package main

import (
	"flag"
	"log"
	"path/filepath"
	"runtime"

	"github.com/lambrospetrou/gomicroblog/gen"
)

func main() {
	//log.Println("Go Microblog service started!")

	// use all the available cores
	runtime.GOMAXPROCS(runtime.NumCPU())

	var dirSite = flag.String("site", "", "specify a directory that contains the site to generate")
	flag.Parse()

	//log.Println("site:", *dirSite)
	if len(*dirSite) > 0 {
		err := gen.GenerateSite(*dirSite, filepath.Join(*dirSite, "config.json"))
		if err != nil {
			log.Fatalln(err)
			panic(err)
		}
		return
	}
	log.Fatalln("Site source directory not given")
	return
}
