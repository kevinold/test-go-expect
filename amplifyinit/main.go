package main

import (
	"log"
	"os"
	"os/exec"
	"time"

	expect "github.com/Netflix/go-expect"
)

func main() {
	c, err := expect.NewConsole(expect.WithStdout(os.Stdout))
	if err != nil {
		log.Fatal(err)
	}
	defer c.Close()

	cmd := exec.Command("amplify", "init")
	cmd.Stdin = c.Tty()
	cmd.Stdout = c.Tty()
	cmd.Stderr = c.Tty()

	go func() {
		c.ExpectEOF()
	}()

	err = cmd.Start()
	if err != nil {
		log.Fatal(err)
	}

	time.Sleep(time.Second)
	c.Send("\r")
	time.Sleep(time.Second)
	c.Send("Y\r")
	// time.Sleep(time.Second)
	// c.Send("\033[B") // down arrow
	// time.Sleep(time.Second)
	// c.Send("\033[A") // up arrow
	time.Sleep(time.Second)
	c.SendLine("\r")
	time.Sleep(time.Second)
	c.SendLine("\r")
	time.Sleep(time.Second)
	c.Send("Y\r")

	err = cmd.Wait()
	if err != nil {
		log.Fatal(err)
	}
}
