package main

import (
	"fmt"
	"github.com/mattn/goemon"
	"os"
)

var defaultConf = `# Generated by goemon -g
livereload: :35730
tasks:
- match: './assets/*.js'
  commands:
  - minifyjs -m -i ${GOEMON_TARGET_FILE} > ${GOEMON_TARGET_DIR}/${GOEMON_TARGET_NAME}.min.js
  - :livereload /
- match: './assets/*.html'
  commands:
  - :livereload /
- match: '*.go'
  commands:
  - go build
  - :restart
  - :livereload /
`

func usage() {
	fmt.Printf("Usage of %s [command] [args...]\n", os.Args[0])
	fmt.Println(" goemon -g : generate default configuration")
	os.Exit(1)
}

func main() {
	switch len(os.Args) {
	case 1:
		usage()
	default:
		switch os.Args[1] {
		case "-h":
			usage()
		case "-g":
			fmt.Println(defaultConf)
			return
		}
	}

	goemon.NewWithArgs(os.Args[1:]).Run()
}
