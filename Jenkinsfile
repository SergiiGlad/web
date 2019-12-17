#!/usr/bin/env groovy

/**
 * This pipeline describes a multi container job, running Maven and Golang builds
 */

podTemplate(yaml: """
apiVersion: v1
kind: Pod
spec:
  containers:
  - name: golang
    image: golang:1.8.0
    command: ['cat']
    tty: true
  - name: docker-dind
    image: docker:stable-dind
    securityContext:
      privileged: true
    env:
      - name: DOCKER_TLS_CERTDIR
        value: ""  
  - name: helm
    image: lachlanevenson/k8s-helm:v2.16.1  
    tty: true
    command:['cat']
 """
  ) {

  node(POD_LABEL) {
    stage('Build Golang project') {
    
      container('golang') {
        echo "Build Golang project"
        sh 'printenv | sort'
      }
    }

    stage('Build docker') {
      
      container('docker') {
        echo "docker build"
      }
    }

  }
}

