pipeline {
    agent any

    environment {
        DOCKER_USER = "gaurav1374"
        IMAGE_NAME = "calculator"
        IMAGE_TAG = "latest"
    }

    stages {

        stage('Checkout Source Code') {
            steps {
                git branch: 'main', url: 'https://github.com/Gaurav-Rajpurohit/SPE_mini_project.git'
            }
        }

        stage('Run Unit Tests') {
            steps {
                echo "Running unit tests..."
            }
        }

        stage('Docker Login') {
            steps {
                bat "docker login -u %DOCKER_USER% -p YOUR_DOCKER_TOKEN"
            }
        }

        stage('Build Docker Image') {
            steps {
                bat "docker build -t %DOCKER_USER%/%IMAGE_NAME%:%IMAGE_TAG% ."
            }
        }

        stage('Push To Docker Hub') {
            steps {
                bat "docker push %DOCKER_USER%/%IMAGE_NAME%:%IMAGE_TAG%"
            }
        }

        stage('Check Ansible in WSL') {
            steps {
                bat "wsl ansible --version"
            }
        }

        stage('Deploy with Ansible (WSL)') {
            steps {
                bat "wsl ansible-playbook deploy.yml"
            }
        }
    }

    post {
        always {
            echo 'Build Finished'
        }

        success {
            echo 'Build Successful'
        }

        failure {
            echo 'Build Failed'
        }
    }
}
