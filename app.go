package main

import (
	"changeme/rs485"
	"context"
	"fmt"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

type H map[string]interface{}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

// 打开串口
func (a *App) OpenUsb() any {
	if rs485.UsbStatus {
		return H{"status": "success", "msg": "usb已经打开"}
	}
	_client := rs485.GetRs485Client()
	if _client == nil {
		return H{"status": "error", "msg": "usb打开失败"}
	}
	return H{"status": "success", "msg": "usb打开成功"}
}

// 关闭串口
func (a *App) CloseUsb() any {
	if !rs485.UsbStatus {
		return H{"status": "success", "msg": "usb已经关闭"}
	}
	rs485.CloseRs485Client()
	return H{"status": "success", "msg": "usb关闭成功"}
}

// 重连
func (a *App) ReConnect() any {
	err := rs485.ReConnect()
	if err != nil {
		return H{"status": "error", "msg": err.Error()}
	}
	return H{"status": "success", "msg": "usb重连成功"}
}

// 取得串口状态
func (a *App) GetUsbStatus() any {
	return H{"status": "success", "UsbStatus": rs485.UsbStatus}
}

// 东阀蝶阀设置地址
func (a *App) DfddfSetAddress(oldAddress, newAddress uint32) any {
	_client := rs485.GetRs485Client()
	if _client == nil {
		return H{"status": "error", "msg": "usb打开失败"}
	}
	err := _client.WriteSingleRegister(byte(oldAddress), 0x0000, uint16(newAddress))
	if err != nil {
		return H{"status": "error", "msg": err.Error()}
	}

	return H{"status": "success", "msg": "发送命令成功"}
}

// 扫描东阀蝶阀地址1-100
func (a *App) DfddfScanAddress() any {
	_client := rs485.GetRs485Client()
	if _client == nil {
		return H{"status": "error", "msg": "usb打开失败"}
	}

	addresses := make([]uint16, 0)

	var address uint16
	for address = 1; address <= 99; address++ {
		res, err := _client.ReadHoldingRegisters(byte(address), 0x0000, 1)
		fmt.Println(res, err)
		if err == nil {
			addresses = append(addresses, address)
		}
	}

	return H{"status": "success", "msg": "扫描完成", "addresses": addresses}
}

// 测试东阀蝶阀的一个地址
func (a *App) DfddfTestAddress(address uint16) any {
	rs485.Rs485Lock.Lock()
	defer rs485.Rs485Lock.Unlock()

	_client := rs485.GetRs485Client()
	if _client == nil {
		return H{"status": "error", "msg": "usb打开失败"}
	}

	res, err := _client.ReadHoldingRegisters(byte(address), 0x0000, 1)
	fmt.Println(res, err)
	if err != nil {
		return H{"status": "error", "msg": err.Error()}
	}

	return H{"status": "success", "msg": "测试成功"}
}
