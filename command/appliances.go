package command

import "fmt"

// Light 厂家自己的灯的对象
type Light struct{}

// On 开灯的操作
func (Light) On() {
	fmt.Println(">>> Light on...")
}

// Off 关灯的操作
func (Light) Off() {
	fmt.Println(">>> Light Off...")
}

// Stereo 厂家自己实现的印象对象
type Stereo struct{}

// On 打开音响
func (Stereo) On() {
	fmt.Println(">>> Stereo On...")
}

// SetCD 设置cd
func (Stereo) SetCD() {
	fmt.Println(">>> Stereo SetCD...")
}

// SetVolume 设置音量
func (Stereo) SetVolume(v int) {
	fmt.Printf(">>> Stereo SetVolume to %d...\n", v)
}

// Off 关闭音响
func (Stereo) Off() {
	fmt.Println(">>> Stereo Off...")
}
