currDir=$(pwd)
GOBIN=~/go/bin
export GOPATH=$GOPATH:$currDir
export PATH=$PATH:$GOBIN
#
mkdir -p src/example
protoc --proto_path=src/ --go_out=src/example --go_opt=paths=source_relative src/msg.proto
go build -o sender src/sender.go src/utils.go
go build -o receiver src/receiver.go src/utils.go
