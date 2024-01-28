# Simulator gRPC definitions

This repository holds the API description of the opencl-go simulator.

**This is work in progress**

**The repository is archived and unmaintained.
Please see [Maintenance Notice](https://github.com/opencl-go/opencl-go.github.io/discussions/25) for further details.**

## Update definitions

### Install compiler (once)

Install `protoc` and the Go plugin.
See here: https://developers.google.com/protocol-buffers/docs/gotutorial .

### Update Go code

run
```sh
go generate generate.go
```
(The compiler and Go plugins may need to be reachable through `PATH`.)


## License

This project is based on the MIT License. See `LICENSE` file.
