#!/bin/bash
SRC_DIR="src"
DST_DIR="build"

protoc -I=$SRC_DIR --go_out=$DST_DIR $SRC_DIR/*.proto