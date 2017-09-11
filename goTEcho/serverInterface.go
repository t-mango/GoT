package goTEcho

//ServerInterface 
type ServerInterface interface{

	InitServer() *ServerInterface 
	Start()
}