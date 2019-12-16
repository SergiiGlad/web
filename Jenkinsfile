pipeline {
    
    agent{
        kubernetes {
            yamlFile 'podTemplWorker.yaml'
        }
    }

     environment {
            //be sure to replace "sergeyglad" with your own Docker Hub username
            DOCKER_IMAGE_NAME = "sergeyglad/wiki"
        }

    stages {
      

        stage('Build Golang project') {
          steps{
            sh 'echo "go build"'
          }
        }
        stage('Build Dockerfile') {
       
            steps {
                container('docker') {
                     
                    sh 'echo "Building Dockerfile"'
                    // docker.Image.build
                    // DOCKER_BUILDKIT=1 

                    sh 'pwd'
                    sh 'ls'
                    sh 'echo Branch Name: ${BRANCH_NAME}'
                    sh 'echo Change ID: ${CHANGE_ID}'
                    sh 'docker build . -t ${DOCKER_IMAGE_NAME} --cache-from ${DOCKER_IMAGE_NAME}'
                }
            }    
        }            
        
        stage('Push when') {

                // isPRMergeBuild
             
                //expression { BRANCH_NAME ==~  /^PR-\d+$/ }
           
                    
                //changeRequest()
        
                steps {

                    echo "Build docker image"
                    container('docker') {
                        withDockerRegistry([credentialsId: 'docker-api-key', url: 'https://index.docker.io/v1/']) {
                            script {
                                if ( RANCH_NAME ==~  /^PR-\d+$/ ) {
                                    sh 'echo It is pull request'
                                } 
                                else {
                                    sh 'echo push to branch'
                                }
                                    //sh 'docke push ${DOCKER_IMAGE_NAME}'
                            }
                        }

                    sh 'echo Branch Name: ${BRANCH_NAME}'
                    sh 'echo Change ID: ${CHANGE_ID}'

                 }   
                }
            }
        
        stage('Test') {
            steps {
                container('docker') {
                    echo 'go testing...'
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
