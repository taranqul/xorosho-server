import requests
import os

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
    
    os.remove(file_name)

def upload_file_existed(presigned_url: str, file_name: str):

    with open(file_name, "rb") as f:
        response = requests.put(presigned_url, data=f)

    if response.status_code == 200:
        print(f"Файл '{file_name}' успешно загружен!")
    else:
        print(f"Ошибка загрузки: {response.status_code} - {response.text}")

if __name__ == "__main__":

    presigned_url = "http://127.0.0.1:9000/upload/b11eea7a-47ec-4d6a-ae0b-b14ae94e2e57_edit.mp4?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=tarantul%2F20260122%2Fus-east-1%2Fs3%2Faws4_request&X-Amz-Date=20260122T172916Z&X-Amz-Expires=3600&X-Amz-SignedHeaders=host&X-Amz-Signature=7c4507fe81972778ba50a9ee563f2ada4c715f575b55174a6240e94e5e4a7abd"
    file_name = "b11eea7a-47ec-4d6a-ae0b-b14ae94e2e57_edit.mp4"

    upload_file_existed(presigned_url, file_name)
