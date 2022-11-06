# TODO for で回せないかな。。。
# @ref https://zenn.dev/wim/articles/terraform_loop
resource "google_cloudbuild_trigger" "core_trigger" {
  name     = "${var.project_prefix}-${var.project_env}-core-cloud-build-trigger"
  location = "global"

  # 事前にcloud buildに該当githubリポジトリを接続しておくこと
  # @ref https://cloud.google.com/architecture/managing-infrastructure-as-code?hl=ja#granting_permissions_to_your_cloud_build_service_account
  github {
    owner = var.git_repository_owner
    name  = var.git_repository_name
    push {
      branch = "^main$"
    }
  }

  include_build_logs = "INCLUDE_BUILD_LOGS_WITH_STATUS"

  substitutions = {
    _API_KEY                  = google_secret_manager_secret_version.secret_core_api_key_version.name
    _DATABASE_URL             = google_secret_manager_secret_version.secret_core_database_url_version.name
    _REVISION_SERVICE_ACCOUNT = google_service_account.cloud_run_invoker.email
    _SERVICE_NAME             = "${var.project_prefix}-${var.project_env}-core-cloud-run"
  }

  filename = var.core_cloud_build_file_path

  service_account = google_service_account.cloud_build_invoker.id

  depends_on = [
    google_service_account.cloud_build_invoker,
    google_service_account.cloud_run_invoker
  ]
}
