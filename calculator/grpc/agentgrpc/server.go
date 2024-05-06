package agentgrpc

import (
	"Calculator/internal/calculator"
	"Calculator/internal/task"
	"context"
	"errors"
	calcv1 "github.com/Akishy/yacalculator/proto/gen/calculator"
	"go/constant"
	"go/token"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"regexp"
	"strconv"
	"time"
)

type serverAPI struct {
	calcv1.UnimplementedTaskServer
	calcv1.UnimplementedAgentServer
	calculatorServer *calculator.Calculator
}

func Register(gRPC *grpc.Server, calculator *calculator.Calculator) {
	calcv1.RegisterAgentServer(gRPC, &serverAPI{calculatorServer: calculator})
	calcv1.RegisterTaskServer(gRPC, &serverAPI{calculatorServer: calculator})

}

func (s *serverAPI) Create(ctx context.Context, req *calcv1.CreateRequest) (*calcv1.CreateResponse, error) {
	if req.GetOwnerId() <= 0 {
		err := status.Error(codes.InvalidArgument, "owner id is required")
		log.Println(err)
		return nil, err
	}

	agentId, err := s.calculatorServer.InitAgent(req.GetOwnerId())
	if err != nil {
		log.Println(err)
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &calcv1.CreateResponse{AgentId: agentId}, nil
}

func (s *serverAPI) Publish(ctx context.Context, req *calcv1.PublishRequest) (*calcv1.PublishResponse, error) {
	if err := validateOperand(req.GetSubexpression().GetLeftOperand()); err != nil {
		log.Println(err)
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	if err := validateOperand(req.GetSubexpression().GetRightOperand()); err != nil {
		log.Println(err)
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	var left, right constant.Value
	var errleft, errright error
	switch req.GetSubexpression().GetLeftOperandType() {
	case calcv1.OperandType_int:
		var leftInt int64
		if leftInt, errleft = strconv.ParseInt(req.GetSubexpression().GetLeftOperand(), 10, 64); errleft != nil {
			log.Println(errleft)
			return nil, status.Error(codes.InvalidArgument, errleft.Error())
		}

		left = constant.MakeInt64(leftInt)
	case calcv1.OperandType_float:
		var leftFloat float64
		if leftFloat, errleft = strconv.ParseFloat(req.GetSubexpression().GetLeftOperand(), 64); errleft != nil {
			log.Println(errleft)
			return nil, status.Error(codes.InvalidArgument, errleft.Error())
		}
		left = constant.MakeFloat64(leftFloat)
	}

	// Парсинг правого операнда
	switch req.GetSubexpression().GetRightOperandType() {
	case calcv1.OperandType_int:
		var rightInt int64
		if rightInt, errright = strconv.ParseInt(req.GetSubexpression().GetRightOperand(), 10, 64); errright != nil {
			log.Println(errright)
			return nil, status.Error(codes.InvalidArgument, errright.Error())
		}
		right = constant.MakeInt64(rightInt)
	case calcv1.OperandType_float:
		var rightFloat float64
		if rightFloat, errright = strconv.ParseFloat(req.GetSubexpression().GetRightOperand(), 64); errright != nil {
			log.Println(errright)
			return nil, status.Error(codes.InvalidArgument, errright.Error())
		}
		right = constant.MakeFloat64(rightFloat)
	}

	_task := task.NewTask(left, right, convertOpToToken(req.GetSubexpression().GetOp()), time.Duration(req.TimeToCalc))
	taskResult, taskOperandType, err := s.calculatorServer.ManageTask(ctx, _task)
	if err != nil {
		log.Println(err)
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &calcv1.PublishResponse{
		UserId:                  req.UserId,
		SubExpressionId:         int64(_task.Id),
		ResultOfCalculation:     taskResult.String(),
		ResultOfCalculationType: taskOperandType,
	}, nil

}

func validateOperand(operand string) error {
	// Регулярное выражение, которое допускает только цифры и точки.
	// Важно, что мы не проверяем здесь количество точек или их расположение,
	// только допустимые символы.
	validRegex := regexp.MustCompile(`^[0-9.]*$`)

	// Проверяем, соответствует ли входная строка допустимому шаблону.
	if validRegex.MatchString(operand) {
		return nil // Строка валидна
	} else {
		return errors.New("invalid characters in operand")
	}
}

func convertOpToToken(operatorType calcv1.OperatorType) token.Token {
	// I use go stdlib library token to calculate basic binary operations
	switch operatorType {
	case calcv1.OperatorType_add:
		return 12
	case calcv1.OperatorType_sub:
		return 13
	case calcv1.OperatorType_mul:
		return 14
	case calcv1.OperatorType_quo:
		return 15
	}
	// else return token that leads to error
	return -1
}
