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

    presigned_url = "http://localhost:9000/upload/ea52affd-83b8-4145-acad-105a17736f30_edit.mp4?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=tarantul%2F20260119%2Fus-east-1%2Fs3%2Faws4_request&X-Amz-Date=20260119T170203Z&X-Amz-Expires=600&X-Amz-SignedHeaders=host&X-Amz-Signature=5bf2a1d969ac152849d68be9b3365bc01b80ab814675ffc6c375c4673c43560b"
    file_name = "ea52affd-83b8-4145-acad-105a17736f30_edit.mp4"

    upload_file_existed(presigned_url, file_name)
