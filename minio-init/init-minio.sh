set -e

mc alias set myminio http://minio:9000 tarantul tarantul123

mc mb myminio/upload || true
mc mb myminio/results || true
mc event add myminio/upload arn:minio:sqs::PRIMARY:webhook --event put,delete