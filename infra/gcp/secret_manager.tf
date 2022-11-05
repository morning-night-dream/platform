resource "google_secret_manager_secret" "secret_mndp_core_api_key" {
  secret_id = "mndp-core-api-key"

  labels = {
    label = "mndp"
  }

  replication {
    automatic = true
  }
}

resource "google_secret_manager_secret_version" "secret_mndp_core_api_key_version" {
  secret = google_secret_manager_secret.secret_mndp_core_api_key.id

  secret_data = var.secret_core_api_key
}

resource "google_secret_manager_secret" "secret_mndp_core_database_url" {
  secret_id = "mndp-core-database-url"

  labels = {
    label = "mndp"
  }

  replication {
    automatic = true
  }
}

resource "google_secret_manager_secret_version" "secret_mndp_core_database_url_version" {
  secret = google_secret_manager_secret.secret_mndp_core_database_url.id

  secret_data = var.secret_core_database_url
}
