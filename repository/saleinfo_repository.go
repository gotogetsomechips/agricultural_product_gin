package repository

import (
	"database/sql"
	"fmt"
	"log"
	"strings"
	"time"

	"agricultural_product_gin/model"
)

// SaleInfoRepository 销售信息数据仓库
type SaleInfoRepository struct {
	DB *sql.DB
}

// NewSaleInfoRepository 创建销售信息仓库
func NewSaleInfoRepository(db *sql.DB) *SaleInfoRepository {
	return &SaleInfoRepository{DB: db}
}

// Save 保存销售信息
func (r *SaleInfoRepository) Save(saleInfo *model.SaleInfo) (int, error) {
	query := "INSERT INTO sale_info(logistics_id, sale_place_id, si_description, sale_time) VALUES(?, ?, ?, ?)"
	result, err := r.DB.Exec(query, saleInfo.LogisticsID, saleInfo.SalePlaceID, saleInfo.Description, saleInfo.SaleTime)
	if err != nil {
		log.Println("保存销售信息失败:", err)
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		log.Println("获取销售信息ID失败:", err)
		return 0, err
	}

	return int(id), nil
}

// Update 更新销售信息
func (r *SaleInfoRepository) Update(saleInfo *model.SaleInfo) error {
	query := "UPDATE sale_info SET logistics_id = ?, sale_place_id = ?, si_description = ?, sale_time = ? WHERE si_id = ?"
	_, err := r.DB.Exec(query, saleInfo.LogisticsID, saleInfo.SalePlaceID, saleInfo.Description, saleInfo.SaleTime, saleInfo.ID)
	if err != nil {
		log.Println("更新销售信息失败:", err)
		return err
	}
	return nil
}

// Delete 删除销售信息
func (r *SaleInfoRepository) Delete(id int) error {
	query := "DELETE FROM sale_info WHERE si_id = ?"
	_, err := r.DB.Exec(query, id)
	if err != nil {
		log.Println("删除销售信息失败:", err)
		return err
	}
	return nil
}

// GetByID 根据ID获取销售信息
func (r *SaleInfoRepository) GetByID(id int) (*model.SaleInfoVO, error) {
	query := `
        SELECT 
            si.si_id, si.logistics_id, si.sale_place_id, si.si_description, si.sale_time,
            pd.pd_name, sp.sp_address, sp.sp_administrator,
            log.start_location, log.destination
        FROM sale_info si
        LEFT JOIN sale_place sp ON sp.sp_id = si.sale_place_id
        LEFT JOIN logistics log ON log.log_id = si.logistics_id
        LEFT JOIN product_info pi ON pi.pi_id = log.product_info_id
        LEFT JOIN product pd ON pd.pd_id = pi.product_id
        WHERE si.si_id = ?`

	row := r.DB.QueryRow(query, id)

	saleInfo := &model.SaleInfoVO{}
	err := row.Scan(
		&saleInfo.ID, &saleInfo.LogisticsID, &saleInfo.SalePlaceID, &saleInfo.Description, &saleInfo.SaleTime,
		&saleInfo.ProductName, &saleInfo.SalePlace, &saleInfo.Administrator,
		&saleInfo.StartLocation, &saleInfo.Destination,
	)

	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		log.Println("获取销售信息失败:", err)
		return nil, err
	}

	return saleInfo, nil
}

// FindAll 查找所有销售信息
func (r *SaleInfoRepository) FindAll() ([]*model.SaleInfoVO, error) {
	query := `
        SELECT 
            si.si_id, si.logistics_id, si.sale_place_id, si.si_description, si.sale_time,
            pd.pd_name, sp.sp_address, sp.sp_administrator,
            log.start_location, log.destination
        FROM sale_info si
        LEFT JOIN sale_place sp ON sp.sp_id = si.sale_place_id
        LEFT JOIN logistics log ON log.log_id = si.logistics_id
        LEFT JOIN product_info pi ON pi.pi_id = log.product_info_id
        LEFT JOIN product pd ON pd.pd_id = pi.product_id`

	rows, err := r.DB.Query(query)
	if err != nil {
		log.Println("查询销售信息失败:", err)
		return nil, err
	}
	defer rows.Close()

	var saleInfos []*model.SaleInfoVO
	for rows.Next() {
		saleInfo := &model.SaleInfoVO{}
		err := rows.Scan(
			&saleInfo.ID, &saleInfo.LogisticsID, &saleInfo.SalePlaceID, &saleInfo.Description, &saleInfo.SaleTime,
			&saleInfo.ProductName, &saleInfo.SalePlace, &saleInfo.Administrator,
			&saleInfo.StartLocation, &saleInfo.Destination,
		)
		if err != nil {
			log.Println("读取销售信息数据失败:", err)
			return nil, err
		}
		saleInfos = append(saleInfos, saleInfo)
	}

	return saleInfos, nil
}

// PageQuery 分页查询销售信息
func (r *SaleInfoRepository) PageQuery(query *model.SaleInfoPageQuery) ([]*model.SaleInfoVO, int64, error) {
	// 构建查询条件
	conditions := []string{}
	args := []interface{}{}

	if query.SaleInfoID > 0 {
		conditions = append(conditions, "si.si_id = ?")
		args = append(args, query.SaleInfoID)
	}

	if query.ProductName != "" {
		conditions = append(conditions, "pd.pd_name LIKE ?")
		args = append(args, "%"+query.ProductName+"%")
	}

	if query.SalePlace != "" {
		conditions = append(conditions, "sp.sp_address LIKE ?")
		args = append(args, "%"+query.SalePlace+"%")
	}

	if !query.SaleTime.IsZero() {
		startTime := query.SaleTime
		endTime := query.SaleTime.Add(24 * time.Hour)
		conditions = append(conditions, "si.sale_time BETWEEN ? AND ?")
		args = append(args, startTime, endTime)
	}

	// 构建条件子句
	whereClause := ""
	if len(conditions) > 0 {
		whereClause = " WHERE " + strings.Join(conditions, " AND ")
	}

	// 查询总记录数
	countQuery := fmt.Sprintf(`
        SELECT COUNT(*) 
        FROM sale_info si
        LEFT JOIN sale_place sp ON sp.sp_id = si.sale_place_id
        LEFT JOIN logistics log ON log.log_id = si.logistics_id
        LEFT JOIN product_info pi ON pi.pi_id = log.product_info_id
        LEFT JOIN product pd ON pd.pd_id = pi.product_id
        %s`, whereClause)

	var total int64
	err := r.DB.QueryRow(countQuery, args...).Scan(&total)
	if err != nil {
		log.Println("查询销售信息总数失败:", err)
		return nil, 0, err
	}

	// 查询当前页数据
	offset := (query.Page - 1) * query.Size
	dataQuery := fmt.Sprintf(`
        SELECT 
            si.si_id, si.logistics_id, si.sale_place_id, si.si_description, si.sale_time,
            pd.pd_name, sp.sp_address, sp.sp_administrator,
            log.start_location, log.destination
        FROM sale_info si
        LEFT JOIN sale_place sp ON sp.sp_id = si.sale_place_id
        LEFT JOIN logistics log ON log.log_id = si.logistics_id
        LEFT JOIN product_info pi ON pi.pi_id = log.product_info_id
        LEFT JOIN product pd ON pd.pd_id = pi.product_id
        %s
        LIMIT ? OFFSET ?`, whereClause)

	queryArgs := append(args, query.Size, offset)

	rows, err := r.DB.Query(dataQuery, queryArgs...)
	if err != nil {
		log.Println("分页查询销售信息失败:", err)
		return nil, 0, err
	}
	defer rows.Close()

	var saleInfos []*model.SaleInfoVO
	for rows.Next() {
		saleInfo := &model.SaleInfoVO{}
		err := rows.Scan(
			&saleInfo.ID, &saleInfo.LogisticsID, &saleInfo.SalePlaceID, &saleInfo.Description, &saleInfo.SaleTime,
			&saleInfo.ProductName, &saleInfo.SalePlace, &saleInfo.Administrator,
			&saleInfo.StartLocation, &saleInfo.Destination,
		)
		if err != nil {
			log.Println("读取销售信息数据失败:", err)
			return nil, 0, err
		}
		saleInfos = append(saleInfos, saleInfo)
	}

	return saleInfos, total, nil
}
