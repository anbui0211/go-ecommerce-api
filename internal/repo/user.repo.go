package repo

import (
	"ecommerce/global"
	"ecommerce/internal/database"
)

type IUserRepository interface {
	GetUserByEmail(email string) bool
}

type userRepository struct {
	sqlc *database.Queries
}

func NewUserRepository() IUserRepository {
	return &userRepository{
		sqlc: database.New(global.Mdbc),
	}
}

func (up *userRepository) GetUserByEmail(email string) bool {
	// SELECT * FROM user WHERE email = '??' ORDER BY email
	// row := global.Mdb.Table(TableNameGoCrmUser).Where("user_email = ?", email).First(&model.GoCrmUser{}).RowsAffected

	// user, err := up.sqlc.GetUserByEmailSQLC(ctx, email)
	// if err != nil {
	// 	return false
	// }
	return 1 != 0
}
