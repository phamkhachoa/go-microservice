terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 5.0"
    }
  }
}

provider "aws" {
  region = var.region
}

module "vpc" {
  source = "./modules/vpc"

  vpc_cidr             = var.vpc_cidr
  availability_zones   = var.availability_zones
  private_subnet_cidrs = var.private_subnet_cidrs
  public_subnet_cidrs  = var.public_subnet_cidrs
  cluster_name         = var.cluster_name
}

module "eks" {
  source = "./modules/eks"

  cluster_name    = var.cluster_name
  cluster_version = var.cluster_version
  vpc_id          = module.vpc.vpc_id
  subnet_ids      = module.vpc.private_subnet_ids
  node_groups     = var.node_groups
}

# Add this to your existing main.tf file to use the database module

module "database" {
  source = "./modules/database"
  region = var.region
  subnet_ids = module.vpc.public_subnet_ids
  vpc_id = module.vpc.vpc_id
}

# Now use the outputs from EKS module as inputs to the ALB module
module "alb" {
  source = "./modules/alb"

  region      = var.region
  cluster_name = module.eks.cluster_name
  vpc_id      = module.vpc.vpc_id

  # These values should come from your EKS module
  cluster_endpoint = module.eks.cluster_endpoint
  cluster_certificate_authority_data = module.eks.cluster_certificate_authority_data
  oidc_provider_url = module.eks.oidc_provider_url
  eks_cluster = module.eks.eks_cluster

  # Make sure the ALB module depends on the EKS module
  # depends_on = [module.eks]
}

module "elasticache" {
  source = "./modules/elasticache"
  regison = var.region
  subnet_ids = module.vpc.private_subnet_ids
  vpc_cidr =  module.vpc.private_subnet_ids
  vpc_id = module.vpc.vpc_id
}


