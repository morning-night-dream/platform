resource "google_service_account" "cloud_build_invoker" {
  account_id   = "${var.project_prefix}-cloud-build-invoker"
  display_name = "A service account for cloud build invoke"
}

# google_project_iam_bindingは既存のサービスアカウントのロールを剥奪する恐れがあるため、google_project_iam_memberを使う
# @ref https://zenn.dev/ptiringo/articles/7dd246fcaa73da19d5fb
resource "google_project_iam_member" "cloud_build_invoker_builder" {
  project = var.project_id
  role    = "roles/cloudbuild.builds.builder"
  member  = "serviceAccount:${google_service_account.cloud_build_invoker.email}"
}

resource "google_project_iam_member" "cloud_build_invoker_service_account_user" {
  project = var.project_id
  role    = "roles/iam.serviceAccountUser"
  member  = "serviceAccount:${google_service_account.cloud_build_invoker.email}"
}

resource "google_project_iam_member" "cloud_build_invoker_service_log_writer" {
  project = var.project_id
  role    = "roles/logging.logWriter"
  member  = "serviceAccount:${google_service_account.cloud_build_invoker.email}"
}

resource "google_project_iam_member" "cloud_build_invoker_service_run_admin" {
  project = var.project_id
  role    = "roles/run.admin"
  member  = "serviceAccount:${google_service_account.cloud_build_invoker.email}"
}

resource "google_service_account" "cloud_run_invoker" {
  account_id   = "${var.project_prefix}-cloud-run-invoker"
  display_name = "A service account for cloud run invoke"
}

resource "google_project_iam_member" "cloud_run_invoker_secret_accessor" {
  project = var.project_id
  role    = "roles/secretmanager.secretAccessor"
  member  = "serviceAccount:${google_service_account.cloud_run_invoker.email}"
}
