# Deploy the initial deployment
kubectl apply -f nginx-deploy-example.yaml

# Verify deployment
kubectl get deployments
kubectl rollout status deployment/nginx-deploy

# Upgrade to nginx:1.17
kubectl set image deployment/nginx-deploy nginx=nginx:1.17 --record

# Monitor the rolling update
kubectl rollout status deployment/nginx-deploy
kubectl get pods -w

# Check rollout history
kubectl rollout history deployment/nginx-deploy

# Rollback to previous version
kubectl rollout undo deployment/nginx-deploy

# Verify rollback
kubectl rollout status deployment/nginx-deploy
kubectl describe deployment nginx-deploy | grep Image
