package modificadorAcceso

import "fmt"

//Struct con acceso publico
type CarPublic struct {
	//Public access
	Brand string
	//Private access
	year int
}

type carPrivate struct {
	brand string
	year  int
}

func PrintMessage(text string) {
	fmt.Println(text)
}
