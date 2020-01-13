#!/usr/bin/env groovy

/**
 * This pipeline describes a CI/CD process for running Golang app to multi stages environment
 */

def podLabel = "jenkins-worker-${UUID.randomUUID().toString()}"
def host = "173-193-102-57.nip.io"
def dockerImage = 'sergeyglad/wiki'


golangTemplate(podLabel) {
  node(podLable) {

    stage('Checkout SCM') {
        checkout scm
    }

    stage('Build  Golang app') {
        container('golang') {
            echo "Build Golang app"
            sh 'CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -ldflags="-w -s" -o main .'
        }
    }

    stage ('Unit test Golang app')  {
        container('golang') {
            echo "Unit test Golang app"
            sh 'CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go test -v .'
        }
    }
      
    //
    // BRANCH_NAME = master  - push to master
    // BRANCH_NAME = PR-1    - pull request
    // BRANCH_NAME = develop - push to other branch
    // BRANCH_NAME = 0.0.1  - git tag
    //

    GIT_COMMIT = sh(returnStdout: true, script: "git rev-parse HEAD").trim()
    echo "GIT_COMMIT: $GIT_COMMIT"
    
    def shortCommit = GIT_COMMIT.take(7)
    echo "shortCommit: $shortCommit"

    def dockerTag = env.BRANCH_NAME 
    
   // if ( isMaster() ) { dockerTag = shortCommit}

    echo "dockerTag: $dockerTag"

    stage('Docker build') {
      container('docker-dind') {
           sh """
              docker build . -t $dockerImage:$dockerTag    
           """
           //
        }
    }

    if ( isPullRequest() ) {
        // exitAsSuccess()
        echo "It's pull request and we don't push image to docker hub"
        currentBuild.result = 'SUCCESS';  
        return 0
    }
   
    stage ('Docker push') {
        container('docker-dind') {

          sh 'docker image ls'
          withDockerRegistry([credentialsId: 'docker-api-key', url: 'https://index.docker.io/v1/']) {
                sh """
                    docker push $dockerImage:$dockerTag
                """
                
          }
        }
    }

    if ( ! isMaster() && ! isBuildingTag() ) {
        // exitAsSuccess()
        currentBuild.result = 'SUCCESS';  
        return 0
    }

    
    stage('Deploy') {
     build job: 'web-delivery', wait: true, 
     parameters: [string(name: 'BRANCH_NAME', value: env.BRANCH_NAME),
                  string(name: 'GIT_COMMIT', value: GIT_COMMIT)]
    } 

  
  }// node
} //podTemplate

// is Push to master branch
def isMaster() {
    return (env.BRANCH_NAME == "master" )
}

def isPullRequest() {
    return (env.BRANCH_NAME ==~  /^PR-\d+$/)
}

def isBuildingTag() {

    // add check that  is branch master?
    return ( env.BRANCH_NAME ==~ /^\d{1}.\d{1}.\d{1}$/ )
}

def isPushtoFeatureBranch() {
    return ( ! isMaster() && ! isBuildingTag() && ! isPullRequest() )
}





