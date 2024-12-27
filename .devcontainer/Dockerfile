# Base image
FROM mcr.microsoft.com/devcontainers/go:1-1.23-bullseye as development

# Install dependencies
RUN apt-get update 

# Install Protobuf
RUN curl -LO https://github.com/protocolbuffers/protobuf/releases/download/v27.1/protoc-27.1-linux-x86_64.zip && \
    unzip protoc-27.1-linux-x86_64.zip -d /usr/local && \
    export PATH="$PATH:/usr/local" && \
    protoc --version

# Install Go tools
RUN GOOS=linux go install go.uber.org/mock/mockgen@latest && \
    GOOS=linux go install golang.org/x/tools/go/analysis/passes/fieldalignment/cmd/fieldalignment@latest

# Install golangci-lint
RUN curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.59.1 && \
    golangci-lint --version

# Install protoc-gen-go and protoc-gen-go-grpc
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@latest \
    && go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

# Set Git to ignore file mode changes
RUN git config --global core.fileMode false
RUN git config --global core.autocrlf true

# # Ensure correct permissions inside the container
RUN chmod -R a+w /go/pkg


