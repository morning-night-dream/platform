#!/bin/bash

# トリガー名からトリガーIDを取得する
get_trigger_value() {
    local trigger_name=$1
    local project=$2
    local ret_value
    
    ret_value=$(gcloud beta builds triggers describe "${trigger_name}" --format "value(id)" --project "${project}")
    echo "${ret_value}"
}

# 最後にビルドが成功した時のcommit hashを取得する
nth_successful_commit() {
  local limit=$1
  local apply_trigger_name=$2
  local project=$3
  local apply_trigger_id
  local nth_successful_build
  local nth_successful_commit

  apply_trigger_id=$(get_trigger_value "$apply_trigger_name" "$project")
  nth_successful_build=$(gcloud builds list --filter "buildTriggerId=$apply_trigger_id AND STATUS=(SUCCESS)" --format "value(id)" --limit=$limit --project "$project" | awk "NR==1")

  if [ -z "$nth_successful_build" ]; then
    exit 1
  fi

  nth_successful_commit=$(gcloud builds describe "$nth_successful_build" --format "value(substitutions.COMMIT_SHA)" --project "$project") || exit 1
  
  echo "${nth_successful_commit}"
}

deploy() {
  local service_name=$1
  local gcr_repo=$2
  local image_name=$3
  local image_tag=$4
  local revision_service_account=$5
  local api_key=$6
  local database_url=$7

  gcloud run deploy $service_name \
    --image $gcr_repo/$image_name:$image_tag \
    --region asia-northeast1 \
    --cpu 1 \
    --memory 128Mi \
    --max-instances 1 \
    --service-account $revision_service_account \
    --set-secrets API_KEY=$api_key,DATABASE_URL=$database_url \
    --allow-unauthenticated
}

service_name=$1
gcr_repo=$2
image_name=$3
image_tag=$4
revision_service_account=$5
api_key=$6
database_url=$7
project_id=$8
build_trigger_name=$9
trigger_build_config_path=$10

# 直近50回のビルドから、実行中のビルドトリガーの最後に成功したcommitを見つける
last_successed_commit_hash=$(nth_successful_commit 50 "$build_trigger_name" "$project_id")

if [ -z "$last_successed_commit_hash" ]; then
  # 最後にビルドが成功したcommitが見つけられなかった場合は無条件でデプロイを実行する
  deploy $service_name $gcr_repo $image_name $image_tag $revision_service_account $api_key $database_url
else
  diff=$(git diff --name-only $last_successed_commit_hash HEAD -- $trigger_build_config_path)

  # trigger_build_config_pathのファイルに差分がないときはデプロイしない
  if [ -z $(echo $diff | grep $trigger_build_config_path) ]; then
    exit 1
  fi

  deploy $service_name $gcr_repo $image_name $image_tag $revision_service_account $api_key $database_url
fi
