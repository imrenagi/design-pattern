package command

type Waiter struct {
  Commands []Interface
}

func (w *Waiter) AddOrder(c Interface) {
  w.Commands = append(w.Commands, c)
}

func (w Waiter) Finalize() {
  for _, c := range w.Commands {
    c.Execute()
  }
}