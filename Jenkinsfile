#!/usr/bin/env groovy
pipeline {
    
    stages {
        stage('Build') {
            steps {
                container('jnlp-slave-docker') {
                    sh 'echo Building'
                    sh 'docker info'
                }    
            }
        }
        stage('Test') {
            steps {
                container('jnlp-slave-docker') {
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
