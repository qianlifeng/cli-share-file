set export

default:
  @just --list

release tag:
  git tag -a $tag -m $tag
  git push origin $tag