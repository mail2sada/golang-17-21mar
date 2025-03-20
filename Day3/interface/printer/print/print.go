package print

type Printer interface {
	Preview()
	PageSetup()
	Print()
}
