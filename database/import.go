package database

import (
	"fmt"

	"github.com/tealeg/xlsx"
)

func ImportExcel(filename string) error {
	f, err := xlsx.OpenFile(filename)
	if err != nil {
		return err
	}
	imported := 0
	for ridx, r := range f.Sheets[0].Rows {
		if ridx == 0 {
			continue
		}
		myTicket := &Ticket{}
		err := r.ReadStruct(myTicket)
		if err != nil {
			fmt.Println("Row ", ridx, err.Error())
			fmt.Println("data", myTicket)
			return err
		}
		sticket := &GTicket{}
		search := db.Where("ticket_number = ?", myTicket.TicketNumber).First(&sticket)
		if search.RowsAffected > 0 {
			// record exists, we can move along, nothing to see here
			continue
		}

		gticket := &GTicket{
			Ticket: *myTicket,
		}
		if result := db.Save(gticket); result.Error != nil {
			fmt.Println("Row ", ridx, err.Error())
			fmt.Println("data", myTicket)
			return err
		}
		imported++

	}
	fmt.Printf("Imported %d", imported)
	return nil
}
