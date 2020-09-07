package web

func StartWebServer() {
	web := NewServer()
	web.Start()
	web.Destroy()
}
