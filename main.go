package main

import (
	"archive/tar"
	"compress/gzip"
	"io"
	"log"
	"os"
	"path/filepath"
)

// Extracts .tar.gzip and .tar.gz files
func ExtractTarball(targz, target string) error {
	targzReader, err := os.Open(targz)
	if err != nil {
		return err
	}
	defer targzReader.Close()

	tarArchive, err := gzip.NewReader(targzReader)
	if err != nil {
		return err
	}
	defer tarArchive.Close()

	tarReader := tar.NewReader(tarArchive)

	for {
		header, err := tarReader.Next()
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}

		path := filepath.Join(target, header.Name)
		info := header.FileInfo()
		if info.IsDir() {
			if err = os.MkdirAll(path, info.Mode()); err != nil {
				return err
			}
			continue
		}

		file, err := os.OpenFile(path, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, info.Mode())
		if err != nil {
			return err
		}
		defer file.Close()
		_, err = io.Copy(file, tarReader)
		if err != nil {
			return err
		}
	}
	return nil
}

func main() {
	if len(os.Args) != 3 {
		log.Fatal("Pass exactly 2 arguments - ./extract-tar-ball <tar-ball-path> <extracted-tar-ball-directory-path>")
	}
	tarball := os.Args[1]
	destination := os.Args[2]

	err := ExtractTarball(tarball, destination)
	if err != nil {
		log.Fatalf("An error occurred while extract tar file at %s. Error: %s", "temp.tar", err)
	}
}
