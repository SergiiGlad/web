#!/usr/bin/env groovy
pipeline {
    agent{
      kubernetes {
       yamlFile 'podTemplWorker.yaml'
    }
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
                    sh 'docker build -t sergeyglad/wiki .'

                    withDockerRegistry([credentialsId: 'docker-api-key', url: 'https://index.docker.io/v1/']) {
                        sh 'docker push sergeyglad/wiki'
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
            }
        }
    }
}
