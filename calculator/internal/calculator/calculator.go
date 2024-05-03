package calculator

import (
	"Calculator/internal/agent"
	"context"
	"database/sql"
	"log"
)

func New(db *sql.DB) *Calculator {
	return &Calculator{
		DB:     db,
		Agents: make(map[int]*agent.Agent),
	}
}

func (calc *Calculator) Start(ctx context.Context) {
	err := calc.DB.PingContext(ctx)
	if err != nil {
		panic(err)
	}
}

func (calc *Calculator) Manage(ctx context.Context) {

}

func (calc *Calculator) InitAgent(ownerId int) error {
	ctx := context.TODO()
	newAgent := agent.NewAgent(ownerId)
	agentId, err := newAgent.Insert(ctx, calc.DB)
	if err != nil {
		log.Println(err)
		return err
	}
	calc.Agents[int(agentId)] = newAgent
}
