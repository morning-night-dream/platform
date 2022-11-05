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
