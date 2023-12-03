FROM golang

WORKDIR /tg-bot

# Download Go modules
COPY go.mod go.sum ./
RUN go mod download

# Copy sources
COPY *.go ./
COPY ./misisapi ./misisapi

# Build
RUN CGO_ENABLED=0 GOOS=linux go build

EXPOSE 801/tcp
CMD [ "./tell-my-server-bot" ]