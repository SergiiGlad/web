#!/usr/bin/env bash

name="wiki-dev"
ns="develop" 
tag="master"
ver="v0.0.0"
host="dev.184-172-214-143.nip.io"

helm upgrade --install $name  --debug  ./wiki \
    --force \
    --wait \
   --namespace $ns \
            --set image.tag=$tag \
            --set appVer=$tag \
            --set-string ingress.hostName="${name}.${host}" \
            --set-string ingress.tls[0].hosts[0]="${name}.${host}" \
            --set-string ingress.tls[0].secretName="acme-${name}-tls" 
   