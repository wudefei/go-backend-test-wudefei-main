package dto

import (
	"user-api/internal/data"
	"user-api/internal/model"
)

type DtoAssembler struct {
	logID string
}

func NewDtoAssembler(logID string) *DtoAssembler {
	return &DtoAssembler{
		logID: logID,
	}
}

func (d *DtoAssembler) AssembleQueryUser(user *model.User) (data.QueryUserRspData, error) {
	return data.QueryUserRspData{
		ID:          user.ID,
		Name:        user.Name,
		Birth:       user.Birth.Format("2006-01-02"),
		Address:     user.Address,
		Description: user.Description,
		CreatedAt:   user.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt:   user.UpdatedAt.Format("2006-01-02 15:04:05"),
	}, nil
}
