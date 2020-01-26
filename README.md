# GET EXECUTABLE BINARY FILES
```
https://yt-dl.org/downloads/latest/youtube-dl
https://www.ffmpeg.org/download.html
```

# MAKE EXECUTABLE BINARY FILES
```
sudo chmod a+rx /usr/local/bin/youtube-dl
sudo chmod a+rx /usr/local/bin/ffmpeg
```

# BUILD THIS FILE
```
go build downloader.go
sudo cp -r downloader /usr/local/bin/
sudo chmod a+rx /usr/local/bin/downloader
```
