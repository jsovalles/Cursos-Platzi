package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jsovalles/Cursos_Platzi/src/arrayslice"
	"github.com/jsovalles/Cursos_Platzi/src/channels"
	"github.com/jsovalles/Cursos_Platzi/src/ciclos"
	"github.com/jsovalles/Cursos_Platzi/src/goget"
	"github.com/jsovalles/Cursos_Platzi/src/goroutines"
	"github.com/jsovalles/Cursos_Platzi/src/interfaces"
	"github.com/jsovalles/Cursos_Platzi/src/keywords"
	"github.com/jsovalles/Cursos_Platzi/src/maps"
	"github.com/jsovalles/Cursos_Platzi/src/math"
	"github.com/jsovalles/Cursos_Platzi/src/modificadorAcceso"
	"github.com/jsovalles/Cursos_Platzi/src/performancetest"
	"github.com/jsovalles/Cursos_Platzi/src/rangeSlice"
	"github.com/jsovalles/Cursos_Platzi/src/rangecloseselect"
	"github.com/jsovalles/Cursos_Platzi/src/structs"
	"github.com/jsovalles/Cursos_Platzi/src/structsypunteros"
	"net/http"
)

func main() {
	if !true {
		mathclass.MathExamples()
		ciclos.CiclosExample()
		keywords.KeywordExample()
		maps.MapsExample()
		rangeSlice.RangeSliceExample()
		arrayslice.ArraySliceExample()
		structs.StructsExample()
		modificadorAcceso.PrintMessage("hola mundo")
		structsypunteros.StructsyPunterosExample()
		interfaces.FuncionesExample()
		goroutines.GoRoutinesExample()
		channels.ChannelExample()
		rangecloseselect.ChannelsExample()
		router := mux.NewRouter().StrictSlash(true)
		router.HandleFunc("/", goget.HomePage)
		http.ListenAndServe(":8088", router)
		fmt.Printf("2, %T", 2)
	}
	performancetest.PerformanceScript()
}
