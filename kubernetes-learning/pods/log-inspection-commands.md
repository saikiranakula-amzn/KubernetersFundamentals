# Inspect logs for container log-x and redirect warnings to file
kubectl logs dev-pod-dind-878516 -c log-x | grep -i warning > /opt/dind-878516_logs.txt

# Alternative: Get all logs and filter warnings
kubectl logs dev-pod-dind-878516 -c log-x --tail=-1 | grep -i warning > /opt/dind-878516_logs.txt

# Verify the file was created
cat /opt/dind-878516_logs.txt
