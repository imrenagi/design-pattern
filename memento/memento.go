package memento

type Memento interface {
  GetText() string
}

type Snapshot struct {
  Text string
}

func (s Snapshot) GetText() string {
  return s.Text
}
