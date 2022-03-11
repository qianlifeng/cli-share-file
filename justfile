set export

default:
  @just --list

release tag:
  git tag -a $tag
  git push origin $tag