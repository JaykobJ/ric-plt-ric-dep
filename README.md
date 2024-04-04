This repository was created as part of my master's thesis Assessing Openness and Practical Challenges of Machine Learning Pipeline Development in OpenRAN. A part of the thesis project was to create a light weight platform to create and test E2E ML models as xapps in near-real-time RIC.

The repository extends the ORAN OSC [near-real-time RIC deployment](https://github.com/o-ran-sc/ric-plt-ric-dep) with Kubeflow.

**NOTE!**

*Kubeflow is running in the same cluster as the near-real-time RIC but is not communicating with the near-real-time RIC. This part is missing due to time constraints, but the project can be freely used and extended.*

&nbsp;

# Installation instructions

We will install near-real-time RIC (E-release), ns-O-RAN and Kubeflow to the same one node cluster

*Original near-real-time RIC installation instructions can be found* [**here**](https://docs.o-ran-sc.org/projects/o-ran-sc-ric-plt-ric-dep/en/latest/installation-guides.html#modify-the-deployment-recipe)

&nbsp;

## Install required software

### Prerequisites

**NOTE!** OSC instructions assume a clean installation of **Ubuntu 20.04** (no k8s, no docker, no helm)

Start with loggin in as root `sudo -s` and cd to root folder `cd ~`

Make folder to keep all related files etc. for near-real-time RIC and other projects (ns-O-RAN & kubeflow) and navigate to the folder

```
mkdir ml_ric && cd ml_ric
```

&nbsp;

### Clone near-rt RIC repository

This repository contains all the deployment scripts and support files for near-real-time RIC.

**NOTE**: use latest branch and set up kube cluster by the E-release recipe (*explained below*)

```
git clone "https://gerrit.o-ran-sc.org/r/ric-plt/ric-dep"
```

&nbsp;

### Get required software and set up infra

Move to the RIC setup directory
```
cd ric-dep/bin
```

Install kubernetes, kubernetes-CNI, helm and docker.
```
./install_k8s_and_helm.sh
```

Check that kubernetes system pods are up and ready
```
kubectl get pods -n kube-system
```

&nbsp;

&nbsp;

## Install Kubeflow

Move to project root folder
```
cd ~/ml_ric
```

Clone Kubeflow repository
```
git clone -b v1.8-branch https://github.com/kubeflow/manifests.git
```

Move to manifest folder
```
cd manifests
```

Install Kubeflow (There will be a lot of error messages, but let it run. It will take some time to finish)
```
while ! kustomize build example | kubectl apply -f -; do echo "Retrying to apply resources"; sleep 10; done
```

Check that all pods are running
```
kubectl get pods -A
```

&nbsp;

&nbsp;

## Install O-RAN Software Community (OSC) near-rt RIC

Move to the RIC setup directory
```
cd ~/ml_ric/ric-dep/bin
```

Copy IP address of your VM or server running the platform
```
ip a
```

Paste the IP as a value in the recipe file for *ricip* AND *auxip* (They contain defualt ip 10.0.0.1)
```
vim ../RECIPE_EXAMPLE/example_recipe_oran_e_release.yaml
```

Install chartmuseum into helm and add ric-common templates
```
./install_common_templates_to_helm.sh
```

Now we can install all required images and set the RIC up and running.
```
./install -f ../RECIPE_EXAMPLE/example_recipe_oran_e_release.yaml
```

#### ! Make sure everything correct and running before connecting ns-3 RAN nodes to RIC !

List helm releases
```
helm list -A
```
#### Results should look like this:
><pre>
>NAME             	NAMESPACE	REVISION	UPDATED                                 	STATUS  	CHART               	APP VERSION
>r4-a1mediator    	ricplt   	1       	2024-04-03 14:34:37.549710262 +0300 EEST	deployed	a1mediator-3.0.0    	1.0        
>r4-alarmmanager  	ricplt   	1       	2024-04-03 14:35:12.471025344 +0300 EEST	deployed	alarmmanager-5.0.0  	1.0        
>r4-appmgr        	ricplt   	1       	2024-04-03 14:34:01.821119241 +0300 EEST	deployed	appmgr-3.0.0        	1.0        
>r4-dbaas         	ricplt   	1       	2024-04-03 14:33:53.158439245 +0300 EEST	deployed	dbaas-2.0.0         	1.0        
>r4-e2mgr         	ricplt   	1       	2024-04-03 14:34:19.632705015 +0300 EEST	deployed	e2mgr-3.0.0         	1.0        
>r4-e2term        	ricplt   	1       	2024-04-03 14:34:28.494874957 +0300 EEST	deployed	e2term-3.0.0        	1.0        
>r4-infrastructure	ricplt   	1       	2024-04-03 14:33:40.071328763 +0300 EEST	deployed	infrastructure-3.0.0	1.0        
>r4-o1mediator    	ricplt   	1       	2024-04-03 14:35:03.634028221 +0300 EEST	deployed	o1mediator-3.0.0    	1.0        
>r4-rtmgr         	ricplt   	1       	2024-04-03 14:34:10.922848413 +0300 EEST	deployed	rtmgr-3.0.0         	1.0        
>r4-submgr        	ricplt   	1       	2024-04-03 14:34:46.288478599 +0300 EEST	deployed	submgr-3.0.0        	1.0        
>r4-vespamgr      	ricplt   	1       	2024-04-03 14:34:55.010015324 +0300 EEST	deployed	vespamgr-3.0.0      	1.0
></pre>
&nbsp;

**The setting up of all the RIC platform pods can take a while. Follow the process with *watch* command**

Check pods of the RIC platform
```
watch -n 10 kubectl get pods -n ricplt
```
#### Results should look like this:
><pre>
>NAME                                                         READY   STATUS    RESTARTS       AGE
>deployment-ricplt-a1mediator-6f4984949f-dszvw                1/1     Running   0              79m
>deployment-ricplt-alarmmanager-97547b6b9-tvj2c               1/1     Running   0              78m
>deployment-ricplt-appmgr-d4cd89bb7-jl5s8                     1/1     Running   0              79m
>deployment-ricplt-e2mgr-76f8b966f7-jfw9g                     1/1     Running   0              79m
>deployment-ricplt-e2term-alpha-74b4f57ddd-q8b4t              1/1     Running   0              79m
>deployment-ricplt-o1mediator-6b9d766b55-rdzws                1/1     Running   0              78m
>deployment-ricplt-rtmgr-559bd7d64c-dsqch                     1/1     Running   16 (11m ago)   79m
>deployment-ricplt-submgr-5bcc5bbbc8-x54hs                    1/1     Running   0              79m
>deployment-ricplt-vespamgr-6fd4d9cd4b-jcqpz                  1/1     Running   0              78m
>r4-infrastructure-kong-747d695785-mbsxp                      2/2     Running   1 (11m ago)    80m
>r4-infrastructure-prometheus-alertmanager-6d64cfb4c4-nsf49   2/2     Running   0              80m
>r4-infrastructure-prometheus-server-76b464f7b7-fs52g         1/1     Running   0              80m
>statefulset-ricplt-dbaas-server-0                            1/1     Running   0              79m
></pre>
&nbsp;

Check pods of the RIC infra
```
watch -n 10 kubectl get pods -n ricinfra
```
#### Results should look like this:
><pre>
>NAME                                        READY   STATUS      RESTARTS   AGE
>deployment-tiller-ricxapp-6cd8c99b6-qklns   1/1     Running     0          80m
>tiller-secret-generator-r2zmq               0/1     Completed   0          80m
></pre>

&nbsp;

### Setup xapp installation requirements

1. **Install python3**
2. **Create a local helm repository with a port other than 8080 on host**
```
docker run --rm -u 0 -it -d -p 8090:8080 -e DEBUG=1 -e STORAGE=local -e STORAGE_LOCAL_ROOTDIR=/charts -v $(pwd)/charts:/charts chartmuseum/chartmuseum:latest
```
3. **Set up the environment variables for CLI connection using the same port as used above**
```
export CHART_REPO_URL=http://0.0.0.0:8090
```
4. **Install dms_cli tool**

>> Move to project root folder
```
cd ~/ml_ric
```

>> Clone app manager repository
```
git clone "https://gerrit.o-ran-sc.org/r/ric-plt/appmgr"
```

>> Change directory to xapp_onboarder
```
cd appmgr/xapp_orchestrater/dev/xapp_onboarder
```

>> If pip3 is not installed
```
apt update && apt install python3-pip
```

>> In case dms_cli binary is already installed, it can be uninstalled using following command
```
pip3 uninstall xapp_onboarder
```

>> Install xapp_onboarder
```
pip3 install ./
```

&nbsp;

&nbsp;

## Install ns-O-RAN

Create folder for nsoran related folders and files
```
cd ~/ml_ric && mkdir nsoran && cd nsoran
```

Copy text from [Dockerfile](https://github.com/wineslab/colosseum-near-rt-ric/blob/ns-o-ran/Dockerfile) and create dockerfile to setup ns-O-RAN ns-3 simulator
```
vim Dockerfile
```

Build docker Image (For close debugging I'm using log level 3)
```
docker build -t ns3 -f Dockerfile --build-arg log_level_e2sim=3 . --no-cache
```

Export image as .tar
```
docker save -o ~/ml_ric/nsoran/ns3.tar ns3
```

Import image into containerd in kubernets namespace
```
ctr -n=k8s.io images import ns3.tar
```

Create yaml file for deployment
```
vim ns-o-ran-pod.yaml
```

Add the following text to the yaml file:
```
apiVersion: v1
kind: Pod
metadata:
  name: ns-3-pod
  namespace: ricplt
spec:
  containers:
  - name: ns-3-pod
    image: docker.io/library/ns3:latest
    imagePullPolicy: Never
    command: ["sleep", "infinity"]
```

Create a ns3 pod to near-real-time RIC
```
kubectl apply -f ns-o-ran-pod.yaml
```

Check that ns3 is deployed
```
kubectl get pods -n ricplt
```

&nbsp;

### Install KPIMON xapp

This xapp will get KPI from the created ns3 RAN nodes

Create folder for xapps
```
cd ~/ml_ric && mkdir xapps && cd xapps
```

Clone kpimon xapp repository
```
git clone https://github.com/wineslab/ns-o-ran-scp-ric-app-kpimon.git kpimon -b libe2proto
```

Cd to the KPIMON xapp folder
```
cd kpimon
```

local Docker registry is supposed to be running at 127.0.0.1:5000
```
docker run -d -p 5000:5000 --name registry registry:2
```

Edit the installation script
```
vim ./launch_app.sh
```

Overwrite row 12 with the following text
```
dms_cli onboard ./xapp-descriptor/config.json ./xapp-descriptor/schema.json
```

Setup the xapp into the RIC platform
```
./launch_app.sh
```

&nbsp;

&nbsp;

## Interact with Kubeflow

The default way of accessing Kubeflow is via port-forward. This enables you to get started quickly without imposing any requirements on your environment. Run the following to port-forward Istio's Ingress-Gateway to local port 8080:
```
kubectl port-forward svc/istio-ingressgateway -n istio-system 8080:80
```

#### **NOTE** If the cluster is running on remote/cloud server/VM, use the following to acces Kubeflow dashboard

Setup a SOCKS5 tunnel via ssh (*for more elegant solution see [here](https://ketzacoatl.github.io/posts/2017-01-10-SOCKS-proxy-with-SSH-to-save-the-day.html)*)
```
ssh -D9999 user@remote-server
```

Test if you can access Kubeflow dashboard. (You should get 302 Found redirect status response code with location: "/dex/auth?client_id=kubeflow-oidc-authservice&redirect_uri=...")
```
curl --proxy socks5h://localhost:9999 -v http://localhost:8080
```

In case you run into issues, check that `/etc/ssh/sshd_config` has `AllowTcpForwarding yes` (or `AllowTcpForwarding local` for service under localhost) to allow the proxy.

I'm using Chrome as my main browser so I set up Firefow to proxy SOCKS5 at localhost:9999. *More information on [proxying localhost](https://www.developsec.com/2020/05/29/proxying-localhost-on-firefox/)*

1. Open Firefox
2. Navigate to Settings
3. Select 'Settings...' under Network Settings
4. Select 'Manual proxy configuration'
5. Set the following parameters: SOCKS Host = localhost, Port = 9999
6. Select 'Proxy DNS when using SOCKS v5'
7. On navigation panel type 'about:config'
8. On search panel type 'network.proxy'
9. Set 'network.proxy.allow_hijacking_localhost' to true
10. Now you should be able to access Kubeflow dashboard on 'localhost:8080'
11. Default user's credentials are 'user\<at\>example.com' and '12341234'


&nbsp;

&nbsp;

## Run ns3

Get E2Term IP addr from the following output
```
kubectl get pods -n ricplt -o wide | grep "e2term" | awk '{print $6}'
```

Open bash in ns3 container
```
kubectl exec -ti -n ricplt ns-3-pod bash
```

Navigate to ns3-mmwave-oran directory
```
cd ns3-mmwave-oran/
```

**NOTE** to run ns3 as a stand alone mode for data collection use the following command. This mode is not talking to the RIC. It outputs data in ceparate file for each RAN unit
```
./waf --run "scratch/scenario-zero.cc --simTime=10 --enableE2FileLogging=1"
```

Run ns3
```
NS_LOG="RicControlMessage" ./waf --run "scratch/scenario-zero.cc --simTime=40 --e2TermIp=<input-E2term-IP>"
```

&nbsp;

&nbsp;

## Run KPIMON xapp

Open bash in KPIMON xapp container
```
kubectl exec -ti -n ricxapp ricxapp-xappkpimon-6988bc5f56-cgpvz bash
```

Run xApp
```
./kpimon -f /opt/ric/config/config-file.json
```
