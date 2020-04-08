package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"os/user"
	"strings"
)

/*
与Chromium文件放在同级目录
D:.
├─.idea
└─Chromium
    └─User Data
        ├─Avatars
		...
└─application.go
└─chromeUserDataLink.exe
*/

func main() {
	u, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	// C:\Users\wang
	var homeDir string = u.HomeDir
	// C:\Users\wangh\AppData\Local\Chromium
	var fullpath = homeDir + "\\AppData\\Local\\Chromium"
	b, err := PathExists(fullpath)
	if err != nil {
		log.Fatal(err)
	}
	if b {
		// 如果存在, 删除旧文件夹
		fmt.Println("文件存在, 删除 \"" + fullpath + "\"")
		err = os.RemoveAll(fullpath)
		if err != nil {
			log.Fatal(err)
		}
	}

	// 获取此程序路径
	// D:\Project\goodogs\src\airforce1\src\chromeUserDataLink
	dir := getExecutePath3()
	dir = dir + "\\Chromium"

	// 获取盘符
	panfu := strings.Split(dir, ":")[0]
	aaa := fmt.Sprintf("/C mklink /%s %s %s", panfu, fullpath, dir)
	fmt.Println(fmt.Sprintf("创建软链: %s", aaa))
	cmd := exec.Command("cmd", aaa)
	if err = cmd.Run(); err != nil {
		fmt.Println(err)
	}
	var a string
	fmt.Printf("输入任意键退出: ")
	_, _ = fmt.Scanln(&a)
}

// 判断路径是否存在
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

// 获取程序执行路径
func getExecutePath3() string {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}
	return dir
}
