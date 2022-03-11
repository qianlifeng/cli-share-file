set dotenv-load

default:
  @just --list

release:
  goreleaser release