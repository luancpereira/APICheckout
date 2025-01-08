package service

import (
	"context"
	"encoding/json"
	"fmt"
	"math"
	"net/http"
	"time"

	"github.com/jinzhu/copier"
	"github.com/luancpereira/APICheckout/core/database"
	"github.com/luancpereira/APICheckout/core/database/sqlc"
	coreError "github.com/luancpereira/APICheckout/core/errors"
)

type Checkout struct{}

/*****
funcs for creations
******/

func (c Checkout) CreateTransaction(description string, transaction_date time.Time, transaction_value float64) (ID int64, err error) {

	err = c.validateDescription(description)
	if err != nil {
		return
	}

	err = c.validateTrasactionValue(transaction_value)
	if err != nil {
		return
	}

	if math.Round(transaction_value*100) != transaction_value*100 {
		transaction_value = math.Round(transaction_value*100) / 100
	}

	params := sqlc.InsertTransactionParams{
		Description:      description,
		TransactionDate:  transaction_date,
		TransactionValue: transaction_value,
	}

	ID, err = database.DB_QUERIER.InsertTransaction(context.Background(), params)
	if err != nil {
		err = database.Utils{}.CoreErrorDatabase(err)
		return
	}

	return
}

/*****
funcs for creations
******/

/*****
funcs for gets
******/

func (Checkout) GetByID(transactionID int64) (transaction sqlc.SelectTransactionByIDRow, err error) {

	transaction, err = database.DB_QUERIER.SelectTransactionByID(context.Background(), transactionID)
	if err != nil {
		err = database.Utils{}.CoreErrorDatabase(err)
		return
	}

	// var response Response
	// err = GetEntity("https://api.fiscaldata.treasury.gov/services/api/fiscal_service/v1/accounting/od/rates_of_exchange?filter=country:eq:Brazil,record_calendar_year:eq:2024", map[string]string{}, &response)
	// if err != nil {
	// 	return
	// }

	return
}

func (Checkout) GetList(filters map[string]string, limit, offset int64) (models []sqlc.SelectTransactionsRow, total int64, err error) {

	params := sqlc.SelectTransactionsParams{
		Column1: limit,
		Column2: offset,
		MinDate: filters["min_date"],
		MaxDate: filters["max_date"],
	}

	models, err = database.DB_QUERIER.SelectTransactions(context.Background(), params)
	if err != nil {
		err = database.Utils{}.CoreErrorDatabase(err)
		return
	}

	var totalParams sqlc.SelectTransactionsTotalParams
	copier.Copy(&totalParams, params)

	total, err = database.DB_QUERIER.SelectTransactionsTotal(context.Background(), totalParams)
	if err != nil {
		err = database.Utils{}.CoreErrorDatabase(err)
		return
	}

	return
}

/*****
funcs for gets
******/

/*****
funcs for validations
******/

func (Checkout) validateDescription(description string) (err error) {
	if len(description) == 0 {
		err = coreError.New("error.description.empty")
		return
	}

	if len(description) > 50 {
		err = coreError.New("error.description.too.long")
		return
	}

	return
}

func (Checkout) validateTrasactionValue(value float64) (err error) {
	if value <= 0 {
		err = coreError.New("error.value.not.positive")
		return
	}

	return
}

/*****
funcs for validations
******/

/*****
other funcs
******/

type Record struct {
	RecordDate            string `json:"record_date"`
	Country               string `json:"country"`
	Currency              string `json:"currency"`
	CountryCurrencyDesc   string `json:"country_currency_desc"`
	ExchangeRate          string `json:"exchange_rate"`
	EffectiveDate         string `json:"effective_date"`
	SrcLineNbr            string `json:"src_line_nbr"`
	RecordFiscalYear      string `json:"record_fiscal_year"`
	RecordFiscalQuarter   string `json:"record_fiscal_quarter"`
	RecordCalendarYear    string `json:"record_calendar_year"`
	RecordCalendarQuarter string `json:"record_calendar_quarter"`
	RecordCalendarMonth   string `json:"record_calendar_month"`
	RecordCalendarDay     string `json:"record_calendar_day"`
}

type Meta struct {
	Count int `json:"count"`
}

type Response struct {
	Data []Record `json:"data"`
	Meta Meta     `json:"meta"`
}

func GetEntity(url string, headers map[string]string, target interface{}) error {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return fmt.Errorf("erro ao criar a requisição: %w", err)
	}

	for key, value := range headers {
		req.Header.Set(key, value)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("erro ao fazer a requisição: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("requisição falhou com status %d", resp.StatusCode)
	}

	if err := json.NewDecoder(resp.Body).Decode(target); err != nil {
		return fmt.Errorf("erro ao decodificar a resposta JSON: %w", err)
	}

	return nil
}

/*****
other funcs
******/
