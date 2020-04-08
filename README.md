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
    --volume $(pwd):/opt/app-root/src \
    --volume ~/.bashrc:/opt/app-root/src/.bashrc \
  registry.access.redhat.com/ubi8/go-toolset bash
```
```
go run main.go
```
