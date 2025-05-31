package dependencies

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"

	"github.com/saullbrandao/ytmusiclist/utils"
)

func EnsureYTDLP() (string, error) {
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
		return "", errors.New(fmt.Sprintf("Error creating the directory \"%s\", %v", userConfigDir, err))
	}

	fileName := "yt-dlp"
	switch runtime.GOOS {
	case "windows":
		fileName += ".exe"
	case "linux":
		fileName += "_linux"
	default:
		fmt.Println("OS not supported")
		utils.GracefulExit()
	}

	// Check if the binary is already present
	ytdlpPath := filepath.Join(binDir, fileName)
	_, err = os.Stat(ytdlpPath)
	if err == nil {
		return ytdlpPath, nil
	}

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
		return "", errors.New(fmt.Sprintf("Error giving execution permission for the yt-dlp binary: %v", err))
	}

	return ytdlpPath, nil
}
