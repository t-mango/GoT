package goTEcho

func (self *WebService) initRoute() {

	//静态文件

	self.Static("/", "public/index.html")

	//webscoket 定义

	self.GET("/nb-iot", self.wsNBiot) //

	//jwt 令牌发放  先在一个池子里吗

	//self.WebService.GetToken("/login/token")

}
