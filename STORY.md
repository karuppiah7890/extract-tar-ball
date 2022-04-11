https://duckduckgo.com/?t=ffab&q=extract+tar+ball+in+golang&ia=web

https://stackoverflow.com/questions/32254757/extract-from-tar-file-in-go

https://go.dev/play/p/3B7F_axr-i

https://golangdocs.com/tar-gzip-in-golang

```bash
extract-tar-ball $ cp ../download-tar/tce-darwin-amd64-v0.11.0.tar.gz .
extract-tar-ball $ ls
README.md			go.mod				tce-darwin-amd64-v0.11.0.tar.gz
STORY.md			main.go

extract-tar-ball $ go run main.go tce-darwin-amd64-v0.11.0.tar.gz tce-darwin-amd64-v0.11.0

extract-tar-ball $ ls
README.md			go.mod				tce-darwin-amd64-v0.11.0.tar.gz
STORY.md			main.go

extract-tar-ball $ go run main.go tce-darwin-amd64-v0.11.0.tar.gz
2022/04/06 15:14:39 Pass exactly 2 arguments - ./extract-tar-ball <tar-ball-path> <extracted-tar-ball-directory-path>
exit status 1

extract-tar-ball $ go run main.go tce-darwin-amd64-v0.11.0.tar.gz tce-darwin-amd64-v0.11.0
2022/04/06 15:15:36 An error occurred while untarring the file at tce-darwin-amd64-v0.11.0.tar.gz. Error: archive/tar: invalid tar header
exit status 1

```

```bash
extract-tar-ball $ go run main.go tce-darwin-amd64-v0.11.0.tar.gz tce-darwin-amd64-v0.11.0
2022/04/11 19:25:05 An error occurred while untarring the file at tce-darwin-amd64-v0.11.0.tar.gz. Error: archive/tar: invalid tar header
exit status 1
extract-tar-ball $ go run main.go tce-darwin-amd64-v0.11.0.tar.gz tce-darwin-amd64-v0.11.0
extract-tar-ball $ ls tce-darwin-amd64-v0.11.0.tar.gz 
tce-darwin-amd64-v0.11.0.tar.gz
extract-tar-ball $ go run main.go tce-darwin-amd64-v0.11.0.tar.gz tce-darwin-amd64-v0.11.0
extract-tar-ball $ ls
README.md			go.mod				tce-darwin-amd64-v0.11.0
STORY.md			main.go				tce-darwin-amd64-v0.11.0.tar.gz
extract-tar-ball $ ls tce-darwin-amd64-v0.11.0
tce-darwin-amd64-v0.11.0
extract-tar-ball $ ls tce-darwin-amd64-v0.11.0
tce-darwin-amd64-v0.11.0
extract-tar-ball $ file tce-darwin-amd64-v0.11.0
tce-darwin-amd64-v0.11.0: POSIX tar archive (GNU)
extract-tar-ball $ go run main.go tce-darwin-amd64-v0.11.0.tar.gz tce-darwin-amd64-v0.11.0

extract-tar-ball $ 

---

extract-tar-ball $ go run main.go tce-darwin-amd64-v0.11.0.tar.gz tce-darwin-amd64-v0.11.0
extract-tar-ball $ rm -rfv tce-darwin-amd64-v0.11.0
tce-darwin-amd64-v0.11.0
extract-tar-ball $ go run main.go tce-darwin-amd64-v0.11.0.tar.gz tce-darwin-amd64-v0.11.0
extract-tar-ball $ ls tce-darwin-amd64-v0.11.0
tce-darwin-amd64-v0.11.0
extract-tar-ball $ ls tce-darwin-amd64-v0.11.0
tce-darwin-amd64-v0.11.0/        tce-darwin-amd64-v0.11.0.tar.gz  
extract-tar-ball $ ls tce-darwin-amd64-v0.11.0/tce-darwin-amd64-v0.11.0/
default-local	install.sh	tanzu		uninstall.sh
extract-tar-ball $ ls tce-darwin-amd64-v0.11.0/
extract-tar-ball $ 
```
