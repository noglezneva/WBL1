package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

type Worker struct {
	stopFlag bool          // Флаг для остановки выполнения
	stopCh   chan struct{} // Канал для передачи сигнала остановки
	wg       sync.WaitGroup
	ctx      context.Context
	cancel   context.CancelFunc
}

// Остановка горутины по флагу
func (w *Worker) stopByFlag() {
	defer w.wg.Done()
	for !w.stopFlag {
		//Выполнение работы горутины
		fmt.Println("Выполняется работа горутины(stopByFlag)")
		time.Sleep(time.Second)
	}
	fmt.Println("Горутина остановлена(stopByFlag)")
}

// Остановка горутины по каналу
func (w *Worker) stopByChannel() {
	defer w.wg.Done()
	for {
		select {
		case <-w.stopCh:
			fmt.Println("Горутина остановлена(stopByChannel)")
			return
		default:
			//Выполнение работы горутина
			fmt.Println("Выполняется работа горутины (stopByChannel)")
			time.Sleep(time.Second)
		}
	}
}

// Остановка горутины по контексту
func (w *Worker) stopByContext() {
	defer w.wg.Done()
	for {
		select {
		case <-w.ctx.Done():
			fmt.Println("Горутина остановлена (stopByContext)")
			return
		default:
			// Выполнение работы горутины
			fmt.Println("Выполняется работа горутины (stopByContext)")
			time.Sleep(time.Second)
		}
	}
}

// Остановка горутины по синхронизации
func (w *Worker) stopBySync() {
	defer w.wg.Done()
	for {
		if w.isStopped() {
			fmt.Println("Горутина остановлена (stopBySync)")
			return
		}

		// Выполнение работы горутины
		fmt.Println("Выполняется работа горутины (stopBySync)")
		time.Sleep(time.Second)
	}
}

// Проверка флага остановки
func (w *Worker) isStopped() bool {
	select {
	case <-w.stopCh:
		return true
	default:
		return false
	}
}
func main() {
	worker := &Worker{
		stopCh: make(chan struct{}),
	}
	fmt.Println("Запуск горутин")
	// Запуск горутин со всеми различными способами остановки
	worker.wg.Add(4)
	go worker.stopByFlag()
	go worker.stopByChannel()

	// Использование контекста для остановки
	worker.ctx, worker.cancel = context.WithCancel(context.Background())
	go worker.stopByContext()

	go worker.stopBySync()

	// Демонстрация работы горутин в течение нескольких секунд
	time.Sleep(3 * time.Second)

	fmt.Println("Останавливаем горутины")

	// Остановка горутин разными способами
	worker.stopFlag = true
	close(worker.stopCh)
	worker.cancel()

	worker.wg.Wait()

	fmt.Println("Главная горутина завершена")
}

// В этой программе создается структура Worker, которая содержит все возможные способы остановки горутины.
// Внутри Worker определены методы stopByFlag(), stopByChannel(), stopByContext(), stopBySync(), которые выполняют работу горутины
// и останавливают ее в зависимости от выбранного способа.

// В функции main() создается экземпляр Worker и запускаются горутины для каждого способа остановки.
// После некоторого ожидания, горутины останавливаются путем установки флага, закрытия канала
// и вызова функций отмены контекста и ожидания, в зависимости от способа.

// Поэтому программа демонстрирует все возможные способы остановки горутины и выводит соответствующие сообщения
// о завершении каждой горутины.
