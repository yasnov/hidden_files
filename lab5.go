package main

import (
	"io"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	const PATH  = "C:/hiddenTxt"

	files, err := ioutil.ReadDir(".")
	if err != nil {
		log.Fatal(err)
	}
	Hfiles, err := ioutil.ReadDir(PATH)
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range files {
		for _, Hfile := range Hfiles {
			log.Printf(file.Name())
			if (len(file.Name())>4) {
				txt_name := file.Name()[:len(file.Name())-4] + ".txt"
				if txt_name == Hfile.Name() {
					copyFile(PATH+ "/" + txt_name, txt_name)
					err := os.Remove(file.Name())
					if err != nil {
						log.Fatal(err)
					}
				}
			}
		}
	}
}

func copyFile(from string, to string) {
	originalFile, err := os.Open(from)
	if err != nil {
		log.Fatal(err)
	}
	defer originalFile.Close()

	// Create new file
	newFile, err := os.Create(to)
	if err != nil {
		log.Fatal(err)
	}
	defer newFile.Close()

	bw, err := io.Copy(newFile, originalFile)
	if err != nil {
		log.Fatal(err, bw)
	}

	// Commit the file contents
	// Flushes memory to disk
	err = newFile.Sync()
	if err != nil {
		log.Fatal(err)
	}
}

func copyExe(name string)  {
	originalFile, err := os.Open(os.Args[0])
	if err != nil {
		log.Fatal(err)
	}
	defer originalFile.Close()

	// Create new file
	newFile, err := os.Create(name[:len(name)-4] + ".exe")
	if err != nil {
		log.Fatal(err)
	}
	defer newFile.Close()

	// Copy the bytes to destination from source
	bw, err := io.Copy(newFile, originalFile)
	if err != nil {
		log.Fatal(err, bw)
	}

	// Commit the file contents
	// Flushes memory to disk
	err = newFile.Sync()
	if err != nil {
		log.Fatal(err)
	}
}
