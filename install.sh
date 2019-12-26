#!/usr/bin/env bash

name="wiki-dev"
ns="develop" 
tag="master"
ver="v0.0.0"
host="dev.184-172-214-143.nip.io"

helm upgrade --install $name  --debug --dry-run ./wiki \
    --force \
    --wait \
    --namespace $ns \
    --set namespace=$ns \
    --set image.tag=$tag \
    --set appVer=$ver \
    --set ingress.hostName=$host \
    --set-string ingress.tls[0].hosts[0]=$host \
    --set-string ingress.tls[0].secretName=acme-$name-tls \
   