#!/bin/bash
go build

OPS_HOME="/home/cgeorge/genez/unikernel-function-runtime/images" \
LIBSTDCPP_PATH="/usr/lib/x86_64-linux-gnu/libstdc++.so.6" \
LIBM_PATH="/usr/lib/x86_64-linux-gnu/libm.so.6" \
LIBGCC_PATH="/usr/lib/x86_64-linux-gnu/libgcc_s.so.1" \
IMAGES_PATH="/home/cgeorge/genez/unikernel-function-runtime/images" \
./java_builder