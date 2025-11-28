set -e

mc alias set myminio http://minio:9000 tarantul tarantul123

mc mb myminio/upload || true