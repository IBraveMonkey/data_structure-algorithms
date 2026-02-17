# 🧩 Unit-тесты (Unit Tests)

Unit-тестирование — это фундамент вашей пирамиды тестов. В Go оно реализовано максимально просто и эффективно через стандартную библиотеку.


# 1. 📦 Основы пакета `testing`

Любой тест в Go — это обычная функция в файле `*_test.go`.

Пример простейшего теста:
```go
package math

import "testing"

func TestAdd(t *testing.T) {
    result := Add(2, 3)
    expected := 5
    if result != expected {
        t.Errorf("Add(2, 3) = %d; want %d", result, expected)
    }
}
```


### Основные методы `testing.T`:
*   `t.Error(args...)`: Сообщает об ошибке, но продолжает выполнение теста.
*   `t.Errorf(format, args...)`: Форматированное сообщение об ошибке.
*   `t.Fatal(args...)`: Сообщает об ошибке и немедленно прекращает выполнение текущего теста.
*   `t.Log(args...)`: Аналог `fmt.Println` для вывода информации только при падении теста или флаге `-v`.


# 2. 📊 Табличные тесты (Table-Driven Tests)

Это стандарт де-факто в сообществе Go. Вместо того чтобы писать 10 функций для разных входных данных, вы описываете структуру с данными и циклом проходите по ней.

```go
func TestAddTableDriven(t *testing.T) {
    tests := []struct {
        name     string
        a, b     int
        expected int
    }{
        {"positive numbers", 2, 3, 5},
        {"negative numbers", -1, -1, -2},
        {"zeroes", 0, 0, 0},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            result := Add(tt.a, tt.b)
            if result != tt.expected {
                t.Errorf("got %d, want %d", result, tt.expected)
            }
        })
    }
}
```
> [!NOTE]
> Метод `t.Run` создает "subtest", что позволяет запускать их по отдельности и видеть результат для каждого кейса.


# 3. 🛠️ Тестовые фикстуры и хелперы (Fixtures and Helpers)

Если вам нужно создать сложный объект или подготовить окружение, используйте вспомогательные функции.

```go
func setupTestDB(t *testing.T) *Database {
    t.Helper() // Указывает, что при ошибке нужно показывать строку вызова этой функции, а не строку внутри helper
    db := NewDatabase()
    // подготовка данных...
    return db
}
```


# 4. 🎭 Моки и дублеры (Mocking)

В Go не принято использовать тяжелые фреймворки для мокинга, если можно обойтись интерфейсами.

**Шаг 1: Опишите интерфейс**
```go
type OrderRepository interface {
    Save(order *Order) error
}
```

**Шаг 2: Создайте мок-структуру в тесте**
```go
type mockRepo struct {
    saveCalled bool
}

func (m *mockRepo) Save(o *Order) error {
    m.saveCalled = true
    return nil
}

func TestOrderService_Create(t *testing.T) {
    repo := &mockRepo{}
    service := NewOrderService(repo)
    // ...
}
```


### Популярные библиотеки:
Если проект большой и моков много, можно использовать:
*   [stretchr/testify](https://github.com/stretchr/testify): Самый популярный набор утилит для ассертов и моков.
*   [golang/mock (gomock)](https://github.com/golang/mock): Генерация моков по интерфейсам.

> [!TIP]
> Старайтесь использовать интерфейсы там, где ожидаете сложную логику или внешние зависимости. Это сделает ваш код "тестируемым" (testable).

<!-- QUIZ_START 

[
    {
        "question": "В чем разница между t.Error и t.Fatal в пакете testing?",
        "options": [
            "t.Error останавливает тест, а t.Fatal — нет",
            "t.Error сообщает об ошибке и продолжает выполнение теста, а t.Fatal немедленно прерывает выполнение текущего теста",
            "Они идентичны",
            "t.Fatal используется только для логов"
        ],
        "correctIndex": 1
    },
    {
        "question": "Какая практика считается стандартом в Go для тестирования функции с множеством разных входных данных?",
        "options": [
            "Написание отдельной функции для каждого случая",
            "Табличные тесты (Table-driven tests) с использованием цикла и t.Run",
            "Использование глобальных переменных",
            "Ручное тестирование через fmt.Println"
        ],
        "correctIndex": 1
    },
    {
        "question": "Для чего нужна функция t.Helper()?",
        "options": [
            "Чтобы вызвать меню помощи",
            "Чтобы при возникновении ошибки Go указывал на строку, где был вызван хелпер, а не на строку внутри самого хелпера",
            "Чтобы ускорить выполнение тестов",
            "Это обязательная функция для любого теста"
        ],
        "correctIndex": 1
    }
]

QUIZ_END -->

