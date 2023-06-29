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
        stage('Build Image') {
            steps {
              // Build Image
                script { 

                // echo "Begin Build Image"
                dockerImage = docker.build("sulaplink001/go-hello-world:latest")
                // sh "docker build -t sulaplink001/go-hello-world:dev-$BUILD_NUMBER . "
                // sh "docker build -t sulaplink001/go-hello-world:latest . "
                // sh "docker push sulaplink001/go-hello-world:latest"

                }
            }
        }

        stage('Push image') {
            steps {
                script{
                    withDockerRegistry([ credentialsId: "dockerhubaccount", url: "" ]) {
                    dockerImage.push()
                    }
                }
            }
        }

        stage('Deploy') {
            steps {
              // Run Docker
                script { 
                
                echo "Begin to Deploy" 
                // sh "docker compose up -d"
                // sh "docker stop hello"
                // sh "docker rm hello"
                // sh "docker run -d -p 5000:5000 --name hello sulaplink001/go-hello-world:dev-$BUILD_NUMBER"

                }
            }
        }
    }
}