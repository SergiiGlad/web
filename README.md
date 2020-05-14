# CI/CD

__Jenkins Pipeline should follow this workflow:__

* Every commit to master branch is a dev release
* Every git tag on a master branch is a QA release

__Production release controlled by a change to production-release.txt file in application repository root, containing a git tag that should be released to production environment__

* Every branch that is not also a PR should have build, test, docker image build, docker image push steps with docker image tag = branch name
* Every PR should have build, test, docker image build, docker image push steps with docker tag = pr-number


## Go Simple Web Server in docker container

### go building example
https://blog.codeship.com/building-minimal-docker-containers-for-go-applications/
We’re disabling cgo which gives us a static binary. We’re also setting the OS to Linux (in case someone builds this on a Mac or Windows) and the -a flag means to rebuild all the packages we’re using, which means all the imports will be rebuilt with cgo disabled. These settings changed in Go 1.4 but I found a workaround in a GitHub Issue. Now we have a static binary! Let’s try it out:

```sh
CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .
```

https://medium.com/@chemidy/create-the-smallest-and-secured-golang-docker-image-based-on-scratch-4752223b7324
But we can optimize it, by removing debug informations and compile only for linux target and disabling cross compilation.
__With go < 1.10__
```
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -ldflags="-w -s" -o /go/bin/hello
```
__With go ≥1.10__
```
RUN GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o /go/bin/hello
```

### docker building
```sh
docker build -t sergeyglad/wiki:1 .

docker push sergeyglad/wiki:1
```

docker run
```sh
docker run --name wiki --rm -d -p 3000:3000 sergeyglad/wiki:1

curl localhost:3000

docker stop wiki
```

https://www.digitalocean.com/community/tutorials/customizing-go-binaries-with-build-tags

# Go (Golang) GOOS and GOARCH

https://gist.github.com/asukakenji/f15ba7e588ac42795f421b48b8aede63#go-golang-goos-and-goarch

## Edit Jenkins password

###### # echo -n password | base64
###### # kubectl edit secret/jenkins
###### # kubectl delete pod/jenkins-name

## A Kubeconfig file to gain access to a Kubernetes cluster


<https://developer.ibm.com/tutorials/configure-a-cicd-pipeline-with-jenkins-on-kubernetes/>

