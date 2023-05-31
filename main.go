package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/ALTA-BE17/Rest-API-Clean-Arch-Test/app/database"
	"github.com/ALTA-BE17/Rest-API-Clean-Arch-Test/app/dependency"
	"github.com/ALTA-BE17/Rest-API-Clean-Arch-Test/app/dependency/container"
	"github.com/ALTA-BE17/Rest-API-Clean-Arch-Test/app/router"
)

func main() {
	container.Run()
	// container.App.Invoke dipanggil untuk menjalankan fungsi -fungsi yang telah terdaftar pada container.dependency.
	err := container.App.Invoke(func(d dependency.Dependency, r router.Routes) {
		database.Migration(d.Config)
		var sigChan = make(chan os.Signal, 1)
		signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
		r.RegisterRoutes()
		go func() {
			if err := d.Echo.Start(fmt.Sprintf(":%v", d.Config.Port)); err != nil {
				d.Logger.Panic("Failed to start server")
				sigChan <- syscall.SIGTERM
			}
		}()
		<-sigChan
		// press ctrl+c to close terminal.
		d.Logger.Info("Shutting down server...")
	})
	if err != nil {
		log.Print(err)
	}
}
