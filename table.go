package cmltest

func (*ModelRole) TableName() string {
	return "quan_role"
}

func (*ModelUserRole) TableName() string {
	return "quan_user_role"
}

func (*ModelRoleResource) TableName() string {
	return "quan_role_resource"
}

func (*ModelPermission) TableName() string {
	return "quan_permission"
}
