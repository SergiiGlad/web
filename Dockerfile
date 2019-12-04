FROM scratch
ADD ./wiki /wiki
EXPOSE 3000
ENTRYPOINT ["/wiki"]
