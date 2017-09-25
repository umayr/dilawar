package main

import (
	"bytes"
	"fmt"
	"strconv"

	"github.com/andlabs/ui"
	humanize "github.com/dustin/go-humanize"
	"github.com/umayr/dilawar"
)

func main() {
	err := ui.Main(func() {
		amountGroup := ui.NewGroup("Amount")
		messageGroup := ui.NewGroup("Message")
		balanceLabel := ui.NewLabel("Your balances: 0")
		amount := ui.NewEntry()
		message := ui.NewEntry()

		debit := ui.NewButton("Debit")
		credit := ui.NewButton("Credit")
		log := ui.NewButton("View log")

		response := ui.NewLabel("")

		transactionGroup := ui.NewGroup("Recent Transactions:")

		amountGroup.SetChild(amount)
		amountGroup.SetMargined(true)
		messageGroup.SetChild(message)
		messageGroup.SetMargined(true)

		actionBx := ui.NewHorizontalBox()
		actionBx.Append(debit, false)
		actionBx.Append(credit, false)
		actionBx.Append(log, false)
		actionBx.SetPadded(true)

		mainBox := ui.NewVerticalBox()
		mainBox.Append(balanceLabel, false)
		mainBox.Append(ui.NewHorizontalSeparator(), false)
		mainBox.Append(amountGroup, false)
		mainBox.Append(messageGroup, false)

		mainBox.Append(actionBx, false)
		mainBox.Append(ui.NewHorizontalSeparator(), false)
		mainBox.Append(response, false)
		mainBox.Append(ui.NewHorizontalSeparator(), false)
		mainBox.Append(ui.NewHorizontalSeparator(), false)
		mainBox.Append(transactionGroup, false)

		window := ui.NewWindow("Dilawar Management!", 300, 300, false)
		window.SetChild(mainBox)

		debit.OnClicked(func(*ui.Button) {
			debit.Disable()
			defer debit.Enable()
			actionHandler(true, amount, message, response)
			updateBalanceLabel(balanceLabel)
			recentTransactions(response, transactionGroup)
		})

		credit.OnClicked(func(*ui.Button) {
			credit.Disable()
			defer credit.Enable()
			actionHandler(false, amount, message, response)
			updateBalanceLabel(balanceLabel)
			recentTransactions(response, transactionGroup)
		})

		log.OnClicked(func(*ui.Button) {
			response.SetText("work in progress")
		})

		window.OnClosing(func(*ui.Window) bool {
			ui.Quit()
			return true
		})

		updateBalanceLabel(balanceLabel)
		recentTransactions(response, transactionGroup)
		window.Show()
	})

	if err != nil {
		panic(err)
	}
}

func actionHandler(isDebit bool, amount *ui.Entry, message *ui.Entry, response *ui.Label) {
	action := "Credit"
	if isDebit {
		action = "Debit"
	}

	i, err := strconv.Atoi(amount.Text())
	if err != nil || i < 0 {
		response.SetText("Error occured! Invalid integer")
	} else {
		if isDebit {
			err = dilawar.Debit(i, message.Text())
		} else {
			err = dilawar.Credit(i, message.Text())
		}
		if err != nil {
			response.SetText("Error occured!")
		} else {
			response.SetText(amount.Text() + fmt.Sprintf(" %s!", action))
		}
	}
}

func updateBalanceLabel(balanceLabel *ui.Label) {
	b, err := dilawar.Balance()
	if err != nil {
		balanceLabel.SetText(fmt.Sprintf("Error occurred while fetching balance"))
	} else {
		balanceLabel.SetText(fmt.Sprintf("Your balance: %d", b))
	}
}

func recentTransactions(response *ui.Label, transactionGroup *ui.Group) {
	transactionGroup.SetChild(getTransactionLogs(true, response))
}

func getTransactionLogs(showLimited bool, response *ui.Label) *ui.Box {
	items, err := dilawar.History()
	if err != nil {
		response.SetText(fmt.Sprintf("error: %s", err.Error()))
		return nil
	}
	rows := ui.NewVerticalBox()
	numOfRecords := len(items) - 1
	for i := numOfRecords; i > 0 && (i > numOfRecords-10 || !showLimited); i-- {
		v := items[i]
		var (
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

		tm = humanize.Time(v.Time)
		if v.Description == "" {
			desc = "N/A"
		} else {
			desc = v.Description
			desc = splitSubN(desc, 127)
		}
		amt = fmt.Sprintf("%s %s", humanize.Comma(int64(v.Amount)), "PKR")

		itemRow := ui.NewHorizontalBox()
		itemRow.SetPadded(true)

		itemRow.Append(ui.NewLabel(amt), false)
		itemRow.Append(ui.NewLabel(abbr), false)
		itemRow.Append(ui.NewLabel(tm), false)
		itemRow.Append(ui.NewLabel(desc), false)

		rows.Append(itemRow, true)
		rows.Append(ui.NewHorizontalSeparator(), false)
	}

	return rows
}

func splitSubN(s string, n int) string {
	sub := ""
	subs := []string{}

	runes := bytes.Runes([]byte(s))
	l := len(runes)
	for i, r := range runes {
		sub = sub + string(r)
		if (i+1)%n == 0 {
			subs = append(subs, sub)
			sub = ""
		} else if (i + 1) == l {
			subs = append(subs, sub)
		}
	}

	if len(subs) > 1 {
		return subs[0] + "..."
	}
	return subs[0]
}
