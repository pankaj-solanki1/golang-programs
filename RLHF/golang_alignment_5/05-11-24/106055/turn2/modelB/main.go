package main

import (
	"bufio"
	"compress/gzip"
	"compress/zlib"
	"io"
	"log"
	"os"
	"runtime"
	"sync"

	"github.com/pierrec/lz4"
)

const (
	chunkSize = 1024 * 1024 // 1MB chunk size
)

func archiveData(source string, destination string, compressor string, compressionLevel int) {
	// Check if source file exists
	_, err := os.Stat(source)
	if os.IsNotExist(err) {
		log.Fatalf("Error: Source file '%s' not found.", source)
	}

	// Create the archive file
	archiveFile, err := os.Create(destination)
	if err != nil {
		log.Fatalf("Error creating archive file: %v", err)
	}
	defer archiveFile.Close()

	var writer io.Writer
	switch compressor {
	case "gzip":
		writer = gzip.NewWriter(archiveFile)
		defer writer.(*gzip.Writer).Close()
		writer.(*gzip.Writer).Level = compressionLevel
	case "zlib":
		writer = zlib.NewWriter(archiveFile)
		defer writer.(*zlib.Writer).Close()
		writer.(*zlib.Writer).Level = compressionLevel
	case "lz4":
		writer = lz4.NewWriter(archiveFile)
		defer writer.(*lz4.Writer).Close()
	default:
		log.Fatalf("Invalid compressor: %s", compressor)
		return
	}

	// Read the source file in chunks and compress concurrently
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		if err := compressChunks(source, writer); err != nil {
			log.Fatalf("Error compressing data: %v", err)
		}
	}()

	wg.Wait()

	log.Printf("Archive created successfully: %s (Compressor: %s, Level: %d)", destination, compressor, compressionLevel)
}

func compressChunks(source string, writer io.Writer) error {
	file, err := os.Open(source)
	if err != nil {
		return err
	}
	defer file.Close()

	reader := bufio.NewReader(file)

	for {
		chunk := make([]byte, chunkSize)
		n, err := reader.Read(chunk)
		if n == 0 {
			if err == io.EOF {
				return nil // End of file
			}
			return err // Other error
		}

		if _, err := writer.Write(chunk[:n]); err != nil {
			return err
		}
	}
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU()) // Use all available CPU cores