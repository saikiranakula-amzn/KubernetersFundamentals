# Node Labeling Commands
# needed for node selector and node affinity

# Add labels to nodes for nodeSelector and affinity
kubectl label nodes <node-name> disktype=ssd
kubectl label nodes <node-name> environment=production
kubectl label nodes <node-name> zone=us-west-1a

# Remove labels (add minus sign at the end)
kubectl label nodes <node-name> disktype-
kubectl label nodes <node-name> environment-

# View node labels
kubectl get nodes --show-labels
kubectl describe node <node-name> | grep Labels
