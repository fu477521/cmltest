package cmltest

func (*ModelRole) TableName() string {
	return "cmltest_role"
}

func (*ModelUserRole) TableName() string {
	return "cmltest_user_role"
}

func (*ModelPermission) TableName() string {
	return "cmltest_permission"
}

func (*ModelMenu) TableName() string {
	return "cmltest_menu"
}

func (*ModelRoleResource) TableName() string {
	return "cmltest_role_resource"
}
