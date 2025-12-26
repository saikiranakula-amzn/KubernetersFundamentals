# Taint Commands for Nodes

# Add taints to nodes (run these commands on your cluster)

# Taint node for production environment
kubectl taint nodes <node-name> environment=production:NoSchedule

# Taint node for GPU workloads
kubectl taint nodes <node-name> node-type=gpu:NoExecute

# Remove taints (add minus sign at the end)
kubectl taint nodes <node-name> environment=production:NoSchedule-
kubectl taint nodes <node-name> node-type=gpu:NoExecute-

# View node taints
kubectl describe node <node-name> | grep Taints
