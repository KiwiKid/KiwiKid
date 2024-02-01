package main

import (
	"context"
	"log"
	"os"
	"io"
	"time"
	"fmt"
	"path/filepath"
)

const outputDir = "./docs"

func main() {
    outputDir := "./docs"
    staticDir := "static"

    if _, err := os.Stat(outputDir); os.IsNotExist(err) {
        err := os.MkdirAll(outputDir, 0755)
        if err != nil {
            log.Fatalf("failed to create output directory: %v", err)
        }
    }

    f, err := os.Create(outputDir + "/index.html")
    if err != nil {
        log.Fatalf("failed to create output file: %v", err)
    }

    err = home(yearsSinceStarted()).Render(context.Background(), f)
    if err != nil {
        log.Fatalf("failed to write output file: %v", err)
    } else {
        log.Printf("created output in %s", outputDir)
    }

    err = filepath.Walk(staticDir, func(path string, info os.FileInfo, err error) error {
        if err != nil {
            return err
        }

        // Skip the root 'static' directory itself
        if path == staticDir {
            return nil
        }

        // Create corresponding subdirectories in outputDir
        if info.IsDir() {
            return os.MkdirAll(filepath.Join(outputDir, filepath.Base(path)), 0755)
        }

        // Copy files with relative path preservation
        relPath, err := filepath.Rel(staticDir, path)
        if err != nil {
            return err
        }
        destPath := filepath.Join(outputDir, relPath)
        return copyFile(path, destPath)
    })

    if err != nil {
        log.Fatalf("failed to copy static files: %v", err)
    } else {
        log.Printf("copied static files to %s", outputDir)
    }
}




func yearsSinceStarted() string {
	// Define the start date.
	startDate := time.Date(2014, time.February, 1, 0, 0, 0, 0, time.UTC)

	// Get the current date.
	currentDate := time.Now().UTC()

	// Calculate the difference in years.
	years := currentDate.Year() - startDate.Year()

	// If this year's anniversary hasn't occurred yet, subtract one year.
	if currentDate.Before(startDate.AddDate(years, 0, 0)) {
		years--
	}

	// Convert the number of years to a string and return.
	return fmt.Sprintf("%d", years)
}

func copyFile(src, dst string) error {
    in, err := os.Open(src)
    if err != nil {
        return err
    }
    defer in.Close()

    if err = os.MkdirAll(filepath.Dir(dst), 0755); err != nil {
        return err
    }

    out, err := os.Create(dst)
    if err != nil {
        return err
    }
    defer out.Close()

    _, err = io.Copy(out, in)
    if err != nil {
        return err
    }

    return out.Close()
}