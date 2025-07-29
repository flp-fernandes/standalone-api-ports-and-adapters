package orchestrator

import (
	"context"
	"log"
	"sync/atomic"

	"github.com/robfig/cron/v3"

	"github.com/flp-fernandes/standalone-api-ports-and-adapters/internal/app"
)

type Orchestrator struct {
	userService  *app.UserService
	orderService *app.OrderService

	userJobRunning  int32
	orderJobRunning int32
}

func NewOrchestrator(userService *app.UserService, orderService *app.OrderService) *Orchestrator {
	return &Orchestrator{
		userService:  userService,
		orderService: orderService,
	}
}

func (o *Orchestrator) Run(ctx context.Context) error {

	c := cron.New(cron.WithSeconds())

	c.AddFunc("*/3 * * * * *", func() {
		if !tryLockAtomic(&o.userJobRunning) {
			log.Println("userJob em execução, pulando...")
			return
		}
		defer atomic.StoreInt32(&o.userJobRunning, 0)

		log.Println("Executando userJob")
		if err := o.userService.PrintAllUsers(ctx); err != nil {
			log.Println("Erro userJob:", err)
		}
	})

	c.AddFunc("*/3 * * * * *", func() {
		if !tryLockAtomic(&o.orderJobRunning) {
			log.Println("orderJob em execução, pulando...")
			return
		}
		defer atomic.StoreInt32(&o.orderJobRunning, 0)

		log.Println("Executando orderJob")
		if err := o.orderService.PrintAllOrders(ctx); err != nil {
			log.Println("Erro orderJob:", err)
		}
	})

	log.Println("⏱️ Scheduler inicializado")
	c.Start()

	<-ctx.Done()
	c.Stop()

	return ctx.Err()
}

func tryLockAtomic(flag *int32) bool {
	return atomic.CompareAndSwapInt32(flag, 0, 1)
}
