package gofig

import (
	"bufio"
	"os"
	"fmt"
	"log"
	"runtime"
	"os/exec"
	"io"
)

func Input(message string) string {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	w.WriteString(fmt.Sprintf("%s> ", message))
	w.Flush()

	cr := byte('\n')
	answer, err := r.ReadBytes(cr)
	if err != nil {
		log.Fatalln("Could not read input:", err)
	}

	penultimateByte := string(answer[len(answer)-2:len(answer)-1])

	lineSepLength := 0
	if penultimateByte != "\r" {
		lineSepLength = 1
	} else {
		lineSepLength = 2
	}

	return string(answer[:len(answer)-lineSepLength])
}

func SendThroughWriter(message string, writer io.Writer) {
	sendString := message + "\n"
	toSend := make([]byte, len(sendString))
	copy(toSend[:], sendString)
	writer.Write(toSend)
}

func Clear() {
	cmd := new(exec.Cmd)
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else if runtime.GOOS == "linux" {
		cmd = exec.Command("clear")
	} else {
		log.Printf("Cannot clear console, OS unsupported: %s", runtime.GOOS)
		return
	}
	cmd.Stdout = os.Stdout
	cmd.Run()
	cmd.Wait()
}
