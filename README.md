# ytmusiclist

A Go CLI tool to download YouTube playlists as MP3 files for offline listening using `yt-dlp`.

## Features

- Automatic `yt-dlp` binary download if not found on PATH.
- Download YouTube playlists as MP3 files.
- Configurable output directory.

## Installation

### Option 1: Download Binary
1. Visit the [Releases](https://github.com/yourusername/ytmusiclist/releases) page.
2. Download the appropriate binary for your OS 

### Option 2: Build from Source
1. **Install Go**: Ensure Go is installed (`go version`).
2. **Clone Repository**:
   ```bash
   git clone https://github.com/saullbrandao/ytmusiclist.git
   cd ytmusiclist
   ```
3. **Build**:
   ```bash
   go build -o ytmusiclist
   ```

## Usage


```bash
ytmusiclist 
```
1. If `ffmpeg` and `yt-dlp` are not on PATH the program will try to install them for you.
2. Follow prompts to enter output directory and playlist URL 

## Requirements

- Internet connection for downloading `yt-dlp`, `ffmpeg` and playlists
- `yt-dlp` (automatically downloaded to `~/config/ytmusiclist/bin/` on Linux and `%APPDATA%/ytmusiclist/bin/` on Windows if not on PATH)
- `ffmpeg` (attempts to install it if not on PATH)

## Acknowledgments

- Built with [Go](https://golang.org/).
- Uses [yt-dlp](https://github.com/yt-dlp/yt-dlp) for downloading.
