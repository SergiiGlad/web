stage('Build and test Golang app') {
    
      container('golang') {
        echo "Build Golang app"
        sh 'ls; pwd; hostname;'
        // sh 'printenv | sort'
       // GOPATH="${WORKSPACE}/go" //${sh(script:'cat production-release.txt',returnStdout: true)}"
       // echo "${GOPATH}"
        withEnv(["GOPATH=${WORKSPACE}"]) {
        sh """
            echo ${env.GOPATH}
            export GOPATH="${env.GOPATH}/go"
            echo $GOPATH
            go version
            go install wiki
            ls ${GOPATH}
            ls ${GOPATH}/bin
        """
        }
      }
    }