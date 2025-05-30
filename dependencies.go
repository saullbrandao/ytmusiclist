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
)

func ensureFFMPEG() error {
	// Check if ffmpeg is available
	_, err := exec.LookPath("ffmpeg")
	if err == nil {
		return nil
	}

	// Download the correct binary based on the OS
	switch runtime.GOOS {
	case "windows":
		err = installFFmpegWindows()
		fmt.Println("\nYou should restart the app to be able to use ffmpeg now")
		fmt.Println("Press any key to exit and open the program again")
		bufio.NewReader(os.Stdin).ReadLine()
		os.Exit(1)
	case "linux":
		err = installFFmpegLinux()
	default:
		fmt.Println("\nOS not supported")
		fmt.Println("Press any key to exit")
		bufio.NewReader(os.Stdin).ReadLine()
		os.Exit(1)
	}
	if err != nil {
		return errors.New(fmt.Sprintf("Error installing ffmpeg: %v", err))
	}

	_, err = exec.LookPath("ffmpeg")
	if err != nil {
		return errors.New("ffmpeg not found after install")
	}

	fmt.Println("FFmpeg installed!")

	return nil
}
func ensureYTDLP() (string, error) {
	// Check if yt-dlp is on PATH
	path, err := exec.LookPath("yt-dlp")
	if err == nil {
		return path, nil
	}

	// Create a directory to store the yt-dlp binary
	userConfigDir, err := os.UserConfigDir()
	if err != nil {
		return "", err
	}

	binDir := filepath.Join(userConfigDir, "ytmusiclist", "bin")
	err = os.MkdirAll(binDir, 0755)
	if err != nil {
		return "", errors.New(fmt.Sprintf("Error creating the directory \"%s\", %s", userConfigDir, err))
	}

	// Set the correct binary for each OS
	fileName := "yt-dlp"
	switch runtime.GOOS {
	case "windows":
		fileName += ".exe"
	case "linux":
		fileName += "_linux"
	default:
		fmt.Println("OS not supported")
		fmt.Println("Press any key to exit")
		bufio.NewReader(os.Stdin).ReadLine()
		os.Exit(1)
	}

	// Check if the binary is already present
	ytdlpPath := filepath.Join(binDir, fileName)
	_, err = os.Stat(ytdlpPath)
	if err == nil {
		return ytdlpPath, nil
	}

	// Download the correct binary
	const YTDLP_URL = "https://github.com/yt-dlp/yt-dlp/releases/latest/download/"
	resp, err := http.Get(YTDLP_URL + fileName)
	if err != nil {
		return "", errors.New(fmt.Sprintf("Error downloading the yt-dlp binary: %v", err))
	}
	defer resp.Body.Close()

	out, err := os.Create(ytdlpPath)
	if err != nil {
		return "", errors.New(fmt.Sprintf("Error creating the yt-dlp file: %v", err))
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return "", errors.New(fmt.Sprintf("Error copying the yt-dlp file: %v", err))
	}

	err = out.Chmod(0755)
	if err != nil {
		return "", errors.New(fmt.Sprintf("Error giving execution permissin for the yt-dlp binary: %v", err))
	}

	return ytdlpPath, nil
}

func installFFmpegLinux() error {
	cmd := exec.Command("sudo", "pacman", "-S", "--noconfirm", "ffmpeg")
	return installFFmpeg(cmd)
}

func installFFmpegWindows() error {
	cmd := exec.Command("winget", "install", "-e", "--id", "Gyan.FFmpeg")
	return installFFmpeg(cmd)
}

func installFFmpeg(cmd *exec.Cmd) error {
	fmt.Println("Attempting to install ffmpeg")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		return errors.New("Failed to install ffmpeg. You have to install ffmpeg manually.")
	}

	return nil
}
