package memento

import "fmt"

func Execute() {

  editor := &Editor{}
  editor.Text = "text pertama"

  controller := Controller{
    Editor:    editor,
  }

  controller.CmdSave()

  fmt.Println(editor.Text)

  editor.Text = "text pertama udah diupdate"

  controller.CmdSave()

  fmt.Println(editor.Text)

  controller.CmdUndo()

  fmt.Println(editor.Text)

  controller.CmdUndo()

  fmt.Println(editor.Text)




}
