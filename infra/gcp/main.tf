provider "google" {
  # Project ID
  project     = var.project_id
  credentials = var.gcp_credentials
  zone        = "asia-northeast1"
}

# Profect prefix
variable "project_prefix" {
  description = "GCP profect prefix"
  type        = string
}

# GCP credentials 
variable "gcp_credentials" {
  description = "GCP Credentials"
  type        = string
}

# Project ID
variable "project_id" {
  description = "GCP project id"
  type        = string
}

# APIキーのシークレット
variable "secret_core_api_key" {
  description = "Secret: API Key"
  type        = string
}

# データベースURLのシークレット
variable "secret_core_database_url" {
  description = "Secret: Database URL"
  type        = string
}

# Git Repository
variable "git_repository_name" {
  description = "Git Repository"
  type        = string
}

# Git Repository Owner
variable "git_repository_owner" {
  description = "Git Repository Owner"
  type        = string
}

# File path to core cloudbuild.yml 
variable "core_cloud_build_file_path" {
  description = "File path to core cloudbuild.yml"
  type        = string
}