#!/usr/bin/env groovy

/**
 * This pipeline describes a CI/CD process for running Golang app to multi stages environment
 */


label = "jenkins-worker-${UUID.randomUUID().toString()}"
host = "184-172-214-143.nip.io"


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
                    docker build . -t $dockerImage
           """
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
                    docker push $dockerImage
                """
          }
        }
    }

    if ( isPushtoFeatureBranch() ) {
            // exitAsSuccess()
            currentBuild.result = 'SUCCESS';  
            return 0
    }

    if ( isMaster()  ) {

        if ( onlyJenkinsfileChangeSet() ) {
             stage('Deploy to Production') {
                echo "Production release controlled by a change to production-release.txt file in application repository root,"
                echo "containing a git tag that should be released to production environment"

                tagDockerImage = '${sh(script:'cat production-release.txt',returnStdout: true)}'
                //? need check is tag exist
                    
                deployHelm( "wiki-prod",                      // name chart release
                            "prod",                           // namespace
                            tagDockerImage )             // image tag from file production-release.txt
                    
                } //stage   

        } else {
            stage('Deploy development version') {
                echo "Every commit to master branch is a dev release"
                echo "Its push to master"
                                
                deployHelm( "wiki-dev",                        // name chart release
                            "develop",                         // namespace
                            env.BRANCH_NAME)                   // image tag = master
                        
                }
            }
    
                       
       
    } //if

    if ( isBuildingTag() ){
        stage('Deploy to QA stage') {
            echo "Every git tag on a master branch is a QA release"

            deployHelm( "wiki-qa",                      // name chart release
                        "qa",                           // namespace
                        env.BRANCH_NAME )               // image tag = 0.0.1 
                    
        }    

               
    }

    printIngress() // ingress info

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

def onlyJenkinsfileChangeSet() {
    def onlyOneFile = false
        currentBuild.changeSets.any { changeSet -> 
        if ( changeSet.items.length == 1 ) { onlyOneFile = true }
        changeSet.items.each { entry ->
            entry.affectedFiles.each { file -> 
                if (file.path.equals("production-release.txt") && onlyOneFile) {
                     return true   
                }    
        }
    }
}   

    return false    
}

def isChangeSet() {

    currentBuild.changeSets.each { changeSet ->  
        changeSet.items.each { entry ->
             entry.affectedFiles.each { file -> 
                  if (file.path.equals("production-release.txt")) {
                    return true
                }
        }        
    }
}    

    return false
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

            echo tag: ${tag}

            helm upgrade --install $name --debug  ./wikiChart \
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
