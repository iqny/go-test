package abstractFactory

type IScore interface {
	Insert() bool
	List() map[string]string
}
