# arb-protos
proto files for microservices usage

# Getting Started

first clone repository inside of root of project to achive this protos directory:
```
git submodule add https://github.com/GoBootCamp-Group1/arb-protos protos
git submodule update --init --recursive
```

in future if you want to push latest changes in your proto services:

```
# Navigate to the submodule directory
cd protos

# Make changes in the submodule, e.g., edit proto files

# Stage and commit changes in the submodule
git add .
git commit -m "Update proto files"

# Push changes in the submodule
git push origin main  # Replace 'main' with the appropriate branch name

# Navigate back to the main repository
cd ..

# Stage the updated submodule reference and commit it in the main repository
git add protos
git commit -m "Update submodule reference"

# Push changes in the main repository
git push origin main  # Replace 'main' with the appropriate branch name
```


# Proto Generation Guide

This document provides instructions on how to generate Go code from the `.proto` files in this project. The generated code includes both standard Go structs and gRPC service definitions.

## Directory Structure

The project's proto files are organized as follows:

```
project-root/
├── protos/
│   ├── ServiceOne/
│   │   └── service_one.proto
│   ├── ServiceTwo/
│   │   └── service_two.proto
│   └── ...
├── cmd/
└── ...
```

## Prerequisites

Ensure you have `protoc` and the necessary Go plugins installed:

1. **Protocol Buffers Compiler (protoc)**:
   - Download and install `protoc` from the [official Protocol Buffers release page](https://github.com/protocolbuffers/protobuf/releases).

2. **Go Protocol Buffers Plugin**:
   ```sh
   go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
   go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
   ```

3. **Add Go Binaries to PATH**:
   Ensure the Go binaries are in your system's `PATH`. Add the following line to your shell configuration file (e.g., `.bashrc`, `.zshrc`):
   ```sh
   export PATH="$PATH:$(go env GOPATH)/bin"
   ```

## Generating Go Code

Run the following command from the root of your project to generate Go code for all proto files:

```sh
protoc --proto_path=./protos --go_out=./protos --go_opt=paths=source_relative --go-grpc_out=./protos --go-grpc_opt=paths=source_relative ./protos/*.proto
```

### Command Breakdown

- `--proto_path=./protos`: Specifies the root directory where the `.proto` files are located.
- `--go_out=./protos`: Specifies the output directory for the generated Go code.
- `--go_opt=paths=source_relative`: Ensures that the generated Go files preserve the directory structure relative to the source.
- `--go-grpc_out=./protos`: Specifies the output directory for the generated gRPC Go code.
- `--go-grpc_opt=paths=source_relative`: Ensures that the generated gRPC Go files preserve the directory structure relative to the source.
- `./protos/*/*.proto`: Specifies the pattern to match all `.proto` files within the service directories.

## Using Makefile

To simplify the code generation process, you can use the provided Makefile. Simply run:

```sh
make proto
```

### Makefile Content

```Makefile
.PHONY: proto
proto:
	protoc --proto_path=./protos --go_out=./protos --go_opt=paths=source_relative --go-grpc_out=./protos --go-grpc_opt=paths=source_relative ./protos/*/*.proto
```

## Contributing

If you add new `.proto` files or modify existing ones, make sure to run the `protoc` command to regenerate the Go code and commit the changes.

## Troubleshooting

If you encounter any issues, ensure that `protoc` and the Go plugins are correctly installed and that the Go binaries are in your system's `PATH`.

For further assistance, refer to the [Protocol Buffers documentation](https://developers.google.com/protocol-buffers) and the [gRPC documentation](https://grpc.io/docs/).

---

By following these instructions, you can efficiently manage and generate Go code from your proto files, ensuring consistent and reliable inter-service communication in your project.
