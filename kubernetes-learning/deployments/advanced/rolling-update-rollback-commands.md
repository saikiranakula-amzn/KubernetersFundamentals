# Rolling Update and Rollback Commands

## Initial Deployment
```bash
kubectl apply -f nginx-rolling-update-rollback.yaml
kubectl rollout status deployment/nginx-rollback-demo
```

## Perform Rolling Update
```bash
# Update to nginx:1.16
kubectl set image deployment/nginx-rollback-demo nginx=nginx:1.16 --record

# Update to nginx:1.17
kubectl set image deployment/nginx-rollback-demo nginx=nginx:1.17 --record

# Update to nginx:latest
kubectl set image deployment/nginx-rollback-demo nginx=nginx:latest --record
```

## Monitor Updates
```bash
kubectl rollout status deployment/nginx-rollback-demo
kubectl get pods -l app=nginx-rollback
kubectl describe deployment nginx-rollback-demo
```

## View Rollout History
```bash
kubectl rollout history deployment/nginx-rollback-demo
kubectl rollout history deployment/nginx-rollback-demo --revision=2
```

## Rollback Operations
```bash
# Rollback to previous version
kubectl rollout undo deployment/nginx-rollback-demo

# Rollback to specific revision
kubectl rollout undo deployment/nginx-rollback-demo --to-revision=2

# Check rollback status
kubectl rollout status deployment/nginx-rollback-demo
```

## Verify Rollback
```bash
kubectl describe deployment nginx-rollback-demo
kubectl get pods -l app=nginx-rollback -o wide
```

## Advanced Rollback Scenarios
```bash
# Pause rollout during update
kubectl rollout pause deployment/nginx-rollback-demo

# Resume paused rollout
kubectl rollout resume deployment/nginx-rollback-demo

# Restart deployment (recreate all pods)
kubectl rollout restart deployment/nginx-rollback-demo
```
