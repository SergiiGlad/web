#!/usr/bin/env groovy

/**
 * This pipeline describes a multi container job, running Docker and Golang builds
 */

def label = "jenkins-worker"
def DOCKER_IMAGE_NAME = "sergeyglad/wiki"

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

    stage('Build and unit test Golang app') {
        container('golang') {
        
        echo "Build Golang app"
        sh 'CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o main .'
        }
    }

    stage('Build docker') {
      
      container('docker-dind') {
        
        echo "Docker build image name ${DOCKER_IMAGE_NAME}:${BRANCH_NAME}"

        sh 'docker build . -t ${DOCKER_IMAGE_NAME}:${BRANCH_NAME}'
        
        withDockerRegistry([credentialsId: 'docker-api-key', url: 'https://index.docker.io/v1/']) {
            sh 'docker push ${DOCKER_IMAGE_NAME}:${BRANCH_NAME}'
        }
     
      }
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

   stage('test') {
         
        echo "TEST"

     
      
    }
}// node
} //podTemplate


