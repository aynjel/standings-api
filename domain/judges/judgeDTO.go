package judges

import (
	"strings"

	"anggi.tabulation/utils/errors"
)

type Judge struct {
	ID        int64  `json:"id"`
	Name      string `json:"name"`
	Phone     string `json:"phone"`
	Email     string `json:"email"`
	Username  string `json:"username"`
	Pin	   		int    `json:"pin"`
	CreatedAt string `json:"created_at"`
}

type JudgeEvent struct {
	ID        int64  `json:"id"`
	EventID   int64  `json:"event_id"`
	JudgeID   int64  `json:"judge_id"`
	CreatedAt string `json:"created_at"`
}

func (judge *Judge) Validate() *errors.RestErr {
	judge.Name = strings.TrimSpace(judge.Name)
	judge.Phone = strings.TrimSpace(judge.Phone)
	judge.Email = strings.TrimSpace(judge.Email)

	if judge.Name == "" {
		return errors.NewBadRequestError("invalid name")
	}
	if judge.Phone == "" {
		return errors.NewBadRequestError("invalid phone")
	}
	if judge.Email == "" {
		return errors.NewBadRequestError("invalid email")
	}
	return nil
}

func (judgeEvent *JudgeEvent) Validate() *errors.RestErr {
	if judgeEvent.EventID == 0 {
		return errors.NewBadRequestError("invalid event id")
	}
	if judgeEvent.JudgeID == 0 {
		return errors.NewBadRequestError("invalid judge id")
	}
	return nil
}