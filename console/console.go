// small library to read input from stdin and write to stdout
package console

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func Start(input chan string) {
	// start goroutine to keep listen what is typed to console input (stdin)
	go func(in chan string) {
		// create new reader from stdin
		reader := bufio.NewReader(os.Stdin)
		// start infinite loop to continuously listen to input
		for {
			// read by one line (enter pressed)
			s, err := reader.ReadString('\n')
			// check for errors
			if err != nil {
				// close channel just to inform others
				close(in)
				log.Println("Error in read string", err)
			}
			in <- strings.TrimSpace(s)
		}
		// pass input channel to closure func
	}(input)
}
