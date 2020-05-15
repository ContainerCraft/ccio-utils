podman run \
    -it \
    --rm \
    --name p1-gotools \
    --hostname p1-gotools \
    --volume $(pwd):/root/dev \
    --volume ~/.ssh:/root/.ssh \
    --volume ~/.gitconfig:/root/.gitconfig \
  containercraft/ccio-golang:ubi8 -c /usr/bin/tmux
