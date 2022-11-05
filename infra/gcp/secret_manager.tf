resource "google_secret_manager_secret" "secret-mndp-core-api-key" {
  secret_id = "mndp-core-api-key"

  labels = {
    label = "mndp"
  }

  replication {
    automatic = true
  }
}

resource "google_secret_manager_secret_version" "secret-mndp-core-api-key-version" {
  secret = google_secret_manager_secret.secret-mndp-core-api-key.id

  secret_data = var.secret_core_api_key
}

resource "google_secret_manager_secret" "secret-mndp-core-database-url" {
  secret_id = "mndp-core-database-url"

  labels = {
    label = "mndp"
  }

  replication {
    automatic = true
  }
}

resource "google_secret_manager_secret_version" "secret-mndp-core-database-url-version" {
  secret = google_secret_manager_secret.secret-mndp-core-database-url.id

  secret_data = var.secret_core_database_url
}
