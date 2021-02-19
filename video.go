package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func framesFromVideo(filename string) (string, error) {
	fmt.Printf("Extracting frames from %s\n", filename)
	ffmpegPath, err := exec.LookPath("ffmpeg")
	if err != nil {
		fmt.Printf("Error: FFMPEG was not found\n%s\n", err)
		return "", err
	}

	fname := strings.Split(strings.Split(filename, "videos/")[1], ".")[0]
	framesOutPath := "frames/" + fname + "%4d.png"
	buildFfmpegCmd := &exec.Cmd {
		Path: ffmpegPath,
		Args: []string{
			ffmpegPath,
			"-ss",
			"00:00:00",
			"-i",
			filename,
			"-r",
			"5.0",
			framesOutPath,
		},
		Stdout: os.Stdout,
		Stderr: os.Stdout,
	}

	fmt.Printf("Running command: %s\n", buildFfmpegCmd.String())
	if err := buildFfmpegCmd.Run(); err != nil {
		fmt.Printf("Error: FFMPEG command failed\n%s\n", err)
		return "", err
	}
	fmt.Println("Frames generated")
	return "", nil
}
