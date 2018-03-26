package main

import (
	"encoding/gob"
	"log"
	"os"
)

func save(dict map[[16]byte]bool) error {
	var (
		err  error
		path = "/opt/dict.gob"
	)
	if len(dict) == 0 {
		log.Println("nil map")
		return nil
	}
	f, err := os.Create(path)
	if err != nil {
		log.Println(err)
		return err
	}

	en := gob.NewEncoder(f)
	if err = en.Encode(dict); err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func loadmap(path string) error {
	var (
		err  error
		dict = make(map[[16]byte]bool)
	)
	f, err := os.Open(path)
	if err != nil {
		log.Println(err)
	}
	dec := gob.NewDecoder(f)
	if err = dec.Decode(&dict); err != nil {
		log.Println(err)
		return err
	}
	log.Printf("dic len:%v\n", len(dict))
	return nil
}
