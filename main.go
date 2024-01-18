package main

import (
	"context"
	"log"
	"os"
)

func main() {
	outputDir := "./docs"

	if _, err := os.Stat(outputDir); os.IsNotExist(err) {
		err := os.MkdirAll(outputDir, 0755) // 0755 is a common permission setting (read and execute access for everyone, full access for the owner)
		if err != nil {
			log.Fatalf("failed to create output directory: %v", err)
		}
	}

	f, err := os.Create(outputDir + "/index.html")
	if err != nil {
		log.Fatalf("failed to create output file: %v", err)
	}

	err = home().Render(context.Background(), f)
	if err != nil {
		log.Fatalf("failed to write output file: %v", err)
	} else {
		log.Printf("created output in %s", outputDir)
	}

}
