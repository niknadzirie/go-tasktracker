package helper

import (
	"fmt"
	"os"
	"tasktracker/internal/model"
	"text/tabwriter"
)

func Display(tasks []model.Task) {
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', 0)

	fmt.Fprintln(w, "ID\tSTATUS\tCREATED AT\tDESCRIPTION")
	fmt.Fprintln(w, "--\t------\t----------\t-----------")

	for _, task := range tasks {
		fmt.Fprintf(w, "%d\t%s\t%s\t%s\n",
			task.ID,
			task.Status,
			task.CreatedAt.Format("2006-01-02 15:04"),
			task.Description,
		)
	}

	w.Flush()
}
