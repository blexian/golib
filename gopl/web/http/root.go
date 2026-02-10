package http

import (
	"fmt"
	"github.com/spf13/cobra"
	"log"
)

type PersistentFlags struct {
	Endpoint string
}

var perFlags *PersistentFlags

func help() {
	hlp := `bhs demo [options], start one demo http server
  example: bhs demo --endpoint 0.0.0.0:8080`
	fmt.Println(hlp)
}

func RootCmd() {
	cmd := &cobra.Command{
		Use:     "bhs",
		Short:   "start http server tools in github.com/blexian/golib/src/web/http",
		Example: "bhs demo --endpoint 0.0.0.0:8080",
		Run: func(cmd *cobra.Command, args []string) {
			help()
		},
	}
	cmd.AddCommand(demoCmd())
	log.Fatal(cmd.Execute())
}

func (p *PersistentFlags) persistentFlags(cmd *cobra.Command) {
	cmd.PersistentFlags().StringVar(&p.Endpoint, "endpoint", "0.0.0.0:8080", "http server endpoint")
}
