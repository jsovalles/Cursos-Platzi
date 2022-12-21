package main

import (
	"github.com/jsovalles/Cursos_Platzi/abstractFactory"
	"github.com/jsovalles/Cursos_Platzi/concurrency"
	"github.com/jsovalles/Cursos_Platzi/finalProject"
	"github.com/jsovalles/Cursos_Platzi/functions"
	"github.com/jsovalles/Cursos_Platzi/introduction"
	"github.com/jsovalles/Cursos_Platzi/oop"
)

func main() {
	if !true {
		introduction.Introduction()
		oop.EmployeeExample()
		oop.InheritanceExample()
		abstractFactory.AbstractFactoryExample()
		functions.FuncionAnonimaExample()
		concurrency.BufferedAndUnbufferedChannelsExample()
		concurrency.WaitGroupExample()
		concurrency.BufferedChannelsAsStoplight()
		concurrency.PipelineExample()
		concurrency.WorkerPoolExample()
		concurrency.MultiplexExample()
	}
	finalProject.Project()
}
