package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
)

type UserInput struct {
	playlistUrl string
	dirName     string
}

func main() {
	ytdlpPath, err := ensureYTDLP()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	for {
		userInput, err := getUserInput()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		err = downloadPlaylist(ytdlpPath, userInput)
		if err != nil {
			fmt.Println(err)
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

func downloadPlaylist(ytdlpPath string, userInput UserInput) error {
	if err := os.MkdirAll(userInput.dirName, 0755); err != nil {
		return err
	}
	cmd := exec.Command(ytdlpPath, "-x", "--audio-format", "mp3", "--download-archive", userInput.dirName+"/downloaded.txt", "-o", userInput.dirName+"/%(title)s [%(id)s].%(ext)s", "--no-post-overwrites", userInput.playlistUrl)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func ensureYTDLP() (string, error) {
	// Check if yt-dlp is on PATH
	path, err := exec.LookPath("yt-dlp")
	if err == nil {
		return path, nil
	}

	// Create a directory to store the yt-dlp binary
	binDir := filepath.Join(os.Getenv("HOME"), ".ytmusiclist", "bin")
	err = os.MkdirAll(binDir, 0755)
	if err != nil {
		return "", err
	}

	// Set the correct binary for each OS
	binName := "yt-dlp"
	ytdlpPath := filepath.Join(binDir)
	switch runtime.GOOS {
	case "windows":
		binName += ".exe"
	case "darwin":
		binName += "_macos"
	case "linux":
		binName += "_linux"
	default:
		fmt.Println("OS not supported")
		os.Exit(1)
	}

	// Check if the binary is already present
	ytdlpPath = filepath.Join(ytdlpPath, binName)
	_, err = os.Stat(ytdlpPath)
	if err == nil {
		return ytdlpPath, nil
	}

	// Download the correct binary
	const YTDLP_URL = "https://github.com/yt-dlp/yt-dlp/releases/latest/download/"
	resp, err := http.Get(YTDLP_URL + binName)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	out, err := os.Create(ytdlpPath)
	if err != nil {
		return "", err
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return "", err
	}

	err = out.Chmod(0755)
	if err != nil {
		return "", err
	}

	return ytdlpPath, nil
}
