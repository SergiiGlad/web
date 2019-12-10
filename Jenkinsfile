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
            sh 'go build'
          }
        }
        stage('Build Dockerfile') {
            steps {
                container('docker') {
                    sh 'echo Building'
                    sh 'docker info'
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
