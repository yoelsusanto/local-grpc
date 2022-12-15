package utils

import (
	"bufio"
	"bytes"
	"context"
	"fmt"
	"log"
	"os/exec"
	"time"
)

func StartSubProcess(ctx context.Context, programName string, args ...string) {
	var out bytes.Buffer

	cmd := exec.CommandContext(ctx, fmt.Sprintf("../%s", programName), args...)

	cmd.Stderr = &out
	cmd.Stdout = &out

	// fmt.Println("Starting server..")

	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}

	go func() {
		reader := bufio.NewReader(&out)

		for keepGoing, timeout := true, ctx.Done(); keepGoing; {
			select {
			case <-timeout:
				fmt.Println("done...")
				keepGoing = false
			default:
				line, err := reader.ReadString('\n')
				if err == nil {
					fmt.Printf("server: %s", line)

				}

				time.Sleep(1 * time.Second)
			}
		}
	}()
}
