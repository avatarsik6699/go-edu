package ctx

import (
	"context"
	"log"
)

/*
	Context - набор метаданных, ассоциированных с каким-либо запросом или процессом.
	По существу представляет собой интерфейс.
	Done() -> chan struct{} можно подписаться и отправить сигнал о завершении (вызвали cancel).
	Err()  -> ошибка (возвращается, когда канал был закрыт).
	cancel -> закрывает внутренний канал + делает CAS для флага close uint32 о закрытии.
					дальнейшие вызовы закрытия канала ни к чему не ведут.


	NB!: Если контекст с withTimeout/withCancel не будет отменен после исполнения кода, тогда
	будет утечка горутин, т.к. таймер запускается в горутине.


*/

type K1 = string
type K2 = string

const k1 K1 = "key"
const k2 K2 = "key"

func Sample() {

	ctx, _ := context.WithCancel(context.Background())
	ctx = context.WithValue(ctx, k1, "b")
	ctx = context.WithValue(ctx, k2, "c")

	log.Println(ctx.Value(k1), ctx.Value(k2))
}
