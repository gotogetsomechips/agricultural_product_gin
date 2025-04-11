package repository

import (
	"database/sql"
	"fmt"
	"log"
	"strings"

	"agricultural_product_gin/model"
)

// CompanyRepository 公司数据仓库
type CompanyRepository struct {
	DB *sql.DB
}

// NewCompanyRepository 创建公司仓库
func NewCompanyRepository(db *sql.DB) *CompanyRepository {
	return &CompanyRepository{DB: db}
}

// Save 保存公司
func (r *CompanyRepository) Save(company *model.Company) (int, error) {
	query := "INSERT INTO company(com_name, com_address, com_administrator, com_phone) VALUES(?, ?, ?, ?)"
	result, err := r.DB.Exec(query, company.Name, company.Address, company.Administrator, company.Phone)
	if err != nil {
		log.Println("保存公司失败:", err)
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		log.Println("获取公司ID失败:", err)
		return 0, err
	}

	return int(id), nil
}

// Update 更新公司
func (r *CompanyRepository) Update(company *model.Company) error {
	query := "UPDATE company SET com_name = ?, com_address = ?, com_administrator = ?, com_phone = ? WHERE com_id = ?"
	_, err := r.DB.Exec(query, company.Name, company.Address, company.Administrator, company.Phone, company.ID)
	if err != nil {
		log.Println("更新公司失败:", err)
		return err
	}
	return nil
}

// Delete 删除公司
func (r *CompanyRepository) Delete(id int) error {
	query := "DELETE FROM company WHERE com_id = ?"
	_, err := r.DB.Exec(query, id)
	if err != nil {
		log.Println("删除公司失败:", err)
		return err
	}
	return nil
}

// GetByID 根据ID获取公司
func (r *CompanyRepository) GetByID(id int) (*model.Company, error) {
	query := "SELECT com_id, com_name, com_address, com_administrator, com_phone FROM company WHERE com_id = ?"
	row := r.DB.QueryRow(query, id)

	company := &model.Company{}
	err := row.Scan(&company.ID, &company.Name, &company.Address, &company.Administrator, &company.Phone)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		log.Println("获取公司失败:", err)
		return nil, err
	}

	return company, nil
}

// FindAll 查找所有公司
func (r *CompanyRepository) FindAll() ([]*model.Company, error) {
	query := "SELECT com_id, com_name, com_address, com_administrator, com_phone FROM company"
	rows, err := r.DB.Query(query)
	if err != nil {
		log.Println("查询公司失败:", err)
		return nil, err
	}
	defer rows.Close()

	var companies []*model.Company
	for rows.Next() {
		company := &model.Company{}
		err := rows.Scan(&company.ID, &company.Name, &company.Address, &company.Administrator, &company.Phone)
		if err != nil {
			log.Println("读取公司数据失败:", err)
			return nil, err
		}
		companies = append(companies, company)
	}

	return companies, nil
}

// PageQuery 分页查询公司
func (r *CompanyRepository) PageQuery(page, pageSize int, name, address, administrator, phone string) ([]*model.Company, int64, error) {
	// 构建查询条件
	conditions := []string{}
	args := []interface{}{}

	if name != "" {
		conditions = append(conditions, "com_name LIKE ?")
		args = append(args, "%"+name+"%")
	}

	if address != "" {
		conditions = append(conditions, "com_address LIKE ?")
		args = append(args, "%"+address+"%")
	}

	if administrator != "" {
		conditions = append(conditions, "com_administrator LIKE ?")
		args = append(args, "%"+administrator+"%")
	}

	if phone != "" {
		conditions = append(conditions, "com_phone LIKE ?")
		args = append(args, "%"+phone+"%")
	}

	// 构建条件子句
	whereClause := ""
	if len(conditions) > 0 {
		whereClause = " WHERE " + strings.Join(conditions, " AND ")
	}

	// 查询总记录数
	countQuery := fmt.Sprintf("SELECT COUNT(*) FROM company%s", whereClause)
	var total int64
	err := r.DB.QueryRow(countQuery, args...).Scan(&total)
	if err != nil {
		log.Println("查询公司总数失败:", err)
		return nil, 0, err
	}

	// 查询当前页数据
	offset := (page - 1) * pageSize
	dataQuery := fmt.Sprintf("SELECT com_id, com_name, com_address, com_administrator, com_phone FROM company%s LIMIT ? OFFSET ?", whereClause)
	queryArgs := append(args, pageSize, offset)

	rows, err := r.DB.Query(dataQuery, queryArgs...)
	if err != nil {
		log.Println("分页查询公司失败:", err)
		return nil, 0, err
	}
	defer rows.Close()

	var companies []*model.Company
	for rows.Next() {
		company := &model.Company{}
		err := rows.Scan(&company.ID, &company.Name, &company.Address, &company.Administrator, &company.Phone)
		if err != nil {
			log.Println("读取公司数据失败:", err)
			return nil, 0, err
		}
		companies = append(companies, company)
	}

	return companies, total, nil
}
