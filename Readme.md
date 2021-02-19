# Show Screenshot Search App

## Description
Split show video file into screenshots and apply text to images.
Provide search functionality to find parts of the video file.

A screenshot for each 0.2s of video is created.

## Requirements
- Golang
- FFMPEG

## Running
Ensure that the dirs are available by running `./setup.sh`  
Place your video file in `videos/`  
Place your transcript in `transcripts/`
### Main Go
```bash
$ go build main.go video.go
$ ./main -f videos/<filename>
```

### API Server

## Other
Similar to [Frinkiac](https://www.frinkiac.com)
