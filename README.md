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

### **ingress/** - Ingress controllers and rules
Advanced traffic routing and SSL termination (examples coming soon).

### **jobs/** - Job and CronJob definitions
Run tasks to completion or on schedules (examples coming soon).

### **volumes/** - Storage configurations (PV, PVC, StorageClass)
Persistent storage for your applications (examples coming soon).

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
