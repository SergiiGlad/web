#!/usr/bin/env groovy

/**
 * This pipeline describes a multi container job, running Docker and Golang builds
 */

def label = "jenkins-worker"
podTemplate(label: label, yaml: """
apiVersion: v1
kind: Pod
spec:
  containers:
  - name: golang
    image: golang:1.13.0-alpine
    command:
      - "cat"
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
    command:
      - "cat"
 """
  ) {

  node(label) {
    stage('Checkout SCM') {
        checkout scm
    } 

    stage('Build and test Golang app') {
        container('golang') {
        
        echo "Build Golang app"
        sh 'CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o main .'
        }
    }

    stage('Build docker') {
      
      container('docker') {
        echo "docker build"
     
      }
    }

    stage('test') {
         
        echo "TEST+++++++++++"

        sh 'echo $BRANCH_NAME'
     
      
    }

    stage('Deploy') {
        container('helm') {
               // isPRMergeBuild
            if ( env.BRANCH_NAME ==~  /^PR-\d+$/ ) {
                sh 'echo It is pull request'
                // is Push to master    
            } else if (env.BRANCH_NAME ==~  /^master$/) {
                sh 'echo Its push to master '
            // isTag    
            } else if (env.BRANCH_NAME =~ /^v\d.\d.\d$/ ){
                sh 'echo qa release with tag : $(BRANCH_NAME)'
            // Other operation    
            } else {
                sh 'echo push to other branch $(BRANCH_NAME)'
            }

        }
    
   } 
}// node
} //podTemplate


