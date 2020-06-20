@echo off
protoc -I=.  --go_out=plugins=grpc:./ ./test.proto
echo "Update code complete!"
pause