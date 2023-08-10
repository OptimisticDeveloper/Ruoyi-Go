package utils

import "os/exec"

// 打开浏览器
func OpenWin(uri string) {
	exec.Command(`cmd`, `/c`, `start`, uri).Start()
}
func OpenMac(uri string) {
	exec.Command("open", uri).Run()
}
