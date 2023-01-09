#!/bin/bash

# Install kotlin compiler
curl -LJO https://github.com/JetBrains/kotlin/releases/download/v1.7.21/kotlin-compiler-1.7.21.zip
unzip kotlin-compiler-1.7.21.zip
rm kotlin-compiler-1.7.21.zip

# Install nanos tool
curl https://ops.city/get.sh -sSfL | sh
mkdir ops
cp ~/.ops/bin/ops ./ops/

# Grab OSv repo
mkdir osv
# todo: grab osv kernel