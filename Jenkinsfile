properties([pipelineTriggers([githubPush()])]) 
pipeline {
    agent any
    stages {
        stage('Testing') {
            steps {
                script { 
                echo "Begin Testing"
                } 
            }
        }
        stage('Build') {
            steps {
              // Build Image
                script { 

                echo "Begin Build"
                sh "docker build -t sulaplink001/go-hello-world:dev-$BUILD_NUMBER . "

                }
            }
        }
        stage('Deploy') {
            steps {
              // Run Docker
                script { 
                
                echo "Begin to Deploy" 
                sh "docker stop hello"
                sh "docker rm hello"
                sh "docker run -d -p 5000:5000 --name hello sulaplink001/go-hello-world:dev-$BUILD_NUMBER"

                }
            }
        }
    }
}