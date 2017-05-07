package main

import "fmt"
import "time"
import "math/rand"

type (
	// Subject 若干个观察者所观察的对象所需要实现的方法
	Subject interface {
		RegisterObserver(Observer)
		RemoveObserver(Observer)
		NotifyObserver()
	}

	// Observer 每一个观察者所需要实现的对象
	Observer interface {
		Update(temp, humidity, pressure float32)
	}

	// Displayer 展示数据的方法
	Displayer interface {
		Display()
	}
)

// WeatherData 被订阅的主题，也就是若干个观察者所观察的对象
type WeatherData struct {
	observers  []Observer
	Temprature float32
	Humidity   float32
	Pressure   float32
}

// RegisterObserver 注册一个观察者
func (wd *WeatherData) RegisterObserver(o Observer) {
	wd.observers = append(wd.observers, o)
}

// RemoveObserver 移除一个观察者
func (wd *WeatherData) RemoveObserver(o Observer) {
	for i, v := range wd.observers {
		if o == v {
			wd.observers = append(wd.observers[:i], wd.observers[i+1:]...)
		}
	}
}

// NotifyObserver 通知所有的观察者数据的变化
func (wd WeatherData) NotifyObserver() {
	for _, v := range wd.observers {
		v.Update(wd.Temprature, wd.Humidity, wd.Pressure)
	}
}

// MeasurementsChanged 当数据变化的时候需要通知观察者
func (wd WeatherData) MeasurementsChanged() {
	wd.NotifyObserver()
}

// SetMeasurements 设置新的数据
func (wd WeatherData) SetMeasurements(temp, humidity, pressure float32) {
	wd.Temprature = temp
	wd.Humidity = humidity
	wd.Pressure = pressure

	wd.MeasurementsChanged()
}

// CurrentCoditions 不同的展示面板类
type CurrentCoditions struct {
	weatherData Subject
	temprature  float32
	humidity    float32
}

// NewCurrentCoditions 新建一个布告板
func NewCurrentCoditions(weatherData Subject) *CurrentCoditions {
	cc := &CurrentCoditions{
		weatherData: weatherData,
	}

	weatherData.RegisterObserver(cc)

	return cc
}

// Update 实现观察者的接口，当被观察的对象有数据变化的时候
// 被观察者会主动调用该方法通知观察者
func (cc *CurrentCoditions) Update(temp, humidity, pressure float32) {
	cc.temprature = temp
	cc.humidity = humidity
}

// Display 模式数据展示
func (cc *CurrentCoditions) Display() {
	for {
		time.Sleep(2 * time.Second)
		fmt.Printf("Current Coditions: %.2f F degrees and %.2f humidity\n",
			cc.temprature,
			cc.humidity,
		)
	}
}

func main() {
	weatherData := &WeatherData{}
	ccDisplay := NewCurrentCoditions(weatherData)
	go ccDisplay.Display()

	for {
		time.Sleep(1 * time.Second)
		weatherData.SetMeasurements(rand.Float32()*100, rand.Float32()*100, rand.Float32()*100)
	}
}
