#
kubectl apply -f chaining.yaml
#
helm repo add cilium https://helm.cilium.io/
helm install cilium cilium/cilium --version 1.16.4 \
  --namespace=kube-system \
  --set cni.chainingMode=generic-veth \
  --set cni.customConf=true \
  --set cni.configMap=cni-configuration \
  --set routingMode=native \
  --set enableIPv4Masquerade=false \
  --set hubble.relay.enabled=true \
  --set hubble.ui.enabled=true

# 手动修改镜像、配置

helm install cilium cilium/cilium --version 1.16.4 \
    --namespace kube-system \
    --set kubeProxyReplacement=true



helm template cilium cilium/cilium --version 1.16.4 \
  --namespace=kube-system \
  --set cni.chainingMode=generic-veth \
  --set cni.customConf=true \
  --set cni.configMap=cni-configuration \
  --set routingMode=native \
  --set kubeProxyReplacement=true \
  --set enableIPv4Masquerade=false \
  --set devices="ovn0 genev_sys_6081 vxlan_sys_4789" \
  > cilium-1164-debug2.yaml
