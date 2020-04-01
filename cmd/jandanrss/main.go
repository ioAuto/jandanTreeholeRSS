package main

import (
	"flag"
	"github.com/iochen/jandanTreeholeRSS"
	"io/ioutil"
	"log"
)

func main() {
	FlagNum := flag.Uint("n", 50, "")
	FlagOut := flag.String("o", "rss.xml", "")
	flag.Parse()

	var feed string
	var err error
	for i := 0; i < 5; i++ {
		feed, err = jandanTreeholeRSS.GetFeed(int(*FlagNum))
		if err == nil {
			if err := ioutil.WriteFile(*FlagOut, []byte(feed), 0644); err != nil {
				log.Println(err)
			}
			return
		}
		log.Println(err)
		log.Println("retry")
	}
	log.Panic(err)
}
