package loader

import (
	"encoding/csv"
	"os"
	"strconv"

	"github.com/amandaanjolin/desafio-go-web/model"
)

func LoadTickets(filePath string) ([]model.Ticket, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	_, err = reader.Read() // ignora o header
	if err != nil {
		return nil, err
	}

	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	var tickets []model.Ticket
	for _, r := range records {
		price, _ := strconv.ParseFloat(r[5], 64)
		tickets = append(tickets, model.Ticket{
			ID:          r[0],
			Name:        r[1],
			Email:       r[2],
			Destination: r[3],
			Hour:        r[4],
			Price:       price,
		})
	}

	return tickets, nil
}
