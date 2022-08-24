#!/bin/bash

#set -e

read -p "istio revision: " REVISION 

NS=test-${REVISION}-`date +%Y%m%d`

kubectl create ns $NS
kubectl label ns $NS istio.io/rev=${REVISION} --overwrite

istioctl manifest generate --set revision=$REVISION -f <(sed -e 's/changeme/'"$NS"'/' -e 's/REVISION/'"$REVISION"'/' gw-override.yaml) | kubectl apply -f -


kubectl -n $NS apply -f <(sed -e 's/test-gateway/test-gateway-'"$REVISION"'/' objects.yaml)
