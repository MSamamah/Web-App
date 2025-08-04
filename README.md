# DevOpsify – Web Application

This is a simple Golang-based web application developed **as the base app for a DevOps demonstration project**.

---

## 1.0 Containerization

We use a **multi-stage Dockerfile** to build a minimal and secure image of our Go web application. Multi-stage builds help reduce image size by separating the build environment from the runtime image, improving both performance and security.

### 📄 Dockerfile (multi-stage build)
The `Dockerfile` compiles the Go app in one stage, and copies the binary to a clean, smaller base image in the second stage.

### 🔧 Commands
```bash
# Build the Docker image
docker build -t <your-dockerhub-username>/web-app:latest .

# Run locally (optional)
docker run -p 8080:8080 <your-dockerhub-username>/web-app:latest
```

---

## 2.0 Kubernetes Manifests

We define raw Kubernetes manifests inside the `k8s/` directory to deploy the application on a Kubernetes cluster.

### Includes:
- `Deployment`: Creates and manages Pods for the application
- `Service`: Exposes the deployment as a network service
- `Ingress`: Routes external traffic to the service
- `Namespace`, `ConfigMap`, and `Secrets`: Manages configuration and sensitive data

### 🔧 Commands
```bash
kubectl apply -f k8s/namespace.yaml
kubectl apply -f k8s/configmap.yaml
kubectl apply -f k8s/secret.yaml
kubectl apply -f k8s/deployment.yaml
kubectl apply -f k8s/service.yaml
kubectl apply -f k8s/ingress.yaml
```

---

## 3.0 Continuous Integration (CI) – GitHub Actions

CI is set up using GitHub Actions with a workflow file `.github/workflows/cicd.yaml`.

### ✅ CI Workflow Stages:
- Checkout code
- Lint Go code using `golangci-lint`
- Run unit tests
- Build Docker image
- Push Docker image to DockerHub
- Update Helm chart `values.yaml` with new image tag using `${{ github.run_id }}`

---

## 4.0 Continuous Deployment (CD) – Argo CD (GitOps)

Argo CD automates Kubernetes deployments using GitOps principles. Once changes are pushed to the GitHub repository, Argo CD detects them and syncs them to the cluster automatically.

### 🛠 Install Argo CD:
```bash
kubectl create namespace argocd
kubectl apply -n argocd -f https://raw.githubusercontent.com/argoproj/argo-cd/stable/manifests/install.yaml

# Expose Argo CD UI via LoadBalancer
kubectl patch svc argocd-server -n argocd -p '{"spec": {"type": "LoadBalancer"}}'
```

After that, you can access Argo CD via the LoadBalancer’s external IP, log in, and connect your Git repository to monitor and sync deployments.

---

## 5.0 Kubernetes Cluster on AWS EKS

The application is deployed to an Amazon EKS (Elastic Kubernetes Service) cluster, which acts as the production environment.

### 📦 Prerequisites
- `kubectl`
- `eksctl`
- `AWS CLI`

### 🔧 Commands
```bash
# Create a new EKS cluster
eksctl create cluster --name demo-cluster --region us-east-1

# Delete the EKS cluster
eksctl delete cluster --name demo-cluster --region us-east-1
```

---

## 6.0 Helm Chart for Multi-Environment Deployments

Helm is used to package Kubernetes resources into reusable and configurable charts. This helps the DevOps team deploy the application to multiple environments like `dev`, `qa`, and `prod` without changing core manifest files.

### 🔧 Commands
```bash
# Install or upgrade the app using Helm
helm upgrade --install web-app ./helm/web-app-chart --namespace web-app --create-namespace
```

By modifying `values.yaml`, you can customize:
- Image version
- Replica count
- Resource limits
- Hostnames
- Environment-specific configs

---

## 7.0 Ingress Controller Setup

To expose the app externally (outside the Kubernetes cluster), we configure an **Ingress Controller**, such as NGINX or AWS ALB.

- The `ingress.yaml` defines the rules and hosts for routing HTTP traffic.
- Ingress Controller dynamically provisions an external **Load Balancer**.
- Requests from users on the internet are routed to specific **services** inside the cluster, which connect to the **pods** running the application.

This setup ensures scalability, SSL termination, and cleaner URL paths.

---

## 8.0 Summary of Technologies

| Category         | Tool/Service           |
|------------------|------------------------|
| Language         | Go                     |
| Containerization | Docker (Multi-stage)   |
| Orchestration    | Kubernetes (EKS)       |
| CI               | GitHub Actions         |
| CD               | Argo CD (GitOps)       |
| Package Mgmt     | Helm                   |
| Hosting          | AWS                    |

---

## 👤 Author

**Muhammad Samamah**  
[LinkedIn](https://www.linkedin.com/in/muhammad-samamah-698a6a304)  
[GitHub](https://github.com/MSamamah)
