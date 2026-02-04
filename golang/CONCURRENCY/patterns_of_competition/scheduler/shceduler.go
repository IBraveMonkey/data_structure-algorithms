package scheduler

import (
	"errors"
	"sync"
	"time"
)

/*
 * Паттерн Scheduler (Планировщик) используется для планирования выполнения функций в будущем.
 * Он позволяет установить таймер для выполнения функции и отменить выполнение, если это необходимо.
 * Пример использования: выполнение задач по расписанию, например, отправка уведомлений или очистка кэша.
 */

// Scheduler управляет отложенным выполнением функций
type Scheduler struct {
	mutex   sync.Mutex          // Защита от гонки данных при доступе к actions и closed
	closed  bool                // Флаг состояния планировщика (закрыт или нет)
	actions map[int]*time.Timer // Карта активных таймеров по ключам
}

// NewScheduler создает новый экземпляр планировщика
func NewScheduler() *Scheduler {
	return &Scheduler{
		actions: make(map[int]*time.Timer), // Инициализируем карту таймеров
	}
}

// SetTimeout запускает функцию после указанной задержки
func (s *Scheduler) SetTimeout(key int, delay time.Duration, action func()) error {
	// Проверяем корректность задержки
	if delay < 0 {
		return errors.New("invalid delay")
	}

	// Проверяем, что функция передана
	if action == nil {
		return errors.New("invalid action")
	}

	s.mutex.Lock()         // Блокируем для безопасного доступа
	defer s.mutex.Unlock() // Разблокируем при выходе

	// Проверяем, не закрыт ли планировщик
	if s.closed {
		return errors.New("scheduler is closed")
	}

	// Если таймер с таким ключом уже существует, останавливаем его
	if timer, found := s.actions[key]; found {
		timer.Stop() // Останавливаем старый таймер
	}

	// Создаем новый таймер, который выполнит action через delay
	s.actions[key] = time.AfterFunc(delay, action)
	return nil
}

// CancelTimeout отменяет запланированное выполнение функции
func (s *Scheduler) CancelTimeout(key int) {
	s.mutex.Lock()         // Блокируем для безопасного доступа
	defer s.mutex.Unlock() // Разблокируем при выходе

	// Ищем таймер по ключу
	if timer, found := s.actions[key]; found {
		timer.Stop()           // Останавливаем таймер
		delete(s.actions, key) // Удаляем из карты
	}
}

// Close останавливает все запланированные функции и закрывает планировщик
func (s *Scheduler) Close() {
	s.mutex.Lock()         // Блокируем для безопасного доступа
	defer s.mutex.Unlock() // Разблокируем при выходе

	s.closed = true // Помечаем планировщик как закрытый

	// Останавливаем и удаляем все активные таймеры
	for key, timer := range s.actions {
		timer.Stop()           // Останавливаем таймер
		delete(s.actions, key) // Удаляем из карты
	}
}
