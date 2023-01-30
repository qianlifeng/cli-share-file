file share service that optimized for command line

## Features

- Optimized for shell use, just `curl` and `wget` to share file
- Share file anonymous, no login required
- File will be deleted after 7 days by default

## Start
- ### Docker
   - ```docker pull qianlifeng/tshare```
    
- ### Manual
   - Download executable binary from [release](https://github.com/qianlifeng/tshare/releases) and start it

## Config
- add following shell script to your `~/.zshrc` or `~/.bashrc`
   ```
   share() { curl --progress-bar -F "file=@$1" http://[yourdomain or ip] | tee /dev/null }   
   ```
- use `share [file]` to upload file
- you will get a link that can be downloaded by `wget [link]`

