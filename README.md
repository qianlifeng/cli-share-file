File share service that optimized for command line

## Features

- Optimized for shell use, just `curl` and `wget` to share file
- Share file anonymous, no login required
- File will be deleted after 7 days by default

## Start
- ### Docker
   - `docker run -p 3000:3000 -v ./data:/.tshare --restart always qianlifeng/tshare`
   - docker-compose.yml
   
      ```
      version: "3"
      services:
        portainer:
          container_name: tshare
          image: qianlifeng/tshare
          ports:
            - 3000:3000
          volumes:
            - ./data:/.tshare
          restart: always
      ```
    
- ### Manual
   - Download executable binary from [release](https://github.com/qianlifeng/tshare/releases) and start it

## Config
- Add following shell script to your `~/.zshrc` or `~/.bashrc`

   ```
   share() { curl --progress-bar -F "file=@$1" http://[yourdomain or ip] | tee /dev/null }   
   ```
- Use `share [file]` to upload file, you will get a link that can be downloaded by `wget [link]`

