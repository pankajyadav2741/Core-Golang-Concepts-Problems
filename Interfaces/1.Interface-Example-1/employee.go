package main

import "fmt"

type salaryCalculator interface {
	calculateSalary() float64
}

type permanent struct {
	empID       int
	basicPay    float64
	variablePay float64
	pf          float64
}

type contract struct {
	empID       int
	basicPay    float64
	variablePay float64
}

type freelance struct {
	empID       int
	ratePerHour float64
	hours       float64
}

func (p permanent) calculateSalary() float64 {
	return p.basicPay + p.variablePay + p.pf
}

func (c contract) calculateSalary() float64 {
	return c.basicPay + c.variablePay
}

func (f freelance) calculateSalary() float64 {
	return f.ratePerHour * f.hours
}

func totalExpenses(salary []float64) float64 {
	var total float64
	for _, v := range salary {
		total += v
	}
	return total
}

func main() {
	pEmp := permanent{
		empID:       23,
		basicPay:    30000.00,
		variablePay: 10000.00,
		pf:          2000.00,
	}

	cEmp := contract{
		empID:       47,
		basicPay:    20000.00,
		variablePay: 15000.00,
	}

	fEmp := freelance{
		empID:       34,
		ratePerHour: 200.50,
		hours:       9.5,
	}

	expense := totalExpenses([]float64{pEmp.calculateSalary(), cEmp.calculateSalary(), fEmp.calculateSalary()})
	fmt.Println("Total Company Expenses in a month: ", expense)
}
