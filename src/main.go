package main

import (
	"course/src/interfaces/config"
	"course/src/interfaces/controller"
	"github.com/shenyisyn/goft-gin/goft"
)

func main() {
	goft.Ignite().
		Config(config.NewDBConfig(), config.NewCourseServiceConfig()).
		Mount("v1", controller.NewCourseController()).
		Launch()
}
