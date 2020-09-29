package main

import (
	"fmt"
	"os/exec"
	"time"
)

func waitForClipboardChange(clipboardString string) {

	//	writeToClipboard(clipboardString)
	for clipboardString == readClipboard() {
		time.Sleep(3 * time.Second)
	}
}

func clearClipboard() {
	for i := 10; i > 0; i-- {
		time.Sleep(1 * time.Second)
		fmt.Printf("\rClearing clipboard in %d Seconds...", i)
	}
	//writeToClipboard("")
	fmt.Println("\rClipboard cleared.                 ")
}
func writeToClipboard(output string) {
	//wlcopy = "wl-copy"
	//wlpaste = "wl-paste"
	//pasteCmdArgs = []string{wlpaste, "--no-newline"}
	//copyCmdArgs = []string{wlcopy}
	// copyCmd := exec.Command("wl-copy")

	// in, err := copyCmd.StdinPipe()

	// if err != nil {
	// 	log.Fatal(err)
	// }

	// if err := copyCmd.Start(); err != nil {
	// 	log.Fatal(err)
	// }

	// if _, err := in.Write([]byte(output)); err != nil {
	// 	log.Fatal(err)
	// }

	// if err := in.Close(); err != nil {
	// 	log.Fatal(err)
	// }

	// copyCmd.Wait()
	clearClipboard()

}

func readClipboard() string {

	pasteCmd := exec.Command("wl-paste")

	out, err := pasteCmd.Output()
	if err != nil {
		return ""
	}
	return string(out)
}
