package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"

	"github.com/google/uuid"
	usg "github.com/luckyComet55/backend-trainee-assignment-2023/usersegment"
)

func createUserReport(data []usg.UserSegmentActions) (string, error) {
	filename, err := uuid.NewRandom()
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	filepath := fmt.Sprintf("./data/%s.csv", filename.String())
	file, err := os.Create(filepath)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	defer file.Close()
	csvWriter := csv.NewWriter(file)
	err = csvWriter.Write([]string{"user_id", "segment_name", "date", "operation"})
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	for _, v := range data {
		err = csvWriter.Write([]string{strconv.Itoa(v.UserId), v.SegmentName, v.Date.String(), v.OperationType})
		if err != nil {
			fmt.Println(err)
			return "", err
		}
	}
	csvWriter.Flush()
	return filename.String(), nil
}
