package structsypunteros

import "fmt"

type pc struct {
	ram   int
	disk  int
	brand string
}

func (myPc pc) ping() {
	fmt.Println(myPc.brand, "pong")
}

func (myPc *pc) duplicateRAM() {
	myPc.ram = myPc.ram * 2
}

//This function will print this custom message on any fmt.println
func (myPc pc) String() string {
	return fmt.Sprintf("Computer specs: %d GB of RAM, %d GB of Disk and is of the brand %s", myPc.ram, myPc.disk, myPc.brand)
}

func StructsyPunterosExample() {
	a := 50
	b := &a

	fmt.Println(a)
	fmt.Println(*b)

	*b = 100
	fmt.Println(a)

	myPc := pc{ram: 16, disk: 260, brand: "msi"}
	fmt.Println(myPc)
	myPc.ping()

	myPc.duplicateRAM()
	fmt.Println(myPc)
}
