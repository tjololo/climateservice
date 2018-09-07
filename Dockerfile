FROM golang:1.11
WORKDIR /go/src/github.com/tjololo/climateservice/
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GO111MODULE=on go build .

FROM scratch
COPY --from=0 /go/src/github.com/tjololo/climateservice/climateservice .
ENTRYPOINT ["/climateservice"]