#!/usr/bin/env bash
#
# This is a helper script that uses the AWS CLI configuration to construct an
# AWS ProviderConfig.

cd "$(dirname "$0")"

set -e -o pipefail

AWS_PROFILE=default && echo -e "[default]\naws_access_key_id = $(aws configure get aws_access_key_id --profile $AWS_PROFILE)\naws_secret_access_key = $(aws configure get aws_secret_access_key --profile $AWS_PROFILE)" > creds.conf
kubectl create secret generic aws-creds -n crossplane-system --from-file=creds=./creds.conf --dry-run=client -o yaml | kubectl apply -f -

kubectl apply -f secret.yaml