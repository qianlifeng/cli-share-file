file share service that optimized for command line

## how to use

1. download from release and start it
2. add following shell script to your `~/.zshrc` or `~/.bashrc`
   ```
   share() { curl --progress-bar -F "file=@$1" http://[yourdomain]/u | tee /dev/null }   
   ```
3. use `share [file]` to upload file
4. you will get a link that can be downloaded by `wget [link]`

## features

1. optimized for shell use, just `curl` and `wget` to share file
2. share file anonymous, no login required
3. file will be deleted after 7 days by default