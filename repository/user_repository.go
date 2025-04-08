package repository

import (
	"database/sql"
	"log"

	"agricultural_product_gin/model"
)

// UserRepository 用户数据仓库
type UserRepository struct {
	DB *sql.DB
}

// NewUserRepository 创建用户仓库
func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{DB: db}
}

// FindByUsername 根据用户名查找用户
func (r *UserRepository) FindByUsername(username string) (*model.User, error) {
	query := `SELECT id, username, password, 
              sex, 
              COALESCE(name, '') as name, 
              COALESCE(phone, '') as phone 
              FROM user WHERE username = ?`
	row := r.DB.QueryRow(query, username)

	user := &model.User{}
	err := row.Scan(&user.ID, &user.Username, &user.Password, &user.Sex, &user.Name, &user.Phone)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		log.Println("查询用户失败:", err)
		return nil, err
	}

	return user, nil
}

// Save 保存用户
func (r *UserRepository) Save(username, password string) error {
	query := "INSERT INTO user(username, password) VALUES(?, ?)"
	_, err := r.DB.Exec(query, username, password)
	if err != nil {
		log.Println("保存用户失败:", err)
		return err
	}
	return nil
}

// Update 更新用户信息
func (r *UserRepository) Update(user *model.User) error {
	query := `UPDATE user SET 
              username = CASE WHEN ? != '' THEN ? ELSE username END,
              sex = CASE WHEN ? IS NOT NULL THEN ? ELSE sex END,
              name = CASE WHEN ? IS NOT NULL THEN ? ELSE name END,
              phone = CASE WHEN ? IS NOT NULL THEN ? ELSE phone END
              WHERE id = ?`

	_, err := r.DB.Exec(query,
		user.Username, user.Username,
		user.Sex, user.Sex,
		user.Name, user.Name,
		user.Phone, user.Phone,
		user.ID)

	if err != nil {
		log.Println("更新用户失败:", err)
		return err
	}
	return nil
}

// UpdatePassword 更新密码
func (r *UserRepository) UpdatePassword(userID int, newPassword string) error {
	query := "UPDATE user SET password = ? WHERE id = ?"
	_, err := r.DB.Exec(query, newPassword, userID)
	if err != nil {
		log.Println("更新密码失败:", err)
		return err
	}
	return nil
}

// GetByID 根据ID获取用户
func (r *UserRepository) GetByID(id int) (*model.User, error) {
	query := "SELECT id, username, password, sex, name, phone FROM user WHERE id = ?"
	row := r.DB.QueryRow(query, id)

	user := &model.User{}
	err := row.Scan(&user.ID, &user.Username, &user.Password, &user.Sex, &user.Name, &user.Phone)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		log.Println("获取用户失败:", err)
		return nil, err
	}

	return user, nil
}
