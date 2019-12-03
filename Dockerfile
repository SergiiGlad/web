FROM scratch
COPY ./webwiki /webwiki
COPY ./*.html /
EXPOSE 3333
ENTRYPOINT ["/webwiki"]
