import requests
import sys

def upload_file(presigned_url: str, file_name: str, content: str = "Hello, MinIO!"):
    # Создаем файл и записываем содержимое
    with open(file_name, "w") as f:
        f.write(content)

    # Открываем файл в бинарном режиме для отправки
    with open(file_name, "rb") as f:
        response = requests.put(presigned_url, data=f)

    if response.status_code == 200:
        print(f"Файл '{file_name}' успешно загружен!")
    else:
        print(f"Ошибка загрузки: {response.status_code} - {response.text}")

if __name__ == "__main__":

    presigned_url = "http://localhost:9000/upload/20a5c14d-43fc-4aad-ac86-dbdce5406f80_edit2.txt?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=tarantul%2F20251219%2Fus-east-1%2Fs3%2Faws4_request&X-Amz-Date=20251219T041337Z&X-Amz-Expires=600&X-Amz-SignedHeaders=host&X-Amz-Signature=1eb4f1fb894fe127636dfb624dac345e44b4c8e7d82bfc6d3dfa5040fc7aefb5"
    file_name = "20a5c14d-43fc-4aad-ac86-dbdce5406f80_edit2.txt"

    upload_file(presigned_url, file_name)
