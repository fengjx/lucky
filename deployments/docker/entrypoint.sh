#!/bin/bash -e

echo "[`date`] WORK_DIR: ${WORK_DIR}"

APP_ENV=${APP_ENV:-test}

export APP_CONFIG=${WORK_DIR}/conf/app-${APP_ENV}.yml

echo "[`date`] Starting server..."
echo "[`date`] Custom config: ${APP_CONFIG}"
./${APP_NAME} ${APP_CONFIG}
