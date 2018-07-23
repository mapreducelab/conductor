package worker

// Worker listens and reacts to events
type Worker interface {
	Run() error
}

type worker struct{}
