package disclaimers

import (
	"fmt"
	"gopkg.in/ini.v1"
	"os"
	"time"
)

// 免责声明
const license = `
免责声明

1、所分享课程和资料目的只用于教学，切勿使用课程中的技术进行违法活动，学员利用课程中的技术进行违法活动，造成的后果与讲师本人及讲师所属机构无关。
2、所有视频、资料、工具仅面向拥有合法授权的渗透测试安全人员及进行常规操作的网络运维人员，用户可在取得足够合法授权且非商用的前提下进行下载、复制、传播或使用。
3、在使用本工具的过程中，您应确保自己的所有行为符合当地法律法规，且不得将此软件用于违反中国人民共和国相关法律的活动。本工具所有作者、贡献者、分享者，不承担用户擅自使用本工具从事任何违法活动所产生的任何责任。
4、讲师和讲师所在机构，对于软件的适用性、可靠性、安全性等均不作任何明示或默示的保证。因使用本软件所产生的直接或间接损失，如中毒，有后门等。讲师和讲师所在公司概不负责。
5、讲师和讲师所有在公司，发布的课程，资料，工具仅为个人学习测试使用，请在下载后24小时内删除，不得用于任何商业用途，否则后果自负。 

我们倡导维护网络安全，人人有责，共同维护网络文明和谐。
同意[Y/y/yes]继续或不同意[N/n/no]退出。
`

func License() {
	licensePath := "./.agree.ini"
	_, err := os.Stat(licensePath)
	if err == nil {
		cfg, _ := ini.Load("./.agree.ini")
		agreement, _ := cfg.Section("Disclaimer clause").Key("agree").Int()
		if agreement == 1 {
			return
		}
	} else {
		// 打印许可声明
		fmt.Println(license)
		// 接受用户输入
		var input string
		fmt.Print("您同意上述条款吗? [Y/N]: ")
		fmt.Scanln(&input)

		// 检查用户输入
		if input != "Y" && input != "y" && input != "yes" {
			fmt.Println("您已选择不同意免责条款。程序将在 3 秒后退出 ...")
			exitWithDelay(3)
		} else {
			cfg := ini.Empty()
			agreeSection, err := cfg.NewSection("Disclaimer clause")
			if err != nil {
				fmt.Println(err)
				os.Exit(0)
			}
			agreeSection.NewKey("agree", "1")
			err = cfg.SaveTo(".agree.ini")
			if err != nil {
				fmt.Println(err)
				os.Exit(0)
			}
		}
	}
}

func exitWithDelay(seconds int) {
	for seconds > 0 {
		fmt.Printf("%d...", seconds)
		time.Sleep(1 * time.Second)
		seconds--
	}
	os.Exit(0)
}
