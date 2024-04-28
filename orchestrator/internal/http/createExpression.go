package http

import (
	"context"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"orchestrator/internal/auth"
	"orchestrator/internal/expression"
	"time"
)

type createExpressionRequest struct {
	Expression  string        `json:"expression"`
	TimeToCalc  time.Duration `json:"time_to_calc"`
	StartToCalc bool          `json:"start_to_calc"`
}

func (osrv *OrchServer) createExpression(ctx context.Context) func(c echo.Context) error {
	return func(c echo.Context) error {
		// забираем уникальный идентификатор "username" из токена
		userInfo := c.Get("userInfo").(*jwt.Token)
		claims := userInfo.Claims.(*jwtCustomClaims)
		name := claims.Name

		// забираем выражение, время вычисления подвыражения, и надо ли его вычислить сразу после создания
		var payload *createExpressionRequest
		err := (&echo.DefaultBinder{}).BindBody(c, &payload)
		if err != nil {
			log.Println(err)
			return c.JSON(http.StatusBadRequest, echo.Map{"status": "failed to parse expression data: " + err.Error()})
		}

		user, err := auth.GetUserInfoByName(ctx, osrv.DB, name)
		if err != nil {
			log.Println(err)
			return echo.ErrInternalServerError
		}

		expr, err := expression.NewExpression(ctx, user.ID, payload.Expression)
		if err != nil {
			log.Println(err)
			return c.JSON(http.StatusBadRequest, echo.Map{"status": "failed to create new expression: " + err.Error()})
		}

		exprID, err := expr.InsertExpression(ctx, osrv.DB)
		if err != nil {
			log.Println(err)
			return c.JSON(http.StatusBadRequest, echo.Map{"status": "failed to create new expression: " + err.Error()})
		}

		if payload.StartToCalc {
			//TODO: сразу начать высчитывать выражение
			return c.JSON(http.StatusOK, echo.Map{"status": expr.Status, "exprID": exprID})
		}

		return c.JSON(http.StatusCreated, echo.Map{"status": expr.Status, "exprID": exprID})
	}
}
