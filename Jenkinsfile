def image

pipeline {
    agent any

    stages {
        stage('Build') {
            tools {
                go '1.17.5'
            }
            steps {
                sh 'CGO_ENABLED=0 go build'
            }
        }

        stage('Docker Build') {
         steps {
                script {
                     image = docker.build("3l0w/improvised:latest")
                }
            }
        }

        stage('Deploy') {
            steps {
                script {
                    docker.withRegistry('https://index.docker.io/v1/', 'dockerhub') {
                        image.push()
                    }
                }
            }
        }
    }
}