// Code generated by goctl. DO NOT EDIT.
package types

type PathID struct {
	ID int `uri:"id"`
}

type ParamID struct {
	ID int `form:"id"`
}

type IDAndName struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type NameAndValue struct {
	Name  string      `json:"name"`
	Value interface{} `json:"value"`
}

type CommonProjectListResponse struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Desc string `json:"desc"`
}

type CommonProjectParamsResponse struct {
	Name    string   `json:"name"`
	Value   string   `json:"value"`
	Options []string `json:"options"`
}

type CodeResponse struct {
	ID    string `json:"id"`    // 在登录时提交
	Image string `json:"image"` // 验证码 base64 格式
}

type LoginRequest struct {
	Username string `json:"username" binding:"required" label:"用户名"`      // 用户名
	Password string `json:"password" binding:"required,min=6" label:"密码"` // 密码
	Code     string `json:"code" binding:"required,len=4" label:"验证码"`    // 验证码
	CodeID   string `json:"codeId" binding:"required" label:"获取验证码时的ID"`  // 获取验证码时的ID
}

type LoginResponse struct {
	Jwt string `json:"jwt"` // jwt 凭证
}

type UserListRequest struct {
	Page     int `form:"page" label:"分页"`       // 分页
	PageSize int `form:"pageSize" label:"每页条数"` // 每页条数
}

type UserListResponse struct {
	Total int        `json:"total"` // 总条数
	Data  []UserList `json:"data"`  // 具体数据
}

type UserList struct {
	ID       int    `json:"id"`
	Username string `json:"username"` // 用户名
	Realname string `json:"realname"` // 真实姓名
	Rolename string `json:"rolename"` // 角色名
	Status   int    `json:"status"`   // 状态 0:未启用 1:正常
	Phone    string `json:"phone"`    // 手机号
}

type UserDetailResponse struct {
	ID       int    `json:"id"`
	Username string `json:"username"` // 用户名
	Realname string `json:"realname"` // 真实姓名
	Phone    string `json:"phone"`    // 手机号
	RoleID   int    `json:"roleId"`   // 角色ID
	Status   int    `json:"status"`   // 状态 0:未启用 1:正常
}

type UserAddRequest struct {
	Username string `json:"username" binding:"required" label:"用户名"`              // 用户名
	Realname string `json:"realname" binding:"required" label:"真实姓名"`             // 真实姓名
	Phone    string `json:"phone" binding:"required" label:"手机号"`                 // 手机号
	RoleID   int    `json:"roleId" binding:"required" label:"角色ID"`               // 角色ID
	Password string `json:"password" binding:"required,min=6,max=255" label:"密码"` // 密码
	Status   int    `json:"status" binding:"oneof=0 1" label:"状态 0:未启用 1:正常"`     // 状态 0:未启用 1:正常
}

type UserEditRequest struct {
	PathID
	Username string `json:"username" binding:"required" label:"用户名"`          // 用户名
	Realname string `json:"realname" binding:"required" label:"真实姓名"`         // 真实姓名
	Phone    string `json:"phone" binding:"required" label:"手机号"`             // 手机号
	RoleID   int    `json:"roleId" binding:"required" label:"角色ID"`           // 角色ID
	Password string `json:"password" label:"密码"`                              // 密码
	Status   int    `json:"status" binding:"oneof=0 1" label:"状态 0:未启用 1:正常"` // 状态 0:未启用 1:正常
}

type RoleListRequest struct {
	Page     int `form:"page" label:"分页"`       // 分页
	PageSize int `form:"pageSize" label:"每页条数"` // 每页条数
}

type RoleListResponse struct {
	Total int        `json:"total"` // 总条数
	Data  []RoleList `json:"data"`  // 具体数据
}

type RoleList struct {
	ID   int    `json:"id"`
	Name string `json:"name"` // 角色名
}

type RoleDetailResponse struct {
	ID   int    `json:"id"`
	Name string `json:"name"` // 角色名
	Auth []int  `json:"auth"` // 权限ID
}

type RoleAddRequest struct {
	Name string `json:"name" binding:"required" label:"角色名"`        // 角色名
	Auth []int  `json:"auth" binding:"required,min=1" label:"权限ID"` // 权限ID
}

type RoleEditRequest struct {
	PathID
	RoleAddRequest
}

type AuthListResponse struct {
	ID       int                `json:"id"`
	Pid      int                `json:"pid"`
	Key      string             `json:"key"`      // 权限标识
	Name     string             `json:"name"`     // 节点名
	IsMenu   int                `json:"isMenu"`   // 是否是菜单栏 0：否 1：是
	API      string             `json:"api"`      // 接口
	Action   string             `json:"action"`   // 操作方法
	Children []AuthListResponse `json:"children"` // 子节点
}

type AuthAddRequest struct {
	Pid    int    `json:"pid" binding:"min=0"`
	Name   string `json:"name" binding:"required" label:"节点名"`               // 节点名
	Key    string `json:"key" binding:"required" label:"权限标识"`               // 权限标识
	IsMenu int    `json:"isMenu" binding:"oneof=0 1" label:"是否是菜单栏 0：否 1：是"` // 是否是菜单栏 0：否 1：是
	API    string `json:"api" label:"接口"`                                    // 接口
	Action string `json:"action" label:"操作方法"`                               // 操作方法
}

type AuthEditRequest struct {
	PathID
	AuthAddRequest
}

type BaseAuthResponse struct {
	ID       int                `json:"id"`
	Pid      int                `json:"pid"`
	Name     string             `json:"name"`     // 节点名
	Children []BaseAuthResponse `json:"children"` // 子节点
}

type UserInfoResponse struct {
	ID       int    `json:"id"`       // 用户ID
	Username string `json:"username"` // 用户名
	Realname string `json:"realname"` // 真实姓名
	Phone    string `json:"phone"`    // 手机号
	Rolename string `json:"rolename"` // 角色名
	Authkeys string `json:"authkeys"` // 角色权限KEY
}

type ChangeSelfPwdRequest struct {
	OldPassword     string `json:"oldPassword" binding:"required,min=6,max=255" label:"老密码"`                     // 老密码
	NewPassword     string `json:"newPassword" binding:"required,nefield=OldPassword,min=6,max=255" label:"新密码"` // 新密码
	ConfirmPassword string `json:"confirmPassword" binding:"required,eqfield=NewPassword" label:"确认密码"`          // 确认密码
}

type NamespaceListResponse struct {
	Name string `json:"name"`
}

type NamespaceAddRequest struct {
	Name string `json:"name" binding:"required"`
}

type NamespaceDelRequest struct {
	Name string `json:"name" binding:"required"`
}

type ProjectListRequest struct {
	Page     int    `form:"page"`
	PageSize int    `form:"pageSize"`
	Name     string `form:"name"`
}

type ProjectListResponse struct {
	Total int           `json:"total"`
	Data  []ProjectList `json:"data"`
}

type ProjectList struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	UserName string `json:"userName"`
	Desc     string `json:"desc"`
	Git      string `json:"git"`
	Token    string `json:"token"`
	UseTag   int    `json:"useTag"`
}

type ProjectAddRequest struct {
	Name     string `json:"name"`
	Desc     string `json:"desc"`
	Git      string `json:"git"`
	UserName string `json:"userName"`
	Token    string `json:"token"`
	UseTag   int    `json:"useTag"`
}

type ProjectEditRequest struct {
	PathID
	ProjectAddRequest
}

type TemplateListRequest struct {
	Page     int    `form:"page"`
	PageSize int    `form:"pageSize"`
	Name     string `form:"name"`
}

type TemplateListResponse struct {
	Total int            `json:"total"`
	Data  []TemplateList `json:"data"`
}

type TemplateList struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

type TemplateAddRequest struct {
	Name    string `json:"name"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

type TemplateEditRequest struct {
	PathID
	TemplateAddRequest
}

type DeployListRequest struct {
	Page     int    `form:"page"`
	PageSize int    `form:"pageSize"`
	Name     string `form:"name"`
}

type DeployListResponse struct {
	Total int          `json:"total"`
	Data  []DeployList `json:"data"`
}

type DeployList struct {
	ID           int    `json:"id"`
	DeployName   string `json:"deployName"`
	ProjectName  string `json:"projectName"`
	TemplateName string `json:"templateName"`
	Status       int    `json:"status"`
	UpdateTime   string `json:"updateTime"`
}

type DeployAddRequest struct {
	Name       string         `json:"name"`
	ProjectID  int            `json:"projectId"`
	TemplateID int            `json:"templateId"`
	Params     []NameAndValue `json:"params"`
}
