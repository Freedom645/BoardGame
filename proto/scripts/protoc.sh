#!/bin/sh

set -xe

SERVER_OUTPUT_DIR=server/chess-backend
CLIENT_OUTPUT_DIR=server/chess-web-frontend/chessGame

protoc --version
protoc --proto_path=proto chessGame.proto \
  --go_out=plugins="grpc:${SERVER_OUTPUT_DIR}" \
  --go_opt=module=github.com/Freedom645/chess-backend \
  --js_out=import_style=commonjs:${CLIENT_OUTPUT_DIR} \
  --grpc-web_out=import_style=typescript,mode=grpcwebtext:${CLIENT_OUTPUT_DIR}
