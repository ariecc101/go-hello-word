properties([pipelineTriggers([githubPush()])]) 
pipeline {
    stages {
        stage('Testing') {
            steps {
                script { 
                echo "Begin Testing Using Sonarqube"
                } 
            }
        }
        stage('Build') {
            steps {
              // Build Image
                script { 

                echo "Begin Build"
                sh "docker build -t arizalsandi/landingpage:dev-$BUILD_NUMBER . "

                }
            }
        }
        stage('Deploy') {
            steps {
              // Run Docker
                script { 
                
                echo "Begin to Deploy" 
                sh "docker image rmi arizalsandi/landingpage:dev-$BUILD_NUMBER"

                }
            }
        }
    }
}