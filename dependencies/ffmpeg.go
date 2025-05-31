package dependencies

import (
	"errors"
	"fmt"
	"github.com/saullbrandao/ytmusiclist/utils"
	"os"
	"os/exec"
	"runtime"
)

func EnsureFFMPEG() error {
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
		utils.GracefulExit()
	case "linux":
		err = installFFmpegLinux()
	default:
		fmt.Println("\nOS not supported")
		fmt.Println("Press any key to exit")
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

func installFFmpegLinux() error {
	// Only arch is being supported right now
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
