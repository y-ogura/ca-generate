## ca-generate

ca-generate is a generate tool for clean architecture in Go.

## Installation
Asuuming you have a working [Go](https://golang.org) setup

```bash
go get -u github.com/y-ogura/ca-generate
```


## Usage

generate files with ca-generate command

```bash
cd $GOPATH/src/project/path
ca-generate [domain]
```

The tool includes contextual help
```
Usage:
  ca-generate [domain]

Flags:
  -c  Create controller flag (default true)
  -u  Create usecase flag (default true)
  -r  Create repository flag (default true)
```

if you wan't generate controller.
use flag `-c=false`

```bash
ca-generate -c=false [domain]
```
