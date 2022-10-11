ARG NODE_VERSION

FROM gcr.io/cloud-builders/npm:node-${NODE_VERSION}

ARG GO_VERSION
ARG AQUA_VERSION
ARG PROTOC_VERSION

RUN apt update && apt install -y wget && \
    curl -sSfL https://raw.githubusercontent.com/aquaproj/aqua-installer/v1.1.0/aqua-installer \
        | sh -s -- -i /usr/local/bin/aqua -v $AQUA_VERSION && \
    wget https://github.com/protocolbuffers/protobuf/releases/download/${PROTOC_VERSION}/protoc-$(echo ${PROTOC_VERSION} | sed 's/v//')-linux-x86_64.zip \
        -O protobuf.zip && \
    unzip protobuf.zip -d /usr/local/bin/protobuf && \
    rm protobuf.zip && \
    chmod -R 755 /usr/local/bin/protobuf/* && \
    wget https://go.dev/dl/go${GO_VERSION}.linux-amd64.tar.gz && \
    rm -rf /usr/local/go && \
    tar -C /usr/local -xzf go${GO_VERSION}.linux-amd64.tar.gz

ENV GOBIN /usr/local/go/bin

ENV PATH $PATH:$GOBIN:/usr/local/bin/protobuf/bin:/builder/home/.local/share/aquaproj-aqua/bin
