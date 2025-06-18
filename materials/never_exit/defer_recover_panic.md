# Defer, recover(), panic

В Go каждая горутина имеет свой собственный стек. Когда функция вызывается, для неё создается новый фрейм стека (stack frame), который содержит:
  
  - Локальные переменные
  - Параметры функции
  - Адрес возврата
  - Другие служебные данные

defer не создает отдельный стек - он работает в рамках того же стека, что и основная функция. Давайте разберем это на примере:

```go 
  func example() (result int) {
    defer func() {
        result = 1
    }()
    defer func() {
        result = 2
    }()
    defer func() {
        result = 3
    }()
    return 0
  }

  1. Создание фрейма стека для example()
    +------------------+
    | result = 0       |  // Именованный возвращаемый параметр
    | defer function   |  // Указатель на отложенную функцию
    | ...              |
    +------------------+

  2. При выполнении return 0:
      - Go не возвращает управление сразу
      - Вместо этого он выполняет все отложенные функции (в порядке LIFO 3 -> 2 -> 1 (result = 1))
          + defer функции имеют доступ к род. фрейму через замыкание (поэтому могут изменять значения во фрейме и перехватывать панику).

          func example() {
              x := 1
              defer func() {
                  x := 2  // Это новая переменная в новом фрейме
                  fmt.Println("defer x:", x)  // Выведет 2
              }()
              fmt.Println("main x:", x)  // Выведет 1
          }

      - Только после этого происходит реальный возврат
```

### Defer функции

- Создают свои собственные фреймы стека
- Имеют свои собственные области видимости
- Могут иметь свои локальные переменные
- Имеют доступ к переменным родительского фрейма через замыкания

### Отличие от обычных функций
- Время выполнения (отложенное)
- Порядок выполнения (LIFO)
- Гарантированное выполнение при выходе из функции

### Особенности
- Могут изменять именованные возвращаемые значения
- Могут перехватывать панику
- Выполняются даже при панике
- Не выполняются при os.Exit()

Таким образом, defer функции - это полноценные функции со своими фреймами стека, просто с особым временем и порядком выполнения.


### Examples:

```go
  // Как будет выглядеть stack frame?
  func level3() {
    panic("something went wrong")
  }

  func level2() {
    defer fmt.Println("level2 cleanup")
    level3()
  }

  func level1() {
    defer fmt.Println("level1 cleanup")
    level2()
  }

  func main() {
    defer fmt.Println("main cleanup")
    level1()
  }

  // Как будет выглядеть stack frame?
  func example() {
    defer func() {
        fmt.Println("First defer")
    }()
    
    defer func() {
        if r := recover(); r != nil {
            fmt.Println("Recovered:", r)
        }
    }()
    
    defer func() {
        fmt.Println("Last defer")
    }()
    
    panic("something went wrong")
  }

  // Кто поймает панику?
  func level3() {
    panic("level3 panic")
  }

  func level2() {
    defer func() {
        if r := recover(); r != nil {
            fmt.Println("Recovered in level2:", r)
        }
    }()
    level3()
  }

  func level1() {
    defer func() {
        if r := recover(); r != nil {
            fmt.Println("Recovered in level1:", r)
        }
    }()
    level2()
  }
```

# Panic, recover(), gorutines

Паника в одной горутине не влияет на выполнение других горутин. Это важное заблуждение, которое нужно развеять. Давайте разберем это на примерах:

## Основные принципы
- Программа завершается, если есть необработанная паника в любой горутине
- Main функция не приостанавливается, а завершается вместе с программой  
- Это поведение по умолчанию в Go

```go
  func main() {
    // Запускаем горутину, которая вызовет панику
    go func() {
        panic("goroutine panic")
    }()

    // Основная горутина продолжает работать
    fmt.Println("Main goroutine is running")
    time.Sleep(time.Second)

    // Не успеет отработать т.к. паника в горутине завершает программу в целом, а не main gorutine. 
    fmt.Println("Main goroutine is still running") 
  }
  // Output:
  // Main goroutine is running
  // panic: goroutine panic

  // Можно ли в одной горутиние поймать панику из другой горутины?
  // Ответ: нельзя, т.к. у каждой горутины свой call stack.
  func main() {
    
    defer func() {
      if v := recover(); v != nil {
        log.Println("captured panic!")
      }
    }()

    go func() {
      panic("u cannot captured me")
    }()
    
    time.Sleep(time.Second)
  }
```