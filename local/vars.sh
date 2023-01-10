#!/bin/bash
export GOPRIVATE="gitlab.tn.ru"
export GONOSUMDB="gitlab.tn.ru"

###
### Backend
###
export TN_LOGLEVEL="debug"
export TN_LOG_OUTPUT="console"
export TN_IP=0.0.0.0
export TN_GRPC_PORT=8091
export TN_HTTP_PORT=3015

###
### Postgres
###
export TN_POSTGRES_DSN="user=postgres password=postgres host=0.0.0.0 port=5556 database=postgres sslmode=disable"
export TN_MIGRATION_MIGRATIONS=$(pwd)/"internal/datastore/migrations"
export TN_MIGRATION_VERSION_TABLE="public.schema_version"

###
### Profile
###
export TN_PROFILE_EMPLOYEE_API_URL="https://tn-profile-stage.tjdev.ru:9090"
export TN_PROFILE_EMPLOYEE_API_TOKEN="wvrabvlttxhnr0l0tvrrmk5pmdbnv1f6tfdjne0ywxrnr05rt0drefpqrxlnvgxtq2c9pqo"
export TN_PROFILE_CLIENT_ID="tn-booking-mini-app"
export TN_PROFILE_CLIENT_SECRET="TnpFMVpXWmhZakV0TXpVM09TMDBaREl4TFdJNE5HWXRNV0V3WmpFMk1qUmhNRGcwQ2c9PQ"
export TN_PROFILE_AUTH_URL="https://tn-profile-stage.tjdev.ru/api/v1/oauth/authorize"
export TN_PROFILE_TOKEN_URL="https://tn-profile-stage.tjdev.ru:8070/api/v1/oauth/token"
export TN_PROFILE_REDIRECT_URL="https://tn-booking-dev.tages.dev"
export TN_PROFILE_SCOPES="phone.read,employee.read,email.read"

###
### Sessions
###
export TN_ACCESS_EXPIRY="80000h"
export TN_REFRESH_EXPIRY="10000h"
export TN_SECRET="wv2doa4vpHVtbNbUv0wUvm01tm60nml"

###
### S3
###
export TN_AWS_S3_REGION="ru-central1"
export TN_AWS_S3_ENDPOINT="storage.yandexcloud.net"
export TN_AWS_S3_FILES_BUCKET="tn-booking"
export TN_AWS_S3_ACCESS_KEY_ID="SvnrD3tQYZqQt84NIkb4"
export TN_AWS_S3_SECRET_ACCESS_KEY="XAHUCKRjIA6gsMw-mSIj870PTLIvMmP1Ho8_83Yv"
export TN_AWS_S3_TEMPORARY_FILES_PREFIX="tmp/"

###
### Email sender
###
export TN_SKIP_TLS="false"
export TN_EMAIL_HOST="mail.tjump.ru"
export TN_EMAIL_PORT=587
export TN_EMAIL_USER="tn.booking.app@gmail.com"
export TN_EMAIL_PASSWORD="TnPass#10"
export TN_EMAIL_FROM="tn.booking.app@gmail.com"

###
### Notes
###
export TN_REPORTS_CHECKER_SCHEDULER_DAY=23
export TN_REPORTS_CHECKER_SCHEDULER_HOUR=16
export TN_REPORTS_CHECKER_SCHEDULER_MINUTES=42
export TN_MISSED_REPORTS_CHECKER_SCHEDULER_DAY=23
export TN_MISSED_REPORTS_CHECKER_SCHEDULER_HOUR=16
export TN_MISSED_REPORTS_CHECKER_SCHEDULER_MINUTES=42
export TN_ENABLE_CREATE_NOTES="true"
export TN_SEND_EMAIL_NOTES_PERIOD="10m"
export TN_SEND_LIFE_NOTES_PERIOD="10m"
export TN_URL_BASE_NOTE="https://tn-booking-dev.tages.dev"
export TN_BOOKING_VIEW_FRONT_ROUTE_NOTE="/booking/[[uuid]]"
export TN_ROOM_VIEW_FRONT_ROUTE_NOTE="/room/[[uuid]]"
export TN_BOT_API_URL="https://bots-tn-life-dev.tages.dev/bots"
export TN_BOT_ID="a6c11aaf-b133-4256-8c9c-61afa71492a6"
export TN_BOT_TOKEN="Cfuy6LwptrKJgRM7u9kX7l4Bi3EZzyZKmgz4zXGD"