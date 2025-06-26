package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/ichang0301/go-design-patterns/04_structural_patterns2/4-2_decorator/4-2-2_server_middleware/server"
)

func main() {
	fmt.Println("Enter the type number of server you want to launch from the following:")
	fmt.Println("1. Plain server")
	fmt.Println("2. Server with logging")
	fmt.Println("3. Server with logging and authentication")

	var selection int
	fmt.Fscanf(os.Stdin, "%d", &selection)

	var mySuperServer http.Handler

	switch selection {
	case 1:
		mySuperServer = new(server.MyServer)
	case 2:
		mySuperServer = &server.LoggerMiddleware{
			LogWriter: os.Stdout,
			Handler:   new(server.MyServer),
		}
	case 3:
		var username, password string
		fmt.Println("Enter username and password separated by a space:")
		fmt.Fscanf(os.Stdin, "%s %s", &username, &password)

		mySuperServer = &server.LoggerMiddleware{
			Handler: &server.BasicAuthMiddleware{
				Handler:  new(server.MyServer),
				User:     username,
				Password: password,
			},
			LogWriter: os.Stdout,
		}
	default:
		mySuperServer = new(server.MyServer)
	}

	http.Handle("/", mySuperServer)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
