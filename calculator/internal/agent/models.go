package agent

type Status int

const (
	ALIVE   Status = iota // Значит что можно поручить задачку
	WORKING               // Значит что живой, но работает
)

type Agent struct {
	Id         int         // Идентификатор агента
	OwnerID    int         // Идентификатор владельца агента
	Status     Status      // Статус агента
	StatusChan chan Status // Канал для отправки статуса оркестратору
}
