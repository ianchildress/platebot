package alert

type Alerter interface {
	Alert(msg string) error
}
