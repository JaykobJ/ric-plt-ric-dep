#!/bin/bash -x
#
################################################################################
#   Copyright (c) 2019 AT&T Intellectual Property.                             #
#   Copyright (c) 2022 Nokia.                                                  #
#                                                                              #
#   Licensed under the Apache License, Version 2.0 (the "License");            #
#   you may not use this file except in compliance with the License.           #
#   You may obtain a copy of the License at                                    #
#                                                                              #
#       http://www.apache.org/licenses/LICENSE-2.0                             #
#                                                                              #
#   Unless required by applicable law or agreed to in writing, software        #
#   distributed under the License is distributed on an "AS IS" BASIS,          #
#   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.   #
#   See the License for the specific language governing permissions and        #
#   limitations under the License.                                             #
################################################################################


usage() {
    echo "Usage: $0 [ -k <k8s version> -d <docker version> -e <helm version> -c <cni-version>" 1>&2;

    echo "k:    kubernetes version" 1>&2;
    echo "c:    kubernetes CNI  version" 1>&2;
    echo "d:    docker version" 1>&2;
    echo "e:    helm version" 1>&2;
    exit 1;
}


wait_for_pods_running () {
  NS="$2"
  CMD="kubectl get pods --all-namespaces "
  if [ "$NS" != "all-namespaces" ]; then
    CMD="kubectl get pods -n $2 "
  fi
  KEYWORD="Running"
  if [ "$#" == "3" ]; then
    KEYWORD="${3}.*Running"
  fi

  CMD2="$CMD | grep \"$KEYWORD\" | wc -l"
  NUMPODS=$(eval "$CMD2")
  echo "waiting for $NUMPODS/$1 pods running in namespace [$NS] with keyword [$KEYWORD]"
  while [  $NUMPODS -lt $1 ]; do
    sleep 5
    NUMPODS=$(eval "$CMD2")
    echo "> waiting for $NUMPODS/$1 pods running in namespace [$NS] with keyword [$KEYWORD]"
  done 
}


start_ipv6_if () {
  IPv6IF="$1"
  if ifconfig -a $IPv6IF; then
    echo "" >> /etc/network/interfaces.d/50-cloud-init.cfg
    echo "allow-hotplug ${IPv6IF}" >> /etc/network/interfaces.d/50-cloud-init.cfg
    echo "iface ${IPv6IF} inet6 auto" >> /etc/network/interfaces.d/50-cloud-init.cfg
    ifconfig ${IPv6IF} up
  fi
}

KUBEV="1.26.15" #1.16.0
KUBECNIV="1.1.1" #0.7.5
HELMV="3.14.3" #3.5.4
DOCKERV="25.0.4" #20.10.21
KUSTOMIZEV="5.0.3"

echo running ${0}
while getopts ":k:d:e:n:c" o; do
    case "${o}" in
    e)	
       HELMV=${OPTARG}
        ;;
    d)
       DOCKERV=${OPTARG}
        ;;
    k)
       KUBEV=${OPTARG}
       ;;
    c)
       KUBECNIV=${OPTARG}
       ;;
    *)
       usage
       ;;
    esac
done

if [[ ${HELMV} == 2.* ]]; then
  echo "helm 2 ("${HELMV}")not supported anymore" 
  exit -1
fi

set -x
export DEBIAN_FRONTEND=noninteractive
echo "$(hostname -I) $(hostname)" >> /etc/hosts
printenv

IPV6IF=""

rm -rf /opt/config
mkdir -p /opt/config
echo "${DOCKERV}" > /opt/config/docker_version.txt
echo "${KUBEV}" > /opt/config/k8s_version.txt
echo "${KUBECNIV}" > /opt/config/k8s_cni_version.txt
echo "${HELMV}" > /opt/config/helm_version.txt
echo "$(hostname -I)" > /opt/config/host_private_ip_addr.txt
echo "$(curl ifconfig.co)" > /opt/config/k8s_mst_floating_ip_addr.txt
echo "$(hostname -I)" > /opt/config/k8s_mst_private_ip_addr.txt
echo "__mtu__" > /opt/config/mtu.txt
echo "__cinder_volume_id__" > /opt/config/cinder_volume_id.txt
echo "$(hostname)" > /opt/config/stack_name.txt

ISAUX='false'
if [[ $(cat /opt/config/stack_name.txt) == *aux* ]]; then
  ISAUX='true'
fi

apt-get update
apt-get install -y ca-certificates curl apt-transport-https gpg

modprobe -- ip_vs
modprobe -- ip_vs_rr
modprobe -- ip_vs_wrr
modprobe -- ip_vs_sh
modprobe -- nf_conntrack

if [ ! -z "$IPV6IF" ]; then
  start_ipv6_if $IPV6IF
fi

SWAPFILES=$(grep swap /etc/fstab | sed '/^[ \t]*#/ d' | sed 's/[\t ]/ /g' | tr -s " " | cut -f1 -d' ')
if [ ! -z $SWAPFILES ]; then
  for SWAPFILE in $SWAPFILES
  do
    if [ ! -z $SWAPFILE ]; then
      echo "disabling swap file $SWAPFILE"
      if [[ $SWAPFILE == UUID* ]]; then
        UUID=$(echo $SWAPFILE | cut -f2 -d'=')
        swapoff -U $UUID
      else
        swapoff $SWAPFILE
      fi
      sed -i "\%$SWAPFILE%d" /etc/fstab
    fi
  done
fi



echo "### Docker version  = "${DOCKERV}
echo "### k8s version     = "${KUBEV}
echo "### helm version    = "${HELMV}
echo "### k8s cni version = "${KUBECNIV}

KUBEVERSION="${KUBEV}-1.1"
CNIVERSION="${KUBECNIV}-1.1"
DOCKERVERSION="${DOCKERV}"

UBUNTU_RELEASE=$(lsb_release -r | sed 's/^[a-zA-Z:\t ]\+//g')
if [[ ${UBUNTU_RELEASE} == 16.* ]]; then
  echo "Installing on Ubuntu $UBUNTU_RELEASE (Xenial Xerus) host"
  if [ ! -z "${DOCKERV}" ]; then
    DOCKERVERSION="${DOCKERV}-0ubuntu1~16.04.5"
  fi
elif [[ ${UBUNTU_RELEASE} == 18.* ]]; then
  echo "Installing on Ubuntu $UBUNTU_RELEASE (Bionic Beaver)"
  if [ ! -z "${DOCKERV}" ]; then
    DOCKERVERSION="${DOCKERV}-0ubuntu1~18.04.4"
  fi
elif [[ ${UBUNTU_RELEASE} == 20.* ]]; then
  echo "Installing on Ubuntu $UBUNTU_RELEASE (Focal Fossal)"
  if [ ! -z "${DOCKERV}" ]; then
    #DOCKERVERSION="${DOCKERV}-0ubuntu1~20.04.2"  # 20.10.21-0ubuntu1~20.04.2
    DOCKERVERSION="5:${DOCKERV}-1~ubuntu.20.04~focal"
  fi
else
  echo "Unsupported Ubuntu release ($UBUNTU_RELEASE) detected.  Exit."
fi

echo "docker version to use = "${DOCKERVERSION}

# DEPRECATED PACKAGE REPOSITORY FROM 1.1.2024
#curl -s https://packages.cloud.google.com/apt/doc/apt-key.gpg | apt-key add -
#echo 'deb http://apt.kubernetes.io/ kubernetes-xenial main' > /etc/apt/sources.list.d/kubernetes.list

# NEW Kubernetes package repository. Does not work for < v1.24
sudo mkdir -p -m 755 /etc/apt/keyrings
sudo chmod 644 /etc/apt/keyrings/kubernetes-apt-keyring.gpg
curl -fsSL "https://pkgs.k8s.io/core:/stable:/v1.26/deb/Release.key" | sudo gpg --dearmor -o /etc/apt/keyrings/kubernetes-apt-keyring.gpg
echo 'deb [signed-by=/etc/apt/keyrings/kubernetes-apt-keyring.gpg] https://pkgs.k8s.io/core:/stable:/v1.26/deb/ /' | sudo tee /etc/apt/sources.list.d/kubernetes.list


# Add Docker's official GPG key:
sudo curl -fsSL https://download.docker.com/linux/ubuntu/gpg -o /etc/apt/keyrings/docker.asc
sudo chmod a+r /etc/apt/keyrings/docker.asc
# Add the repository to Apt sources:
echo \
  "deb [arch=$(dpkg --print-architecture) signed-by=/etc/apt/keyrings/docker.asc] https://download.docker.com/linux/ubuntu \
  $(. /etc/os-release && echo "$VERSION_CODENAME") stable" | \
  sudo tee /etc/apt/sources.list.d/docker.list > /dev/null


mkdir -p /etc/apt/apt.conf.d
echo "APT::Acquire::Retries \"3\";" > /etc/apt/apt.conf.d/80-retries

apt-get update
RES=$(apt-get install -y curl jq netcat make ipset moreutils socat 2>&1)
if [[ $RES == */var/lib/dpkg/lock* ]]; then
  echo "Fail to get dpkg lock.  Wait for any other package installation"
  echo "process to finish, then rerun this script"
  exit -1
fi

APTOPTS="--allow-downgrades --allow-change-held-packages --allow-unauthenticated --ignore-hold "

for PKG in kubeadm docker-ce-cli; do
  INSTALLED_VERSION=$(dpkg --list |grep ${PKG} |tr -s " " |cut -f3 -d ' ')
  if [ ! -z ${INSTALLED_VERSION} ]; then
    if [ "${PKG}" == "kubeadm" ]; then
      kubeadm reset -f
      apt-get purge -y $APTOPTS kubeadm kubelet kubectl kubernetes-cni

      rm -rf /etc/cni/net.d
      rm -rf ~/.kube
      rm -rf /etc/cni
      rm -rf /etc/kubernetes
      rm -rf /var/lib/etcd
      rm -rf /var/lib/kubelet
    else
      docker stop $(docker ps -a -q)
      docker rm $(docker ps -a -q)
      apt-get purge -y $APTOPTS docker-ce docker-ce-cli containerd.io docker-buildx-plugin docker-compose-plugin

      # Delete Docker directories
      rm -rf /var/lib/docker /etc/docker
      rm /etc/apparmor.d/docker
      rm -rf /var/run/docker.sock
      rm -rf /usr/bin/docker-compose
      rm -rf /var/run/docker
      rm -rf /var/lib/docker
      rm -rf /var/lib/containerd
      systemctl stop docker.socket
    fi
  fi
done
apt-get autoclean -y
apt-get autoremove -y

##
## SETUP CONTAINER RUNTIMES
##

cat <<EOF | sudo tee /etc/modules-load.d/k8s.conf
overlay
br_netfilter
EOF

sudo modprobe overlay
sudo modprobe br_netfilter

# sysctl params required by setup, params persist across reboots
cat <<EOF | sudo tee /etc/sysctl.d/k8s.conf
net.bridge.bridge-nf-call-iptables  = 1
net.bridge.bridge-nf-call-ip6tables = 1
net.ipv4.ip_forward                 = 1
EOF

# Apply sysctl params without reboot
sudo sysctl --system

if [ -z ${DOCKERVERSION} ]; then
  apt-get install -y $APTOPTS docker-ce docker-ce-cli containerd.io docker-buildx-plugin docker-compose-plugin
else
  apt-get install -y $APTOPTS docker-ce=$DOCKERVERSION docker-ce-cli=$DOCKERVERSION containerd.io docker-buildx-plugin docker-compose-plugin
fi
cat > /etc/docker/daemon.json <<EOF
{
  "exec-opts": ["native.cgroupdriver=systemd"],
  "log-driver": "json-file",
  "log-opts": {
    "max-size": "100m"
  },
  "storage-driver": "overlay2"
}
EOF

rm /etc/containerd/config.toml
containerd config default > /etc/containerd/config.toml
sed -i 's/SystemdCgroup \= false/SystemdCgroup \= true/g' /etc/containerd/config.toml
#sed -i 's@sandbox_image = \"registry.k8s.io\/pause:3.6\"@sandbox_image = \"registry.k8s.io\/pause:3.9\"@g' /etc/containerd/config.toml

mkdir -p /etc/systemd/system/docker.service.d
systemctl enable docker.service
systemctl enable containerd.service
systemctl daemon-reload
systemctl restart docker
systemctl restart containerd


if [ -z ${CNIVERSION} ]; then
  apt-get install -y $APTOPTS kubernetes-cni
else
  apt-get install -y $APTOPTS kubernetes-cni=${CNIVERSION}
fi

if [ -z ${KUBEVERSION} ]; then
  apt-get install -y $APTOPTS kubeadm kubelet kubectl
else
  apt-get install -y $APTOPTS kubeadm=${KUBEVERSION} kubelet=${KUBEVERSION} kubectl=${KUBEVERSION}
fi

DOWNLOAD_DIR="/usr/bin"
if [ -z ${KUSTOMIZEV} ]; then
  curl -s "https://raw.githubusercontent.com/kubernetes-sigs/kustomize/master/hack/install_kustomize.sh"  | bash -s -- ${DOWNLOAD_DIR}
else
  curl -s "https://raw.githubusercontent.com/kubernetes-sigs/kustomize/master/hack/install_kustomize.sh"  | bash -s -- ${KUSTOMIZEV} ${DOWNLOAD_DIR}
fi

apt-mark hold kubernetes-cni kubelet kubeadm kubectl docker-ce docker-ce-cli containerd.io docker-buildx-plugin docker-compose-plugin

systemctl enable --now kubelet

kubeadm config images pull --kubernetes-version=${KUBEV}


NODETYPE="master"
if [ "$NODETYPE" == "master" ]; then

  if [[ ${KUBEV} == 1.2* ]]; then
    cat <<EOF > /root/config.yaml
apiVersion: kubeadm.k8s.io/v1beta3
kubernetesVersion: v${KUBEV}
kind: ClusterConfiguration
networking:
  dnsDomain: cluster.local
  podSubnet: 192.168.0.0/16
  serviceSubnet: 172.16.0.0/12
---
apiVersion: kubeproxy.config.k8s.io/v1alpha1
kind: KubeProxyConfiguration
mode: ipvs
ipvs:
  excludeCIDRs: ["10.0.0.0/8"]
---
kind: KubeletConfiguration
apiVersion: kubelet.config.k8s.io/v1beta1
cgroupDriver: systemd
EOF
  else
    echo "Unsupported Kubernetes version requested.  Bail."
    exit
  fi

  cat <<EOF > /root/rbac-config.yaml
apiVersion: v1
kind: ServiceAccount
metadata:
  name: tiller
  namespace: kube-system
---
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRoleBinding
metadata:
  name: tiller
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: ClusterRole
  name: cluster-admin
subjects:
  - kind: ServiceAccount
    name: tiller
    namespace: kube-system
EOF


  kubeadm init --config /root/config.yaml

  cd /root
  rm -rf .kube
  mkdir -p .kube
  cp -i /etc/kubernetes/admin.conf /root/.kube/config
  chown root:root /root/.kube/config
  export KUBECONFIG=/root/.kube/config
  echo "KUBECONFIG=${KUBECONFIG}" >> /etc/environment

  kubectl get nodes --all-namespaces

  # we refer to version 0.18.1 because later versions use namespace kube-flannel instead of kube-system TODO
  #kubectl apply -f "https://raw.githubusercontent.com/flannel-io/flannel/v0.18.1/Documentation/kube-flannel.yml"
  #kubectl apply -f "https://raw.githubusercontent.com/flannel-io/flannel/v0.24.0/Documentation/kube-flannel.yml"
  #kubectl apply -f "https://raw.githubusercontent.com/flannel-io/flannel/master/Documentation/kube-flannel.yml"

  # Latest Calico that uses the kube-system namespace
  kubectl apply -f https://raw.githubusercontent.com/projectcalico/calico/v3.26.0/manifests/calico.yaml

  # Install Calico
  # NOTE: Due to the large size of the CRD bundle, kubectl apply might exceed request limits. Instead, use kubectl create or kubectl replace.
  # 1) Install the Tigera Calico operator and custom resource definitions.
  #kubectl create -f https://raw.githubusercontent.com/projectcalico/calico/v3.27.2/manifests/tigera-operator.yaml
  # 2) Install Calico by creating the necessary custom resource
  #curl -O https://raw.githubusercontent.com/projectcalico/calico/v3.27.2/manifests/custom-resources.yaml
  #sed -i 's@cidr: 192.168.0.0\/16@cidr: 10.244.0.0\/16@g' custom-resources.yaml
  #kubectl create -f custom-resources.yaml
  #kubectl create -f https://raw.githubusercontent.com/projectcalico/calico/v3.27.2/manifests/custom-resources.yaml

  wait_for_pods_running 9 kube-system

  kubectl taint nodes --all node-role.kubernetes.io/control-plane-
  kubectl taint nodes --all node-role.kubernetes.io/master-

  # Use Rancher to utilize local storage for nodes (needed for kubeflow)
  kubectl apply -f https://raw.githubusercontent.com/rancher/local-path-provisioner/v0.0.26/deploy/local-path-storage.yaml
  
  wait_for_pods_running 1 local-path-storage

  # Set as default storage class
  kubectl patch storageclass local-path -p '{"metadata": {"annotations":{"storageclass.kubernetes.io/is-default-class":"true"}}}'

  HELMV=$(cat /opt/config/helm_version.txt)
  HELMVERSION=${HELMV}
  if [ ! -e helm-v${HELMVERSION}-linux-amd64.tar.gz ]; then
    wget https://get.helm.sh/helm-v${HELMVERSION}-linux-amd64.tar.gz
  fi
  cd /root && rm -rf Helm && mkdir Helm && cd Helm
  tar -xvf ../helm-v${HELMVERSION}-linux-amd64.tar.gz
  mv linux-amd64/helm /usr/local/bin/helm

  cd /root

  rm -rf /root/.helm
#  if [[ ${KUBEV} == 1.16.* ]]; then
#    if [[ ${HELMVERSION} == 2.* ]]; then
#       helm init --service-account tiller --override spec.selector.matchLabels.'name'='tiller',spec.selector.matchLabels.'app'='helm' --output yaml > /tmp/helm-init.yaml
#       sed 's@apiVersion: extensions/v1beta1@apiVersion: apps/v1@' /tmp/helm-init.yaml > /tmp/helm-init-patched.yaml
#       kubectl apply -f /tmp/helm-init-patched.yaml
#    fi
#  else
#    if [[ ${HELMVERSION} == 2.* ]]; then
#       helm init --service-account tiller
#    fi
#  fi
#  if [[ ${HELMVERSION} == 2.* ]]; then
#     helm init -c
#     export HELM_HOME="$(pwd)/.helm"
#     echo "HELM_HOME=${HELM_HOME}" >> /etc/environment
#  fi

  while ! helm version; do
    echo "Waiting for Helm to be ready"
    sleep 15
  done

  echo "Preparing a master node (lower ID) for using local FS for PV"
  PV_NODE_NAME=$(kubectl get nodes |grep master | cut -f1 -d' ' | sort | head -1)
  kubectl label --overwrite nodes $PV_NODE_NAME local-storage=enable
  if [ "$PV_NODE_NAME" == "$(hostname)" ]; then
    mkdir -p /opt/data/dashboard-data
  fi

  echo "Done with master node setup"
fi


if [[ ! -z "" && ! -z "" ]]; then 
  echo " " >> /etc/hosts
fi
if [[ ! -z "" && ! -z "" ]]; then 
  echo " " >> /etc/hosts
fi
if [[ ! -z "" && ! -z "helm.ricinfra.local" ]]; then 
  echo " helm.ricinfra.local" >> /etc/hosts
fi

if [[ "1" -gt "100" ]]; then
  cat <<EOF >/etc/ca-certificates/update.d/helm.crt

EOF
fi

if [[ "1" -gt "100" ]]; then
  mkdir -p /etc/docker/certs.d/:
  cat <<EOF >/etc/docker/ca.crt

EOF
  cp /etc/docker/ca.crt /etc/docker/certs.d/:/ca.crt

  service docker restart
  systemctl enable docker.service
  docker login -u  -p  :
  docker pull :/whoami:0.0.1
fi
