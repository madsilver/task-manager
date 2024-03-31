#!/bin/bash

usage() {
    echo -ne "Usage: ./deploy [option]
    options:
        -i : infrastructure deployment
        -a : applications deployment\n"
    exit
}

infraDeploy() {
    kubectl apply -f infra/mysql-secrets.yaml
    kubectl apply -f infra/mysql-deployment.yaml
    kubectl apply -f infra/rabbitmq-secrets.yaml
    kubectl apply -f infra/rabbitmq-deployment.yaml
}

appDeploy() {
  kubectl apply -f deployment.yaml
}

if [ "$@" ]; then
    while getopts "iah" opt; do
        case $opt in
            i)
                infraDeploy
                shift
                ;;
            a)
                appDeploy
                shift
                ;;
            h)
                usage
                exit 0
                ;;
            \?)
                ;;
        esac
    done
else
    usage
    exit 0
fi