package main

//import (
//	"fmt"
//	"os"
//	"sort"
//	"strings"
//)
//
//func main() {
//	args := os.Args
//	if len(args) == 1 {
//		fmt.Println("sub command:")
//		fmt.Println(strings.Join(gCmdNameSlice, "\n"))
//		return
//	}
//	cmdName := args[1]
//	smallCmdName := strings.ToLower(cmdName)
//	for _, item := range commandList {
//		cmdSlice := item.Get()
//		for _, cmd := range cmdSlice {
//			if strings.ToLower(cmd.Name) != smallCmdName {
//				continue
//			}
//			cmd.Logic(args...)
//		}
//	}
//}
//
//var gCmdNameSlice []string
//var gCmdMap = map[string]func(...string){}
//
//func init() {
//	for _, item := range commandList {
//		cmdSlice := item.Get()
//		for _, cmd := range cmdSlice {
//			smallSubCmdName := strings.ToLower(cmd.Name)
//			_, ok := gCmdMap[smallSubCmdName]
//			if ok {
//				panic("[fxztzbqwl6] have same command")
//			}
//			gCmdMap[smallSubCmdName] = cmd.Logic
//			gCmdNameSlice = append(gCmdNameSlice, cmd.Name)
//		}
//	}
//	sort.Strings(gCmdNameSlice)
//}
//
//var commandList = []cmdtool.CMDInterface{
//	gitTool.GitCmd{},
//}
