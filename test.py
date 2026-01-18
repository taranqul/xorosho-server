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

def upload_file_existed(presigned_url: str, file_name: str,):

    with open(file_name, "rb") as f:
        response = requests.put(presigned_url, data=f)

    if response.status_code == 200:
        print(f"Файл '{file_name}' успешно загружен!")
    else:
        print(f"Ошибка загрузки: {response.status_code} - {response.text}")
    
    os.remove(file_name)

if __name__ == "__main__":

    presigned_url = "http://localhost:9000/upload/ba1ffb4b-2be5-4991-b748-d8bff1b5329f_edit2.txt?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=tarantul%2F20260107%2Fus-east-1%2Fs3%2Faws4_request&X-Amz-Date=20260107T210531Z&X-Amz-Expires=600&X-Amz-SignedHeaders=host&X-Amz-Signature=e08c9c2f9bd67e5e1ef58f592be50bc0b393e42b47c592f697c0e4ed78236374"
    file_name = "ba1ffb4b-2be5-4991-b748-d8bff1b5329f_edit2.txt"

    upload_file(presigned_url, file_name)
