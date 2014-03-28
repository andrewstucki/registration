package utils

import (
	"fmt"
	"os/exec"
	"runtime"
)

func YesNo(prompt string) (answer bool) {
	var input string
	fmt.Printf("%s [y/n]: ", prompt)
	fmt.Scanf("%s", &input)
	for (input != "Y") && (input != "N") && (input != "y") && (input != "n") {
		fmt.Printf("Please answer 'y' or 'n': ")
		fmt.Scanf("%s", &input)
	}
	return (input == "Y") || (input == "y")
}

func OpenSite(url string) (err error) {
	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("cmd", "/c", "start", url)
		break
	case "darwin":
		cmd = exec.Command("open", url)
		break
	}
	return cmd.Run()
}
