#!/bin/bash

usage() {
    echo -ne "Usage: ./deploy [option]
    options:
        -d : infra and application deployment
        -a : abort\n"
    exit
}

deploy() {
  kubectl apply -f infra/mysql-secrets.yaml
  kubectl apply -f infra/rabbitmq-secrets.yaml
  kubectl apply -f infra/mysql-deployment.yaml
  kubectl apply -f infra/rabbitmq-deployment.yaml
  kubectl apply -f app/api-deployment.yaml
  kubectl apply -f app/worker-deployment.yaml
}

abort() {
  kubectl delete -f infra/mysql-secrets.yaml
  kubectl delete -f infra/rabbitmq-secrets.yaml
  kubectl delete -f infra/mysql-deployment.yaml
  kubectl delete -f infra/rabbitmq-deployment.yaml
  kubectl delete -f app/api-deployment.yaml
  kubectl delete -f app/worker-deployment.yaml
}

if [ "$@" ]; then
    while getopts "dah" opt; do
        case $opt in
            d)
                deploy
                shift
                ;;
            a)
                abort
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