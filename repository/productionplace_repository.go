package repository

import (
	"database/sql"
	"fmt"
	"log"
	"strings"

	"agricultural_product_gin/model"
)

// ProductionPlaceRepository 生产地仓库
type ProductionPlaceRepository struct {
	DB *sql.DB
}

// NewProductionPlaceRepository 创建生产地仓库
func NewProductionPlaceRepository(db *sql.DB) *ProductionPlaceRepository {
	return &ProductionPlaceRepository{DB: db}
}

// Save 保存生产地信息
func (r *ProductionPlaceRepository) Save(place *model.ProductionPlace) (int, error) {
	query := "INSERT INTO product_place(pp_address, pp_administrator, pp_phone) VALUES(?, ?, ?)"
	result, err := r.DB.Exec(query, place.Address, place.Administrator, place.Phone)
	if err != nil {
		log.Println("保存生产地信息失败:", err)
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		log.Println("获取生产地ID失败:", err)
		return 0, err
	}

	return int(id), nil
}

// Update 更新生产地信息
func (r *ProductionPlaceRepository) Update(place *model.ProductionPlace) error {
	query := "UPDATE product_place SET pp_address = ?, pp_administrator = ?, pp_phone = ? WHERE pp_id = ?"
	_, err := r.DB.Exec(query, place.Address, place.Administrator, place.Phone, place.ID)
	if err != nil {
		log.Println("更新生产地信息失败:", err)
		return err
	}
	return nil
}

// Delete 删除生产地信息
func (r *ProductionPlaceRepository) Delete(id int) error {
	query := "DELETE FROM product_place WHERE pp_id = ?"
	_, err := r.DB.Exec(query, id)
	if err != nil {
		log.Println("删除生产地信息失败:", err)
		return err
	}
	return nil
}

// GetByID 根据ID获取生产地信息
func (r *ProductionPlaceRepository) GetByID(id int) (*model.ProductionPlace, error) {
	query := "SELECT pp_id, pp_address, pp_administrator, pp_phone FROM product_place WHERE pp_id = ?"
	row := r.DB.QueryRow(query, id)

	place := &model.ProductionPlace{}
	err := row.Scan(&place.ID, &place.Address, &place.Administrator, &place.Phone)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		log.Println("获取生产地信息失败:", err)
		return nil, err
	}

	return place, nil
}

// PageQuery 分页查询生产地信息
func (r *ProductionPlaceRepository) PageQuery(
	page, pageSize int,
	id, address, administrator string,
) ([]*model.ProductionPlace, int64, error) {
	// 构建查询条件
	conditions := []string{}
	args := []interface{}{}

	if id != "" {
		conditions = append(conditions, "pp_id = ?")
		args = append(args, id)
	}
	if address != "" {
		conditions = append(conditions, "pp_address LIKE ?")
		args = append(args, "%"+address+"%")
	}
	if administrator != "" {
		conditions = append(conditions, "pp_administrator LIKE ?")
		args = append(args, "%"+administrator+"%")
	}

	// 构建条件子句
	whereClause := ""
	if len(conditions) > 0 {
		whereClause = " WHERE " + strings.Join(conditions, " AND ")
	}

	// 查询总记录数
	countQuery := "SELECT COUNT(*) FROM product_place" + whereClause
	var total int64
	err := r.DB.QueryRow(countQuery, args...).Scan(&total)
	if err != nil {
		log.Println("查询生产地总数失败:", err)
		return nil, 0, err
	}

	// 查询当前页数据
	offset := (page - 1) * pageSize
	dataQuery := fmt.Sprintf(
		"SELECT pp_id, pp_address, pp_administrator, pp_phone FROM product_place%s LIMIT ? OFFSET ?",
		whereClause)
	queryArgs := append(args, pageSize, offset)

	rows, err := r.DB.Query(dataQuery, queryArgs...)
	if err != nil {
		log.Println("分页查询生产地信息失败:", err)
		return nil, 0, err
	}
	defer rows.Close()

	var places []*model.ProductionPlace
	for rows.Next() {
		place := &model.ProductionPlace{}
		err := rows.Scan(&place.ID, &place.Address, &place.Administrator, &place.Phone)
		if err != nil {
			log.Println("读取生产地数据失败:", err)
			return nil, 0, err
		}
		places = append(places, place)
	}

	return places, total, nil
}

// GetAll 获取所有生产地信息
func (r *ProductionPlaceRepository) GetAll() ([]*model.ProductionPlace, error) {
	query := "SELECT pp_id, pp_address, pp_administrator, pp_phone FROM product_place"
	rows, err := r.DB.Query(query)
	if err != nil {
		log.Println("查询所有生产地信息失败:", err)
		return nil, err
	}
	defer rows.Close()

	var places []*model.ProductionPlace
	for rows.Next() {
		place := &model.ProductionPlace{}
		err := rows.Scan(&place.ID, &place.Address, &place.Administrator, &place.Phone)
		if err != nil {
			log.Println("读取生产地数据失败:", err)
			return nil, err
		}
		places = append(places, place)
	}

	return places, nil
}
