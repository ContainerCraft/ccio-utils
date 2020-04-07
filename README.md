# p1-gotools
golang dev tool for p1 init script utilities
```
git clone git@github.com:containercraft/p1-gotools.git && cd p1-gotools    
```
```
podman run -it \
    --hostname p1-gotools \
    --volume ~/.bashrc:/opt/app-root/src/.bashrc \
    --volume $(pwd):/opt/app-root/src/p1-gotools \
  registry.access.redhat.com/ubi8/go-toolset bash
```
