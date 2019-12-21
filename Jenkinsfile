#!/usr/bin/env groovy

/**
 * This pipeline describes a CI/CD process for running Golang app to multi stages environment
 */

env.DOCKER_IMAGE_NAME = 'sergeyglad/wiki'

def label = "jenkins-worker-${UUID.randomUUID().toString()}"


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

    stage('Build  Golang app') {
        container('golang') {
            echo "Build Golang app"
            sh 'CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o main .'
        }    
    }

    stage ('Unit test Golang app')  {
        container('golang') { 
            echo "Unit test Golang app"
        }
    }    

    stage('Docker build') {
      container('docker-dind') {
        
        //
        // Environment variables DOCKER_IMAGE_NAME  set by Jenkins plugins 
        // 
        // BRANCH_NAME = master  - master branch
        // BRANCH_NAME = PR-1    - pull request
        // BRANCH_NAME = develop - other branch
        // BRANCH_NAME = v0.0.1  - git tag
        //

        echo "Docker build image name ${DOCKER_IMAGE_NAME}:${BRANCH_NAME}"
        sh 'docker build . -t ${DOCKER_IMAGE_NAME}:${BRANCH_NAME}'
        }
    }

    if ( isPullRequest() ) {
        // exitAsSuccess() 
        return 0
    }   


    sh 'ps aux'
    

    stage ('Docker push') {
        container('docker-dind') {

          sh 'docker image ls'    
          withDockerRegistry([credentialsId: 'docker-api-key', url: 'https://index.docker.io/v1/']) {
                sh 'docker push ${DOCKER_IMAGE_NAME}:${BRANCH_NAME}'
          }
        }    
    }

    if ( isPushtoFeatureBranch() ) {
            // exitAsSuccess() 
            return 0
    }

    
    stage('Deploy') {
   
            def tagDockerImage
            def nameStage

                             
            if ( isChangeSet() && isMaster() ) {
                
                stage('Deploy to Production')
                    echo "Production release controlled by a change to production-release.txt file in application repository root," 
                    echo "containing a git tag that should be released to production environment"

                    tagDockerImage = "${sh(script:'cat production-release.txt',returnStdout: true)}"
                    //? need check is tag exist

                    nameStage = "wiki-prod"
             
                    container('kubectl') {
                        deploy( tagDockerImage, nameStage )
                    }    

                          
            }  
            
            if ( isMaster() ) {
               stage('Deploy to development version') {
                    echo "Every commit to master branch is a dev release" 
                    echo "Its push to master"
                        
                    tagDockerImage = env.BRANCH_NAME
                    nameStage = "wiki-dev"
              
                    container('kubectl') {
                        deploy( tagDockerImage, nameStage )
                     }   
               }        
               
            } 
            
            if ( isBuildingTag() ){
                stage('Deploy to QA stage') {
                    echo "Every git tag on a master branch is a QA release" 
                
                    tagDockerImage = env.BRANCH_NAME
                    nameStage = "wiki-qa"
               
                    container('kubectl') {
                        deploy( tagDockerImage, nameStage )
                    }

                // integrationTest 
                // stage('approve'){ input "OK to go?" }
                }
            }    

            
            
              

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
    return ( env.BRANCH_NAME ==~ /^v\d.\d.\d$/ )    
}

def isPushtoFeatureBranch() {
    return ( ! isMaster() && ! isBuildingTag() && ! isPullRequest() )
}

def isChangeSet() {
    
    def changeLogSets = currentBuild.changeSets
           for (int i = 0; i < changeLogSets.size(); i++) {
           def entries = changeLogSets[i].items
           for (int j = 0; j < entries.length; j++) {
               def files = new ArrayList(entries[j].affectedFiles)
               for (int k = 0; k < files.size(); k++) {
                   def file = files[k]
                   if (file.path.equals("production-release.txt")) {
                       return true
                   }
               }
            }
    }

    return false
}

def deploy( tagName, appName ) {

        echo "Release image: ${DOCKER_IMAGE_NAME}:$tagName"
        echo "Deploy app name: $appName"
  
        withKubeConfig([credentialsId: 'kubeconfig']) {
        sh"""
           
            kubectl delete deploy ${appName} --wait -n jenkins
            kubectl delete svc ${appName} --wait -n jenkins
            kubectl run ${appName} -n jenkins --image=${DOCKER_IMAGE_NAME}:${tagName} --port=3000 --labels="app=${appName}" --image-pull-policy=Always
            kubectl expose -n jenkins deploy/${appName} --port=3000 --target-port=3000 --type=NodePort 
            kubectl get svc -n jenkins

        """ 
        }  
  
}
