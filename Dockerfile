FROM scratch
ADD ./webwiki /webwiki
EXPOSE 3000
ENTRYPOINT ["/webwiki"]
