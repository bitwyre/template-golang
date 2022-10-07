def useAgent
if (BRANCH_NAME == 'master') {
    useAgent = "production"
} else if(BRANCH_NAME == 'testnet'){
    useAgent = "testnet"
} else if(BRANCH_NAME == 'staging'){
    useAgent = "staging"
} else if(BRANCH_NAME == 'develop'){
    useAgent = "develop"
}else{
    useAgent = "develop"
}

pipeline {
    agent {
        label useAgent
    }

    environment {
        REGISTRY_URL            = "invoker.bitwyre.com/bitwyre"
        SERVICE_NAME            = "${JOB_NAME.split('/')[0]}"
        OVERRIDE_SERVICE_NAME   = "order_service_go"
        IMAGE_TAG               = sh(script: "git rev-parse --short HEAD", returnStdout: true).trim()
        APP_VERSION             = sh(script: """cat package.json | grep version | head -1 | awk -F: '{ print \$2 }' | cut -d '"' -f2""", returnStdout: true).trim()
        DOCKER_PATH             = "gateway/rest/${SERVICE_NAME}/docker/build.Dockerfile"
        HELM_PATH               = "infrastructure/helm_chart/${OVERRIDE_SERVICE_NAME}"
        HELM_NAMESPACE          = "engine"
        IMAGE_NAME              = "${REGISTRY_URL}/${OVERRIDE_SERVICE_NAME}"
    }

    stages {
        stage('Build') {
            steps {
                script {
                    if (GIT_BRANCH == 'master') {
                        echo "[+] Building Docker Image for ${GIT_BRANCH}"
                        sh "docker build -t ${IMAGE_NAME}:${APP_VERSION} -f ${DOCKER_PATH} ."

                        echo "[+] Pushing Image to ${IMAGE_NAME}:${APP_VERSION}"
                        sh "docker push ${IMAGE_NAME}:${APP_VERSION}"
                    } else {
                        echo "[+] Building Image with tag ${IMAGE_NAME}:${IMAGE_TAG}-${GIT_BRANCH}"
                        sh "docker build -t ${IMAGE_NAME}:${IMAGE_TAG}-${GIT_BRANCH} -f ${DOCKER_PATH} ."

                        echo "[+] Pushing Image to ${IMAGE_NAME}:${IMAGE_TAG}-${GIT_BRANCH}"
                        sh "docker push ${IMAGE_NAME}:${IMAGE_TAG}-${GIT_BRANCH}"
                }
            }
        }

        stage('Deploy') {
            steps {
                script {
                    echo "Deploy Dev"
                }
            }
        }
    }

    post {
        success {
            echo 'I succeeded!'
        }
        unstable {
            echo 'I am unstable :('
        }
        failure {
            echo 'Failed :(('
        }
        cleanup {
            echo "Clean up in post workspace."
            cleanWs()
        }
    }
}