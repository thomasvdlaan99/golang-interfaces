package learning

import "fmt"

type engine interface {
	kilometerLeft() uint16
}

type gasEngine struct {
	kpl   uint16
	liter uint16
}

type electricEngine struct {
	kpkwh uint16
	kwh   uint16
}

func (e gasEngine) kilometerLeft() uint16 {
	return e.kpl * e.liter
}

func (e electricEngine) kilometerLeft() uint16 {
	return e.kpkwh * e.kwh
}

func canMakeIt(e engine, kilometers uint16) string {
	if kilometers <= e.kilometerLeft() {
		return "You can make it"
	} else {
		return "You need to fuel up first"
	}

}

func StructsInterfaces() {
	var elecEngine electricEngine = electricEngine{12, 50}
	var gassEngine electricEngine = electricEngine{12, 50}
	fmt.Println(canMakeIt(elecEngine, 1000))
	fmt.Println(canMakeIt(gassEngine, 100))

}
