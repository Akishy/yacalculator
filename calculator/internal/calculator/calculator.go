package calculator

import (
	"Calculator/internal/agent"
	"Calculator/internal/task"
	"context"
	"database/sql"
	"errors"
	calcv1 "github.com/Akishy/yacalculator/proto/gen/calculator"
	"go/constant"
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

func (calc *Calculator) ManageTask(ctx context.Context, _task *task.Task) (constant.Value, calcv1.OperandType, error) {
	for _, calcAgent := range calc.Agents {
		if calcAgent.Status == agent.ALIVE {
			log.Println("found alive agent")
			calcAgent.TaskChan <- *_task
			for result := range calcAgent.ResultChan {
				return result.Value, 0, nil
			}
		}
	}
	return nil, -1, errors.New("agent not alive")
}

func (calc *Calculator) ListAllAgents() {
	log.Println(calc.Agents)
}

func (calc *Calculator) InitAgent(ownerId int64) (int64, error) {
	ctx := context.TODO()
	newAgent := agent.NewAgent(int(ownerId))
	agentId, err := newAgent.Insert(ctx, calc.DB)
	if err != nil {
		log.Println(err)
		return 0, err
	}
	calc.Agents[int(agentId)] = newAgent
	calc.ListAllAgents()
	go newAgent.Run()
	return agentId, nil
}
