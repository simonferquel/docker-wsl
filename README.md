# docker-wsl

A very hacky, unstable, highly-untested way to create WSL 2 distros out of Docker images

## Getting the tool

- install go
- run go get github.com/simonferquel/docker-wsl

## Example 1 - hello world

- docker-wsl create hello-world
- wsl -d hello-world -e /hello

## Example 2 - run a redis specialized distro

- docker-wsl create redis
- wsl -d redis -e redis-server