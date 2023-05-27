package main

import (
	"IM-Chat-Golang/controller"
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/user/login", controller.UserLogin)
	http.HandleFunc("/user/register", controller.UserRegister)

	//添加好友
	http.HandleFunc("/contact/addfriend", controller.Addfriend)
	//加载好友
	http.HandleFunc("/contact/loadfriend", controller.LoadFriend)

	//加载群聊
	http.HandleFunc("/contact/loadcommunity", controller.LoadCommunity)
	//加入群聊
	http.HandleFunc("/contact/joincommunity", controller.JoinCommunity)
	//创建群聊
	http.HandleFunc("/contact/createcommunity", controller.CreateCommunity)
	//连接socket
	http.HandleFunc("/chat", controller.Chat)

	//上传图片
	http.HandleFunc("/attach/upload", controller.Upload)

	//暴露静态资源
	http.Handle("/asset/", http.FileServer(http.Dir(".")))
	http.Handle("/mnt/", http.FileServer(http.Dir(".")))
	RegisterView()

	//启动Web服务器
	http.ListenAndServe(":8080", nil)
}

func RegisterView() {
	//一次解析出全部模板
	tpl, err := template.ParseGlob("view/**/*")
	if nil != err {
		log.Fatal(err)
	}
	//通过for循环做好映射
	for _, v := range tpl.Templates() {
		//
		tplname := v.Name()
		fmt.Println("HandleFunc     " + v.Name())
		http.HandleFunc(tplname, func(w http.ResponseWriter,
			request *http.Request) {
			//
			fmt.Println("parse     " + v.Name() + "==" + tplname)
			err := tpl.ExecuteTemplate(w, tplname, nil)
			if err != nil {
				log.Fatal(err.Error())
			}
		})
	}

}
