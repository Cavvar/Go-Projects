FROM obraun/vss-protoactor-jenkins as builder
COPY . /app
WORKDIR /app
RUN go build -o treecli/main treecli/main.go

FROM iron/go
COPY --from=builder /app/treecli/main /app/treecli
EXPOSE 1338
ENTRYPOINT [ "/app/treecli" ]
