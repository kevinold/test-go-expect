package main

// time.Sleep(time.Second)
// c.Send("\033[B") // down arrow
// time.Sleep(time.Second)
// c.Send("\033[A") // up arrow

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"time"

	expect "github.com/Netflix/go-expect"
)

func initCreateReactApp() {
	t := time.Now()
	dateStr, err := fmt.Printf("%d%02d%02d%02d%02d",
		t.Year(), t.Month(), t.Day(),
		t.Hour(), t.Minute())
	if err != nil {
		log.Fatal(err)
	}
	projectName := fmt.Sprint(dateStr, "-amplify-cli-cra")
	// Help here as I'm can't get the date formatted the right way in the projectName
	// What I want is YYYY-MM-DD HH:MM without any formatting, so YYYYMMDDHHMM
	fmt.Println(dateStr)     // 20220710230412
	fmt.Println(projectName) // 12-amplify-cli-cra

	c, err := expect.NewConsole(expect.WithStdout(os.Stdout))
	if err != nil {
		log.Fatal(err)
	}
	defer c.Close()

	cmd := exec.Command("npx", "create-react-app@latest", projectName)
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
	exec.Command("cd", projectName)

	err = cmd.Wait()
	if err != nil {
		log.Fatal(err)
	}
}

func amplifyInit() {
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

func amplifyAddApiDSAutoMerge() {
	c, err := expect.NewConsole(expect.WithStdout(os.Stdout))
	if err != nil {
		log.Fatal(err)
	}
	defer c.Close()

	cmd := exec.Command("amplify", "add", "api")
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
	c.Send("\033[A") // up arrow - Conflict detection
	time.Sleep(time.Second)
	c.SendLine("\r")
	time.Sleep(time.Second)
	c.Send("Y\r") // Enable conflict detection? Yes
	time.Sleep(time.Second)
	c.SendLine("\r") // Auto Merge
	time.Sleep(time.Second)
	c.SendLine("\r") // Continue
	time.Sleep(time.Second)
	c.Send("\033[B") // down arrow - One-to-many relationship (e.g., “Blogs” with “Posts” and “Comments”)
	time.Sleep(time.Second)
	c.SendLine("\r") // Continue
	time.Sleep(time.Second)
	c.Send("N\r")

	err = cmd.Wait()
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	if _, err := os.Stat("amplify/team-provider-info.json"); err == nil || os.IsExist(err) {
		amplifyAddApiDSAutoMerge()
	} else {
		initCreateReactApp()
		amplifyInit()
	}
}
