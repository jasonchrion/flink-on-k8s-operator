#!/bin/bash

IMG="${IMG:-ghcr.io\/spotify\/flink-operator:latest}"
NS="${NS:-flink-operator-system}"

sed -i '/- \.\.\/crd/d' ../../config/default/kustomization.yaml
kubectl kustomize ../../config/default | tee templates/flink-operator.yaml
sed -i 's/'"${NS}"'/{{ .Values.flinkOperatorNamespace.name }}/g' templates/flink-operator.yaml
sed -i 's/replicas: 1/replicas: {{ .Values.replicas }}/g' templates/flink-operator.yaml
sed -i "s/$IMG/{{ .Values.operatorImage.name }}/g" templates/flink-operator.yaml
sed -i 's/--watch-namespace=/--watch-namespace={{ .Values.watchNamespace.name }}/' templates/flink-operator.yaml
cp ../../config/crd/bases/flinkoperator.k8s.io_flinkclusters.yaml templates/flink-cluster-crd.yaml

git checkout ../../config/default/kustomization.yaml