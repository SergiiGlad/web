#!/usr/bin/env groovy
pipeline {
     agent{
        kubernetes {
            label 'jenkins-slave-docker'
            label jenkins: "slave"
            label jenkins/jenkins-jenkins-slave: "true"
            defaultContainer 'jenkins/jnlp-slave:3.27-1'
        }
    }

    stages {
        stage('Build') {
            steps {
                echo 'Building..'
            }
        }
        stage('Test') {
            steps {
                echo 'Testing..'
            }
        }
        stage('Deploy') {
            steps {
                echo 'Deploying....'
            }
        }
    }
}
