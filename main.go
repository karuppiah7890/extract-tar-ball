package main

import (
	"archive/tar"
	"compress/gzip"
	"io"
	"log"
	"os"
	"path/filepath"
)

// Extracts .gzip files.
// Extracts .tar.gzip and .tar.gz files to .tar
func UnGzip(source, target string) error {
	reader, err := os.Open(source)
	if err != nil {
		return err
	}
	defer reader.Close()

	archive, err := gzip.NewReader(reader)
	if err != nil {
		return err
	}
	defer archive.Close()

	writer, err := os.Create(target)
	if err != nil {
		return err
	}
	defer writer.Close()

	_, err = io.Copy(writer, archive)
	return err
}

// Extracts .tar files
func Untar(tarball, target string) error {
	reader, err := os.Open(tarball)
	if err != nil {
		return err
	}
	defer reader.Close()
	tarReader := tar.NewReader(reader)

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
	err := UnGzip(tarball, "temp.tar")
	if err != nil {
		log.Fatalf("An error occurred while extracting the gzip / gz file at %s. Error: %s", tarball, err)
	}

	err = Untar("temp.tar", destination)
	if err != nil {
		log.Fatalf("An error occurred while extract tar file at %s. Error: %s", "temp.tar", err)
	}
}
