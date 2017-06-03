package command

// CommandCount 定义命令的数量
const CommandCount = 7

// RemoteControl 遥控器对象
type RemoteControl struct {
	onCommands  []Command
	offCommands []Command
	undoCommand Command
}

// NewRemoteControl 初始化一个新的遥控器
func NewRemoteControl() *RemoteControl {
	remoteControl := &RemoteControl{}
	remoteControl.onCommands = make([]Command, CommandCount)
	remoteControl.offCommands = make([]Command, CommandCount)

	for i := 0; i < CommandCount; i++ {
		remoteControl.onCommands[i] = NoCommand{}
		remoteControl.offCommands[i] = NoCommand{}
		remoteControl.undoCommand = NoCommand{}
	}

	return remoteControl
}

// SetCommand 为slot设置开关命名
func (control *RemoteControl) SetCommand(slot int, on, off Command) {
	control.onCommands[slot] = on
	control.offCommands[slot] = off
}

// OnButtonWasPressed 打开的按钮被按下
func (control *RemoteControl) OnButtonWasPressed(slot int) {
	control.onCommands[slot].Execute()
	control.undoCommand = control.onCommands[slot]
}

// OffButtonWasPressed 关闭的按钮被按下
func (control *RemoteControl) OffButtonWasPressed(slot int) {
	control.offCommands[slot].Execute()
	control.undoCommand = control.offCommands[slot]
}

func (control RemoteControl) UndoButtonWasPressed() {
	control.undoCommand.Undo()
}
