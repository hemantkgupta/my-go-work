package main

import (
	"errors"
	"fmt"
	"os"
	"testing"
)

// createFile is a helper function called from multiple tests
func createFile(t *testing.T) (_ string, err error) {
	f, err := os.Create("tempFile")
	if err != nil {
		return "", err
	}
	defer func() {
		err = errors.Join(err, f.Close())
	}()
	// write some data to f
	t.Cleanup(func() {
		os.Remove(f.Name())
	})
	return f.Name(), nil
}

func TestFileProcessing(t *testing.T) {
	fName, err := createFile(t)
	if err != nil {
		t.Fatal(err)
	}
	// do testing, don't worry about cleanup
	fmt.Println(fName)
}

// createFile is a helper function called from multiple tests
func createFile2(tempDir string) (_ string, err error) {
	// Use Temp file
	f, err := os.CreateTemp(tempDir, "tempFile")
	if err != nil {
		return "", err
	}
	defer func() {
		err = errors.Join(err, f.Close())
	}()
	// write some data to f
	return f.Name(), nil
}

func TestFileProcessing2(t *testing.T) {
	tempDir := t.TempDir()
	fName, err := createFile2(tempDir)
	if err != nil {
		t.Fatal(err)
	}
	// do testing, don't worry about cleanup
	fmt.Println(fName)
}
