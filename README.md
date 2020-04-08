# p1-gotools
golang dev tool for p1 init script utilities
```
git clone git@github.com:containercraft/p1-gotools.git && cd p1-gotools    
```
```
podman run \
    -it \
    --rm \
    --name p1-gotools \
    --hostname p1-gotools \
    --volume $(pwd):/root/dev \
    --volume ~/.ssh:/root/.ssh \
    --volume ~/.gitconfig:/root/.gitconfig \
  containercraft/ccio-golang:ubi8 
```
```
go run main.go
```
