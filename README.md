# DevOpsify – Web Application

This is a simple Golang-based web application developed **as the base app for a DevOps demonstration project**.

---

## 1.0 Containerization

We use a **multi-stage Dockerfile** to build a minimal and secure image of our Go web application.

### 📄 Dockerfile (multi-stage build)
Reduces image size and attack surface by separating the build and runtime stages.

### 🔧 Commands
```bash
# Build the Docker image
docker build -t <your-dockerhub-username>/web-app:latest .

# Run locally (optional)
docker run -p 8080:8080 <your-dockerhub-username>/web-app:latest
```

---

## 2.0 Kubernetes Manifests

We define raw Kubernetes manifests for:

- `Deployment`
- `Service`
- `Ingress`
- `Namespace`, `ConfigMap`, and `Secrets`

These are stored in the `k8s/` directory and can be used manually to deploy the app.

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

CI is implemented using GitHub Actions located in `.github/workflows/cicd.yaml`.

### ✅ Stages include:
- Checkout code
- Lint Go code
- Run unit tests
- Build Docker image
- Push Docker image to DockerHub
- Update Helm chart `values.yaml` image tag

---

## 4.0 Continuous Deployment (CD) – Argo CD (GitOps)

We use **Argo CD** to implement GitOps-based continuous deployment.

### 🛠 Install Argo CD:
```bash
kubectl create namespace argocd
kubectl apply -n argocd -f https://raw.githubusercontent.com/argoproj/argo-cd/stable/manifests/install.yaml

# Expose UI via LoadBalancer
kubectl patch svc argocd-server -n argocd -p '{"spec": {"type": "LoadBalancer"}}'

# For Windows (escape characters)
kubectl patch svc argocd-server -n argocd -p '{\"spec\": {\"type\": \"LoadBalancer\"}}'
```

---

## 5.0 Kubernetes Cluster on AWS EKS

We provision the target platform using **Amazon EKS** via `eksctl`.

### 📦 Prerequisites
- [kubectl](https://kubernetes.io/docs/tasks/tools/)
- [eksctl](https://eksctl.io/)
- [AWS CLI](https://docs.aws.amazon.com/cli/latest/userguide/install-cliv2.html)

### 🔧 Commands
```bash
# Create EKS cluster
eksctl create cluster --name demo-cluster --region us-east-1

# Delete EKS cluster
eksctl delete cluster --name demo-cluster --region us-east-1
```

---

## 6.0 Helm Chart for Multi-Environment Deployments

Helm enables dynamic deployments to `dev`, `qa`, and `prod` environments by passing values through `values.yaml`.

### 📁 Folder Structure
```
helm/
└── web-app-chart/
    ├── Chart.yaml
    ├── values.yaml
    └── templates/
        ├── deployment.yaml
        ├── service.yaml
        ├── ingress.yaml
```

### 🔧 Commands
```bash
# Install/upgrade using Helm
helm upgrade --install web-app ./helm/web-app-chart --namespace web-app --create-namespace
```

---

## 7.0 Ingress Controller Setup

We configure an **Ingress Controller** to expose the app to the internet through an AWS LoadBalancer.

Make sure:
- Your cluster has an Ingress Controller (e.g., AWS ALB Controller or NGINX Ingress Controller).
- Your `ingress.yaml` is correctly configured with hostname/rules.

---

## 8.0 Summary of Technologies

| Category        | Tool/Service           |
|----------------|------------------------|
| Language        | Go                     |
| Containerization| Docker (Multi-stage)   |
| Orchestration   | Kubernetes (EKS)       |
| CI              | GitHub Actions         |
| CD              | Argo CD (GitOps)       |
| Package Mgmt    | Helm                   |
| Hosting         | AWS                    |

---

## 👤 Author

**Muhammad Samamah**  
[LinkedIn](https://www.linkedin.com/in/muhammad-samamah-698a6a304)  
[GitHub](https://github.com/MSamamah)

