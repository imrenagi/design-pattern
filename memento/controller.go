package memento

import "fmt"

type Controller struct {
  Editor *Editor
  Histories []Memento

  // ptr is pointer to the current snapshot being used
  ptr int
}

func (c *Controller) CmdSave() {
  m := c.Editor.Save()
  c.Histories = append(c.Histories, m)

  l := len(c.Histories)
  c.ptr = l-1

  fmt.Println("current pointer ", c.ptr)
}

func (c *Controller) CmdUndo() {
    c.ptr -= 1
    m := c.Histories[c.ptr]
    c.Editor.Restore(m)
    fmt.Println("undo current pointer ", c.ptr)
}

func (c *Controller) CmdRedo() {
  l := len(c.Histories)
  if (c.ptr + 1) >= l {
    return
  }
  c.ptr++
  m := c.Histories[c.ptr]
  c.Editor.Restore(m)

  fmt.Println("redo current pointer ", c.ptr)
}
