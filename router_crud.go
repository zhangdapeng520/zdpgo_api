package zdpgo_gin

import (
	"github.com/zhangdapeng520/zdpgo_gin/libs/gin"
	"github.com/zhangdapeng520/zdpgo_code"
)

// RouterCrud 增删改查路由
type RouterCrud interface {
	Add(c *gin.Context)         // 增加数据
	AddMany(c *gin.Context)     // 批量增加数据
	DeleteById(c *gin.Context)  // 根据ID删除数据
	DeleteByIds(c *gin.Context) // 根据ID列表删除数据
	UpdateById(c *gin.Context)  // 修改数据
	UpdateByIds(c *gin.Context) // 根据ID列表修改数据
	FindById(c *gin.Context)    // 根据ID查询数据
	FindByIds(c *gin.Context)   // 根据ID列表查询数据
	FindByPage(c *gin.Context)  // 根据分页查询数据
}

// RegisterCrudRouter 注册CRUD路由
func (g *Gin) RegisterCrudRouter(router *gin.RouterGroup, table string, tableObjects interface{}) {
	// table相关路由
	tableRouter := router.Group(table)
	{
		// 增加数据
		tableRouter.POST("", g.Add(table))

		// 批量增加数据
		tableRouter.POST("many", g.AddMany(table))

		// 根据ID删除数据
		tableRouter.DELETE("", g.DeleteById(table))

		// 根据ID列表删除数据
		tableRouter.DELETE("ids", g.DeleteByIds(table))

		// 根据ID更新数据
		tableRouter.PATCH("", g.UpdateById(table))

		// 根据ID列表更新数据
		tableRouter.PATCH("ids", g.UpdateByIds(table))

		// 根据ID查询数据
		tableRouter.PUT("", g.FindById(table, tableObjects))

		// 根据ID列表查询数据
		tableRouter.PUT("ids", g.FindByIds(table, tableObjects))

		// 根据分页查询数据
		tableRouter.PUT("page", g.FindByPage(table, tableObjects))
	}
}

// AddForm 新增表单
type AddForm struct {
	Columns []string      `form:"columns" json:"columns" binding:"required"` // 字段列表
	Values  []interface{} `form:"values" json:"values" binding:"required"`   // 值列表
}

// AddManyForm 批量增加数据的表单
type AddManyForm struct {
	Columns []string        `form:"columns" json:"columns" binding:"required"` // 字段列表
	Values  [][]interface{} `form:"values" json:"values" binding:"required"`   // 值列表
}

// DeleteByIdForm 根据ID删除数据的表单
type DeleteByIdForm struct {
	ID int64 `form:"id" json:"id" binding:"required"` // id
}

// DeleteByIdsForm 根据ID列表删除数据的表单
type DeleteByIdsForm struct {
	IDS []int64 `form:"ids" json:"ids" binding:"required"` // id列表
}

// UpdateByIdForm 根据ID更新数据的表单
type UpdateByIdForm struct {
	ID      int64         `form:"id" json:"id" binding:"required"`           // id
	Columns []string      `form:"columns" json:"columns" binding:"required"` // 字段列表
	Values  []interface{} `form:"values" json:"values" binding:"required"`   // 值列表
}

// UpdateByIdsForm 根据ID列表修改数据
type UpdateByIdsForm struct {
	IDS     []int64       `form:"ids" json:"ids" binding:"required"`         // id
	Columns []string      `form:"columns" json:"columns" binding:"required"` // 字段列表
	Values  []interface{} `form:"values" json:"values" binding:"required"`   // 值列表
}

// FindByIdForm 根据ID查询数据的表单
type FindByIdForm struct {
	ID      int64    `form:"id" json:"id" binding:"required"`   // id
	Columns []string `form:"columns" json:"columns" binding:""` // 字段列表
}

// FindByIdsForm 根据ID列表查询数据的表单
type FindByIdsForm struct {
	IDS     []int64  `form:"ids" json:"ids" binding:"required"`         // id列表
	Columns []string `form:"columns" json:"columns" binding:"required"` // 字段列表
}

// FindByPageForm 根据分页查询数据的表单
type FindByPageForm struct {
	Columns []string `form:"columns" json:"columns" binding:"required"` // 字段列表
	Page    int      `form:"page" json:"page" binding:"required"`       // 第几页
	Size    int      `form:"size" json:"size" binding:"required"`       // 每页数量
}

// Add 增加数据
func (g *Gin) Add(table string) func(c *gin.Context) {
	return func(c *gin.Context) {
		// 获取参数并校验
		f := AddForm{}
		if err := c.ShouldBind(&f); err != nil {
			g.HandleValidatorError(c, err)
			return
		}

		// 执行添加
		addId, err := g.mysql.Add(table, f.Columns, f.Values)

		// 创建并返回响应
		rsp := NewResponse()
		if err != nil {
			g.log.Error("添加数据失败", "error", err.Error())
			rsp.Code = zdpgo_code.CODE_PARAM_ERROR
			rsp.Message = "添加数据失败"
			g.Success(c, rsp)
			return
		}

		// 创建并返回数据响应
		data := gin.H{
			"id": addId,
		}
		rspData := NewResponseData(data)
		g.SuccessData(c, rspData)
	}
}

// AddMany 批量新增数据
func (g *Gin) AddMany(table string) func(c *gin.Context) {
	return func(c *gin.Context) {
		// 获取参数并校验
		f := AddManyForm{}
		if err := c.ShouldBind(&f); err != nil {
			g.HandleValidatorError(c, err)
			return
		}

		// 执行添加
		addId, err := g.mysql.AddMany(table, f.Columns, f.Values)

		// 创建并返回响应
		rsp := NewResponse()
		if err != nil {
			g.log.Error("批量添加数据失败", "error", err.Error())
			rsp.Code = zdpgo_code.CODE_PARAM_ERROR
			rsp.Message = "批量添加数据失败"
			g.Success(c, rsp)
			return
		}

		// 创建并返回数据响应
		data := gin.H{
			"id": addId,
		}
		rspData := NewResponseData(data)
		g.SuccessData(c, rspData)
	}
}

// DeleteById 根据ID删除数据
func (g *Gin) DeleteById(table string) func(c *gin.Context) {
	return func(c *gin.Context) {
		// 获取参数并校验
		f := DeleteByIdForm{}
		if err := c.ShouldBind(&f); err != nil {
			g.HandleValidatorError(c, err)
			return
		}

		// 执行删除
		deleted, err := g.mysql.DeleteById(table, f.ID)

		// 创建并返回响应
		rsp := NewResponse()
		if err != nil {
			g.log.Error("根据ID删除数据失败", "error", err.Error())
			rsp.Code = zdpgo_code.CODE_PARAM_ERROR
			rsp.Message = "根据ID删除数据失败"
			g.Success(c, rsp)
			return
		}

		// 创建并返回数据响应
		data := gin.H{
			"deleted": deleted,
		}
		rspData := NewResponseData(data)
		g.SuccessData(c, rspData)
	}
}

// DeleteByIds 根据ID列表删除数据
func (g *Gin) DeleteByIds(table string) func(c *gin.Context) {
	return func(c *gin.Context) {
		// 获取参数并校验
		f := DeleteByIdsForm{}
		if err := c.ShouldBind(&f); err != nil {
			g.HandleValidatorError(c, err)
			return
		}

		// 执行删除
		deleted, err := g.mysql.DeleteByIds(table, f.IDS...)

		// 创建并返回响应
		rsp := NewResponse()
		if err != nil {
			g.log.Error("根据ID删除数据失败", "error", err.Error())
			rsp.Code = zdpgo_code.CODE_PARAM_ERROR
			rsp.Message = "根据ID删除数据失败"
			g.Success(c, rsp)
			return
		}

		// 创建并返回数据响应
		data := gin.H{
			"deleted": deleted,
		}
		rspData := NewResponseData(data)
		g.SuccessData(c, rspData)
	}
}

// UpdateById 根据ID更新数据
func (g *Gin) UpdateById(table string) func(c *gin.Context) {
	return func(c *gin.Context) {
		// 获取参数并校验
		f := UpdateByIdForm{}
		if err := c.ShouldBind(&f); err != nil {
			g.HandleValidatorError(c, err)
			return
		}

		// 执行删除
		update, err := g.mysql.UpdateById(table, f.Columns, f.Values, f.ID)

		// 创建并返回响应
		rsp := NewResponse()
		if err != nil {
			g.log.Error("根据ID更新数据失败", "error", err.Error())
			rsp.Code = zdpgo_code.CODE_PARAM_ERROR
			rsp.Message = "根据ID更新数据失败"
			g.Success(c, rsp)
			return
		}

		// 创建并返回数据响应
		data := gin.H{
			"update": update,
		}
		rspData := NewResponseData(data)
		g.SuccessData(c, rspData)
	}
}

// UpdateByIds 根据ID列表修改数据
func (g *Gin) UpdateByIds(table string) func(c *gin.Context) {
	return func(c *gin.Context) {
		// 获取参数并校验
		f := UpdateByIdsForm{}
		if err := c.ShouldBind(&f); err != nil {
			g.HandleValidatorError(c, err)
			return
		}

		// 执行删除
		update, err := g.mysql.UpdateByIds(table, f.Columns, f.Values, f.IDS)

		// 创建并返回响应
		rsp := NewResponse()
		if err != nil {
			g.log.Error("根据ID列表更新数据失败", "error", err.Error())
			rsp.Code = zdpgo_code.CODE_PARAM_ERROR
			rsp.Message = "根据ID列表更新数据失败"
			g.Success(c, rsp)
			return
		}

		// 创建并返回数据响应
		data := gin.H{
			"update": update,
		}
		rspData := NewResponseData(data)
		g.SuccessData(c, rspData)
	}
}

// FindById 根据ID查询数据
func (g *Gin) FindById(table string, objects interface{}) func(c *gin.Context) {
	return func(c *gin.Context) {
		// 获取参数并校验
		f := FindByIdForm{}
		if err := c.ShouldBind(&f); err != nil {
			g.HandleValidatorError(c, err)
			return
		}

		// 执行查询
		err := g.mysql.FindByIdToStruct(table, f.Columns, f.ID, objects)

		// 创建并返回响应
		rsp := NewResponse()
		if err != nil {
			g.log.Error("g.mysql.FindByIdToStruct 根据ID查询数据失败", "error", err.Error())
			rsp.Code = zdpgo_code.CODE_PARAM_ERROR
			rsp.Message = "根据ID查询数据失败"
			g.Success(c, rsp)
			return
		}

		// 创建并返回数据响应
		rspData := NewResponseData(objects)
		g.SuccessData(c, rspData)
	}
}

// FindByIds 根据id列表查询数据
func (g *Gin) FindByIds(table string, objects interface{}) func(c *gin.Context) {
	return func(c *gin.Context) {
		// 获取参数并校验
		f := FindByIdsForm{}
		if err := c.ShouldBind(&f); err != nil {
			g.HandleValidatorError(c, err)
			return
		}

		// 执行查询
		err := g.mysql.FindByIdsToStruct(table, f.Columns, f.IDS, objects)

		// 创建并返回响应
		rsp := NewResponse()
		if err != nil {
			g.log.Error("g.mysql.FindByIdsToStruct 根据ID查询数据失败", "error", err.Error())
			rsp.Code = zdpgo_code.CODE_PARAM_ERROR
			rsp.Message = "根据ID查询数据失败"
			g.Success(c, rsp)
			return
		}

		// 创建并返回数据响应
		rspData := NewResponseData(objects)
		g.SuccessData(c, rspData)
	}
}

// FindByPage 分页查询数据
func (g *Gin) FindByPage(table string, objects interface{}) func(c *gin.Context) {
	return func(c *gin.Context) {
		// 获取参数并校验
		f := FindByPageForm{}
		if err := c.ShouldBind(&f); err != nil {
			g.HandleValidatorError(c, err)
			return
		}

		// 执行查询
		err := g.mysql.FindByPageToStruct(table, f.Columns, f.Page, f.Size, objects)

		// 创建并返回响应
		rsp := NewResponse()
		if err != nil {
			g.log.Error("g.mysql.FindByIdsToStruct 根据ID查询数据失败", "error", err.Error())
			rsp.Code = zdpgo_code.CODE_PARAM_ERROR
			rsp.Message = "根据ID查询数据失败"
			g.Success(c, rsp)
			return
		}

		// 创建并返回数据响应
		rspData := NewResponseData(objects)
		g.SuccessData(c, rspData)
	}
}
