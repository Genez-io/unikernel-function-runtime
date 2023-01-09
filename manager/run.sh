#!/bin/bash
go build
sudo \
IMAGE_PATH="/home/cgeorge/genez/unikernel-function-runtime/images" \
KERNEL_PATH="/home/cgeorge/genez/unikernel-function-runtime/manager/deps/kernels" \
FIRECRACKER_PATH="/usr/local/bin/firecracker" \
./manager