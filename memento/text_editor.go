package memento

type Editor struct {
  Text string
}

func (e *Editor) Save() Memento {
  return Snapshot{Text: e.Text}
}

func (e *Editor) Restore(m Memento) {
  e.Text = m.GetText()
}