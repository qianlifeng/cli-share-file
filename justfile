set export

default:
  @just --list

# just release 0.1.3
release tag:
  # build amd64 on mac m1 chip, refer: https://prinsss.github.io/build-x86-docker-images-on-an-m1-macs/
  # need to run follow commands first
  # - docker buildx create --use --name m1_builder
  # - docker buildx inspect --bootstrap
  docker buildx build --platform linux/amd64 -t qianlifeng/tshare -t qianlifeng/tshare:$tag . --push
  git tag -a v$tag -m v$tag
  git push origin v$tag