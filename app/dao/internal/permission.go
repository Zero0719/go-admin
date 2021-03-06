// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/frame/gmvc"
)

// PermissionDao is the manager for logic model data accessing and custom defined data operations functions management.
type PermissionDao struct {
	gmvc.M                                      // M is the core and embedded struct that inherits all chaining operations from gdb.Model.
	C       permissionColumns // C is the short type for Columns, which contains all the column names of Table for convenient usage.	
	DB      gdb.DB                              // DB is the raw underlying database management object.
	Table   string                              // Table is the underlying table name of the DAO.
}

// PermissionColumns defines and stores column names for table permission.
type permissionColumns struct {
	Id        string //           
    Name      string // 权限名称  
    Uri       string // 权限地址  
    CreateAt  string // 添加时间  
    UpdateAt  string // 编辑时间  
    DeleteAt  string // 删除时间
}

// NewPermissionDao creates and returns a new DAO object for table data access.
func NewPermissionDao() *PermissionDao {
    columns := permissionColumns{
		Id:       "id",         
            Name:     "name",       
            Uri:      "uri",        
            CreateAt: "create_at",  
            UpdateAt: "update_at",  
            DeleteAt: "delete_at",
	}
	return &PermissionDao{
		C: 	   columns,
		M:     g.DB("default").Model("permission").Safe(),
		DB:    g.DB("default"),
		Table: "permission",
	}
}