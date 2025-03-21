package common

import (
	"fmt"
	"os"
	"testing"
	"text/tabwriter"
)

func TestTabWriter(t *testing.T) {
	const format = "%v\t%v\t%v\t%v\t%v\t\n"
	tw := new(tabwriter.Writer).Init(os.Stdout, 0, 8, 2, ' ', 0)
	fmt.Fprintf(tw, format, "Title", "Artist", "Album", "Year", "Length")
	fmt.Fprintf(tw, format, "-----", "------", "-----", "----", "------")
	//for _, t := range tracks {
	//	fmt.Fprintf(tw, format, t.Title, t.Artist, t.Album, t.Year, t.Length)
	//}
	tw.Flush()
}
