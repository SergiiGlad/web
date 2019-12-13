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

                        //docker.Image.build
                    sh 'DOCKER_BUILDKIT=1  docker build . -t ${DOCKER_IMAGE_NAME} --cache-from ${DOCKER_IMAGE_NAME}'
                               

                    withDockerRegistry([credentialsId: 'docker-api-key', url: 'https://index.docker.io/v1/']) {
                        sh 'docker push ${DOCKER_IMAGE_NAME}'
                    }
                    script {
                        sh 'CHANGE_ID: ${env.CHANGE_ID}'
                    }
                }    
            }
        }
        stage('Test') {
            steps {
                container('docker') {
                    echo 'go testing..'
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
