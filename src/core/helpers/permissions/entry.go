package permissions

type Entry interface {
	Subject() string
	Objects() []string
}

type entry struct {
	subject string
	objects []string
}

func NewEntry(sub string, objects []string) Entry {
	return &entry{sub, objects}
}

func (instance *entry) Subject() string {
	return instance.subject
}

func (instance *entry) Objects() []string {
	return instance.objects
}
