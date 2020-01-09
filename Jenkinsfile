#!/usr/bin/env groovy

/**
 * This pipeline describes a CI/CD process for running Golang app to multi stages environment
 */

label = "jenkins-worker-${UUID.randomUUID().toString()}"
host = "173-193-102-57.nip.io"

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
    // BRANCH_NAME = master  - master branch
    // BRANCH_NAME = PR-1    - pull request
    // BRANCH_NAME = develop - other branch
    // BRANCH_NAME = 0.0.1  - git tag
    //
    
    dockerImage = 'sergeyglad/wiki:' + env.BRANCH_NAME
   
    stage('Docker build') {
      container('docker-dind') {
           sh """
                    
           """
           //docker build . -t $dockerImage
        }
    }

    if ( isPullRequest() ) {
        // exitAsSuccess()
        currentBuild.result = 'SUCCESS';  
        return 0
    }

    stage ('Docker push') {
        container('docker-dind') {

          sh 'docker image ls'
          withDockerRegistry([credentialsId: 'docker-api-key', url: 'https://index.docker.io/v1/']) {
                sh """
                    
                """
                // docker push $dockerImage
          }
        }
    }

    if ( isPushtoFeatureBranch() ) {
        // exitAsSuccess()
        currentBuild.result = 'SUCCESS';  
        return 0
    }

    shortCommit = sh(returnStdout: true, script: "git log -n 1 --pretty=format:'%h'").trim() 

    sh 'git rev-parse HEAD > GIT_COMMIT'

    stage('Deploy') {
     build job: 'web-delivery', wait: true, 
     parameters: [string(name: 'BRANCH_NAME1', value: env.BRANCH_NAME),
                  string(name: 'shortCommit1', value: shortCommit),
                  string(name: 'GIT_COMMIT1', value: GIT_COMMIT)]
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
    return ( env.BRANCH_NAME ==~ /^\d.\d.\d$/ )
}

def isPushtoFeatureBranch() {
    return ( ! isMaster() && ! isBuildingTag() && ! isPullRequest() )
}


def printIngress() {
 container('kubectl') {
    withKubeConfig([credentialsId: 'kubeconfig']) {
    
        sh 'kubectl get ing --all-namespaces'
        
        } 
    }     
}

def deployHelm(name, ns, tag) {

     container('helm') {
        withKubeConfig([credentialsId: 'kubeconfig']) {
        sh """    
            echo appVersion: $tag >> ./wikiChart/Chart.yaml

            helm upgrade --install $name ./wikiChart \
            --force \
            --wait \
            --namespace $ns \
            --set image.name=$dockerImage \
            --set appVer=$tag \
            --set ingress.hostName="${name}.${host}" \
            --set ingress.tls[0].hosts[0]="${name}.${host}" \
            --set ingress.tls[0].secretName="acme-${name}-tls" \

            helm ls
        """
    
        }
    }    

}


