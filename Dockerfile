FROM scratch

COPY ./improvised /

EXPOSE 8080

ENTRYPOINT ["/improvised"]