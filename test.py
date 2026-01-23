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

def get_file(presigned_url: str, file_name: str):
    output_path = file_name

    with requests.get(presigned_url, stream=True) as r:
        r.raise_for_status()

        with open(output_path, "wb") as f:
            for chunk in r.iter_content(chunk_size=1024 * 1024):
                if chunk:
                    f.write(chunk)

    print("Файл скачан:", output_path)

if __name__ == "__main__":

    presigned_url = "http://127.0.0.1:9000/results/b11eea7a-47ec-4d6a-ae0b-b14ae94e2e57_result.mkv?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=tarantul%2F20260123%2Fus-east-1%2Fs3%2Faws4_request&X-Amz-Date=20260123T170220Z&X-Amz-Expires=3600&X-Amz-SignedHeaders=host&X-Amz-Signature=af932521279fbae4b2b2d8bd293cd4e36f5047d9c4118b1f534ea56f0d48ab7d"
    file_name = "b11eea7a-47ec-4d6a-ae0b-b14ae94e2e57_result.mkv"

    get_file(presigned_url, file_name)

