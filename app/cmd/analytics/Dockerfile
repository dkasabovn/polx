# keep the golang version in sync with the .tool-version file
FROM polx/builder AS builder

FROM scratch

COPY --from=builder /go/bin/analytics .
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

COPY --from=builder /etc/passwd /etc/passwd
COPY .env .env
USER scratchuser

CMD ["./analytics"]