package repository

import (
	"agricultural_product_gin/model"
	"database/sql"
	"log"
	"strings"
)

// ProductionRepository 生产信息仓库
type ProductionRepository struct {
	DB *sql.DB
}

// NewProductionRepository 创建生产信息仓库
func NewProductionRepository(db *sql.DB) *ProductionRepository {
	return &ProductionRepository{DB: db}
}

// Save 保存生产信息
func (r *ProductionRepository) Save(production *model.ProductionInfo) (int, error) {
	query := `INSERT INTO product_info (
        product_id, product_place_id, seed, pi_description, 
        planting_date, harvest_date
    ) VALUES (?, ?, ?, ?, ?, ?)`

	result, err := r.DB.Exec(query,
		production.ProductID, production.ProductPlaceID, production.SeedSource,
		production.Description, production.PlantingDate, production.HarvestDate)
	if err != nil {
		log.Println("保存生产信息失败:", err)
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		log.Println("获取生产信息ID失败:", err)
		return 0, err
	}

	return int(id), nil
}

// Update 更新生产信息
func (r *ProductionRepository) Update(production *model.ProductionInfo) error {
	query := `UPDATE product_info SET 
        product_id = ?, product_place_id = ?, seed = ?, 
        pi_description = ?, planting_date = ?, harvest_date = ?
        WHERE pi_id = ?`

	_, err := r.DB.Exec(query,
		production.ProductID, production.ProductPlaceID, production.SeedSource,
		production.Description, production.PlantingDate, production.HarvestDate,
		production.ID)
	if err != nil {
		log.Println("更新生产信息失败:", err)
		return err
	}
	return nil
}

// Delete 删除生产信息
func (r *ProductionRepository) Delete(id int) error {
	query := "DELETE FROM product_info WHERE pi_id = ?"
	_, err := r.DB.Exec(query, id)
	if err != nil {
		log.Println("删除生产信息失败:", err)
		return err
	}
	return nil
}

// GetByID 根据ID获取生产信息(带详细信息)
func (r *ProductionRepository) GetByID(id int) (*model.ProductionInfoWithDetails, error) {
	query := `SELECT 
        pi.pi_id, pi.product_id, pi.product_place_id, pi.seed, 
        pi.pi_description, pi.planting_date, pi.harvest_date,
        pp.pp_administrator, pp.pp_phone,
        pd.pd_name, pp.pp_address
    FROM product_info pi
    LEFT JOIN product pd ON pi.product_id = pd.pd_id
    LEFT JOIN product_place pp ON pi.product_place_id = pp.pp_id
    WHERE pi.pi_id = ?`

	row := r.DB.QueryRow(query, id)

	info := &model.ProductionInfoWithDetails{}
	err := row.Scan(
		&info.ID, &info.ProductID, &info.ProductPlaceID, &info.SeedSource,
		&info.Description, &info.PlantingDate, &info.HarvestDate,
		&info.Administrator, &info.Phone,
		&info.ProductName, &info.ProductionPlace)

	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		log.Println("获取生产信息失败:", err)
		return nil, err
	}

	return info, nil
}

// PageQuery 分页查询生产信息
func (r *ProductionRepository) PageQuery(
	page, pageSize int,
	productInfoID, productName, productPlace, seed, administrator string,
) ([]*model.ProductionInfoWithDetails, int64, error) {
	// 构建查询条件
	conditions := []string{}
	args := []interface{}{}

	if productInfoID != "" {
		conditions = append(conditions, "pi.pi_id = ?")
		args = append(args, productInfoID)
	}
	if productName != "" {
		conditions = append(conditions, "pd.pd_name LIKE ?")
		args = append(args, "%"+productName+"%")
	}
	if productPlace != "" {
		conditions = append(conditions, "pp.pp_address LIKE ?")
		args = append(args, "%"+productPlace+"%")
	}
	if seed != "" {
		conditions = append(conditions, "pi.seed LIKE ?")
		args = append(args, "%"+seed+"%")
	}
	if administrator != "" {
		conditions = append(conditions, "pp.pp_administrator LIKE ?")
		args = append(args, "%"+administrator+"%")
	}

	// 构建条件子句
	whereClause := ""
	if len(conditions) > 0 {
		whereClause = " WHERE " + strings.Join(conditions, " AND ")
	}

	// 查询总记录数
	countQuery := `SELECT COUNT(*) FROM product_info pi
        LEFT JOIN product pd ON pi.product_id = pd.pd_id
        LEFT JOIN product_place pp ON pi.product_place_id = pp.pp_id` + whereClause

	var total int64
	err := r.DB.QueryRow(countQuery, args...).Scan(&total)
	if err != nil {
		log.Println("查询生产信息总数失败:", err)
		return nil, 0, err
	}

	// 查询当前页数据
	offset := (page - 1) * pageSize
	dataQuery := `SELECT 
        pi.pi_id, pi.product_id, pi.product_place_id, pi.seed, 
        pi.pi_description, pi.planting_date, pi.harvest_date,
        pp.pp_administrator, pp.pp_phone,
        pd.pd_name, pp.pp_address
    FROM product_info pi
    LEFT JOIN product pd ON pi.product_id = pd.pd_id
    LEFT JOIN product_place pp ON pi.product_place_id = pp.pp_id` +
		whereClause + " LIMIT ? OFFSET ?"

	queryArgs := append(args, pageSize, offset)

	rows, err := r.DB.Query(dataQuery, queryArgs...)
	if err != nil {
		log.Println("分页查询生产信息失败:", err)
		return nil, 0, err
	}
	defer rows.Close()

	var productions []*model.ProductionInfoWithDetails
	for rows.Next() {
		info := &model.ProductionInfoWithDetails{}
		err := rows.Scan(
			&info.ID, &info.ProductID, &info.ProductPlaceID, &info.SeedSource,
			&info.Description, &info.PlantingDate, &info.HarvestDate,
			&info.Administrator, &info.Phone,
			&info.ProductName, &info.ProductionPlace)
		if err != nil {
			log.Println("读取生产信息数据失败:", err)
			return nil, 0, err
		}
		productions = append(productions, info)
	}

	return productions, total, nil
}

// GetAll 获取所有生产信息(用于下拉选择等)
func (r *ProductionRepository) GetAll() ([]*model.ProductionInfoWithDetails, error) {
	query := `SELECT 
        pi.pi_id, pi.product_id, pi.product_place_id, pi.seed, 
        pi.pi_description, pi.planting_date, pi.harvest_date,
        pp.pp_administrator, pp.pp_phone,
        pd.pd_name, pp.pp_address
    FROM product_info pi
    LEFT JOIN product pd ON pi.product_id = pd.pd_id
    LEFT JOIN product_place pp ON pi.product_place_id = pp.pp_id`

	rows, err := r.DB.Query(query)
	if err != nil {
		log.Println("查询所有生产信息失败:", err)
		return nil, err
	}
	defer rows.Close()

	var productions []*model.ProductionInfoWithDetails
	for rows.Next() {
		info := &model.ProductionInfoWithDetails{}
		err := rows.Scan(
			&info.ID, &info.ProductID, &info.ProductPlaceID, &info.SeedSource,
			&info.Description, &info.PlantingDate, &info.HarvestDate,
			&info.Administrator, &info.Phone,
			&info.ProductName, &info.ProductionPlace)
		if err != nil {
			log.Println("读取生产信息数据失败:", err)
			return nil, err
		}
		productions = append(productions, info)
	}

	return productions, nil
}
