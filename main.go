package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

type UserInput struct {
	playlistUrl string
	dirName     string
}

func main() {
	err := ensureFFMPEG()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		fmt.Println("Press any key to exit")
		bufio.NewReader(os.Stdin).ReadLine()
		os.Exit(1)
	}

	ytdlpPath, err := ensureYTDLP()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		fmt.Println("Press any key to exit")
		bufio.NewReader(os.Stdin).ReadLine()
		os.Exit(1)
	}

	for {
		userInput, err := getUserInput()
		if err != nil {
			fmt.Println(err)
			fmt.Println("Press any key to exit")
			bufio.NewReader(os.Stdin).ReadLine()
			os.Exit(1)
		}

		err = downloadPlaylist(ytdlpPath, userInput)
		if err != nil {
			fmt.Println(err)
			fmt.Println("Press any key to exit")
			bufio.NewReader(os.Stdin).ReadLine()
			os.Exit(1)
		}

		fmt.Println("\nPlaylist downloaded!")
	}

}

func getUserInput() (UserInput, error) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("\nEnter the directory name(Leave empty for current directory): ")
	dirName, _ := reader.ReadString('\n')
	dirName = strings.TrimSpace(dirName)

	if dirName == "" {
		dirName = "."
	}

	fmt.Print("Enter YouTube playlist URL: ")
	url, _ := reader.ReadString('\n')
	url = strings.TrimSpace(url)

	if url == "" {
		return UserInput{}, errors.New("Error: URL required")
	}

	return UserInput{playlistUrl: url, dirName: dirName}, nil
}
