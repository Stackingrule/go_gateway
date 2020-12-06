package main

import (
	"github.com/e421083458/gin_scaffold/router"
	"github.com/e421083458/golang_common/lib"
	"os"
	"os/signal"
	"syscall"
)

<<<<<<< HEAD
func main()  {
	lib.InitModule("./conf/dev/",[]string{"base","mysql","redis",})
=======
func main() {
	lib.InitModule("./conf/dev/", []string{"base", "mysql", "redis"})
>>>>>>> 75546dc0f54b4f6d4ece0208a542fdec4d21faa0
	defer lib.Destroy()
	router.HttpServerRun()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGKILL, syscall.SIGQUIT, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	router.HttpServerStop()
<<<<<<< HEAD
}
=======
}
>>>>>>> 75546dc0f54b4f6d4ece0208a542fdec4d21faa0
