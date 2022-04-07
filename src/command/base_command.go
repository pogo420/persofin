package command

type BaseCommand interface {
	Description() string
	Execute(string) int
}
