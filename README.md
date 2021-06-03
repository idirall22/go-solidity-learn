# go-solidity-learn

## Install dependencies

### Install protoc:
1. `curl -LO https://github.com/protocolbuffers/protobuf/releases/download/v3.17.1/protoc-3.17.1-linux-x86_64.zip`
2. `unzip -d protoc protoc-3.17.1-linux-x86_64.zip`
3. `mv protoc/bin/protoc $GOBIN`

### Install SOLC:
1. `sudo add-apt-repository ppa:ethereum/ethereum`
2. `sudo apt-get update`
3. `sudo apt-get install solc`

### Install ABIGEN:
1. `cd $GOPATH/src/github.com/ethereum/go-ethereum/` if you are working without modules.
or
1. `cd $GOPATH/pkg/mod/github.com/ethereum/go-ethereum/`if you are working with modules.
2. `make devtools`
