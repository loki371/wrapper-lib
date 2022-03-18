#!/bin/bash
SRC_DIR="./protobuf"
DST_DIR="../micro-service-game/authentication/src/protobuf"

protoc -I=$SRC_DIR --go_out=$DST_DIR $SRC_DIR/*.proto