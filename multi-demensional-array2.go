package main
import "fmt"

func main(){
	//创建二维数组
	sites := [2][2]string{}

	//向二维数组添加元素
	sites[0][0] = "google"
	sites[0][1] = "runoob"
	sites[1][0] = "taobao"
	sites[1][1] = "weibo"

	//显示结果
	fmt.Println(sites)
}