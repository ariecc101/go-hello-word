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

                echo "Begin Build Image"
                dockerImage = docker.build("sulaplink001/go-hello-world:dev-$BUILD_NUMBER")
                // dockerImage = docker.build("sulaplink001/go-hello-world:VERSIONDEV")
                sh "sed -i 's/VERSIONDEV/$BUILD_NUMBER/g' docker-compose.yml"
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
            //   Run Docker
                script { 
                
                echo "Begin to Deploy" 
                sh "docker compose pull && docker compose up -d"
                }
            }
        }
    }
}