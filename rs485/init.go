package rs485

import (
	"errors"
	"fmt"
	"log"
	"os"
	"runtime"
	"sync"
	"time"

	"github.com/tarm/serial"
	"github.com/wgbbiao/modbus"
)

var client *modbus.Client
var rs485sync sync.Once

// USB连接状态
var UsbStatus = false

// 全局锁
var Rs485Lock sync.Mutex

// 连接失败次数
var connectErrCnt int = 0

func openUsb() {
	config := serial.Config{
		Name: "COM3",
		// Name:  "/dev/ttyUSB0",
		Baud:        9600,
		StopBits:    1,
		Parity:      serial.ParityNone,
		ReadTimeout: time.Second * 2,
	}
	USBs := []string{}
	switch runtime.GOOS {
	case "windows":
		// config.Name = "COM3"
		USBs = []string{"COM3", "COM4", "COM5", "COM6", "COM7", "COM8", "COM9"}
	case "linux":
		// config.Name = "/dev/ttyUSB0"
		USBs = []string{"/dev/ttyUSB0", "/dev/ttyUSB1", "/dev/ttyUSB2"}
	case "darwin":
		USBs = []string{"/dev/tty.usbserial-210", "/dev/tty.usbserial-1420"}
	}
	var err error
	UsbStatus = false
	for _, usb := range USBs {
		// 是否存在usb
		x, _ := os.Stat(usb)
		if x == nil {
			fmt.Println("usb not exist", usb)
			continue
		}
		config.Name = usb
		client, err = modbus.NewClient(&config)
		if err != nil {
			log.Println("open usb error", err)
			// 记录到日志
			UsbStatus = false
			continue
		}
		UsbStatus = true
		connectErrCnt = 0
		break
	}
	if !UsbStatus {
		return
	}
	// client.DelayRtsBeforeSend = 100 * time.Millisecond
	client.EnableLog()

	// // 步进电机控制器读取间隔比较长
	// client.DelayReadTimes[5] = 2
	// client.DelayReadTimes[3] = 8
}

func GetRs485Client() *modbus.Client {
	rs485sync.Do(func() {
		openUsb()
	})
	return client
}

func CloseRs485Client() {
	if client != nil {
		client.Close()
	}
}

func ReConnect() error {
	// 每1分钟重启一次
	client.Close()
	openUsb()
	if UsbStatus {
		return nil
	} else {
		return errors.New("USB连接失败")
	}
}

// 增加错误次数 如果次数超过5，刚重连
func AddConnectErrCnt() {
	connectErrCnt++
	if connectErrCnt > 5 {
		ReConnect()
	}
}

func ResetConnectErrCnt() {
	connectErrCnt = 0
}
