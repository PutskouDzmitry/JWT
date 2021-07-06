pipeline {
    agent any
    tools {
        go 'go1.16.5'
    }
    stages {        
        stage('Pre Test') {
            steps {
                echo 'Installing dependencies'
                sh 'go version'
            }
        }
    }
}
