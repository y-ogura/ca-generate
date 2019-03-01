package generator

import (
	"errors"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

type generator struct {
	path string
	flag []bool
	name string
}

// New create generator
func New(flag []bool, name string) Generator {
	_, err := os.Stat(name)
	if err == nil {
		msg := errors.New(name + " is already exists")
		log.Fatal(msg)
	}
	current, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	rep := regexp.MustCompile(`[\S\s]+?src/`)
	path := rep.ReplaceAllString(current, "")
	return &generator{
		path: path,
		flag: flag,
		name: name,
	}
}

// Generator generator interface
type Generator interface {
	Path() string
	Flag() []bool
	Name() string
	Generate()
}

func (g *generator) Path() string {
	return g.path
}

func (g *generator) Flag() []bool {
	return g.flag
}

func (g *generator) Name() string {
	return g.name
}

func (g *generator) Generate() {
	fmt.Println("start generate files.")
	if g.Flag()[0] {
		fmt.Println("generate controller")
		g.controller()
	}
	if g.Flag()[1] {
		fmt.Println("generate usecase")
		g.usecase()
	}
	if g.Flag()[2] {
		fmt.Println("generate repository")
		g.repository()
	}
	fmt.Println("finish generate files.")
}

func toUpperCamel(str string) string {
	return strings.Replace(str, str[:1], strings.ToUpper(str[:1]), 1)
}
