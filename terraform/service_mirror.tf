# Copyright 2020 the Exposure Notifications Server authors
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#      http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

#
# Create and deploy the service
#

resource "google_service_account" "mirror" {
  project      = data.google_project.project.project_id
  account_id   = "en-mirror-sa"
  display_name = "Exposure Notification Mirror"
}

resource "google_service_account_iam_member" "cloudbuild-deploy-mirror" {
  service_account_id = google_service_account.mirror.id
  role               = "roles/iam.serviceAccountUser"
  member             = "serviceAccount:${data.google_project.project.number}@cloudbuild.gserviceaccount.com"

  depends_on = [
    google_project_service.services["cloudbuild.googleapis.com"],
  ]
}

resource "google_secret_manager_secret_iam_member" "mirror-db" {
  for_each = toset([
    "sslcert",
    "sslkey",
    "sslrootcert",
    "password",
  ])

  secret_id = google_secret_manager_secret.db-secret[each.key].id
  role      = "roles/secretmanager.secretAccessor"
  member    = "serviceAccount:${google_service_account.mirror.email}"
}

resource "google_project_iam_member" "mirror-observability" {
  for_each = toset([
    "roles/cloudtrace.agent",
    "roles/logging.logWriter",
    "roles/monitoring.metricWriter",
    "roles/stackdriver.resourceMetadata.writer",
  ])

  project = var.project
  role    = each.key
  member  = "serviceAccount:${google_service_account.mirror.email}"
}

resource "google_cloud_run_service" "mirror" {
  name     = "mirror"
  location = var.cloudrun_location

  autogenerate_revision_name = true

  metadata {
    annotations = merge(
      local.default_service_annotations,
      var.default_service_annotations_overrides,
      lookup(var.service_annotations, "mirror", {}),
    )
  }
  template {
    spec {
      service_account_name = google_service_account.mirror.email

      containers {
        image = "gcr.io/${data.google_project.project.project_id}/github.com/google/exposure-notifications-server/mirror:initial"

        resources {
          limits = {
            cpu    = "2000m"
            memory = "1G"
          }
        }

        dynamic "env" {
          for_each = merge(
            local.common_cloudrun_env_vars,

            // This MUST come last to allow overrides!
            lookup(var.service_environment, "_all", {}),
            lookup(var.service_environment, "mirror", {}),
          )

          content {
            name  = env.key
            value = env.value
          }
        }
      }

      container_concurrency = 5
      // 30 seconds less than cloud scheduler maximum.
      timeout_seconds = 570
    }

    metadata {
      annotations = merge(
        local.default_revision_annotations,
        var.default_revision_annotations_overrides,
        lookup(var.revision_annotations, "mirror", {}),
      )
    }
  }

  depends_on = [
    google_project_service.services["run.googleapis.com"],
    google_secret_manager_secret_iam_member.mirror-db,
    null_resource.build,
    null_resource.migrate,
  ]

  lifecycle {
    ignore_changes = [
      metadata[0].annotations["client.knative.dev/user-image"],
      metadata[0].annotations["run.googleapis.com/client-name"],
      metadata[0].annotations["run.googleapis.com/client-version"],
      metadata[0].annotations["run.googleapis.com/ingress-status"],
      metadata[0].annotations["serving.knative.dev/creator"],
      metadata[0].annotations["serving.knative.dev/lastModifier"],
      metadata[0].labels["cloud.googleapis.com/location"],
      template[0].metadata[0].annotations["client.knative.dev/user-image"],
      template[0].metadata[0].annotations["run.googleapis.com/client-name"],
      template[0].metadata[0].annotations["run.googleapis.com/client-version"],
      template[0].metadata[0].annotations["serving.knative.dev/creator"],
      template[0].metadata[0].annotations["serving.knative.dev/lastModifier"],
      template[0].spec[0].containers[0].image,
    ]
  }
}


#
# Create scheduler job to invoke the service on a fixed interval.
#

resource "google_service_account" "mirror-invoker" {
  project      = data.google_project.project.project_id
  account_id   = "en-mirror-invoker-sa"
  display_name = "Exposure Notification Mirror Invoker"
}

resource "google_cloud_run_service_iam_member" "mirror-invoker" {
  project  = google_cloud_run_service.mirror.project
  location = google_cloud_run_service.mirror.location
  service  = google_cloud_run_service.mirror.name
  role     = "roles/run.invoker"
  member   = "serviceAccount:${google_service_account.mirror-invoker.email}"
}

resource "google_cloud_scheduler_job" "mirror-invoke" {
  name             = "mirror-invoke"
  region           = var.cloudscheduler_location
  schedule         = "*/5 * * * *"
  time_zone        = var.cloud_scheduler_timezone
  attempt_deadline = "600s"

  retry_config {
    retry_count = 3
  }

  http_target {
    http_method = "GET"
    uri         = "${google_cloud_run_service.mirror.status.0.url}/"
    oidc_token {
      audience              = google_cloud_run_service.mirror.status.0.url
      service_account_email = google_service_account.mirror-invoker.email
    }
  }

  depends_on = [
    google_app_engine_application.app,
    google_cloud_run_service.mirror,
    google_cloud_run_service_iam_member.mirror-invoker,
    google_project_service.services["cloudscheduler.googleapis.com"],
  ]
}
