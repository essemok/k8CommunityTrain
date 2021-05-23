FROM scratch

ENV SERVICE_PORT 8000

EXPOSE $SERVICE_PORT

COPY k8Community /

CMD ["/k8Community"]
