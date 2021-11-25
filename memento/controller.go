package memento

type Controller struct {
  Editor *Editor
  Histories []Memento
}

func (c *Controller) CmdSave() {
  m := c.Editor.Save()
  c.Histories = append(c.Histories, m)
}

func (c *Controller) CmdUndo() {
  l := len(c.Histories)
  if l > 0 {
    m := c.Histories[l-1]
    c.Histories = c.Histories[:l-1]
    c.Editor.Restore(m)
  }
}
//
// func (c *Controller) CmdRedo() {
//
// }
