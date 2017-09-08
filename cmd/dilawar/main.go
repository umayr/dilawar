package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/dustin/go-humanize"
	"github.com/olekukonko/tablewriter"
	"github.com/umayr/dilawar"
	"github.com/urfave/cli"
)

func main() {

	app := cli.NewApp()
	app.Name = "dilawar"
	app.Usage = "tiny finance manager"
	app.Version = "0.1.0"

	app.Action = func(c *cli.Context) error {
		b, err := dilawar.Balance()
		if err != nil {
			return err
		}

		fmt.Printf("%d\n", b)
		return nil
	}

	app.Commands = []cli.Command{
		{
			Name:  "debit",
			Usage: "adds amount in debit",
			Flags: []cli.Flag{
				cli.IntFlag{
					Name:  "amount, a",
					Usage: "amount to be added",
					Value: 0,
				},
				cli.StringFlag{
					Name:  "message, m",
					Usage: "message for the transaction",
				},
			},
			Action: func(c *cli.Context) error {
				if c.Int("amount") <= 0 {
					return cli.NewExitError("provided amount is invalid", 1)
				}

				return dilawar.Debit(c.Int("amount"), c.String("message"))
			},
		},
		{
			Name:  "credit",
			Usage: "adds amount in credit",
			Flags: []cli.Flag{
				cli.IntFlag{
					Name:  "amount, a",
					Usage: "amount to be added",
					Value: 0,
				},
				cli.StringFlag{
					Name:  "message, m",
					Usage: "message for the transaction",
				},
			},
			Action: func(c *cli.Context) error {
				if c.Int("amount") <= 0 {
					return cli.NewExitError("provided amount is invalid", 1)
				}

				return dilawar.Credit(c.Int("amount"), c.String("message"))
			},
		},
		{
			Name:  "history",
			Usage: "shows all history",
			Flags: []cli.Flag{
				cli.BoolFlag{
					Name:  "csv",
					Usage: "outputs as comma seperate values",
				},
			},
			Action: func(c *cli.Context) error {
				items, err := dilawar.History()
				if err != nil {
					return cli.NewExitError(fmt.Sprintf("error: %s", err.Error()), 127)
				}

				data := [][]string{}

				for _, v := range items {
					var (
						id   string
						tm   string
						desc string
						abbr string
						amt  string
					)

					if v.Type == dilawar.TypeCredit {
						abbr = "CR"
					} else if v.Type == dilawar.TypeDebit {
						abbr = "DR"
					}

					id = strconv.Itoa(int(v.ID))
					tm = humanize.Time(v.Time)
					if v.Description == "" {
						desc = "N/A"
					} else {
						desc = v.Description
					}
					amt = fmt.Sprintf("%s %s", humanize.Comma(int64(v.Amount)), "PKR")

					data = append(data, []string{id, desc, tm, abbr, amt})
				}

				b, err := dilawar.Balance()
				if err != nil {
					return cli.NewExitError(fmt.Sprintf("error: %s", err.Error()), 127)
				}

				head := []string{"ID", "Description", "Time", "Type", "Amount"}
				foot := []string{"", "", "", "Total", humanize.Comma(int64(b))}

				if c.Bool("csv") {
					fmt.Println(strings.Join(head, ","))
					for _, line := range data {
						fmt.Println(strings.Join(line, ","))
					}
					fmt.Println(strings.Join(foot, ","))
					return nil
				}

				table := tablewriter.NewWriter(os.Stdout)
				table.SetHeader(head)
				table.SetFooter(foot)
				table.SetBorder(false)
				table.AppendBulk(data)
				table.Render()

				return nil
			},
		},
	}
	app.Run(os.Args)
}
