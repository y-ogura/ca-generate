package generator

import (
	"fmt"
	"log"
	"os"
)

func (g *generator) controller() error {
	tmp := `// Package controller auto generate controller
package controller

import "net/http"

// ` + toUpperCamel(g.Name()) + `Controller ` + g.Name() + ` controller
type ` + toUpperCamel(g.Name()) + `Controller struct {}

// New` + toUpperCamel(g.Name()) + `Controller ` + g.Name() + ` controller
func New` + toUpperCamel(g.Name()) + `Controller (w http.ResponseWriter, r *http.Request) {
	c := &` + toUpperCamel(g.Name()) + `Controller{}
}
`
	if err := os.MkdirAll(g.Name()+"/controller", 0777); err != nil {
		log.Fatal(err)
		return err
	}
	file, err := os.OpenFile(g.Name()+"/controller/"+g.Name()+"_controller.go", os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer file.Close()
	fmt.Fprintln(file, tmp)
	return nil
}
