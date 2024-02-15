package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
    "unicode"
    "strings"
	"time"
)

func main() {
	startTime := time.Now() // Record the start time

    file, err := os.Open("file.txt") 
    if err != nil {
        log.Fatalf("failed to open file: %s", err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)

    // Increase the buffer size for the scanner to handle long lines.
    // The first argument is nil, which means it will use a new buffer.
    // The second argument is the max size. Adjust the size based on your needs.
    const maxBufferSize = 10 * 1024 * 1024 // For example, 10MB
    buffer := make([]byte, maxBufferSize)
    scanner.Buffer(buffer, maxBufferSize)

    wordCount, vowelCount, lineCount, punctuationCount, spaceCount := 0, 0, 0, 0, 0

    for scanner.Scan() {
        lineCount++
        line := scanner.Text()
        
        for _, r := range line {
            switch {
            case unicode.IsSpace(r):
                spaceCount++
            case unicode.IsPunct(r):
                punctuationCount++
            case strings.ContainsRune("aeiouAEIOU", r):
                vowelCount++
            }
        }

        words := strings.Fields(line)
        wordCount += len(words)
    }

    if err := scanner.Err(); err != nil {
        log.Fatalf("error reading file: %s", err)
    }

    fmt.Printf("Total words: %d\n", wordCount)
    fmt.Printf("Total vowels: %d\n", vowelCount)
    fmt.Printf("Total lines: %d\n", lineCount)
    fmt.Printf("Total punctuation marks: %d\n", punctuationCount)
    fmt.Printf("Total spaces: %d\n", spaceCount)



	elapsedTime := time.Since(startTime) // Calculate elapsed time
	fmt.Printf("Execution time: %s\n", elapsedTime)

}
