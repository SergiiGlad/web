#!/usr/bin/env bash

name="wiki-dev"
ns="develop" 
tag="master"
ver="v0.0.0"
host="184-172-214-143.nip.io"

helm upgrade --install $name --dry-run --debug  ./wiki \
    --namespace $ns \
    --set image.tag=$tag \
    --set appVer=$ver \
    --set ingress.hostName=$name.$host \
    --set-string ingress.tls[0].hosts[0]=$name.$host \
    --set-string ingress.tls[0].secretName=acme-$name-tls \
   