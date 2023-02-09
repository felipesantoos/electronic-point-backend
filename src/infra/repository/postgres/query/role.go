package query

type AccountRoleQueryBuilder interface {
	Select() AccountRoleQuerySelectBuilder
}

type accountRoleQueryBuilder struct{}

func AccountRole() AccountRoleQueryBuilder {
	return &accountRoleQueryBuilder{}
}

type AccountRoleQuerySelectBuilder interface {
	All() string
}

type accountRoleQuerySelectBuilder struct{}

func (*accountRoleQueryBuilder) Select() AccountRoleQuerySelectBuilder {
	return &accountRoleQuerySelectBuilder{}
}

func (*accountRoleQuerySelectBuilder) All() string {
	return "SELECT * FROM account_role;"
}

type ActivityTypeQueryBuilder interface {
	Select() ActivityTypeQuerySelectBuilder
}

type activityTypeQueryBuilder struct{}

func ActivityType() ActivityTypeQueryBuilder {
	return &activityTypeQueryBuilder{}
}

type ActivityTypeQuerySelectBuilder interface {
	All() string
}

type activityTypeQuerySelectBuilder struct{}

func (*activityTypeQueryBuilder) Select() ActivityTypeQuerySelectBuilder {
	return &activityTypeQuerySelectBuilder{}
}

func (*activityTypeQuerySelectBuilder) All() string {
	return "SELECT * FROM activity_type;"
}
