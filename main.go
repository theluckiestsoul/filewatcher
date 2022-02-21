package main

import (
	"flag"
	"log"

	"github.com/fsnotify/fsnotify"
)

func main() {
	p := flag.String("p", "", "Path to watch")
	flag.Parse()
	if *p == "" {
		panic("enter path")
	}
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	done := make(chan bool)
	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				log.Println("event:", event)
				if event.Op&fsnotify.Write == fsnotify.Write {
					log.Println("modified file:", event.Name)
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.Println("error:", err)
			}
		}
	}()
	err = watcher.Add(*p)
	if err != nil {
		log.Fatal(err)
	}
	<-done
}
