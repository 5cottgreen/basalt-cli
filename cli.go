package basaltcli

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/5cottgreen/basalt"
)

func HandleCommand() {
	fmt.Println("# Enter command")
loop:
	fmt.Println(">")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	line := scanner.Text()

	space := strings.IndexByte(line, ' ')
	if space == -1 {
		fmt.Printf("# '%s' command entered incorrectly\r\n", line)
		goto loop
	}

	cmd := line[:space]
	switch cmd {
	case "get":
		argline := line[space+1:]
		args := strings.Split(argline, " ")
		l := len(args)

		if l > 3 {
			fmt.Printf("# '%s' too many arguments\r\n", line)
			goto loop
		}

		for _, arg := range args[:l-1] {
			switch arg {
			case "sisyphus":
			case "p9":
			case "p10":
			default:
				fmt.Println("# sisyphus p9 p10 packages only allowed")
				goto loop
			}
		}

		switch args[l-1] {
		case "-s":
			disp := render(basalt.SimpleFetch(args[:l-1]))
			json, err := json.Marshal(disp)
			if err != nil {
				log.Fatal(err)
			}

			fmt.Println(string(json))
			goto loop
		case "-c":
			disp := render(basalt.ComparativeFetch(args[:l-1]))
			json, err := json.Marshal(disp)
			if err != nil {
				log.Fatal(err)
			}

			fmt.Println(string(json))
			goto loop
		case "-cr":
			disp := render(basalt.ReverseComparativeFetch(args[:l-1]))
			json, err := json.Marshal(disp)
			if err != nil {
				log.Fatal(err)
			}

			fmt.Println(string(json))
			goto loop
		case "-cv":
			disp := render(basalt.VersionComparativeFetch(args[:l-1]))
			json, err := json.Marshal(disp)
			if err != nil {
				log.Fatal(err)
			}

			fmt.Println(string(json))
			goto loop
		default:
			fmt.Println("# -s -c -cr -cv flags only allowed")
		}

	}
}

func render(pkgs basalt.Packages) (disp display) {
	return display{Packages: pkgs, Time: time.Now()}
}
