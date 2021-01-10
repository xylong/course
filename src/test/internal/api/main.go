package main

import (
	"course/src/test/internal/api/ctrl"
	"github.com/shenyisyn/goft-gin/goft"
	"log"
)

func main() {
	if err := goft.Ignite().Mount("test/v1", ctrl.Controllers...).Run(); err != nil {
		log.Fatal(err)
	}
}
