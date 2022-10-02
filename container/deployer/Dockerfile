FROM gcr.io/cloud-builders/gcloud

ARG KO_VER

RUN apt update && \
    add-apt-repository ppa:longsleep/golang-backports && \
    apt update && \
    apt install -y golang-go
ENV GOBIN /root/go/bin
ENV PATH $PATH:$GOBIN
RUN go install github.com/google/ko@${KO_VER}
