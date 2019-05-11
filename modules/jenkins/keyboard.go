package jenkins

import "github.com/gdamore/tcell"

func (widget *Widget) initializeKeyboardControls() {
	widget.SetKeyboardChar("/", widget.ShowHelp)
	widget.SetKeyboardChar("j", widget.Next)
	widget.SetKeyboardChar("k", widget.Prev)
	widget.SetKeyboardChar("o", widget.openJob)
	widget.SetKeyboardChar("r", widget.Refresh)

	widget.SetKeyboardKey(tcell.KeyDown, widget.Next)
	widget.SetKeyboardKey(tcell.KeyEnter, widget.openJob)
	widget.SetKeyboardKey(tcell.KeyEsc, widget.Unselect)
	widget.SetKeyboardKey(tcell.KeyUp, widget.Prev)
}