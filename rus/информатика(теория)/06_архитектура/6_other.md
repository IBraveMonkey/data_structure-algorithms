# 🎭 Другие архитектурные паттерны


## 📑 Содержание
1. [Паттерны UI (MVC, MVP, MVVM)](#1-паттерны-ui-mvc-mvp-mvvm)
2. [Паттерны работы с БД (Active Record, Data Mapper)](#2-паттерны-работы-с-бд-active-record-data-mapper)
3. [Разное (Lazy Load, Delegation)](#3-разное)

---


## 1. 🖼️ Паттерны UI (MVC, MVP, MVVM)

Как разделить "картинку" (UI) и логику?


### 🚦 MVC (Model - View - Controller)
Классика веба (особенно серверного, как Spring MVC, Django, Ruby on Rails).

*   **Model**: Данные и бизнес-логика.
*   **View**: Отображение (HTML страница).
*   **Controller**: Принимает ввод пользователя, обновляет модель.

```mermaid
%%{init: { 'theme': 'base', 'themeVariables': { 'lineColor': '#009688', 'primaryColor': '#009688', 'primaryTextColor': '#009688', 'attributeBkg': '#009688', 'attributeTextColor': '#009688', 'signalColor': '#009688', 'actorLineColor': '#009688', 'nodeBorder': '#009688', 'clusterBorder': '#009688', 'textColor': '#009688', 'fontSize': '16px' } } }%%
graph LR
    User -->|Действие| Controller
    Controller -->|Обновляет| Model
    Model -->|Данные| View
    View -->|Показывает| User



linkStyle default stroke:#009688,stroke-width:2px;




```


### 🎤 MVP (Model - View - Presenter)
Часто в десктопных или старых Android приложениях.
*   **Presenter**: Посредник. View сама ничего не знает о Модели. Presenter берет данные из Модели и "кладет" их во View.


### 🧬 MVVM (Model - View - ViewModel)
Современный стандарт для фронтенда (React, Vue, Angular) и мобилок (WPF, Jetpack Compose).
*   **ViewModel**: Специальная модель именно для отображения.
*   **Data Binding**: Магия! Изменил переменную в коде — кнопка на экране сама перекрасилась.

| Паттерн | Кто главный? | Связь View и Model | Где популярен? |
| :--- | :--- | :--- | :--- |
| **MVC** | Controller | Часто прямая | Веб-бэкенд (SSR) |
| **MVP** | Presenter | Нет (через Presenter) | Android (old), Desktop |
| **MVVM** | ViewModel | Data Binding (Авто) | React, Vue, Mobile |

---


## 2. 🗄️ Паттерны работы с БД

Как подружить объекты в коде (OOP) и таблицы в базе (RDBMS)?


### 🏃 Active Record
Объект сам умеет себя сохранять.
*   **Пример**: `user.save()`, `user.delete()`.
*   **Где**: Ruby on Rails, Eloquent (Laravel).
*   **Плюсы**: Очень быстро писать простой код.
*   **Минусы**: Нарушает принцип единственной ответственности (SRP). Объект и данные хранит, и с базой общается.


### 🗺️ Data Mapper
Объект и База полностью разделены. Есть посредник (Mapper/Repository), который их дружит.
*   **Пример**: `repository.save(user)`.
*   **Где**: Hibernate (Java), TypeORM (Node), GORM (Go - частично).
*   **Плюсы**: Чистый код, бизнес-логика не зависит от базы.
*   **Минусы**: Больше кода писать.

```mermaid
%%{init: { 'theme': 'base', 'themeVariables': { 'lineColor': '#009688', 'primaryColor': '#009688', 'primaryTextColor': '#009688', 'attributeBkg': '#009688', 'attributeTextColor': '#009688', 'signalColor': '#009688', 'actorLineColor': '#009688', 'nodeBorder': '#009688', 'clusterBorder': '#009688', 'textColor': '#009688', 'fontSize': '16px' } } }%%
graph TD
    subgraph AR ["Active Record"]
        Obj1["User Object<br/>+Save()<br/>+Delete()"]
    end

    subgraph DM ["Data Mapper"]
        Obj2["User Entity<br/>(Только данные)"]
        Mapper["Repository<br/>+Save(user)<br/>+Delete(user)"]
        Mapper --> Obj2
    end



linkStyle default stroke:#009688,stroke-width:2px;




```

---


## 3. 🧩 Разное


### 💤 Lazy Load (Ленивая загрузка)
"Не грузи то, что не нужно".
*   **Пример**: Загрузили Пользователя, но его список Заказов (который может быть огромным) не грузим. Загрузим только тогда, когда в коде обратимся к `user.getOrders()`.


### 🤝 Delegation (Делегирование)
"Пусть сделает профессионал".
Объект не делает работу сам, а вызывает метод у другого объекта. Это основа паттерна Композиция.


### 📚 Registry (Реестр)
Глобальный объект (или Singleton), где хранятся настройки или сервисы, доступные всему приложению.
*   Сейчас считается анти-паттерном, лучше использовать **Dependency Injection**.

<!-- QUIZ_START 
[
    {
        "question": "В чем главная особенность MVVM паттерна?",
        "options": ["Использование контроллера для связи", "Наличие Presenter", "Двустороннее связывание данных (Data Binding) между View и ViewModel", "Полный отказ от Модели"],
        "correctIndex": 2
    },
    {
        "question": "Почему паттерн Active Record часто критикуют?",
        "options": ["Он слишком медленный", "Он нарушает принцип единственной ответственности (SRP), смешивая бизнес-логику и работу с БД", "Он требует знания SQL", "Он работает только с Ruby"],
        "correctIndex": 1
    },
    {
        "question": "Что позволяет паттерн Lazy Load?",
        "options": ["Ускорить загрузку приложения за счет откладывания загрузки тяжелых данных до момента их реального использования", "Загружать все данные сразу при старте", "Дублировать данные в кеш", "Удалять неиспользуемые объекты"],
        "correctIndex": 0
    }
]
QUIZ_END -->