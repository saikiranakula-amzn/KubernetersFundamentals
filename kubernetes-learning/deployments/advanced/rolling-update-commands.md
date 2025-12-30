# Rolling Update Commands

## Deploy the initial deployment
```bash
kubectl apply -f nginx-rolling-update-rollback.yaml
```

## Check deployment status
```bash
kubectl get deployments
kubectl rollout status deployment/nginx-rollback-demo
```

## Update the deployment (trigger rolling update)
```bash
kubectl set image deployment/nginx-rollback-demo nginx=nginx:1.16
```

## Monitor the rolling update
```bash
kubectl rollout status deployment/nginx-rollback-demo
kubectl get pods -w
```

## Check rollout history
```bash
kubectl rollout history deployment/nginx-rollback-demo
```

## Scale the deployment
```bash
kubectl scale deployment nginx-rollback-demo --replicas=5
```

## Pause and resume rollout
```bash
kubectl rollout pause deployment/nginx-rollback-demo
kubectl rollout resume deployment/nginx-rollback-demo
```

## Clean up
```bash
kubectl delete -f nginx-rolling-update-rollback.yaml
```
