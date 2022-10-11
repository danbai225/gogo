package main

import (
	"fmt"
	"gogo/csgo"
	"gogo/resource"
	"gogo/tinder_rules"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	"github.com/goccy/go-json"
	"github.com/ncruces/zenity"
	"github.com/shirou/gopsutil/v3/process"
)

type Conf struct {
	E5Arena string `json:"e5_arena"` //5EArena安装路径
	R       string `json:"r"`        //r0安装路径
}

var confing *Conf
var mod, dllMod string

func main() {
	//配置相关
	_, err := os.Stat("conf.json")
	if err != nil {
		err := i()
		if err != nil {
			fmt.Println(err)
			return
		}
	}
	mod, err = zenity.ListItems("请选择模式", []string{"5EArena", "官匹", "R0", "重新设置程序路径"}...)
	if mod == "" {
		fmt.Println("请选择")
		return
	}
	if mod == "重新设置程序路径" {
		err := i()
		if err != nil {
			fmt.Println(err)
			return
		}
		return
	}
	err = load()
	if err != nil {
		fmt.Println(err)
		return
	}
	dllMod, _ = zenity.ListItems("功能模式", []string{"单透", "多功能", "多功能增强"}...)
	if mod == "" {
		fmt.Println("请选择")
		return
	}
	homeDir, _ := os.UserHomeDir()
	switch mod {
	case "5EArena":
		//防止5e检测 删除驱动
		_ = os.Remove(fmt.Sprintf(`%s%s`, confing.E5Arena, `\resources\node_modules\vibran\cache\vibran_drv.sys`))
		chDll := fmt.Sprintf("%s%s", confing.E5Arena, `\resources\node_modules\vibran\cache\cheano.dll`)
		file, err := os.ReadFile(chDll)
		if err == nil {
			//备份
			_ = os.WriteFile(strings.ReplaceAll(chDll, "cheano.dll", "cheano.dll.bak"), file, os.ModePerm)
			//移动5e dll文件
			_ = os.WriteFile(`cheano.dll`, file, os.ModePerm)
			err = os.Remove(chDll)
			if err != nil {
				fmt.Println(err)
				return
			}
		}
	}
	baseArr := []string{
		GetCurPath() + "\\*",
		fmt.Sprintf("%s%s", homeDir, `\Documents\*`),
		fmt.Sprintf("%s%s", homeDir, `\Downloads\*`),
		`C:\Windows\Prefetch\*`,
	}
	psArr := make([]string, 0)
	RArr := make([][]string, 0)
	//生成火绒配置文件
	if confing.E5Arena != "" {
		rArr := make([]string, 0)
		rArr = append(rArr, baseArr...)
		rArr = append(rArr, fmt.Sprintf("%s%s", confing.E5Arena, `\resources\node_modules\vibran\cache\vibran_drv.sys`))
		rArr = append(rArr, fmt.Sprintf("%s%s", confing.E5Arena, `\resources\node_modules\vibran\cache\cheano.dll`))
		RArr = append(RArr, rArr)
		psArr = append(psArr, fmt.Sprintf("%s%c%s", confing.E5Arena, os.PathSeparator, "5EArena.exe"))
	}
	if confing.R != "" {
		rArr := make([]string, 0)
		rArr = append(rArr, baseArr...)
		rArr = append(rArr, fmt.Sprintf("%s%s", homeDir, `\*`))
		RArr = append(RArr, rArr)
		psArr = append(psArr, fmt.Sprintf("%s%c%s", confing.R, os.PathSeparator, "jianglao_arena.exe"))
	}
	rules := tinder_rules.NewRules("GOGO防止扫盘", psArr, RArr)
	_ = os.WriteFile("火绒规则文件.json", rules, os.ModePerm)
	//配置写出
	_ = os.MkdirAll(fmt.Sprintf("%s%s", homeDir, `\Documents\Osiris`), os.ModePerm)
	_ = os.WriteFile(fmt.Sprintf("%s%s", homeDir, `\Documents\Osiris\config`), resource.Config, os.ModePerm)
	game()
}

// 初始化配置
func i() error {
	fmt.Println("请选择平台执行文件路径，没有的平台跳过即可")
	conf := Conf{}
	path, _ := zenity.SelectFile(zenity.FileFilters{{"5EArena执行文件", []string{"5EArena.exe"}}})
	conf.E5Arena = filepath.Dir(path)
	path, _ = zenity.SelectFile(zenity.FileFilters{{"R0执行文件", []string{"jianglao_arena.exe"}}})
	conf.R = filepath.Dir(path)
	marshal, _ := json.Marshal(conf)
	return os.WriteFile("conf.json", marshal, os.ModePerm)
}
func load() error {
	file, err := os.ReadFile("conf.json")
	if err != nil {
		return err
	}
	confing = &Conf{}
	return json.Unmarshal(file, confing)
}

/*GetCurPath 获取当前文件执行的路径*/
func GetCurPath() string {
	file, _ := exec.LookPath(os.Args[0])
	path, _ := filepath.Abs(file)
	rst := filepath.Dir(path)
	return rst
}
func game() {
	fmt.Println("欢迎使用大陀螺")
	for {
		//等待csgo运行注入
		processes, err := process.Processes()
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("正在查找csgo进程准备注入")
		var p *process.Process
		for p == nil {
			for _, p2 := range processes {
				if name, _ := p2.Name(); name == "csgo.exe" {
					p = p2
					break
				}
			}
			time.Sleep(5 * time.Second)
			processes, _ = process.Processes()
		}
		size := uint64(0)
		fmt.Println("pid:", p.Pid, "等待加载")
		for size < 629145600 {
			info, _ := p.MemoryInfo()
			size = info.RSS
		}
		//写出文件
		switch dllMod {
		case "单透":
			_ = os.WriteFile("danbai.dll", resource.GOESPDll, os.ModePerm)
		case "多功能":
			_ = os.WriteFile("danbai.dll", resource.OsirisDll, os.ModePerm)
		case "多功能增强":
			_ = os.WriteFile("danbai.dll", resource.NEPSDll, os.ModePerm)
		}

		fmt.Println("pid:", p.Pid, "开始注入")
		switch mod {
		case "5EArena":
			_ = os.WriteFile("danbai.exe", resource.ToolExe, os.ModePerm)
			fmt.Println(csgo.Inject(fmt.Sprintf("%s%s", GetCurPath(), `\danbai.dll`)))
			time.Sleep(time.Second)
			fmt.Println(csgo.Inject(fmt.Sprintf("%s%s", GetCurPath(), `\cheano.dll`)))
		case "官匹", "R0":
			_ = os.WriteFile("danbai2.exe", resource.ToolExe2, os.ModePerm)
			fmt.Println(csgo.Inject2(p.Pid, fmt.Sprintf("%s%s", GetCurPath(), `\danbai.dll`)))
		}
		fmt.Println("注入完成")
		_ = os.Remove("danbai.dll")
		_ = os.Remove("danbai.exe")
		_ = os.Remove("danbai2.exe")
		processes, err = process.Processes()
		for p != nil {
			flag := false
			for _, p2 := range processes {
				if p2.Pid == p.Pid {
					flag = true
					break
				}
			}
			if !flag {
				p = nil
			} else {
				processes, _ = process.Processes()
				time.Sleep(5 * time.Second)
			}
		}
		fmt.Println("检测到游戏结束,等待下一次注入")
	}
}
