# Kubernetes Learning Repository

A structured collection of Kubernetes YAML definitions for learning core concepts. Each folder contains practical examples that you can deploy and experiment with to understand how Kubernetes works.

## Structure

### **pods/** - Pod definitions (basic/advanced)
The fundamental building blocks of Kubernetes. Start here to understand how containers run in Kubernetes.

**Basic Examples:**
- `nginx-pod.yaml` - Simple nginx pod to get you started
- `nginx-with-config-secret.yaml` - Shows how to inject configuration and secrets into pods

**Advanced Examples:**
- `security/nginx-security-context.yaml` - Security hardening with user/group settings and capabilities
- `resources/nginx-resource-limits.yaml` - CPU and memory requests/limits for resource management
- `resources/resource-limit-range.yaml` - Set default and maximum resource limits for containers
- `resources/memory-limit-range.yaml` - Memory-specific limits for better resource control
- `resources/resource-quota.yaml` - Namespace-level resource consumption limits
- `scheduling/nginx-with-toleration.yaml` - How pods can tolerate node taints for specialized scheduling
- `scheduling/nginx-node-selector.yaml` - Simple node selection using labels
- `scheduling/nginx-node-affinity.yaml` - Advanced node and pod scheduling rules with affinity/anti-affinity
- `scheduling/taint-commands.md` - Commands for node taints and scheduling control
- `scheduling/node-labeling-commands.md` - Commands for labeling nodes
- `containers/nginx-with-config-secret.yaml` - Shows how to inject configuration and secrets into pods
- `containers/nginx-with-init-container.yaml` - Init containers that run before main containers start
- `containers/nginx-with-sidecar.yaml` - Sidecar containers that run alongside main containers
- `service-accounts/nginx-service-account.yaml` - Identity for pods to access Kubernetes API
- `service-accounts/nginx-service-account-token.yaml` - Authentication token for service accounts
- `service-accounts/nginx-with-service-account.yaml` - Using service accounts for pod identity and API access

### **replicasets/** - ReplicaSet configurations
Ensures multiple copies of your pods are always running. Think of it as a pod babysitter.
- `nginx-replicaset.yaml` - Maintains 3 nginx pods automatically

### **daemonsets/** - DaemonSet configurations
Ensures a copy of a pod runs on every node in the cluster.
- `fluentd-daemonset.yaml` - Log collection agent running on all nodes

### **deployments/** - Deployment manifests (basic/advanced)
The preferred way to run applications. Deployments manage ReplicaSets and provide rolling updates.

**Basic Examples:**
- `nginx-deployment.yaml` - Standard deployment with 3 replicas and update capabilities
- `rolling-update-deployment.yaml` - Rolling update deployment with strategy configuration and service

**Advanced Examples:**
- `nginx-rolling-update-rollback.yaml` - Rolling update deployment with rollback capabilities
- `rolling-update-commands.md` - Commands for rolling updates and monitoring
- `rolling-update-rollback-commands.md` - Complete guide for rolling updates and rollbacks

### **services/** - Service definitions (basic/advanced)
How pods communicate with each other and the outside world.

**Basic Examples:**
- `nginx-clusterip-service.yaml` - Internal cluster communication only
- `nginx-nodeport-service.yaml` - External access via node ports
- `nginx-headless-service.yaml` - Headless service for StatefulSet with direct pod access

**Advanced Examples:**
- `nginx-loadbalancer-service.yaml` - AWS Load Balancer for production traffic

### **configmaps/** - ConfigMap examples
Store configuration data separately from your container images.
- `nginx-configmap.yaml` - Configuration files and environment variables

### **secrets/** - Secret configurations
Securely store sensitive data like passwords and API keys.
- `nginx-secret.yaml` - Base64 encoded credentials and tokens

### **namespaces/** - Namespace definitions
Organize and isolate resources within your cluster.
- `development-namespace.yaml` - Dev environment isolation
- `production-namespace.yaml` - Prod environment isolation
- `testing-namespace.yaml` - Test environment isolation

### **Resource Management Files**
- `resource-limit-range.yaml` - Set default and maximum resource limits for containers
- `memory-limit-range.yaml` - Memory-specific limits for better resource control
- `resource-quota.yaml` - Namespace-level resource consumption limits

### **Service Accounts & Security**
- `nginx-service-account.yaml` - Identity for pods to access Kubernetes API
- `nginx-service-account-token.yaml` - Authentication token for service accounts
- `taint-commands.md` - Commands for node taints and scheduling control

### **authorization/rbac/** - Role-Based Access Control
Control access to Kubernetes resources with fine-grained permissions.
- `basic-rbac.yaml` - ServiceAccount, Role, and RoleBinding for pod read access
- `cluster-rbac.yaml` - ClusterRole and ClusterRoleBinding for cluster-wide permissions

### **authorization/admission-controllers/** - Admission Controller configurations
Control and modify resource creation with admission controllers.
- `namespace-autoprovision.yaml` - NamespaceAutoProvision admission controller configuration
- `webhook-server.py` - Python webhook server with mutate and validate endpoints
- `webhook-deployment.yaml` - Deployment configuration for webhook server
- `webhook-config.yaml` - MutatingAdmissionWebhook and ValidatingAdmissionWebhook configurations
- `requirements.txt` - Python dependencies for webhook server

Since the kube-apiserver is running as pod you can check the process to see enabled and disabled plugins.

```bash
ps -ef | grep kube-apiserver | grep admission-plugins
```

### **ingress/** - Ingress controllers and network policies
Advanced traffic routing, SSL termination, and network security.
- `nginx-ingress-controller.yaml` - Basic NGINX ingress controller setup with NodePort service
- `nginx-ingress.yaml` - Basic ingress controller with host-based routing
- `ingress-network-policy.yaml` - Ingress network policy allowing specific pod and namespace access
- `egress-network-policy.yaml` - Egress network policy restricting outbound traffic

### **jobs/** - Job and CronJob definitions
Run tasks to completion or on schedules.
- `math-job.yaml` - Basic job that runs a calculation with completions and parallelism
- `math-cronjob.yaml` - CronJob that runs every 2 minutes with job history limits

### **volumes/** - Storage configurations (PV, PVC, StorageClass)
Persistent storage for your applications.
- `persistent-volumes/nginx-pv.yaml` - Basic persistent volume with hostPath storage
- `persistent-volume-claims/nginx-pvc.yaml` - Persistent volume claim requesting 1Gi storage
- `storage-classes/fast-ssd-storageclass.yaml` - StorageClass for dynamic provisioning with AWS EBS
- `nginx-pod-with-pvc.yaml` - Pod using persistent volume claim for nginx html directory

### **monitoring/** - Health checks and monitoring
Pod health monitoring with probes.
- `nginx-readiness-probe.yaml` - Readiness probes with httpGet, tcpSocket, and exec examples
- `nginx-liveness-probe.yaml` - Liveness probes with httpGet, tcpSocket, and exec examples
- `prometheus-monitoring.yaml` - Basic Prometheus setup for metrics collection
- `sample-app-metrics.yaml` - Sample application with Prometheus metrics annotations

**Prometheus Setup:**
```bash
# 1. Deploy Prometheus
kubectl apply -f prometheus-monitoring.yaml

# 2. Deploy sample app with metrics
kubectl apply -f sample-app-metrics.yaml

# 3. Access Prometheus UI
kubectl get svc -n monitoring
# Visit http://node-ip:nodeport

# 4. Check targets in Prometheus
# Go to Status > Targets to see discovered pods

# 5. Query metrics
# Use PromQL: up, node_cpu_seconds_total, etc.
```

**Metrics Server Setup:**
```bash
# Install metrics server for resource metrics
kubectl apply -f https://github.com/kubernetes-sigs/metrics-server/releases/latest/download/components.yaml

# Verify metrics server is running
kubectl get pods -n kube-system | grep metrics-server

# Check node metrics
kubectl top nodes

# Check pod metrics
kubectl top pods
```

**How it works:**
- Prometheus automatically discovers pods with `prometheus.io/scrape: "true"` annotation
- Scrapes metrics from `/metrics` endpoint on specified port
- ConfigMap defines scraping rules and service discovery
- Sample app exposes system metrics via node-exporter

### **crds/** - Custom Resource Definitions
Extend Kubernetes API with custom resources.
- `webapp-crd.yaml` - Custom resource definition for WebApp with validation schema
- `webapp-controller.go` - Basic Go controller for WebApp CRD
- `go.mod` - Go module dependencies for the controller

### **helm-charts/** - Helm Charts
Package and deploy applications with Helm.
- `webapp/` - Basic Helm chart with deployment, service, PV, PVC, and secret

### **kustomize/** - Kustomize configurations
Customize Kubernetes YAML for different environments.
- `base/` - Base nginx deployment configuration
- `overlays/dev/` - Development environment (1 replica)
- `overlays/staging/` - Staging environment (2 replicas)  
- `overlays/prod/` - Production environment (5 replicas)

**Deploy Helm Chart:**
```bash
# Package chart into tar format
helm package ./helm-charts/webapp/

# Install from packaged chart
helm install my-webapp webapp-0.1.0.tgz

# Install the chart from directory
helm install my-webapp ./helm-charts/webapp/

# Check deployment status
helm status my-webapp
kubectl get all -l app=webapp

# Upgrade with custom values
helm upgrade my-webapp ./helm-charts/webapp/ --set replicaCount=3

# Rollback to previous version
helm rollback my-webapp

# Rollback to specific revision
helm rollback my-webapp 1

# View release history
helm history my-webapp

# Uninstall the chart
helm uninstall my-webapp

# Dry run to see generated manifests
helm install my-webapp ./helm-charts/webapp/ --dry-run --debug
```

**Deploy and Use CRDs:**
```bash
# 1. Apply the CRD to extend Kubernetes API
kubectl apply -f webapp-crd.yaml

# 2. Verify CRD is installed
kubectl get crd webapps.example.com

# 3. Create a WebApp custom resource
kubectl apply -f webapp-crd.yaml

# 4. List WebApp resources
kubectl get webapps
kubectl get wa  # using short name

# 5. Describe WebApp resource
kubectl describe webapp my-webapp

# 6. Build and deploy controller (optional)
go mod tidy
go build -o webapp-controller webapp-controller.go
# Create container image and deploy to cluster

# 7. Delete WebApp resource
kubectl delete webapp my-webapp

# 8. Remove CRD (deletes all WebApp resources)
kubectl delete crd webapps.example.com
```

**Operator Framework (Alternative Approach):**
```bash
# Install Operator SDK
curl -LO https://github.com/operator-framework/operator-sdk/releases/download/v1.32.0/operator-sdk_linux_amd64
chmod +x operator-sdk_linux_amd64 && sudo mv operator-sdk_linux_amd64 /usr/local/bin/operator-sdk

# Create new operator project
operator-sdk init --domain=example.com --repo=github.com/example/webapp-operator

# Create API and controller
operator-sdk create api --group=apps --version=v1 --kind=WebApp --resource --controller

# Build and deploy operator
make docker-build docker-push IMG=webapp-operator:latest
make deploy IMG=webapp-operator:latest

# Create custom resource
kubectl apply -f config/samples/apps_v1_webapp.yaml
```

## Imperative vs Declarative Commands

While you would be working mostly the declarative way – using definition files, imperative commands can help in getting one time tasks done quickly, as well as generate a definition template easily. This would help save considerable amount of time during your exams.

Before we begin, familiarize with the two options that can come in handy while working with the below commands:

**--dry-run**: By default as soon as the command is run, the resource will be created. If you simply want to test your command, use the `--dry-run=client` option. This will not create the resource, instead, tell you whether the resource can be created and if your command is right.

**-o yaml**: This will output the resource definition in YAML format on screen.

Use the above two in combination to generate a resource definition file quickly, that you can then modify and create resources as required, instead of creating the files from scratch.

### POD
Create an NGINX Pod:
```bash
kubectl run nginx --image=nginx
```

Generate POD Manifest YAML file (-o yaml). Don't create it(–dry-run):
```bash
kubectl run nginx --image=nginx --dry-run=client -o yaml
```

### Deployment
Create a deployment:
```bash
kubectl create deployment --image=nginx nginx
```

Generate Deployment YAML file (-o yaml). Don't create it(–dry-run):
```bash
kubectl create deployment --image=nginx nginx --dry-run=client -o yaml
```

Generate Deployment with 4 Replicas:
```bash
kubectl create deployment nginx --image=nginx --replicas=4
```

You can also scale a deployment using the kubectl scale command:
```bash
kubectl scale deployment nginx --replicas=4
```

Another way to do this is to save the YAML definition to a file and modify:
```bash
kubectl create deployment nginx --image=nginx --dry-run=client -o yaml > nginx-deployment.yaml
```

### Service
Create a Service named redis-service of type ClusterIP to expose pod redis on port 6379:
```bash
kubectl expose pod redis --port=6379 --name redis-service --dry-run=client -o yaml
```
(This will automatically use the pod's labels as selectors)

Or:
```bash
kubectl create service clusterip redis --tcp=6379:6379 --dry-run=client -o yaml
```
(This will not use the pods labels as selectors, instead it will assume selectors as app=redis. You cannot pass in selectors as an option. So it does not work very well if your pod has a different label set. So generate the file and modify the selectors before creating the service)

Create a Service named nginx of type NodePort to expose pod nginx's port 80 on port 30080 on the nodes:
```bash
kubectl expose pod nginx --port=80 --name nginx-service --type=NodePort --dry-run=client -o yaml
```
(This will automatically use the pod's labels as selectors, but you cannot specify the node port. You have to generate a definition file and then add the node port in manually before creating the service with the pod.)

Or:
```bash
kubectl create service nodeport nginx --tcp=80:80 --node-port=30080 --dry-run=client -o yaml
```
(This will not use the pods labels as selectors)

Both the above commands have their own challenges. While one of it cannot accept a selector the other cannot accept a node port. I would recommend going with the `kubectl expose` command. If you need to specify a node port, generate a definition file using the same command and manually input the nodeport before creating the service.

### Secret
Create a Secret with literal values:
```bash
kubectl create secret generic db-secret --from-literal=DB_Host=sql01 --from-literal=DB_User=root --from-literal=DB_Password=password123
```

## Usage

Deploy any YAML file:
```bash
kubectl apply -f <path-to-yaml-file>
```

View resources:
```bash
kubectl get pods
kubectl get deployments
kubectl get services
```

Clean up:
```bash
kubectl delete -f <path-to-yaml-file>
```

## Helm Installation

**Install Helm:**
```bash
# macOS
brew install helm

# Linux
curl https://raw.githubusercontent.com/helm/helm/main/scripts/get-helm-3 | bash

# Windows (using Chocolatey)
choco install kubernetes-helm

# Verify installation
helm version
```

**Basic Helm Commands:**
```bash
# Add a repository
helm repo add bitnami https://charts.bitnami.com/bitnami

# Update repositories
helm repo update

# Search for charts
helm search repo nginx

# Install a chart
helm install my-nginx bitnami/nginx

# List releases
helm list

# Upgrade a release
helm upgrade my-nginx bitnami/nginx

# Uninstall a release
helm uninstall my-nginx

# Create your own chart
helm create my-chart
```
