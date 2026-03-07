pipeline {
    agent any

    environment {
        DOCKER_IMAGE = "gaurav1374/calculator"
    }

    stages {

        stage('Checkout Source Code') {
            steps {
                git url: 'https://github.com/Gaurav-Rajpurohit/SPE_mini_project.git', branch: 'main'
            }
        }

        stage('Run Unit Tests') {
            steps {
                sh '''
                go mod tidy
                go test -v ./...
                '''
            }
        }

        stage('Docker Login') {
            steps {
                withCredentials([usernamePassword(
                    credentialsId: 'docker-creds',
                    usernameVariable: 'DOCKER_USER',
                    passwordVariable: 'DOCKER_PASS'
                )]) {
                    sh '''
                    echo $DOCKER_PASS | docker login -u $DOCKER_USER --password-stdin
                    '''
                }
            }
        }

        stage('Build Docker Image') {
            steps {
                sh '''
                docker build -t $DOCKER_IMAGE:$BUILD_NUMBER .
                docker tag $DOCKER_IMAGE:$BUILD_NUMBER $DOCKER_IMAGE:latest
                '''
            }
        }

        stage('Push To Docker Hub') {
            steps {
                sh '''
                docker push $DOCKER_IMAGE:$BUILD_NUMBER
                docker push $DOCKER_IMAGE:latest
                '''
            }
        }

        stage('Check Ansible Connection') {
            steps {
                sh '''
                ansible -i inventory.ini all -m ping
                '''
            }
        }

        stage('Deploy with Ansible') {
            steps {
                sh '''
                ansible-playbook -i inventory.ini deploy.yml
                '''
            }
        }
    }

    post {

        success {
            echo "Build Successful"
        }

        failure {
            echo "Build Failed"
        }

        always {
            echo "Build Finished"
        }
    }
}
