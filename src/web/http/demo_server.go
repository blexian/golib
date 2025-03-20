package http

import (
	"context"
	"github.com/spf13/cobra"
	"log"
	"net/http"
)

type DemoFlags struct {
	GetTest  string
	Endpoint string
}

var dFlags *DemoFlags

func getHandlerConstructor(s string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte(s))
		if err != nil {
			log.Fatal(err)
		}
	}
}

func demoServer(endpoint string, handlerMap map[string]http.HandlerFunc, cancel <-chan struct{}) {
	for url, handler := range handlerMap {
		http.HandleFunc(url, handler)
	}
	http.HandleFunc("/health", getHandlerConstructor("ok"))
	server := &http.Server{Addr: endpoint, Handler: http.DefaultServeMux}
	go func() {
		log.Fatal(server.ListenAndServe())
	}()
	<-cancel
	_ = server.Shutdown(context.Background())
}

func demoCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use: "demo",
		Run: func(cmd *cobra.Command, args []string) {
			handlerMap := map[string]http.HandlerFunc{}
			if dFlags.GetTest != "" {
				handlerMap["/test"] = getHandlerConstructor(dFlags.GetTest)
			}
			demoServer(dFlags.Endpoint, handlerMap, nil)
		},
	}
	dFlags = &DemoFlags{}
	dFlags.addFlags(cmd)
	return cmd
}

func (d *DemoFlags) addFlags(cmd *cobra.Command) {
	cmd.Flags().StringVar(&d.GetTest, "get-test", "", "create one get handler in http server, and return value that you assign when request")
	cmd.Flags().StringVar(&d.Endpoint, "endpoint", "0.0.0.0:8080", "http server endpoint")
}
