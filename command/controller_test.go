package command

import (
	"fmt"
	"testing"
)

func TestRemoteControl(t *testing.T) {
	remoteControl := NewRemoteControl()
	remoteControl.OnButtonWasPressed(0)
	remoteControl.OffButtonWasPressed(0)
	remoteControl.UndoButtonWasPressed()

	fmt.Println("-------------")

	// 定义命令
	lightOnCmd := NewLightOnCommand(Light{})
	lightOffCmd := NewLightOffCommand(Light{})
	stereoOnCmd := NewStereoOnWitchCDCommand(Stereo{})
	stereoOffCmd := NewStereoOffCommand(Stereo{})

	remoteControl.SetCommand(0, lightOnCmd, lightOffCmd)
	remoteControl.OnButtonWasPressed(0)
	remoteControl.OffButtonWasPressed(0)
	remoteControl.UndoButtonWasPressed()

	fmt.Println("-------------")

	remoteControl.SetCommand(1, stereoOnCmd, stereoOffCmd)
	remoteControl.OnButtonWasPressed(1)
	remoteControl.OffButtonWasPressed(1)
	remoteControl.UndoButtonWasPressed()

	fmt.Println("-------------")

	// 测试宏命令
	cmds := []Command{
		lightOnCmd,
		stereoOnCmd,
		// .....
	}
	macroCmd := NewMacroCommand(cmds)
	macroCmd.Execute()
	macroCmd.Undo()

}
