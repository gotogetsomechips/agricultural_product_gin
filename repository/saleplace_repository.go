package repository

import (
	"database/sql"
	"fmt"
	"log"
	"strings"

	"agricultural_product_gin/model"
)

// SalePlaceRepository 销售地数据仓库
type SalePlaceRepository struct {
	DB *sql.DB
}

// NewSalePlaceRepository 创建销售地仓库
func NewSalePlaceRepository(db *sql.DB) *SalePlaceRepository {
	return &SalePlaceRepository{DB: db}
}

// Save 保存销售地
func (r *SalePlaceRepository) Save(salePlace *model.SalePlace) (int, error) {
	query := "INSERT INTO sale_place(sp_address, sp_administrator, sp_phone) VALUES(?, ?, ?)"
	result, err := r.DB.Exec(query, salePlace.Address, salePlace.Administrator, salePlace.Phone)
	if err != nil {
		log.Println("保存销售地失败:", err)
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		log.Println("获取销售地ID失败:", err)
		return 0, err
	}

	return int(id), nil
}

// Update 更新销售地
func (r *SalePlaceRepository) Update(salePlace *model.SalePlace) error {
	query := "UPDATE sale_place SET sp_address = ?, sp_administrator = ?, sp_phone = ? WHERE sp_id = ?"
	_, err := r.DB.Exec(query, salePlace.Address, salePlace.Administrator, salePlace.Phone, salePlace.ID)
	if err != nil {
		log.Println("更新销售地失败:", err)
		return err
	}
	return nil
}

// Delete 删除销售地
func (r *SalePlaceRepository) Delete(id int) error {
	query := "DELETE FROM sale_place WHERE sp_id = ?"
	_, err := r.DB.Exec(query, id)
	if err != nil {
		log.Println("删除销售地失败:", err)
		return err
	}
	return nil
}

// GetByID 根据ID获取销售地
func (r *SalePlaceRepository) GetByID(id int) (*model.SalePlace, error) {
	query := "SELECT sp_id, sp_address, sp_administrator, sp_phone FROM sale_place WHERE sp_id = ?"
	row := r.DB.QueryRow(query, id)

	salePlace := &model.SalePlace{}
	err := row.Scan(&salePlace.ID, &salePlace.Address, &salePlace.Administrator, &salePlace.Phone)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		log.Println("获取销售地失败:", err)
		return nil, err
	}

	return salePlace, nil
}

// FindAll 查找所有销售地
func (r *SalePlaceRepository) FindAll() ([]*model.SalePlace, error) {
	query := "SELECT sp_id, sp_address, sp_administrator, sp_phone FROM sale_place"
	rows, err := r.DB.Query(query)
	if err != nil {
		log.Println("查询销售地失败:", err)
		return nil, err
	}
	defer rows.Close()

	var salePlaces []*model.SalePlace
	for rows.Next() {
		salePlace := &model.SalePlace{}
		err := rows.Scan(&salePlace.ID, &salePlace.Address, &salePlace.Administrator, &salePlace.Phone)
		if err != nil {
			log.Println("读取销售地数据失败:", err)
			return nil, err
		}
		salePlaces = append(salePlaces, salePlace)
	}

	return salePlaces, nil
}

// PageQuery 分页查询销售地
func (r *SalePlaceRepository) PageQuery(page, pageSize int, id, address, administrator, phone string) ([]*model.SalePlace, int64, error) {
	// 构建查询条件
	conditions := []string{}
	args := []interface{}{}

	if id != "" {
		conditions = append(conditions, "sp_id = ?")
		args = append(args, id)
	}

	if address != "" {
		conditions = append(conditions, "sp_address LIKE ?")
		args = append(args, "%"+address+"%")
	}

	if administrator != "" {
		conditions = append(conditions, "sp_administrator LIKE ?")
		args = append(args, "%"+administrator+"%")
	}

	if phone != "" {
		conditions = append(conditions, "sp_phone LIKE ?")
		args = append(args, "%"+phone+"%")
	}

	// 构建条件子句
	whereClause := ""
	if len(conditions) > 0 {
		whereClause = " WHERE " + strings.Join(conditions, " AND ")
	}

	// 查询总记录数
	countQuery := fmt.Sprintf("SELECT COUNT(*) FROM sale_place%s", whereClause)
	var total int64
	err := r.DB.QueryRow(countQuery, args...).Scan(&total)
	if err != nil {
		log.Println("查询销售地总数失败:", err)
		return nil, 0, err
	}

	// 查询当前页数据
	offset := (page - 1) * pageSize
	dataQuery := fmt.Sprintf("SELECT sp_id, sp_address, sp_administrator, sp_phone FROM sale_place%s LIMIT ? OFFSET ?", whereClause)
	queryArgs := append(args, pageSize, offset)

	rows, err := r.DB.Query(dataQuery, queryArgs...)
	if err != nil {
		log.Println("分页查询销售地失败:", err)
		return nil, 0, err
	}
	defer rows.Close()

	var salePlaces []*model.SalePlace
	for rows.Next() {
		salePlace := &model.SalePlace{}
		err := rows.Scan(&salePlace.ID, &salePlace.Address, &salePlace.Administrator, &salePlace.Phone)
		if err != nil {
			log.Println("读取销售地数据失败:", err)
			return nil, 0, err
		}
		salePlaces = append(salePlaces, salePlace)
	}

	return salePlaces, total, nil
}
