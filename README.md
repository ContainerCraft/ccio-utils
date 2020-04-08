# p1-gotools
golang dev tool for p1 init script utilities
```
git clone git@github.com:containercraft/p1-gotools.git && chmod 755 -R p1-gotools && cd p1-gotools    
```
```
podman run \
    -it \
    --rm \
    --name p1-gotools \
    --hostname p1-gotools \
    --volume $(pwd):/root/dev \
    --volume ~/.ssh:/root/.ssh \
  containercraft/ccio-golang bash
```
```
go run main.go
```
