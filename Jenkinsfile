def useAgent
if (BRANCH_NAME == 'main') {
    useAgent = "production"
} else if(BRANCH_NAME == 'testnet'){
    useAgent = "testnet"
} else if(BRANCH_NAME == 'staging'){
    useAgent = "staging"
} else if(BRANCH_NAME == 'develop'){
    useAgent = "develop"
}

pipeline {
    agent {
        label useAgent
    }

    environment {
        CI                      = 'true'
        REGISTRY_URL            = "invoker.bitwyre.com/bitwyre"
        SERVICE_NAME            = "${JOB_NAME.split('/')[0]}"
        OVERRIDE_SERVICE_NAME   = "private-api"
        IMAGE_TAG               = sh(script: "git rev-parse --short HEAD", returnStdout: true).trim()
        APP_VERSION             = sh(script: """cat package.json | grep version | head -1 | awk -F: '{ print \$2 }' | cut -d '"' -f2""", returnStdout: true).trim()
        DOCKER_PATH             = "gateway/rest/${SERVICE_NAME}/Dockerfile"
        HELM_PATH               = "infrastructure/helm_chart/${OVERRIDE_SERVICE_NAME}"
        HELM_NAMESPACE          = "gateway"
        IMAGE_NAME              = "${REGISTRY_URL}/${OVERRIDE_SERVICE_NAME}"
    }

    stages {
        stage('Build') {
            when {
                anyOf { changeset "**/${SERVICE_NAME}/**"; changeset "**/shared/**";  }
            }

            steps {
                script {
                    if (GIT_BRANCH == 'master') {
                        echo "[+] Building Docker Image for ${GIT_BRANCH}"
                        sh "docker build -t ${IMAGE_NAME}:${APP_VERSION} -f ${DOCKER_PATH} ."

                        echo "[+] Pushing Image to ${IMAGE_NAME}:${APP_VERSION}"
                        sh "docker push ${IMAGE_NAME}:${APP_VERSION}"
                    } else if(GIT_BRANCH == 'testnet'){
                        echo "[+] Building Image with tag ${IMAGE_NAME}:${IMAGE_TAG}-testnet"
                        sh "docker build -t ${IMAGE_NAME}:${IMAGE_TAG}-testnet -f ${DOCKER_PATH} ."

                        echo "[+] Pushing Image to ${IMAGE_NAME}:${IMAGE_TAG}-testnet"
                        sh "docker push ${IMAGE_NAME}:${IMAGE_TAG}-testnet"
                    } else if(GIT_BRANCH == 'staging'){
                        echo "[+] Building Image with tag ${IMAGE_NAME}:${IMAGE_TAG}-stg"
                        sh "docker build -t ${IMAGE_NAME}:${IMAGE_TAG}-stg -f ${DOCKER_PATH} ."

                        echo "[+] Pushing Image to ${IMAGE_NAME}:${IMAGE_TAG}-stg"
                        sh "docker push ${IMAGE_NAME}:${IMAGE_TAG}-stg"
                    } else if(GIT_BRANCH == 'develop'){
                        echo "[+] Building Image with tag ${IMAGE_NAME}:${IMAGE_TAG}-dev"
                        sh "docker build -t ${IMAGE_NAME}:${IMAGE_TAG}-dev -f ${DOCKER_PATH} ."

                        echo "[+] Pushing Image to ${IMAGE_NAME}:${IMAGE_TAG}-dev"
                        sh "docker push ${IMAGE_NAME}:${IMAGE_TAG}-dev"
                    }
                }
            }
        }

        stage('Deploy Development') {
            when {
                anyOf { changeset "**/${SERVICE_NAME}/**"; changeset "**/shared/**";  }

                branch 'develop'
            }
            steps {
                script {
                    echo "[+] Deploying Helm Chart for ${GIT_BRANCH}"
                    sh "kubectl config use-context development"
                    sh "helm upgrade --install ${OVERRIDE_SERVICE_NAME} -f ${HELM_PATH}/values-dev.yaml --set image.imagetag=${IMAGE_TAG}-dev -n ${HELM_NAMESPACE} ${HELM_PATH}/"
                }
            }
            post {
                always {
                    script {
                        echo "[+] Deployment status"
                        sh "kubectl rollout status deployment/${OVERRIDE_SERVICE_NAME} -n ${HELM_NAMESPACE} --timeout=1000s"
                    }
                }
            }
        }

        stage('Deploy Staging') {
            when {
                anyOf { changeset "**/${SERVICE_NAME}/**"; changeset "**/shared/**";  }

                branch 'staging'
            }
            steps {
                script {
                    echo "[+] Deploying Helm Chart for ${GIT_BRANCH}"
                    sh "kubectl config use-context staging"
                    sh "helm upgrade --install ${OVERRIDE_SERVICE_NAME} -f ${HELM_PATH}/values-staging.yaml --set image.imagetag=${IMAGE_TAG}-stg -n ${HELM_NAMESPACE} ${HELM_PATH}/"
                }
            }
            post {
                always {
                    script {
                        echo "[+] Deployment status"
                        sh "kubectl rollout status deployment/${OVERRIDE_SERVICE_NAME} -n ${HELM_NAMESPACE} --timeout=1000s"
                    }
                }
            }
        }

        stage('Deploy Testnet') {
            when {
                anyOf { changeset "**/${SERVICE_NAME}/**"; changeset "**/shared/**";  }

                branch 'testnet'
            }
            steps {
                script {
                    echo "[+] Deploying Helm Chart for ${GIT_BRANCH}"
                    sh "kubectl config use-context testnet"
                    sh "helm upgrade --install ${OVERRIDE_SERVICE_NAME} -f ${HELM_PATH}/values-testnet.yaml --set image.imagetag=${IMAGE_TAG}-testnet -n ${HELM_NAMESPACE} ${HELM_PATH}/"
                }
            }
            post {
                always {
                    script {
                        echo "[+] Deployment status"
                        sh "kubectl rollout status deployment/${OVERRIDE_SERVICE_NAME} -n ${HELM_NAMESPACE} --timeout=1000s"
                    }
                }
            }
        }

        stage('Deploy Production') {
            when {
                anyOf { changeset "**/${SERVICE_NAME}/**"; changeset "**/shared/**";  }
                branch 'master'
            }
            steps {
                script {
                    echo "[+] Deploying Helm Chart for ${GIT_BRANCH}"
                    sh "kubectl config use-context production"
                    sh "helm upgrade --install ${OVERRIDE_SERVICE_NAME} -f ${HELM_PATH}/values-production.yaml --set image.imagetag=${APP_VERSION} -n ${HELM_NAMESPACE} ${HELM_PATH}/"
                }
            }
            post {
                always {
                    script {
                        echo "[+] Deployment status"
                        sh "kubectl rollout status deployment/${OVERRIDE_SERVICE_NAME} -n ${HELM_NAMESPACE} --timeout=1000s"
                    }
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