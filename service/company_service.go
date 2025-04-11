package service

import (
	"log"

	"agricultural_product_gin/dto"
	"agricultural_product_gin/model"
	"agricultural_product_gin/repository"
)

// CompanyService 公司服务
type CompanyService struct {
	CompanyRepo *repository.CompanyRepository
}

// NewCompanyService 创建公司服务
func NewCompanyService(companyRepo *repository.CompanyRepository) *CompanyService {
	return &CompanyService{CompanyRepo: companyRepo}
}

// CreateCompany 创建公司
func (s *CompanyService) CreateCompany(companyDTO *dto.CompanyDTO) *dto.Result {
	// 转换DTO为模型
	company := &model.Company{
		Name:          companyDTO.Name,
		Address:       companyDTO.Address,
		Administrator: companyDTO.Administrator,
		Phone:         companyDTO.Phone,
	}

	// 保存公司
	id, err := s.CompanyRepo.Save(company)
	if err != nil {
		log.Println("创建公司失败:", err)
		return errorResult(500, "创建公司失败")
	}

	// 返回创建成功的公司ID
	return successResult("添加成功", id)
}

// UpdateCompany 更新公司
func (s *CompanyService) UpdateCompany(companyDTO *dto.CompanyDTO) *dto.Result {
	// 检查公司是否存在
	existingCompany, err := s.CompanyRepo.GetByID(companyDTO.ID)
	if err != nil {
		log.Println("查询公司失败:", err)
		return errorResult(500, "系统错误")
	}

	if existingCompany == nil {
		return errorResult(404, "公司不存在")
	}

	// 转换DTO为模型
	company := &model.Company{
		ID:            companyDTO.ID,
		Name:          companyDTO.Name,
		Address:       companyDTO.Address,
		Administrator: companyDTO.Administrator,
		Phone:         companyDTO.Phone,
	}

	// 更新公司
	err = s.CompanyRepo.Update(company)
	if err != nil {
		log.Println("更新公司失败:", err)
		return errorResult(500, "更新失败")
	}

	return successResult("更新成功", nil)
}

// DeleteCompany 删除公司
func (s *CompanyService) DeleteCompany(id int) *dto.Result {
	// 检查公司是否存在
	existingCompany, err := s.CompanyRepo.GetByID(id)
	if err != nil {
		log.Println("查询公司失败:", err)
		return errorResult(500, "系统错误")
	}

	if existingCompany == nil {
		return errorResult(404, "公司不存在")
	}

	// 删除公司
	err = s.CompanyRepo.Delete(id)
	if err != nil {
		log.Println("删除公司失败:", err)
		return errorResult(500, "删除失败")
	}

	return successResult("删除成功", nil)
}

// GetCompanyByID 根据ID获取公司
func (s *CompanyService) GetCompanyByID(id int) *dto.Result {
	company, err := s.CompanyRepo.GetByID(id)
	if err != nil {
		log.Println("获取公司失败:", err)
		return errorResult(500, "系统错误")
	}

	if company == nil {
		return errorResult(404, "公司不存在")
	}

	// 直接返回数据，不显示弹窗
	return &dto.Result{
		Code: 200,
		Data: company,
	}
}

// GetAllCompanies 获取所有公司
func (s *CompanyService) GetAllCompanies() *dto.Result {
	companies, err := s.CompanyRepo.FindAll()
	if err != nil {
		log.Println("获取所有公司失败:", err)
		return errorResult(500, "系统错误")
	}

	// 直接返回数据，不显示弹窗
	return &dto.Result{
		Code: 200,
		Data: companies,
	}
}

// PageQueryCompanies 分页查询公司
func (s *CompanyService) PageQueryCompanies(queryDTO *dto.CompanyPageQueryDTO) *dto.Result {
	// 参数校验
	if queryDTO.Page <= 0 {
		queryDTO.Page = 1
	}
	if queryDTO.PageSize <= 0 {
		queryDTO.PageSize = 10
	}

	// 分页查询
	companies, total, err := s.CompanyRepo.PageQuery(
		queryDTO.Page,
		queryDTO.PageSize,
		queryDTO.Name,
		queryDTO.Address,
		queryDTO.Administrator,
		queryDTO.Phone,
	)
	if err != nil {
		log.Println("分页查询公司失败:", err)
		return errorResult(500, "系统错误")
	}

	// 封装分页结果
	pageResult := dto.NewPageResult(total, companies, queryDTO.Page, queryDTO.PageSize)

	// 直接返回数据，不显示弹窗
	return &dto.Result{
		Code: 200,
		Data: pageResult,
	}
}
