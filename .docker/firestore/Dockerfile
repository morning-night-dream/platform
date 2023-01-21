ARG NODE_VERSION

FROM node:${NODE_VERSION}-alpine

WORKDIR /firestore

RUN apk add openjdk11 && \
    npm install -g firebase-tools

EXPOSE 8080

CMD ["firebase", "emulators:start", "--project=emulator"]
