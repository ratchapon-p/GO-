package prices

import (
	// "errors"
	"fmt"

	"example.com/calculator/conversion"
	// "example.com/calculator/filemanager"
	"example.com/calculator/iomanager"
)

type TaxIncludedPriceJob struct {
	IOManager	iomanager.IOManager `json:"-"`
	TaxRate          float64 `json:"tax_rate"`
	InputPrices      []float64 `json:"input_prices"`
	TaxIncludedPrics map[string]string `json:"tax_include_prices"`
}

//---------------------------------------------------------------------------------------

func (job *TaxIncludedPriceJob) LoadData() error {

	lines, err := job.IOManager.ReadLines()

	if err != nil{
		fmt.Println(err)
		return err
	}


	prices, err := conversion.StringToFloat(lines)

	if err != nil{
		fmt.Println(err)
		return err
	}

	job.InputPrices = prices

	return nil
}

//---------------------------------------------------------------------------------------

func (job TaxIncludedPriceJob) Process() error {

	err := job.LoadData()

	if err != nil {
		return err
		
	}



	result := make(map[string]string)

	for _, price := range job.InputPrices {
		taxIncludedPrics := price * (1 + job.TaxRate)
		result[fmt.Sprintf("%.2f", price)] = fmt.Sprintf("%.2f", taxIncludedPrics)
	}

	job.TaxIncludedPrics = result
	return job.IOManager.WriteJSON(job)

}

//---------------------------------------------------------------------------------------

func NewTaxIncludedPriceJob(iom iomanager.IOManager, taxRate float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		IOManager: iom,
		InputPrices: []float64{10, 20, 30},
		TaxRate:     taxRate,
	}

}
