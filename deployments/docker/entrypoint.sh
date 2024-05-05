#!/bin/bash -e

echo "[`date`] WORK_DIR: ${WORK_DIR}"
ls -la ${WORK_DIR}

APP_ENV=${APP_ENV:-test}

export LUCHEN_LOG_DIR=${WORK_DIR}/logs
export APP_CONFIG=${WORK_DIR}/conf/app-${APP_ENV}.yml

echo "[`date`] starting server..."
echo "[`date`] custom config: ${APP_CONFIG}"
echo "[`date`] log dir: ${LUCHEN_LOG_DIR}"
./${APP_NAME} ${APP_CONFIG}
