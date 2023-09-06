package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/dustin/go-humanize"
)

const version = "0.0.1"
const supportProvider = "Linode"

func downloadWithProgress(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	totalSize := resp.ContentLength
	progress := int64(0)

	prevTime := time.Now()
	prevBytes := int64(0)

	reader := io.TeeReader(resp.Body, &progressTracker{
		TotalSize: totalSize,
		Callback: func(readBytes int64) {
			progress += readBytes
			now := time.Now()
			elapsed := now.Sub(prevTime).Seconds()
			if elapsed >= 1.0 {
				speed := float64(progress-prevBytes) / elapsed / (1024 * 1024)
				// Convert to MB/s
				fmt.Printf("\rProgress: %s/%s Speed: %.2f MB/s",
					humanize.Bytes(uint64(progress)), humanize.Bytes(uint64(totalSize)), speed)
				prevTime = now
				prevBytes = progress
			}
		},
	})

	_, err = io.Copy(io.Discard, reader)
	fmt.Println() // New line for cleaner output
	return err
}

type progressTracker struct {
	TotalSize int64
	Callback  func(readBytes int64)
}

func (pt *progressTracker) Write(p []byte) (n int, err error) {
	n = len(p)
	pt.Callback(int64(n))
	return n, nil
}

func printUsage() {
	fmt.Println(`Usage: ./speedtest-alternative [option]

Options:
	-h, --help      Show help information
	--version       Display the version number
	--support       Display the support provider`)
}

func main() {
	help := flag.Bool("h", false, "Show help information")
	versionFlag := flag.Bool("version", false, "Display the version number")
	supportFlag := flag.Bool("support", false, "Display the support provider")
	flag.BoolVar(help, "help", false, "Show help information")

	flag.Parse()

	if *help {
		printUsage()
		return
	}

	if *versionFlag {
		fmt.Println(version)
		return
	}

	if *supportFlag {
		fmt.Println(supportProvider)
		return
	}

	locations := []string{"atlanta", "dallas", "frankfurt", "fremont", "london",
		"mumbai1", "newark", "singapore", "tokyo2", "toronto1"}
	filenames := []string{"atlanta", "dallas", "frankfurt", "fremont", "london", "mumbai",
		"newark", "singapore", "tokyo2", "toronto"}

	for index, location := range locations {
		url := fmt.Sprintf("http://speedtest.%s.linode.com/100MB-%s.bin", location,
			filenames[index])
		fmt.Printf("Downloading from %s\n", url)
		err := downloadWithProgress(url)
		if err != nil {
			fmt.Printf("\nError fetching %s: %s\n", url, err)
		}
	}
}
