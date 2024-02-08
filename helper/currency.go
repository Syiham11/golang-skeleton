package helper

func CalculateTotalFee(realTotalFee float64, discount float64, taxRate float32) float64 {
	totalFee := realTotalFee - discount + (realTotalFee * (float64(taxRate) / 100))
	return float64(totalFee)
}
