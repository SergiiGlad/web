#!/usr/bin/env groovy

/**
 * This pipeline describes a CI/CD process for running Golang app to multi stages environment
 */


def podLabel = "jenkins-worker-${UUID.randomUUID().toString()}"
def host = "173-193-102-57.nip.io"
def dockerImage = 'sergeyglad/wiki'


podTemplate(label: podLabel, yaml: """
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
  - name: helm
    image: lachlanevenson/k8s-helm:v2.16.1
    tty: true
    command:
      - "cat"
"""){
	node(podLabel) {

		stage('Checkout application SCM') {
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
					
		// BRANCH_NAME = master  - push to master
		// BRANCH_NAME = PR-1    - pull request
		// BRANCH_NAME = develop - push to other branch
		// BRANCH_NAME = 0.0.1  - git tag
		def dockerTag = env.BRANCH_NAME
						
		if ( isMaster() ) dockerTag = sh(returnStdout: true, script: "git rev-parse HEAD").trim().take(7) //short commit

				
		stage('Docker build') {
			container('docker-dind') {
				sh "docker build . -t $dockerImage:$dockerTag"
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
					sh "docker push $dockerImage:$dockerTag"
				}
		  }
		}

		if ( isMaster() || isBuildingTag() ) {
			stage('Deploy') {
				build job: 'web-delivery',
				parameters: [string(name: 'dockerTag', value: dockerTag, description: 'git tag or short commit')]
			}
		} 

			
	}// node
} //podTemplate

def isMaster() {
  return (env.BRANCH_NAME  == "master") 
}

def isPullRequest() {
  return (env.BRANCH_NAME ==~  /^PR-\d+$/)
}

def isBuildingTag() {
  return ( env.BRANCH_NAME ==~ /^\d+\.\d+\.\d+$/ )
}





