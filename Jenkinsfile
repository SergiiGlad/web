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
  - name: kubectl
    image: lachlanevenson/k8s-kubectl:v1.16.4
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

    stage('Build and push image') {
      
      container('docker-dind') {
        
        // Environment variables DOCKER_IMAGE_NAME  set by Jenkins plugins 
        echo "Docker build image name ${DOCKER_IMAGE_NAME}:${BRANCH_NAME}"

        sh 'docker build . -t ${DOCKER_IMAGE_NAME}:${BRANCH_NAME}'
        
        withDockerRegistry([credentialsId: 'docker-api-key', url: 'https://index.docker.io/v1/']) {
            sh 'docker push ${DOCKER_IMAGE_NAME}:${BRANCH_NAME}'
        }
     
      }
    }


    stage('Deploy') {
          
                echo "Deploying...."
                container('helm') {
                 withKubeConfig([credentialsId: 'kubeconfig']) {
                    sh 'helm version'
                 } 
              
            }
        }


    
    stage('Deploy') {
            container('kubectl') {
 
            if ( isPullRequest() ) {
                // isPRMergeBuild
               echo "It is pull request"
               echo "Every PR should have build, test, docker image build, docker image push steps with docker tag = pr-number"
               echo "docker image ${DOCKER_IMAGE_NAME}:${BRANCH_NAME} has push"
            } else if ( isMaster() ) {
               // is Push to master
               echo "Its push to master"
               echo "Every commit to master branch is a dev release"

               // deploy dev release  
               devRelease()
            } else if ( isBuildingTag() ){
                echo "Every git tag on a master branch is a QA release" 
                
                // deploy to test env

                // deployToQA()
                // integrationTest()

             
            } else {
                // Other operation   
                sh 'echo push to other branch $(BRANCH_NAME)'
            }

        }
    
   } 

   stage('test') {
         
        echo "TEST"

     
      
    }
}// node
} //podTemplate

def isMaster() {
    return (env.BRANCH_NAME ==~  /^master$/)
}

def isPullRequest() {
    return (env.BRANCH_NAME ==~  /^PR-\d+$/)
}

def isBuildingTag() {
    return ( env.BRANCH_NAME ==~ /^v\d.\d.\d$/ )    
}

def devRelease() {
    stage ('Dev release') {
    withKubeConfig([credentialsId: 'kubeconfig']) {
                    sh 'kubectl rollout restart deploy/wiki-dev -n jenkins'
                }   
    }            
}

