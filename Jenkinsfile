node {
    stage('Prepare') {
        checkout scm
        sh "make clean"
    }
    stage('Build') {
        checkout scm
        sh "./scripts/ci-test.sh all"
    }
    stage('Test') {
        checkout scm
        sh "./scripts/ci-test.sh test"
    }
    stage('Lint') {
        checkout scm
        sh "./scripts/ci-test.sh lint"
    }
}
