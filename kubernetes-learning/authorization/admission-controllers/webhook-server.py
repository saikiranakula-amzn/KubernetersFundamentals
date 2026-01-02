#!/usr/bin/env python3
import json
import base64
from flask import Flask, request, jsonify

app = Flask(__name__)

@app.route('/mutate', methods=['POST'])
def mutate():
    admission_review = request.get_json()
    
    # Extract the object from the admission request
    obj = admission_review["request"]["object"]
    
    # Create mutation patches
    patches = []
    
    # Add a label if it doesn't exist
    if "labels" not in obj.get("metadata", {}):
        patches.append({
            "op": "add",
            "path": "/metadata/labels",
            "value": {"mutated": "true"}
        })
    else:
        patches.append({
            "op": "add",
            "path": "/metadata/labels/mutated",
            "value": "true"
        })
    
    # Encode patches as base64
    patch_bytes = json.dumps(patches).encode()
    patch_b64 = base64.b64encode(patch_bytes).decode()
    
    # Return admission response
    return jsonify({
        "apiVersion": "admission.k8s.io/v1",
        "kind": "AdmissionReview",
        "response": {
            "uid": admission_review["request"]["uid"],
            "allowed": True,
            "patchType": "JSONPatch",
            "patch": patch_b64
        }
    })

@app.route('/validate', methods=['POST'])
def validate():
    admission_review = request.get_json()
    
    # Extract the object from the admission request
    obj = admission_review["request"]["object"]
    
    # Validation logic
    allowed = True
    message = ""
    
    # Example: Reject pods without resource limits
    if obj["kind"] == "Pod":
        containers = obj.get("spec", {}).get("containers", [])
        for container in containers:
            if "resources" not in container or "limits" not in container["resources"]:
                allowed = False
                message = f"Container {container['name']} must have resource limits"
                break
    
    # Return admission response
    return jsonify({
        "apiVersion": "admission.k8s.io/v1",
        "kind": "AdmissionReview",
        "response": {
            "uid": admission_review["request"]["uid"],
            "allowed": allowed,
            "status": {"message": message} if message else {}
        }
    })

@app.route('/health', methods=['GET'])
def health():
    return jsonify({"status": "healthy"})

if __name__ == '__main__':
    app.run(host='0.0.0.0', port=8443, ssl_context='adhoc')
