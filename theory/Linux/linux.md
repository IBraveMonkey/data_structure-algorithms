# Linux: Основы работы с системой

## Содержание
1. [Введение в Linux](#1-введение-в-linux)
2. [Архитектура Linux](#2-архитектура-linux)
3. [Файловая система](#3-файловая-система)
4. [Управление файлами](#4-управление-файлами)
5. [Работа с текстом](#5-работа-с-текстом)
6. [Управление процессами](#6-управление-процессами)
7. [Права доступа](#7-права-доступа)
8. [Сетевые утилиты](#8-сетевые-утилиты)
9. [Управление пользователями](#9-управление-пользователями)
10. [Практические рецепты](#10-практические-рецепты)

---

## 1. Введение в Linux

**Linux** — это семейство Unix-подобных операционных систем с открытым исходным кодом, основанных на ядре Linux.

### Ключевые особенности
- **Открытый исходный код**: Любой может изучать, модифицировать и распространять код
- **Многопользовательская система**: Поддержка множества пользователей одновременно
- **Многозадачность**: Параллельное выполнение множества процессов
- **Безопасность**: Продуманная система прав и изоляции процессов
- **Стабильность**: Работает месяцами и годами без перезагрузки

### Популярные дистрибутивы

```mermaid
graph TD
    Linux[Linux Kernel]
    Linux --> Debian[Debian]
    Linux --> RedHat[Red Hat]
    Linux --> Arch[Arch Linux]
    
    Debian --> Ubuntu[Ubuntu]
    Debian --> Mint[Linux Mint]
    
    RedHat --> Fedora[Fedora]
    RedHat --> CentOS[CentOS]
    
    Ubuntu --> Server[Ubuntu Server]
    Ubuntu --> Desktop[Ubuntu Desktop]
    
    style Linux fill:#f9d71c,stroke:#000,stroke-width:3px
    style Ubuntu fill:#e95420,stroke:#000
    style Debian fill:#d70a53,stroke:#000
```

> [!NOTE]
> **Ubuntu** и **Debian** наиболее популярны для серверов. **Fedora** и **Arch** — для разработчиков и энтузиастов.

---

## 2. Архитектура Linux

Linux состоит из нескольких уровней абстракции:

```mermaid
graph TB
    subgraph UserSpace["User Space (Пользовательское пространство)"]
        Apps[Приложения: bash, vim, docker]
        LibC[Системные библиотеки: glibc]
    end
    
    subgraph KernelSpace["Kernel Space (Пространство ядра)"]
        SysCalls[System Calls API]
        ProcessMgmt[Управление процессами]
        MemMgmt[Управление памятью]
        FileSys[Файловые системы]
        Network[Сетевой стек]
        Drivers[Драйверы устройств]
    end
    
    Hardware[Оборудование: CPU, RAM, Диски, Сеть]
    
    Apps --> SysCalls
    LibC --> SysCalls
    SysCalls --> ProcessMgmt
    SysCalls --> MemMgmt
    SysCalls --> FileSys
    SysCalls --> Network
    Drivers --> Hardware
    
    style UserSpace fill:#d4edda
    style KernelSpace fill:#fff3cd
    style Hardware fill:#f8d7da
```

### Основные компоненты

| Компонент | Назначение |
|:---|:---|
| **Kernel** | Ядро системы, управляет ресурсами |
| **Shell** | Интерпретатор команд (bash, zsh) |
| **System Libraries** | Библиотеки для взаимодействия с ядром |
| **System Utilities** | Утилиты для администрирования |

---

## 3. Файловая система

В Linux все является файлом: устройства, процессы, сокеты.

### Структура каталогов

```mermaid
graph TD
    Root["/  (корень)"]
    
    Root --> bin["/bin - Базовые команды"]
    Root --> etc["/etc - Конфигурация"]
    Root --> home["/home - Домашние каталоги"]
    Root --> var["/var - Переменные данные"]
    Root --> tmp["/tmp - Временные файлы"]
    Root --> usr["/usr - Пользовательские программы"]
    Root --> dev["/dev - Устройства"]
    Root --> proc["/proc - Процессы"]
    
    var --> log["/var/log - Логи"]
    var --> lib["/var/lib - Базы данных"]
    
    usr --> usrbin["/usr/bin - Программы"]
    usr --> usrlib["/usr/lib - Библиотеки"]
    
    style Root fill:#e95420,color:#fff
    style etc fill:#fff3cd
    style var fill:#d4edda
    style usr fill:#cfe2ff
```

| Каталог | Описание |
|:---|:---|
| `/` | Корневая директория |
| `/bin` | Основные исполняемые файлы (ls, cp, mv) |
| `/sbin` | Системные команды для администратора |
| `/etc` | Конфигурационные файлы системы |
| `/home` | Домашние каталоги пользователей |
| `/root` | Домашний каталог суперпользователя |
| `/var` | Изменяемые данные (логи, почта, очереди) |
| `/tmp` | Временные файлы (очищаются при перезагрузке) |
| `/usr` | Установленные программы и библиотеки |
| `/dev` | Файлы устройств |
| `/proc` | Виртуальная ФС с информацией о процессах |
| `/sys` | Виртуальная ФС с информацией о системе |
| `/boot` | Файлы загрузчика и ядра |
| `/lib` | Системные библиотеки |
| `/opt` | Опциональное ПО |
| `/mnt` | Точки монтирования |
| `/media` | Съемные носители |

---

## 4. Управление файлами

### Основные команды

| Команда | Описание | Пример |
|:---|:---|:---|
| `ls` | Просмотр содержимого каталога | `ls -la` |
| `cd` | Смена каталога | `cd /home/user` |
| `pwd` | Текущий путь | `pwd` |
| `mkdir` | Создание каталога | `mkdir -p dir/subdir` |
| `touch` | Создание файла | `touch file.txt` |
| `cp` | Копирование | `cp -r source/ dest/` |
| `mv` | Перемещение/переименование | `mv old.txt new.txt` |
| `rm` | Удаление | `rm -rf directory/` |
| `cat` | Вывод содержимого | `cat file.txt` |
| `file` | Определение типа файла | `file archive.tar.gz` |
| `ln` | Создание ссылок | `ln -s /path/to/file link` |
| `find` | Поиск файлов | `find /home -name "*.log"` |
| `locate` | Быстрый поиск по базе | `locate nginx.conf` |
| `du` | Размер файлов/каталогов | `du -sh *` |
| `df` | Использование дисков | `df -h` |

> [!TIP]
> Опция `-h` (human-readable) делает вывод размеров более читаемым: `1.5G` вместо `1572864`.

### Работа с архивами

```bash
# Создание tar.gz архива
tar -czf archive.tar.gz directory/

# Извлечение
tar -xzf archive.tar.gz

# Просмотр содержимого без извлечения
tar -tzf archive.tar.gz
```

---

## 5. Работа с текстом

### Утилиты для анализа текста

| Команда | Назначение | Пример |
|:---|:---|:---|
| `cat` | Вывод файла | `cat file.txt` |
| `less` | Постраничный просмотр | `less large.log` |
| `head` | Первые строки | `head -n 20 file.txt` |
| `tail` | Последние строки | `tail -f /var/log/syslog` |
| `grep` | Поиск по шаблону | `grep "error" app.log` |
| `sed` | Потоковое редактирование | `sed 's/old/new/g' file.txt` |
| `awk` | Обработка столбцов | `awk '{print $1}' data.txt` |
| `sort` | Сортировка | `sort -n numbers.txt` |
| `uniq` | Удаление дубликатов | `uniq -c sorted.txt` |
| `wc` | Подсчет строк/слов/байт | `wc -l file.txt` |
| `diff` | Сравнение файлов | `diff file1.txt file2.txt` |

### Мощные комбинации Pipeline

```bash
# Топ-10 IP адресов в логах Nginx
cat /var/log/nginx/access.log | cut -d' ' -f1 | sort | uniq -c | sort -rn | head -10

# Найти самые большие файлы
find / -type f -exec du -h {} + 2>/dev/null | sort -rh | head -20

# Количество процессов по пользователям
ps aux | awk '{print $1}' | sort | uniq -c | sort -rn
```

> [!NOTE]
> Оператор `|` (pipe) передает вывод одной команды на вход другой, позволяя строить мощные цепочки обработки.

---

## 6. Управление процессами

### Просмотр процессов

```bash
# Список всех процессов
ps aux

# Интерактивный мониторинг
top       # Классический
htop      # Улучшенная версия (требует установки)
```

### Сигналы и завершение процессов

```mermaid
graph LR
    Process[Процесс] --> SIGTERM[SIGTERM 15<br/>Мягкое завершение]
    Process --> SIGKILL[SIGKILL 9<br/>Принудительное завершение]
    Process --> SIGHUP[SIGHUP 1<br/>Перезагрузка конфига]
    Process --> SIGINT[SIGINT 2<br/>Прерывание Ctrl+C]
    
    style SIGTERM fill:#d4edda
    style SIGKILL fill:#f8d7da
```

| Команда | Описание | Пример |
|:---|:---|:---|
| `kill` | Отправка сигнала по PID | `kill -15 1234` |
| `killall` | Завершение по имени | `killall nginx` |
| `pkill` | Завершение по паттерну | `pkill -f "python.*script"` |
| `xkill` | Завершение по клику (GUI) | `xkill` |

> [!WARNING]
> `kill -9` (SIGKILL) не дает процессу времени на корректное завершение. Используйте только если `kill -15` не помог.

### Фоновые задачи

```bash
# Запуск в фоне
command &

# Отправка в фон (Ctrl+Z, затем)
bg

# Возврат на передний план
fg

# Список задач
jobs

# Запуск с отключением от терминала
nohup command &
```

---

## 7. Права доступа

### Структура прав

```
-rwxr-xr--  1 user group 4096 Jan 30 12:00 file.txt
 │││││││││
 │││└┴┴┴┴┴─ Права для: владельца, группы, остальных
 ││└─────── Количество ссылок
 │└──────── Тип (- файл, d каталог, l ссылка)
```

### Права в символьном виде

| Символ | Право | Значение |
|:---:|:---|:---|
| `r` | Read | Чтение (4) |
| `w` | Write | Запись (2) |
| `x` | Execute | Выполнение (1) |
| `-` | Нет права | 0 |

### Права в восьмеричной системе

```mermaid
graph LR
    subgraph Расчет["Расчет прав"]
        A["r=4, w=2, x=1"]
        B["rwx = 4+2+1 = 7"]
        C["r-x = 4+0+1 = 5"]
        D["r-- = 4+0+0 = 4"]
    end
    
    Example["chmod 754 file<br/>Владелец: rwx(7)<br/>Группа: r-x(5)<br/>Остальные: r--(4)"]
    
    Расчет --> Example
```

### Изменение прав

```bash
# Восьмеричный формат
chmod 755 script.sh    # rwxr-xr-x
chmod 644 file.txt     # rw-r--r--
chmod 600 secret.key   # rw-------

# Символьный формат
chmod u+x script.sh    # Добавить выполнение владельцу
chmod g-w file.txt     # Убрать запись у группы
chmod o= file.txt      # Убрать все права у остальных
chmod a+r file.txt     # Добавить чтение всем

# Рекурсивно
chmod -R 755 directory/
```

### Изменение владельца

```bash
# Сменить владельца
chown user file.txt

# Сменить владельца и группу
chown user:group file.txt

# Рекурсивно
chown -R user:group directory/
```

> [!IMPORTANT]
> Только суперпользователь (root) может изменять владельца файлов.

---

## 8. Сетевые утилиты

### Основные команды

| Команда | Назначение | Пример |
|:---|:---|:---|
| `ip` | Управление сетью | `ip addr show` |
| `ping` | Проверка доступности | `ping -c 4 google.com` |
| `traceroute` | Трассировка маршрута | `traceroute google.com` |
| `netstat` | Сетевые соединения (устарела) | `netstat -tulpn` |
| `ss` | Сетевые сокеты (замена netstat) | `ss -tulpn` |
| `curl` | HTTP запросы | `curl -I https://api.com` |
| `wget` | Скачивание файлов | `wget https://file.tar.gz` |
| `nslookup` | DNS запросы | `nslookup google.com` |
| `dig` | Подробная DNS информация | `dig google.com` |
| `tcpdump` | Перехват пакетов | `tcpdump -i eth0 port 80` |

### Мониторинг сетевого трафика

```bash
# Показать процессы, использующие сеть
nethogs

# Статистика по интерфейсам
iftop
```

---

## 9. Управление пользователями

```bash
# Создание пользователя
useradd -m -s /bin/bash username

# Удаление пользователя
userdel -r username

# Изменение пользователя
usermod -aG sudo username    # Добавить в группу sudo

# Смена пароля
passwd username

# Переключение пользователя
su - username

# Выполнение от имени root
sudo command
```

### Конфигурационные файлы

| Файл | Описание |
|:---|:---|
| `/etc/passwd` | База пользователей |
| `/etc/shadow` | Зашифрованные пароли |
| `/etc/group` | Группы пользователей |
| `~/.bashrc` | Настройки оболочки пользователя |

---

## 10. Практические рецепты

### Копирование между серверами

```bash
# Копирование файла на удаленный сервер
scp file.txt user@host:/path/to/destination

# Копирование каталога
scp -r directory/ user@host:/path/

# Копирование с сервера
scp user@host:/path/file.txt ./
```

### Мониторинг логов в реальном времени

```bash
# Следить за логом
tail -f /var/log/syslog

# С поиском по шаблону
tail -f app.log | grep ERROR

# Последние 100 строк с обновлением
tail -n 100 -f /var/log/nginx/access.log
```

### Топ IP адресов в логах

```bash
# Nginx access log
awk '{print $1}' /var/log/nginx/access.log | sort | uniq -c | sort -rn | head -10

# С использованием cut
cut -d' ' -f1 /var/log/nginx/access.log | sort | uniq -c | sort -rn | head -10
```

### Поиск файлов и выполнение команд

```bash
# Найти и удалить старые логи
find /var/log -name "*.log" -mtime +30 -delete

# Найти большие файлы (>100MB)
find / -type f -size +100M -exec ls -lh {} \; 2>/dev/null

# Изменить права для всех .sh файлов
find . -name "*.sh" -exec chmod +x {} \;
```

### Анализ использования диска

```bash
# Топ-10 самых больших каталогов
du -h / 2>/dev/null | sort -rh | head -10

# Размер каждого каталога в текущей директории
du -sh */ | sort -rh

# Использование дисков
df -h
```

### Работа с процессами

```bash
# Найти PID процесса по имени
pgrep nginx

# Полная информация о процессе
ps aux | grep nginx

# Убить все процессы python
pkill -f python

# Время работы процесса
ps -p PID -o etime=
```

---

## Полезные переменные окружения

```bash
# Просмотр всех переменных
env

# Установка переменной
export VAR_NAME="value"

# Добавить в PATH
export PATH="$PATH:/new/path"

# Сделать постоянной (добавить в ~/.bashrc)
echo 'export VAR="value"' >> ~/.bashrc
source ~/.bashrc
```

---

## Горячие клавиши Bash

| Комбинация | Действие |
|:---|:---|
| `Ctrl+C` | Прервать команду |
| `Ctrl+Z` | Приостановить команду |
| `Ctrl+D` | Выход (EOF) |
| `Ctrl+L` | Очистить экран |
| `Ctrl+A` | В начало строки |
| `Ctrl+E` | В конец строки |
| `Ctrl+U` | Удалить от курсора до начала |
| `Ctrl+K` | Удалить от курсора до конца |
| `Ctrl+R` | Поиск по истории команд |
| `Tab` | Автодополнение |
| `!!` | Повторить последнюю команду |

---

> [!TIP]
> **Изучайте man-страницы**: `man <команда>` — лучший источник информации о любой команде Linux.

> [!CAUTION]
> Команды с `rm -rf /`, `dd`, форматирование дисков могут **безвозвратно уничтожить данные**. Всегда дважды проверяйте перед выполнением!