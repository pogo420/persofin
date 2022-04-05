package command

type BaseCommand interface {
	Description() string
	Execute() int
}
