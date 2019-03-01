package generator

import (
	"fmt"
	"log"
	"os"
)

func (g *generator) usecase() error {
	tmp := `// Package usecase auto generate usecase
package usecase

type ` + g.Name() + `Usecase struct {}

// New` + toUpperCamel(g.Name()) + `Usecase create new ` + g.Name() + ` usecase
func New` + toUpperCamel(g.Name()) + `Usecase () ` + toUpperCamel(g.Name()) + `Usecase {
	return &` + g.Name() + `Usecase{}
}

// ` + toUpperCamel(g.Name()) + `Usecase usecase interface
type ` + toUpperCamel(g.Name()) + `Usecase interface {
	// add methods here.
}
`

	if err := os.MkdirAll(g.Name()+"/usecase", 0777); err != nil {
		log.Fatal(err)
		return err
	}
	file, err := os.OpenFile(g.Name()+"/usecase/"+g.Name()+"_usecase.go", os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer file.Close()
	fmt.Fprintln(file, tmp)
	return nil
}
