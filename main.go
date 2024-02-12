package main

import (
	"encoding/gob"
	"fmt"
	"os"
)

const (
	cifarDataDir = "data/cifar-10-batches-py/"
)

func main() {
	// Load CIFAR dataset
	trainData, trainLabels := loadCIFAR(cifarDataDir + "data_batch_1")

	// Preprocess the data
	trainImages := preprocess(trainData)

	// Print some information about the dataset
	fmt.Printf("Number of training images: %d\n", len(trainImages))
	fmt.Printf("Image dimensions: %dx%d\n", len(trainImages[0]), len(trainImages[0][0]))
	fmt.Printf("Number of training labels: %d\n", len(trainLabels))
}

func loadCIFAR(filename string) ([][]uint8, []uint8) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file:", err)
		os.Exit(1)
	}
	defer file.Close()

	decoder := gob.NewDecoder(file)

	var batch Batch
	err = decoder.Decode(&batch)
	if err != nil {
		fmt.Println("Error decoding file:", err)
		os.Exit(1)
	}

	return batch.Images, batch.Labels
}

func preprocess(images [][]uint8) [][][]uint8 {
	preprocessed := make([][][]uint8, len(images))
	for i := range images {
		image := make([][]uint8, 32)
		for j := range image {
			image[j] = make([]uint8, 32)
			copy(image[j], images[i][j*32:(j+1)*32])
		}
		preprocessed[i] = image
	}

	return preprocessed
}

type Batch struct {
	Labels []uint8
	Images [][]uint8
}
