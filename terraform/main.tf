
provider "kubernetes" {
  host                   = data.aws_eks_cluster.cluster.endpoint
  cluster_ca_certificate = base64decode(data.aws_eks_cluster.cluster.certificate_authority[0].data)
  exec {
    api_version = "client.authentication.k8s.io/v1beta1"
    args = ["--name", "sandbox-de-4-v121-blue", "--profile", "pd-testing"]
    command     = "../bin/eks-token-darwin-arm64"
  }
}
provider "aws" {
  profile = "pd-testing"
  region  = "eu-west-1"
}

data "aws_eks_cluster" "cluster" {
  name = "sandbox-de-4-v121-blue"
}

data "kubernetes_all_namespaces" "allns" {}

output "all-ns" {
  value = data.kubernetes_all_namespaces.allns.namespaces
}
