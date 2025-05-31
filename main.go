package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/saullbrandao/ytmusiclist/dependencies"
	"github.com/saullbrandao/ytmusiclist/utils"
)

type UserInput struct {
	playlistUrl string
	dirName     string
}

func main() {
	err := dependencies.EnsureFFMPEG()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		utils.GracefulExit()
	}

	ytdlpPath, err := dependencies.EnsureYTDLP()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		utils.GracefulExit()
	}

	for {
		userInput, err := getUserInput()
		if err != nil {
			fmt.Println(err)
			utils.GracefulExit()
		}

		err = downloadPlaylist(ytdlpPath, userInput)
		if err != nil {
			fmt.Println(err)
			utils.GracefulExit()
		}

		fmt.Println("\nPlaylist downloaded!")
	}

}

func getUserInput() (UserInput, error) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("\nEnter the directory name you want to use(Leave empty for current directory): ")
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
