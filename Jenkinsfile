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
        bat 'docker login -u %DOCKER_USER% -p dckr_pat_3G0TaZht3yhLW1pB80ThnHhTog8'
    }
}

        stage('Build Docker Image') {
            steps {
                bat 'docker build -t %DOCKER_USER%/%IMAGE_NAME%:%IMAGE_TAG% .'
            }
        }

        stage('Push To Docker Hub') {
            steps {
                bat 'docker push %DOCKER_USER%/%IMAGE_NAME%:%IMAGE_TAG%'
            }
        }

        stage('Check Ansible Connection') {
            steps {
                bat 'wsl ansible --version'
            }
        }

        stage('Deploy with Ansible') {
            steps {
                bat 'wsl ansible-playbook deploy.yml'
            }
        }
    }

    post {
        success {
            emailext(
                subject: "Jenkins Build SUCCESS: ${JOB_NAME} #${BUILD_NUMBER}",
                body: """
                Build Successful!

                Job Name: ${JOB_NAME}
                Build Number: ${BUILD_NUMBER}
                Build URL: ${BUILD_URL}

            
                """,
                to: "1762003rajpurohitgaurav@gmail.com",
                from: "1762003rajpurohitgaurav@gmail.com"
            )
        }

        failure {
            emailext(
                subject: "Jenkins Build FAILED: ${JOB_NAME} #${BUILD_NUMBER}",
                body: """
                Build Failed!

                Job Name: ${JOB_NAME}
                Build Number: ${BUILD_NUMBER}
                Build URL: ${BUILD_URL}

                Please check Jenkins logs.
                """,
                to: "1762003rajpurohitgaurav@gmail.com",
                from: "1762003rajpurohitgaurav@gmail.com"
            )
        }

        always {
            echo "Build Finished at: ${new Date()}"
        }
    }
    }
}
