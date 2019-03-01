package generator

import (
	"fmt"
	"log"
	"os"
)

func (g *generator) repository() error {
	tmp := `// Package repository auto generate repository
package repository

type ` + g.Name() + `Repository struct {}

// New` + toUpperCamel(g.Name()) + `Repository create new ` + g.Name() + ` repository
func New` + toUpperCamel(g.Name()) + `Repository () ` + toUpperCamel(g.Name()) + `Repository {
	return &` + g.Name() + `Repository{}
}

// ` + toUpperCamel(g.Name()) + `Repository repository interface
type ` + toUpperCamel(g.Name()) + `Repository interface {
	// add methods here.
}
`

	if err := os.MkdirAll(g.Name()+"/repository", 0777); err != nil {
		log.Fatal(err)
	}
	file, err := os.OpenFile(g.Name()+"/repository/"+g.Name()+"_repository.go", os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer file.Close()
	fmt.Fprintln(file, tmp)
	return nil
}
