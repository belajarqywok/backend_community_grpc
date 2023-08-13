package main

import (
	"os"
	"context"
	b64 "encoding/base64"
	"golang.org/x/oauth2/google"
	log "github.com/sirupsen/logrus"

	"google.golang.org/api/sheets/v4"
	"google.golang.org/api/option"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Error(err)
		return
	}

	// create api context
	ctx := context.Background()

	// get bytes from base64 encoded google service accounts key
	credBytes, err := b64.StdEncoding.DecodeString(os.Getenv("KEY_JSON_BASE64"))
	if err != nil {
		log.Error(err)
		return
	}

	// authenticate and get configuration
	config, err := google.JWTConfigFromJSON(credBytes, "https://www.googleapis.com/auth/spreadsheets")
	if err != nil {
		log.Error(err)
		return
	}

	// create client with config and context
	client := config.Client(ctx)

	// create new service using client
	srv, err := sheets.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		log.Error(err)
		return
	}




	// https://docs.google.com/spreadsheets/d/<SPREADSHEETID>/edit#gid=<SHEETID>
	sheetId := 0
	spreadsheetId := "101qeQkjTtX6E5hxrZnLQwDSX6GkJySgoFVOVxDaTnkw"

	// Convert sheet ID to sheet name.
	response1, err := srv.Spreadsheets.Get(spreadsheetId).Fields("sheets(properties(sheetId,title))").Do()
	if err != nil || response1.HTTPStatusCode != 200 {
		log.Error(err)
		return
	}

	sheetName := ""
	for _, v := range response1.Sheets {
		prop := v.Properties
		if prop.SheetId == int64(sheetId) {
			sheetName = prop.Title
			break
		}
	}

	//Append value to the sheet.
	row := &sheets.ValueRange {
		Values: [][]interface{}{{"ABC", "abc@gmail.com"}},
	}

	response2, err := srv.Spreadsheets.Values.Append(spreadsheetId, sheetName, row).ValueInputOption("USER_ENTERED").InsertDataOption("INSERT_ROWS").Context(ctx).Do()
	if err != nil || response2.HTTPStatusCode != 200 {
		log.Error(err)
		return
	}

	log.Info(response2)
}