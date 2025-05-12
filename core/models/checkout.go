package models

import "github.com/luancpereira/APICheckout/core/database/sqlc"

type TransactionDetail struct {
	sqlc.SelectTransactionByIDRow
	ExchangeRate                            float64
	TransactionValueConvertedToWishCurrency float64
}

type TransactionDetailList struct {
	sqlc.SelectTransactionsRow
	ExchangeRate                            float64
	TransactionValueConvertedToWishCurrency float64
}

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
