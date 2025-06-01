FROM golang:1.24 as build

WORKDIR /go/src/app
COPY . .

RUN go mod download
RUN go build -o /bin/server ./cmd/sudoku-solver

# FROM busybox as busybox

FROM gcr.io/distroless/base


# Adding shell and shell commands for debugging
# COPY --from=busybox /bin/sh /bin/sh
# COPY --from=busybox /bin/ls /bin/ls
# COPY --from=busybox /bin/cat /bin/cat


WORKDIR /src

COPY --from=build /bin/server /src/server


ENTRYPOINT [ "/src/server" ]
