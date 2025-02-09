helm repo add deepflow https://deepflowio.github.io/deepflow
helm repo update deepflow # use `helm repo update` when helm < 3.7.0
cat << EOF > values-custom.yaml
global:
  allInOneLocalStorage: true
  image:
    repository: deepflowce
EOF
helm install deepflow -n deepflow deepflow/deepflow --version 6.6.018 --create-namespace -f values-custom.yaml
