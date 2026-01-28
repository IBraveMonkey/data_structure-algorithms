**Docker** — платформа для контейнеризации, позволяющая упаковывать приложения и их зависимости в легковесные контейнеры. Контейнеры работают изолированно, но используют ядро хоста, что делает их быстрее виртуальных машин. Этот конспект объясняет, как Docker работает под капотом, его компоненты (**daemon**, **volumes**, **networks**), ключевые команды и **docker-compose**.

---

## Как работает Docker: Под капотом

Docker использует технологии Linux (или других ОС) для создания изолированных сред — контейнеров. Основные строительные блоки: **образы**, **контейнеры**, **Docker Daemon**, **CLI**, и подсистемы ядра.

### Основные концепции

- **Образ (Image)**: Неизменяемый шаблон (аналог чертежа), содержащий приложение, библиотеки и конфигурации. Например, образ nginx включает веб-сервер Nginx.
- **Контейнер**: Запущенный экземпляр образа, изолированная среда с собственными процессами, файловой системой и сетью.
- **Docker Daemon**: Сервер, управляющий контейнерами, образами, сетями и томами.
- **Docker CLI**: Интерфейс командной строки для взаимодействия с daemon’ом (docker run, docker pull).

### Архитектура Docker

1. **Клиент (Docker CLI)**: Пользователь отправляет команды (например, docker run).
2. **Docker Daemon (dockerd)**: Сервер, выполняющий операции (запуск контейнеров, управление образами).
3. **Containerd**: Runtime-слой, управляющий жизненным циклом контейнеров.
4. **runc**: Низкоуровневый runtime для запуска контейнеров с использованием Linux-технологий.
5. **Ядро ОС**: Предоставляет изоляцию через **namespaces**, **cgroups**, и **union filesystems**.

### Технологии Linux

- **Namespaces**: Изолируют ресурсы (PID, network, mount, user), создавая "песочницу" для контейнера.
- **Cgroups**: Ограничивают ресурсы (CPU, память, диск) для контейнеров.
- **Union Filesystems (OverlayFS)**: Обеспечивают слоистую файловую систему для образов, где слои кэшируются и переиспользуются.

### Жизненный цикл контейнера

1. Пользователь запускает docker run <image>.
2. CLI отправляет запрос daemon’у.
3. Daemon через containerd и runc создает контейнер:
   - Настраивает namespaces/cgroups.
   - Монтирует файловую систему из образа.
   - Запускает команду (например, /bin/bash).

4. Контейнер выполняется, пока не завершится или не будет остановлен.

---

## Ключевые компоненты Docker

### 1. Docker Daemon (dockerd)

- **Что это**: Фоновый процесс, управляющий всеми операциями Docker (создание контейнеров, управление сетями, томами).
- **Как работает**:
  - Слушает запросы от CLI через UNIX-сокет (/var/run/docker.sock) или TCP.
  - Взаимодействует с containerd для запуска контейнеров.
  - Управляет образами, кэшем и метаданными.
- **Пример**:
  ```bash
  sudo systemctl status docker
  # Проверяет статус daemon’а
  ```

### 2. Образы (Images)

- **Что это**: Слоистый архив, содержащий приложение и зависимости. Каждый слой — изменения файловой системы.
- **Как создаются**:
  - Через Dockerfile (инструкции для сборки).
  - Слои кэшируются для скорости.
- **Пример**:
  ```dockerfile
  FROM ubuntu:20.04
  RUN apt-get update && apt-get install -y python3
  CMD ["python3", "--version"]
  ```
  ```bash
  docker build -t my-python .
  # Создает образ my-python
  ```

### 3. Контейнеры

- **Что это**: Запущенные образы, изолированные среды с процессами.
- **Как работают**:
  - Используют namespaces для изоляции.
  - Cgroups ограничивают ресурсы.
  - OverlayFS монтирует файловую систему.
- **Пример**:
  ```bash
  docker run -it ubuntu:20.04 /bin/bash
  # Запускает контейнер с интерактивной оболочкой
  ```

### 4. Volumes (Тома)

- **Что это**: Механизм для хранения данных вне контейнера, сохраняющий их между запусками.
- **Типы**:
  - **Named volumes**: Управляются Docker’ом.
  - **Bind mounts**: Привязка к хостовой директории.
  - **tmpfs**: Временные данные в памяти.
- **Как работают**:
  - Тома монтируются в контейнер (например, /data).
  - Данные сохраняются даже после удаления контейнера.
- **Пример**:
  ```bash
  docker volume create my-volume
  docker run -v my-volume:/data ubuntu:20.04
  # Монтирует том my-volume в /data
  docker run -v /host/data:/data ubuntu:20.04
  # Bind mount: /host/data → /data
  ```

### 5. Networks (Сети)

- **Что это**: Механизм для связи контейнеров и хоста.
- **Типы сетей**:
  - **Bridge**: Локальная сеть для контейнеров (по умолчанию).
  - **Host**: Контейнер использует сеть хоста.
  - **None**: Без сети.
  - **Overlay**: Для связи между хостами (в Docker Swarm).
- **Как работают**:
  - Docker создает виртуальные сети с DNS для контейнеров.
  - Контейнеры общаются по именам или IP.
- **Пример**:
  ```bash
  docker network create my-network
  docker run --network my-network --name app1 -d nginx
  docker run --network my-network --name app2 -d redis
  # app1 и app2 могут общаться по именам
  ```

---

## Основные команды Docker

### Управление образами

- **Скачать образ**:
  ```bash
  docker pull nginx:latest
  ```
- **Создать образ**:
  ```bash
  docker build -t my-app:1.0 .
  ```
- **Просмотр образов**:
  ```bash
  docker images
  ```

### Управление контейнерами

- **Запустить контейнер**:
  ```bash
  docker run -d --name my-app -p 8080:80 nginx
  # -d: фоновый режим, -p: проброс порта
  ```
- **Просмотр контейнеров**:
  ```bash
  docker ps     # Активные
  docker ps -a  # Все
  ```
- **Остановить/удалить**:
  ```bash
  docker stop my-app
  docker rm my-app
  ```

### Управление томами

- **Создать том**:
  ```bash
  docker volume create my-volume
  ```
- **Просмотр томов**:
  ```bash
  docker volume ls
  ```
- **Удалить том**:
  ```bash
  docker volume rm my-volume
  ```

### Управление сетями

- **Создать сеть**:
  ```bash
  docker network create my-network
  ```
- **Просмотр сетей**:
  ```bash
  docker network ls
  ```
- **Подключить контейнер**:
  ```bash
  docker network connect my-network my-app
  ```

### Общие команды

- **Информация о системе**:
  ```bash
  docker info
  ```
- **Очистка**:
  ```bash
  docker system prune
  # Удаляет неиспользуемые образы, контейнеры, сети
  ```

---

## Docker Compose

### Что это

**Docker Compose** — инструмент для определения и запуска нескольких контейнеров с помощью YAML-файла. Упрощает управление приложениями, состоящими из нескольких сервисов (например, веб-сервер + база данных).

### Как работает

- **docker-compose.yml**: Описывает сервисы, сети, тома.
- **Команда** docker-compose: Запускает/управляет всеми сервисами.

### Пример: Веб-приложение с Nginx и Redis

**docker-compose.yml**:

```yaml
version: "3.8"
services:
  web:
    image: nginx:latest
    ports:
      - "8080:80"
    volumes:
      - web-data:/usr/share/nginx/html
    networks:
      - app-network
  redis:
    image: redis:latest
    networks:
      - app-network
volumes:
  web-data:
networks:
  app-network:
    driver: bridge
```

**Команды**:

```bash
docker-compose up -d
# Запускает сервисы в фоновом режиме
docker-compose ps
# Показывает статус
docker-compose down
# Останавливает и удаляет
```

**Объяснение**:

- **web**: Контейнер Nginx, порт 8080 → 80, том web-data для данных.
- **redis**: Контейнер Redis, подключен к той же сети.
- **app-network**: Bridge-сеть для связи сервисов.
- **web-data**: Том для хранения данных Nginx.

### Преимущества

- Упрощает запуск многоконтейнерных приложений.
- Автоматически создает сети и тома.
- Легко масштабируется (docker-compose up --scale web=3).

### Когда использовать

- Локальная разработка и тестирование.
- Микросервисные приложения.
- CI/CD с несколькими сервисами.

---

## Как Docker обрабатывает запросы: Пошагово

1. **Пользователь вводит команду**:

   ```bash
   docker run -d -p 8080:80 nginx
   ```

2. **CLI отправляет запрос** daemon’у через /var/run/docker.sock.
3. **Daemon проверяет образ**:
   - Если nginx отсутствует, скачивает из Docker Hub.

4. **Containerd создает контейнер**:
   - runc настраивает namespaces, cgroups, OverlayFS.
   - Запускает команду (например, nginx -g 'daemon off;').

5. **Daemon настраивает сеть**:
   - Пробрасывает порт 8080 → 80.
   - Подключает контейнер к bridge-сети.

6. **Контейнер работает**, логи доступны через docker logs.

---

## Примеры сценариев

### 1. Запуск веб-сервера

```bash
docker run -d --name web -p 80:80 nginx
# Запускает Nginx, доступен на localhost:80
```

### 2. Приложение с базой данных

**docker-compose.yml**:

```yaml
version: "3.8"
services:
  app:
    image: node:16
    volumes:
      - ./app:/app
    working_dir: /app
    command: npm start
    ports:
      - "3000:3000"
    depends_on:
      - db
  db:
    image: postgres:13
    environment:
      POSTGRES_PASSWORD: example
    volumes:
      - db-data:/var/lib/postgresql/data
volumes:
  db-data:
```

```bash
docker-compose up -d
# Запускает Node.js приложение и PostgreSQL
```

### 3. Управление томами

```bash
docker volume create data
docker run -v data:/app/data -d my-app
# Данные сохраняются в томе data
```

---

## Рекомендации

### Лучшие практики

- **Минимизируйте образы**:
  - Используйте легкие базовые образы (например, alpine).
  - Удаляйте временные файлы в Dockerfile:
    ```dockerfile
    RUN apt-get update && apt-get install -y package && rm -rf /var/lib/apt/lists/*
    ```
- **Используйте .dockerignore**:
  - Исключайте ненужные файлы (.git, node_modules).
- **Ограничивайте ресурсы**:
  ```bash
  docker run --memory="512m" --cpus="0.5" nginx
  ```
- **Логируйте правильно**:
  - Используйте docker logs или централизованные системы (ELK, Fluentd).

### Полезные команды

- **Остановка всех контейнеров**:
  ```bash
  docker stop $(docker ps -q)
  ```
- **Удаление неиспользуемых ресурсов**:
  ```bash
  docker system prune -a
  ```
- **Инспекция контейнера**:
  ```bash
  docker inspect my-app
  ```

### Когда использовать docker-compose

- Для локальной разработки (например, веб + база данных).
- Для тестирования микросервисов.
- Избегайте в продакшене (используйте Kubernetes или Swarm).

---

## Как это работает в Go-приложении

**Dockerfile для Go**:

```dockerfile
FROM golang:1.20-alpine AS builder
WORKDIR /app
COPY . .
RUN go build -o myapp main.go

FROM alpine:latest
COPY --from=builder /app/myapp /usr/local/bin/myapp
CMD ["myapp"]
```

**main.go**:

```go
package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello, Docker!")
	})
	http.ListenAndServe(":8080", nil)
}
```

**Запуск**:

```bash
docker build -t my-go-app .
docker run -d -p 8080:8080 my-go-app
# Доступно на localhost:8080
```

**docker-compose.yml**:

```yaml
version: "3.8"
services:
  app:
    image: my-go-app
    build:
      context: .
      dockerfile: Dockerfile
    ports:
      - "8080:8080"
    volumes:
      - app-data:/data
volumes:
  app-data:
```

```bash
docker-compose up -d
```

---

## Примечания

- **Daemon**: Центральный компонент, но containerd и runc делают основную работу.
- **Volumes/Networks**: Критичны для хранения данных и связи контейнеров.
- **Docker Compose**: Упрощает разработку, но не для продакшена.
- **Go**: Пример показывает, как контейнеризировать Go-приложение.
- **Код**: Минималистичные примеры для понимания. В продакшене добавьте логирование, health checks.
