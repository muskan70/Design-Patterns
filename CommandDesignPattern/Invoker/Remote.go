package invoker

import command "comm/Command"

type RemoteControl struct {
	CurCommand command.ICommand
	commands   []command.ICommand
}

func (r *RemoteControl) SetCommand(c command.ICommand) {
	r.CurCommand = c
}

func (r *RemoteControl) PressButton() {
	r.CurCommand.Execute()
	r.commands = append(r.commands, r.CurCommand)
}

func (r *RemoteControl) Undo() {
	r.commands = r.commands[:len(r.commands)-1]
	r.CurCommand = r.commands[len(r.commands)-1]
	r.CurCommand.Execute()
}
