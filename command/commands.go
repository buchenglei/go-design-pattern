package command

import "fmt"

// Command 定义一个命令必须要实现的方法
type Command interface {
	Execute()
	Undo()
}

// NoCommand 实现一个默认的空命令
type NoCommand struct{}

// Execute 什么都不执行
func (NoCommand) Execute() {
	fmt.Println(">>> Execute nothing...")
}

func (NoCommand) Undo() {
	fmt.Println(">>> Undo nothing...")
}

// MacroCommand 一次性执行多个命令
type MacroCommand struct {
	commands []Command
}

// NewMacroCommand 添加命令
func NewMacroCommand(commands []Command) *MacroCommand {
	return &MacroCommand{
		commands: commands,
	}
}

// Execute 依次执行每一个命令
func (cmd *MacroCommand) Execute() {
	for _, cmd := range cmd.commands {
		cmd.Execute()
	}
}

// Undo 依次撤销每一个命令
func (cmd *MacroCommand) Undo() {
	for _, cmd := range cmd.commands {
		cmd.Undo()
	}
}

// LightOnCommand 开灯的命令
type LightOnCommand struct {
	light Light
}

// NewLightOnCommand 创建一个开灯命令
func NewLightOnCommand(light Light) *LightOnCommand {
	return &LightOnCommand{
		light: light,
	}
}

// Execute 执行开灯的命令
func (command LightOnCommand) Execute() {
	command.light.On()
}

func (command LightOnCommand) Undo() {
	command.light.Off()
}

// LightOffCommand 关灯的命令
type LightOffCommand struct {
	light Light
}

// NewLightOffCommand 创建一个关灯的命令
func NewLightOffCommand(light Light) *LightOffCommand {
	return &LightOffCommand{
		light: light,
	}
}

// Execute 执行关灯的命令
func (command LightOffCommand) Execute() {
	command.light.Off()
}

func (command LightOffCommand) Undo() {
	command.light.On()
}

// StereoOnWitchCDCommand 打开音响的命令
type StereoOnWitchCDCommand struct {
	stereo Stereo
}

// NewStereoOnWitchCDCommand 创建一个打开音响的命令
func NewStereoOnWitchCDCommand(stereo Stereo) *StereoOnWitchCDCommand {
	return &StereoOnWitchCDCommand{
		stereo: stereo,
	}
}

// Execute 执行关灯的命令
func (command StereoOnWitchCDCommand) Execute() {
	command.stereo.On()
	command.stereo.SetCD()
	command.stereo.SetVolume(11)
}

func (command StereoOnWitchCDCommand) Undo() {
	command.stereo.Off()
}

// StereoOffCommand 关闭音响的命令
type StereoOffCommand struct {
	stereo Stereo
}

// NewStereoOffCommand 创建一个关闭音响的命令
func NewStereoOffCommand(stereo Stereo) *StereoOffCommand {
	return &StereoOffCommand{
		stereo: stereo,
	}
}

// Execute 执行关灯的命令
func (command StereoOffCommand) Execute() {
	command.stereo.Off()
}

func (command StereoOffCommand) Undo() {
	command.stereo.On()
	command.stereo.SetCD()
	command.stereo.SetVolume(11)
}
