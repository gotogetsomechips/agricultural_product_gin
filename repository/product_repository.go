package repository

import (
	"database/sql"
	"fmt"
	"log"
	"strings"

	"agricultural_product_gin/model"
)

// ProductRepository 产品数据仓库
type ProductRepository struct {
	DB *sql.DB
}

// NewProductRepository 创建产品仓库
func NewProductRepository(db *sql.DB) *ProductRepository {
	return &ProductRepository{DB: db}
}

// Save 保存产品
func (r *ProductRepository) Save(product *model.Product) (int, error) {
	query := "INSERT INTO product(pd_name, type, image, pd_description, unit_price) VALUES(?, ?, ?, ?, ?)"
	result, err := r.DB.Exec(query, product.Name, product.Type, product.Image, product.Description, product.UnitPrice)
	if err != nil {
		log.Println("保存产品失败:", err)
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		log.Println("获取产品ID失败:", err)
		return 0, err
	}

	return int(id), nil
}

// Update 更新产品
func (r *ProductRepository) Update(product *model.Product) error {
	query := "UPDATE product SET pd_name = ?, type = ?, image = ?, pd_description = ?, unit_price = ? WHERE pd_id = ?"
	_, err := r.DB.Exec(query, product.Name, product.Type, product.Image, product.Description, product.UnitPrice, product.ID)
	if err != nil {
		log.Println("更新产品失败:", err)
		return err
	}
	return nil
}

// Delete 删除产品
func (r *ProductRepository) Delete(id int) error {
	query := "DELETE FROM product WHERE pd_id = ?"
	_, err := r.DB.Exec(query, id)
	if err != nil {
		log.Println("删除产品失败:", err)
		return err
	}
	return nil
}

// GetByID 根据ID获取产品
func (r *ProductRepository) GetByID(id int) (*model.Product, error) {
	query := "SELECT pd_id, pd_name, type, image, pd_description, unit_price FROM product WHERE pd_id = ?"
	row := r.DB.QueryRow(query, id)

	product := &model.Product{}
	err := row.Scan(&product.ID, &product.Name, &product.Type, &product.Image, &product.Description, &product.UnitPrice)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		log.Println("获取产品失败:", err)
		return nil, err
	}

	return product, nil
}

// FindAll 查找所有产品
func (r *ProductRepository) FindAll() ([]*model.Product, error) {
	query := "SELECT pd_id, pd_name, type, image, pd_description, unit_price FROM product"
	rows, err := r.DB.Query(query)
	if err != nil {
		log.Println("查询产品失败:", err)
		return nil, err
	}
	defer rows.Close()

	var products []*model.Product
	for rows.Next() {
		product := &model.Product{}
		err := rows.Scan(&product.ID, &product.Name, &product.Type, &product.Image, &product.Description, &product.UnitPrice)
		if err != nil {
			log.Println("读取产品数据失败:", err)
			return nil, err
		}
		products = append(products, product)
	}

	return products, nil
}

// FindByCondition 根据条件查询产品
func (r *ProductRepository) FindByCondition(name, productType string) ([]*model.Product, error) {
	// 构建查询条件
	conditions := []string{}
	args := []interface{}{}

	if name != "" {
		conditions = append(conditions, "pd_name LIKE ?")
		args = append(args, "%"+name+"%")
	}

	if productType != "" {
		conditions = append(conditions, "type = ?")
		args = append(args, productType)
	}

	// 构建SQL
	query := "SELECT pd_id, pd_name, type, image, pd_description, unit_price FROM product"
	if len(conditions) > 0 {
		query += " WHERE " + strings.Join(conditions, " AND ")
	}

	// 执行查询
	rows, err := r.DB.Query(query, args...)
	if err != nil {
		log.Println("条件查询产品失败:", err)
		return nil, err
	}
	defer rows.Close()

	var products []*model.Product
	for rows.Next() {
		product := &model.Product{}
		err := rows.Scan(&product.ID, &product.Name, &product.Type, &product.Image, &product.Description, &product.UnitPrice)
		if err != nil {
			log.Println("读取产品数据失败:", err)
			return nil, err
		}
		products = append(products, product)
	}

	return products, nil
}

// PageQuery 分页查询产品
func (r *ProductRepository) PageQuery(page, pageSize int, name, productType string) ([]*model.Product, int64, error) {
	// 构建查询条件
	conditions := []string{}
	args := []interface{}{}

	if name != "" {
		conditions = append(conditions, "pd_name LIKE ?")
		args = append(args, "%"+name+"%")
	}

	if productType != "" {
		conditions = append(conditions, "type = ?")
		args = append(args, productType)
	}

	// 构建条件子句
	whereClause := ""
	if len(conditions) > 0 {
		whereClause = " WHERE " + strings.Join(conditions, " AND ")
	}

	// 查询总记录数
	countQuery := fmt.Sprintf("SELECT COUNT(*) FROM product%s", whereClause)
	var total int64
	err := r.DB.QueryRow(countQuery, args...).Scan(&total)
	if err != nil {
		log.Println("查询产品总数失败:", err)
		return nil, 0, err
	}

	// 查询当前页数据 - 添加 unit_price 字段
	offset := (page - 1) * pageSize
	dataQuery := fmt.Sprintf("SELECT pd_id, pd_name, type, image, pd_description, unit_price FROM product%s LIMIT ? OFFSET ?", whereClause)
	queryArgs := append(args, pageSize, offset)

	rows, err := r.DB.Query(dataQuery, queryArgs...)
	if err != nil {
		log.Println("分页查询产品失败:", err)
		return nil, 0, err
	}
	defer rows.Close()

	var products []*model.Product
	for rows.Next() {
		product := &model.Product{}
		// 添加 unit_price 字段到 Scan 方法
		err := rows.Scan(&product.ID, &product.Name, &product.Type, &product.Image, &product.Description, &product.UnitPrice)
		if err != nil {
			log.Println("读取产品数据失败:", err)
			return nil, 0, err
		}
		products = append(products, product)
	}

	return products, total, nil
}

// GetProductTypes 获取所有产品类型
func (r *ProductRepository) GetProductTypes() ([]string, error) {
	query := "SELECT DISTINCT type FROM product"
	rows, err := r.DB.Query(query)
	if err != nil {
		log.Println("查询产品类型失败:", err)
		return nil, err
	}
	defer rows.Close()

	var types []string
	for rows.Next() {
		var productType string
		err := rows.Scan(&productType)
		if err != nil {
			log.Println("读取产品类型失败:", err)
			return nil, err
		}
		types = append(types, productType)
	}

	return types, nil
}
