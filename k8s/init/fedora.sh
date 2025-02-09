#! /bin/bash

function error_exit {
    echo "$1" >&2
    exit 1
}

echo "1. install base tools"
yum update -y
yum install -y net-tools jq iproute-tc NetworkManager-tui bridge-utils bind-utils tcpdump  vim nano wget bash-comp* htop

echo "2. 将 SELinux 的状态设置为宽松模式"
sed -ri 's/SELINUX=enforcing/SELINUX=disabled/' /etc/selinux/config
setenforce 0

echo "3. 关闭防火墙"
systemctl stop firewalld
systemctl disable firewalld
systemctl disable firewalld.service

echo "4. 关闭 swap"
sudo dnf remove -y zram-generator-defaults
sudo swapoff -a

echo "5. 安装docker"
yum install -y yum-utils
yum-config-manager --add-repo http://mirrors.aliyun.com/docker-ce/linux/centos/docker-ce.repo
sed -i 's/$releasever/8/g' /etc/yum.repos.d/docker-ce.repo
yum makecache
sudo yum list docker-ce --showduplicates | sort -r
yum install -y docker-ce
cat <<EOF > daemon.json
{
  "registry-mirrors": ["https://hdi5v8p1.mirror.aliyuncs.com"],
  "exec-opts": ["native.cgroupdriver=systemd"],
	"insecure-registries" : [ "ykl.io:40443"]
}
EOF
mv daemon.json /etc/docker/
systemctl enable docker.service
systemctl restart docker.service

echo "6. 安装k8s"
cat > /etc/sysctl.d/k8s.conf << EOF
   net.bridge.bridge-nf-call-ip6tables = 1
   net.bridge.bridge-nf-call-iptables = 1
EOF
sysctl --system

modprobe br_netfilte
lsmod | grep br_netfilter

modprobe ip_tables
modprobe ip6_tables
cat > /etc/modules-load.d/iptables.conf << EOF
ip_tables
ip6_tables
EOF

cat > /etc/yum.repos.d/kubernetes.repo << EOF
[kubernetes]
name=Kubernetes
baseurl=https://mirrors.aliyun.com/kubernetes/yum/repos/kubernetes-el7-x86_64/
enabled=1
gpgcheck=0
repo_gpgcheck=0
gpgkey=https://mirrors.aliyun.com/kubernetes/yum/doc/yum-key.gpg https://mirrors.aliyun.com/kubernetes/yum/doc/rpm-package-key.gpg
EOF

yum makecache
yum list kubectl --showduplicates | sort -r
yum -y install  kubeadm-1.23.17-0  kubelet-1.23.17-0 kubectl-1.23.17-0
systemctl enable kubelet
kubeadm config images pull --image-repository=registry.aliyuncs.com/google_containers --kubernetes-version=v1.23.17
