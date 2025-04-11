package repository

import (
	"agricultural_product_gin/model"
	"database/sql"
	"fmt"
	"log"
	"strings"
	"time"
)

// LogisticsRepository 物流数据仓库
type LogisticsRepository struct {
	DB *sql.DB
}

// NewLogisticsRepository 创建物流仓库
func NewLogisticsRepository(db *sql.DB) *LogisticsRepository {
	return &LogisticsRepository{DB: db}
}

// Save 保存物流信息
func (r *LogisticsRepository) Save(logistics *model.Logistics) (int, error) {
	query := "INSERT INTO logistics(product_info_id, company_id, start_location, destination, start_time, end_time) VALUES(?, ?, ?, ?, ?, ?)"

	var endTimeValue interface{}
	if logistics.EndTime != nil {
		endTimeValue = logistics.EndTime
	} else {
		endTimeValue = nil
	}

	result, err := r.DB.Exec(query, logistics.ProductInfoID, logistics.CompanyID, logistics.StartLocation, logistics.Destination, logistics.StartTime, endTimeValue)
	if err != nil {
		log.Println("保存物流信息失败:", err)
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		log.Println("获取物流ID失败:", err)
		return 0, err
	}

	return int(id), nil
}

// Update 更新物流信息
func (r *LogisticsRepository) Update(logistics *model.Logistics) error {
	query := `UPDATE logistics 
			SET product_info_id = ?, company_id = ?, start_location = ?, 
			destination = ?, start_time = ?, end_time = ? 
			WHERE log_id = ?`

	var endTimeValue interface{}
	if logistics.EndTime != nil {
		endTimeValue = logistics.EndTime
	} else {
		endTimeValue = nil
	}

	_, err := r.DB.Exec(query, logistics.ProductInfoID, logistics.CompanyID, logistics.StartLocation, logistics.Destination, logistics.StartTime, endTimeValue, logistics.ID)
	if err != nil {
		log.Println("更新物流信息失败:", err)
		return err
	}
	return nil
}

// Delete 删除物流信息
func (r *LogisticsRepository) Delete(id int) error {
	query := "DELETE FROM logistics WHERE log_id = ?"
	_, err := r.DB.Exec(query, id)
	if err != nil {
		log.Println("删除物流信息失败:", err)
		return err
	}
	return nil
}

// GetByID 根据ID获取物流信息
func (r *LogisticsRepository) GetByID(id int) (*model.Logistics, error) {
	query := `SELECT l.log_id, l.product_info_id, l.company_id, l.start_location, l.destination, 
			l.start_time, l.end_time, p.pd_name, c.com_name, c.com_administrator, c.com_phone
			FROM logistics l
			LEFT JOIN product_info pi ON l.product_info_id = pi.pi_id
			LEFT JOIN product p ON pi.product_id = p.pd_id
			LEFT JOIN company c ON l.company_id = c.com_id
			WHERE l.log_id = ?`

	row := r.DB.QueryRow(query, id)

	logistics := &model.Logistics{}
	var endTime sql.NullTime

	err := row.Scan(
		&logistics.ID, &logistics.ProductInfoID, &logistics.CompanyID,
		&logistics.StartLocation, &logistics.Destination, &logistics.StartTime, &endTime,
		&logistics.ProductName, &logistics.CompanyName, &logistics.Administrator, &logistics.Phone,
	)

	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		log.Println("获取物流信息失败:", err)
		return nil, err
	}

	if endTime.Valid {
		logistics.EndTime = &endTime.Time
	} else {
		logistics.EndTime = nil
	}

	return logistics, nil
}

// FindAll 查找所有物流信息
func (r *LogisticsRepository) FindAll() ([]*model.Logistics, error) {
	query := `SELECT l.log_id, l.product_info_id, l.company_id, l.start_location, l.destination, 
			l.start_time, l.end_time, p.pd_name, c.com_name, c.com_administrator, c.com_phone
			FROM logistics l
			LEFT JOIN product_info pi ON l.product_info_id = pi.pi_id
			LEFT JOIN product p ON pi.product_id = p.pd_id
			LEFT JOIN company c ON l.company_id = c.com_id`

	rows, err := r.DB.Query(query)
	if err != nil {
		log.Println("查询物流信息失败:", err)
		return nil, err
	}
	defer rows.Close()

	var logisticsList []*model.Logistics
	for rows.Next() {
		logistics := &model.Logistics{}
		var endTime sql.NullTime

		err := rows.Scan(
			&logistics.ID, &logistics.ProductInfoID, &logistics.CompanyID,
			&logistics.StartLocation, &logistics.Destination, &logistics.StartTime, &endTime,
			&logistics.ProductName, &logistics.CompanyName, &logistics.Administrator, &logistics.Phone,
		)

		if err != nil {
			log.Println("读取物流数据失败:", err)
			return nil, err
		}

		if endTime.Valid {
			logistics.EndTime = &endTime.Time
		} else {
			logistics.EndTime = nil
		}

		logisticsList = append(logisticsList, logistics)
	}

	return logisticsList, nil
}

// PageQuery 分页查询物流信息
func (r *LogisticsRepository) PageQuery(dto *model.LogisticsPageQueryDTO) ([]*model.Logistics, int64, error) {
	// 构建查询条件
	conditions := []string{}
	args := []interface{}{}

	if dto.LogisticsId > 0 {
		conditions = append(conditions, "l.log_id = ?")
		args = append(args, dto.LogisticsId)
	}

	if dto.ProductName != "" {
		conditions = append(conditions, "p.pd_name LIKE ?")
		args = append(args, "%"+dto.ProductName+"%")
	}

	if dto.CompanyName != "" {
		conditions = append(conditions, "c.com_name LIKE ?")
		args = append(args, "%"+dto.CompanyName+"%")
	}

	if dto.StartLocation != "" {
		conditions = append(conditions, "l.start_location LIKE ?")
		args = append(args, "%"+dto.StartLocation+"%")
	}

	if dto.Destination != "" {
		conditions = append(conditions, "l.destination LIKE ?")
		args = append(args, "%"+dto.Destination+"%")
	}

	if dto.Administrator != "" {
		conditions = append(conditions, "c.com_administrator LIKE ?")
		args = append(args, "%"+dto.Administrator+"%")
	}

	if dto.StartTime != "" {
		// 将日期转换为时间段
		startDate, err := time.Parse("2006-01-02", dto.StartTime)
		if err == nil {
			beginTime := startDate.Format("2006-01-02 00:00:00")
			endTime := startDate.Format("2006-01-02 23:59:59")

			conditions = append(conditions, "l.start_time >= ? AND l.start_time <= ?")
			args = append(args, beginTime, endTime)
		} else {
			log.Println("日期格式转换失败:", err)
		}
	}

	// 构建条件子句
	whereClause := ""
	if len(conditions) > 0 {
		whereClause = " WHERE " + strings.Join(conditions, " AND ")
	}

	// 查询总记录数
	baseQuery := `FROM logistics l
	LEFT JOIN product_info pi ON l.product_info_id = pi.pi_id
	LEFT JOIN product p ON pi.product_id = p.pd_id
	LEFT JOIN company c ON l.company_id = c.com_id`

	countQuery := fmt.Sprintf("SELECT COUNT(*) %s%s", baseQuery, whereClause)

	var total int64
	err := r.DB.QueryRow(countQuery, args...).Scan(&total)
	if err != nil {
		log.Println("查询物流信息总数失败:", err)
		return nil, 0, err
	}

	// 查询当前页数据
	offset := (dto.Page - 1) * dto.Size
	dataQuery := fmt.Sprintf(`SELECT l.log_id, l.product_info_id, l.company_id, l.start_location, l.destination, 
		l.start_time, l.end_time, p.pd_name, c.com_name, c.com_administrator, c.com_phone
		%s%s 
		LIMIT ? OFFSET ?`, baseQuery, whereClause)

	queryArgs := append(args, dto.Size, offset)

	rows, err := r.DB.Query(dataQuery, queryArgs...)
	if err != nil {
		log.Println("分页查询物流信息失败:", err)
		return nil, 0, err
	}
	defer rows.Close()

	var logisticsList []*model.Logistics
	for rows.Next() {
		logistics := &model.Logistics{}
		var endTime sql.NullTime

		err := rows.Scan(
			&logistics.ID, &logistics.ProductInfoID, &logistics.CompanyID,
			&logistics.StartLocation, &logistics.Destination, &logistics.StartTime, &endTime,
			&logistics.ProductName, &logistics.CompanyName, &logistics.Administrator, &logistics.Phone,
		)

		if err != nil {
			log.Println("读取物流数据失败:", err)
			return nil, 0, err
		}

		if endTime.Valid {
			logistics.EndTime = &endTime.Time
		} else {
			logistics.EndTime = nil
		}

		logisticsList = append(logisticsList, logistics)
	}

	return logisticsList, total, nil
}
