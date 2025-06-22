# Разница между t.Fatal и t.Error

## t.Error
- Помечает тест как неудачный, но продолжает выполнение теста
- Полезно когда нужно проверить несколько условий в одном тесте
- Позволяет увидеть все ошибки в тесте

## t.Fatal  
- Помечает тест как неудачный и немедленно останавливает выполнение теста
- Эквивалентно вызову `t.Error + t.FailNow()`
- Используется когда дальнейшее выполнение теста бессмысленно

## Когда использовать

### t.Error
- Для проверки множественных условий, где нужно увидеть все ошибки
- Когда важно проверить несколько аспектов в одном тесте

### t.Fatal
- Когда дальнейшее выполнение теста бессмысленно (например, не удалось инициализировать объект)
- При критических ошибках, которые делают остальные проверки невозможными

## Пример улучшения
В коде `t.Fail()` можно заменить на `t.Fatal("eventManager should not be nil")` для более информативного сообщения об ошибке.


Структура тестов
func TestFunctionName(t *testing.T) {
    // Arrange (подготовка)
    input := "test"
    expected := "expected"
    
    // Act (действие)
    result := FunctionName(input)
    
    // Assert (проверка)
    if result != expected {
        t.Errorf("Expected %s, got %s", expected, result)
    }
}

Table-driven tests (рекомендуемый подход)
func TestFunctionName_TableDriven(t *testing.T) {
    tests := []struct {
        name     string
        input    string
        expected string
    }{
        {"empty string", "", ""},
        {"normal string", "hello", "HELLO"},
        {"with numbers", "test123", "TEST123"},
    }
    
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            result := FunctionName(tt.input)
            if result != tt.expected {
                t.Errorf("Expected %s, got %s", tt.expected, result)
            }
        })
    }
}

Тестирование конкурентного кода
Основные проблемы конкурентности
Race Conditions - когда результат зависит от порядка выполнения
Data Races - одновременный доступ к данным без синхронизации
Deadlocks - взаимная блокировка горутин
Starvation - некоторые горутины не получают ресурсы

# Запуск с race detection
go test -race ./pkg/events

# Стресс-тестирование с race detection
go test -race -count=1000 ./pkg/events

# Параллельное выполнение тестов
go test -race -parallel=8 ./pkg/events

# Benchmark с параллельным выполнением
go test -bench=BenchmarkConcurrent -benchmem ./pkg/events

## Чек-лист для тестирования

### Синхронный код:
- [ ] Все публичные функции покрыты тестами
- [ ] Используются table-driven tests
- [ ] Тестируются граничные случаи
- [ ] Тестируется обработка ошибок
- [ ] Тесты изолированы и детерминированы

### Конкурентный код:
- [ ] Используется race detector (-race)
- [ ] Стресс-тестирование с множественными итерациями
- [ ] Тестирование с разным количеством горутин
- [ ] Проверка отсутствия deadlocks
- [ ] Тестирование таймаутов
- [ ] Benchmark тесты для производительности