# Install EKS

> Please follow the prerequisites doc before this.

## Install an EKS cluster with `eksctl`

```bash
eksctl create cluster --name demo-cluster --region us-east-1 
```

## Delete the cluster

```bash
eksctl delete cluster --name demo-cluster --region us-east-1
```
