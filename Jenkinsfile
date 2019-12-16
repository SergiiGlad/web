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
                    //sh 'docker build . -t ${DOCKER_IMAGE_NAME} --cache-from ${DOCKER_IMAGE_NAME}'
                }
            }    
        }            

        stage ('TAG') {

           when {
                        buildingTag()
                        }
            steps {
                sh 'echo env.tag' 
            }            

        }
        
        stage('Push to Docker hub') {

                    
                steps {

                    echo "Build docker image"
                    container('docker') {
                        withDockerRegistry([credentialsId: 'docker-api-key', url: 'https://index.docker.io/v1/']) {
                            
                        }

                    sh 'echo Branch Name: ${BRANCH_NAME}'
                    sh 'echo Change ID: ${CHANGE_ID}'
                    sh 'echo env GIT_TAG_NAME: ${GIT_TAG_NAME}'

                 }   

                 script {
                                
                                // isPRMergeBuild
                                if ( env.BRANCH_NAME ==~  /^PR-\d+$/ ) {
                                    sh 'echo It is pull request'
                                } 
                                else if buildingTag() {
                                    sh 'echo push to master BRANCH_NAME: $(BRANCH_NAME)'
                                } else if (env.GIT_TAG_NAME =~ /^.$/ ){
                                    sh 'echo qa release with tag : $(GIT_TAG_NAME)'
                                    sh 'echo GIT TAG MESSAGE: $(GIT_TAG_MESSAGE) '
                                }
                                else {
                                    sh 'echo push to other branch $(BRANCH_NAME)'

                                }
                                    //sh 'docke push ${DOCKER_IMAGE_NAME}'
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
