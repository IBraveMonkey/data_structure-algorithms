package batcher

import (
	"errors"
	"sync"
	"time"
)

/*
 * Паттерн Batcher (Батчер) используется для объединения нескольких операций в одну группу (батч) с целью уменьшения количества вызовов или улучшения производительности.
 * Он накапливает сообщения до достижения определенного размера или интервала времени, а затем выполняет действие над всеми накопленными со&#1089;общениями.
 * Пример использования: отправка логов на сервер, где логи собираются в батчи и отправляются раз в несколько секунд или когда набирается определенное количество логов.
 */

// Batcher управляет накоплением и групповой обработкой сообщений
type Batcher struct {
	maxSize       int            // Максимальный размер батча (сколько сообщений накопить)
	flushInterval time.Duration  // Интервал времени для автоматической отправки
	flushAction   func([]string) // Функция, которая обрабатывает батч сообщений
	ticker        *time.Ticker   // Таймер для периодической отправки

	mutex    sync.Mutex // Защита от гонки данных при доступе к messages
	messages []string   // Накопленные сообщения (текущий батч)

	closeCh     chan struct{} // Канал для сигнала закрытия
	closeDoneCh chan struct{} // Канал для подтверждения завершения работы
}

// NewBatcher создает новый экземпляр батчера с валидацией параметров
func NewBatcher(action func([]string), size int, interval time.Duration) (*Batcher, error) {
	// Проверяем, что передана функция обработки
	if action == nil {
		return nil, errors.New("invalid action")
	}

	// Проверяем корректность размера батча
	if size <= 0 {
		return nil, errors.New("invalid size")
	}

	// Проверяем корректность интервала
	if interval <= 0 {
		return nil, errors.New("invalid interval")
	}

	// Создаем и возвращаем новый батчер
	return &Batcher{
		maxSize:       size,
		flushAction:   action,
		flushInterval: interval,
		closeCh:       make(chan struct{}), // Инициализируем канал закрытия
		closeDoneCh:   make(chan struct{}), // Инициализируем канал подтверждения
	}, nil
}

// Append добавляет сообщение в батч
func (b *Batcher) Append(message string) error {
	b.mutex.Lock()         // Блокируем для безопасного доступа к messages
	defer b.mutex.Unlock() // Разблокируем при выходе

	// Проверяем, не закрыт ли батчер (неблокирующая проверка)
	select {
	case <-b.closeCh:
		return errors.New("batcher is close")
	default:
	}

	// Добавляем сообщение в батч
	b.messages = append(b.messages, message)

	// Если батч достиг максимального размера, отправляем немедленно
	if len(b.messages) == b.maxSize {
		b.flushLocked()                 // Отправляем батч (мьютекс уже захвачен)
		b.ticker.Reset(b.flushInterval) // Сбрасываем таймер для следующего батча
	}

	return nil
}

// Run запускает воркер для периодической отправки батчей
func (b *Batcher) Run() {
	// Предотвращаем множественный запуск
	if b.ticker != nil {
		return
	}

	// Создаем таймер для периодической отправки
	b.ticker = time.NewTicker(b.flushInterval)

	// Запускаем горутину-воркер
	go func() {
		defer close(b.closeDoneCh) // Сигнализируем о завершении при выходе

		for {
			// Первый select: проверяем закрытие без блокировки
			select {
			case <-b.closeCh:
				b.flush() // Отправляем оставшиеся сообщения
				return
			default:
			}

			// Второй select: ждем событий (таймер или закрытие)
			select {
			case <-b.closeCh:
				b.flush() // Отправляем оставшиеся сообщения
				return
			case <-b.ticker.C:
				b.flush() // Отправляем батч по таймеру
			}
		}
	}()
}

// flush отправляет накопленные сообщения (с блокировкой)
func (b *Batcher) flush() {
	b.mutex.Lock()         // Блокируем для безопасного доступа
	defer b.mutex.Unlock() // Разблокируем при выходе

	b.flushLocked() // Вызываем версию без блокировки
}

// flushLocked отправляет накопленные сообщения (без блокировки, мьютекс уже захвачен)
func (b *Batcher) flushLocked() {
	// Если батч пустой, ничего не делаем
	if len(b.messages) == 0 {
		return
	}

	messages := b.messages // Сохраняем текущий батч
	b.messages = nil       // Очищаем батч для новых сообщений

	// Запускаем обработку в отдельной горутине (не блокируем Append)
	go b.flushAction(messages)
}

// Close дожидается завершения воркера и отправляет оставшиеся сообщения
func (b *Batcher) Close() {
	// Проверяем, не закрыт ли уже батчер (неблокирующая проверка)
	select {
	case <-b.closeCh:
		return // Уже закрыт, выходим
	default:
	}

	// Блокируем и закрываем канал сигнала
	b.mutex.Lock()
	close(b.closeCh) // Сигнализируем воркеру о необходимости завершения
	b.mutex.Unlock()

	<-b.closeDoneCh // Ждем завершения воркера
	b.ticker.Stop() // Останавливаем таймер
}
