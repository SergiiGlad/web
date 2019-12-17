pipeline {
    
    agent{
        kubernetes {
            yamlFile 'podTemplWorker.yaml'
        }
    }

     environment {
            //be sure to replace "sergeyglad" with your own Docker Hub username
            DOCKER_IMAGE_NAME = "sergeyglad/wiki"
            PROD = 'prod'
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
                sh 'echo building tag' 
            }            

        }

        stage ('Change file production-release.txt ')  {
            when {              
                    
                    changeset pattern: "production-release.txt"
             
            }

            steps {
                echo "echo production release"
                //PROD environment var
                script {
                    PROD="ooooo"
                    echo "script ${PROD}"
                    //$(cat production-release.txt) ) 
                }
                //sh 'printenv | sort'
                echo "${PROD}"
            }
        }
        
        stage('Push to Docker hub') {

                    
                steps {

                    echo "Build docker image"
                    container('docker') {
                        withDockerRegistry([credentialsId: 'docker-api-key', url: 'https://index.docker.io/v1/']) {

                           // sh 'docker push ${DOCKER_IMAGE_NAME}'
                        }

                    sh 'echo Branch Name: ${BRANCH_NAME}'
                    sh 'echo Change ID ${CHANGE_ID}'

                 }   

                 script{
                                // isPRMergeBuild
                                if ( env.BRANCH_NAME ==~  /^PR-\d+$/ ) {
                                    sh 'echo It is pull request'
                                // is Push to master    
                                } else if (env.BRANCH_NAME ==~  /^master$/) {
                                    sh 'echo Its push to master '
                                // isTag    
                                } else if (env.BRANCH_NAME =~ /^v\d.\d.\d$/ ){
                                    sh 'echo qa release with tag : $(BRANCH_NAME)'
                                // Other operation    
                                } else {
                                    sh 'echo push to other branch $(BRANCH_NAME)'
                                }
                                   
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
                sh 'echo Deploying....'
                container('helm') {
                 withKubeConfig([credentialsId: 'kubeconfig']) {
                 sh 'helm version'
                 } 
               }
            }
        }
    }
}
