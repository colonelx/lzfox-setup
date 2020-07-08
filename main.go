package main

import (
	"log"
	"os"
	"time"

	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
	"go.bug.st/serial.v1"
)

const gladeUI = "lzfox_setup.glade"

func timeTick(lbl *gtk.Label) {

	for true {
		currentTime := time.Now()
		lbl.SetText(currentTime.Format("2006-01-02 15:04:05"))
		time.Sleep(time.Second)
	}
}

func main() {
	const appID = "com.benchparty.lzfox_setup"
	app, err := gtk.ApplicationNew(appID, glib.APPLICATION_FLAGS_NONE)
	// Initialize GTK without parsing any command line arguments.
	gtk.Init(&os.Args)

	// Create a new toplevel window, set its title, and connect it to the
	// "destroy" signal to exit the GTK main loop when it is destroyed.
	builder, err := gtk.BuilderNewFromFile(gladeUI)
	if err != nil {
		log.Fatal("Unable to create window:", err)
		panic(err)
	}

	windowObj, err := builder.GetObject("lzfox_main")
	window, _ := windowObj.(*gtk.Window)
	if err != nil {
		log.Fatal("Unable to get the window:", err)
		panic(err)
	}

	// win.SetTitle("Simple Example")
	_, err = window.Connect("destroy", func() {
		gtk.MainQuit()
	})

	if err != nil {
		log.Fatal("Unable to destroy window:", err)
		panic(err)
	}

	// c := time.Tick(1 * time.Second)
	// progress := &Progress{}

	// go longJob(progress)
	// for {
	// 	select {
	// 	case <-c:
	// 		fmt.Println(progress.Get())
	// 	}
	// }

	// Detect ports
	selTtyObj, err := builder.GetObject("selTty")
	selTty, _ := selTtyObj.(*gtk.ComboBoxText)

	ports := detectTtys()

	for _, port := range ports {
		selTty.Append(port, port)
	}

	lblSystemTimeObj, err := builder.GetObject("lblSystemTime")
	lblSystemTime, _ := lblSystemTimeObj.(*gtk.Label)
	go timeTick(lblSystemTime)

	window.SetDefaultSize(600, 400)
	window.ShowAll()
	gtk.Main()
	os.Exit(app.Run(os.Args))
}

func detectTtys() []string {
	ports, err := serial.GetPortsList()

	if err != nil {
		log.Fatal(err)
	}
	if len(ports) == 0 {
		log.Fatal("No serial ports found!")
	}
	return ports
}
