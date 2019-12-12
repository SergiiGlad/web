#!/usr/bin/env groovy
pipeline {
    agent{
      kubernetes {
       yamlFile 'podTemplWorker.yaml'
    }
  }
    environment {
            //be sure to replace "sergiiglad" with your own Docker Hub username
            DOCKER_IMAGE_NAME = "sergiiglad/wiki"
        }
 
    stages {
        stage('Build Golang project') {
          steps{
            sh 'echo "go test"'
          }
        }
        stage('Build Dockerfile') {
            steps {
                container('docker') {
                    sh 'echo Building Dockerfile'

                    //docker.Image.build
                    sh 'docker build -t ${DOCKER_IMAGE_NAME} .'

                    withDockerRegistry([credentialsId: 'docker-api-key', url: 'https://index.docker.io/v1/']) {
                        sh 'docker push ${DOCKER_IMAGE_NAME}'
                    }
                    
                    
                }    
            }
        }
        stage('Test') {
            steps {
                container('docker') {
                    echo 'go testing..'
                }    
            }
        }
        stage('Deploy') {
            steps {
                echo 'Deploying....'
                container('helm') {
                 withKubeConfig([credentialsId: 'kubeconfig']) {
                 sh 'helm version'
                 } 
               }
            }
        }
    }
}
