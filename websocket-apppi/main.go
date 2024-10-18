package main

import (
    "github.com/micmonay/keybd_event"
    "time"
)

func main() {
    // Ініціалізуємо бібліотеку для емуляції натискань клавіатури
    kb, err := keybd_event.NewKeyBonding()
    if err != nil {
        panic(err)
    }

    // Вказуємо клавішу "a"
    kb.SetKeys(keybd_event.VK_A)

    // Канал для таймера зупинки через 2 хвилини
    stop := time.After(2 * time.Minute)

    for {
        select {
        case <-stop:
            return // Завершення програми після 2 хвилин
        default:
            // Натискання клавіші "a"
            err = kb.Launching()
            if err != nil {
                panic(err)
            }
            // Затримка між натисканнями (тут 100 мс)
            time.Sleep(100 * time.Millisecond)
        }
    }
}
