package gitTool

import (
	"fmt"
	"github.com/xzf/cmdtool"
)

type GitCmd struct {
}

func (g GitCmd) Get() []cmdtool.CMD {
	return []cmdtool.CMD{
		gitPushCommand(),
	}
}

func gitPushCommand() cmdtool.CMD {
	return cmdtool.CMD{
		Name: "GitPush",
		Logic: func(args ...string) {
			fmt.Println("len(args)", ":", len(args))
			fmt.Println("args", ":", args)

			//			if len(args) < 1 {
			//				fmt.Println(`e.g. :
			//dt GitPush  (commit message will be "no commit message")
			//or
			//dt GitPush {commit message part} {commit message part} {commit message part}...`)
			//			}
			//			if len(args) == 1 {
			//				cmd := `git add --all && git commit -m "%s" && git push`
			//				output, err := tool.RunCmd("git", "add", "--all")
			//				if err != nil {
			//					fmt.Println(runtime.Caller(0))
			//					fmt.Println("error:", err)
			//					fmt.Println("output:", output)
			//					return
			//				}
			//				tool.RunCmd("git", "commit", "--allow-empty-message", "-m", `""`)
			//				return
			//			}

		},
	}
}
