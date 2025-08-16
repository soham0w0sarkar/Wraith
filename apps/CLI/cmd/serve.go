package cmd

import (
	"fmt"
	"log"
	"net/http"
	"os/exec"

	"github.com/creack/pty"
	"github.com/fsnotify/fsnotify"
	"github.com/gorilla/websocket"
	"github.com/spf13/cobra"
)



var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start local Wraith IDE server",
	Run: func(cmd *cobra.Command, args []string) {
		runServer()
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
}

func runServer() {
	
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				fmt.Println("File change:", event)
				
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				fmt.Println("Watcher error:", err)
			}
		}
	}()

	err = watcher.Add(".")
	if err != nil {
		log.Fatal(err)
	}

	upgrader := websocket.Upgrader{}
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Println("upgrade:", err)
			return
		}
		go handleWS(conn)
	})

	fs := http.FileServer(http.Dir("./frontend/dist"))
	http.Handle("/", fs)

	fmt.Println("🚀 Wraith running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleWS(conn *websocket.Conn) {
	cmd := exec.Command("bash")
	ptmx, err := pty.Start(cmd)
	if err != nil {
		log.Println("failed to start pty:", err)
		return
	}
	defer func() { _ = ptmx.Close() }()

	// read from PTY -> WS
	go func() {
		buf := make([]byte, 1024)
		for {
			n, err := ptmx.Read(buf)
			if err != nil {
				return
			}
			conn.WriteMessage(websocket.TextMessage, buf[:n])
		}
	}()

	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			return
		}
		ptmx.Write(msg)
	}
}

