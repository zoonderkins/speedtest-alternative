## What is this?

This is a simple speedtest tool of collection of VPS providers.

## Todo

- [x] Linode
- [ ] Vultr
- [ ] DigitalOcean
- [ ] Hetzner
- [ ] Alwyzon
- [ ] Contabo
- [ ] HostHatch
- [ ] AWS
- [ ] GCP
- [ ] Azure

## Installation

```bash
# Init module
go mod init speedtest-alternative

# Install dependencies
go get github.com/dustin/go-humanize

```

## Build

### Via Makefile

```bash
# to build for the current platform.
make build 
# to cross-compile for Linux with AMD64.
make build-linux-amd64 
# to cross-compile for Linux with ARM64.
make build-darwin-amd64
# to cross-compile for macOS with ARM64 (M1/M2 chips).
make build-darwin-arm64 
```


### Via command line

```bash
GOOS=linux GOARCH=amd64 go build -o build/linode-speedtest_amd64 ./app.go

GOOS=darwin GOARCH=amd64 go build -o build/linode-speedtest_darwin_amd64 ./app.go

GOOS=darwin GOARCH=arm64 go build -o build/linode-speedtest_darwin_arm64 ./app.go

```

## Usage

```bash

--help     Displays usage instructions.
--version  Shows the version number.
--support  Displays the support provider.

```

## License

Apache License 2.0