FROM registry.access.redhat.com/ubi9/go-toolset:1.21 as build
WORKDIR /app
COPY . .
USER 0
RUN make bin

FROM registry.access.redhat.com/ubi9/ubi-micro
WORKDIR /app
COPY --from=build /app/bin/flightctl-acm-addon .
CMD ["/app/flightctl-acm-addon"]
