package main

import (
   ch "./ch"
   "fmt"
)


func main () {
    var tt ch.Obj2node
    tt.Init(100000,1000)

tt.AddNode("192.168.1.117")
tt.AddNode("192.168.1.149")
tt.AddNode("192.168.1.157")
tt.AddNode("192.168.1.150")
tt.AddNode("192.168.1.136")
tt.AddNode("192.168.1.154")
tt.AddNode("192.168.1.148")
tt.AddNode("192.168.1.158")
tt.AddNode("192.168.1.151")
tt.AddNode("192.168.1.130")
tt.AddNode("192.168.1.153")
tt.AddNode("192.168.1.135")
tt.AddNode("192.168.1.147")
tt.AddNode("192.168.1.132")
tt.AddNode("192.168.1.162")
tt.AddNode("192.168.1.163")
tt.AddNode("192.168.1.164")
tt.AddNode("192.168.1.165")
tt.AddNode("192.168.1.178")
tt.AddNode("192.168.1.159")
tt.AddNode("192.168.1.189")
tt.AddNode("192.168.1.190")
tt.AddNode("192.168.1.191")
tt.AddNode("192.168.1.192")
tt.AddNode("192.168.1.193")
tt.AddNode("192.168.1.194")
tt.AddNode("192.168.1.195")
tt.AddNode("192.168.1.196")
tt.AddNode("192.168.1.197")
tt.AddNode("192.168.1.198")
tt.AddNode("192.168.1.199")
tt.AddNode("192.168.1.200")
tt.AddNode("192.168.1.201")
tt.AddNode("192.168.1.202")
tt.AddNode("192.168.1.203")
tt.AddNode("192.168.1.204")
tt.AddNode("192.168.1.205")
tt.AddNode("192.168.1.206")
tt.AddNode("192.168.1.207")
tt.AddNode("192.168.1.208")
tt.AddNode("192.168.1.184")

    tt.Active()   //调用Add ,del后,必须调用该函数,才可以使用Get. 否则程序将退出.
    fmt.Println(tt.Get("jdsfjsf"))
    fmt.Println(tt.Get("123456"))
    fmt.Println(tt.Get("45345132"))
    fmt.Println(tt.Get("8760060"))
    fmt.Println(tt.Get("564636345345"))
    //tt.Delete("192.168.1.1")
    
    //tt.Active()
    tt.Debug()
}














