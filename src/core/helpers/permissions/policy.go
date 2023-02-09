package permissions

type Policy interface {
	Subject() string
	Object() string
	Action() string
}

type policy struct {
	subject string
	object  string
	action  string
}

func NewPolicy(sub, obj, act string) Policy {
	return &policy{sub, obj, act}
}

func (instance *policy) Subject() string {
	return instance.subject
}

func (instance *policy) Object() string {
	return instance.object
}

func (instance *policy) Action() string {
	return instance.action
}
