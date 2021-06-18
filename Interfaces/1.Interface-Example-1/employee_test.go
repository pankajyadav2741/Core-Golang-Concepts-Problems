package main

import "testing"

func Test_permanent_calculateSalary(t *testing.T) {
	tests := []struct {
		name string
		p    permanent
		want float64
	}{
		{
			name: "Permanent Employee",
			p: permanent{
				empID:       1,
				basicPay:    600.00,
				variablePay: 350.00,
				pf:          50.00,
			},
			want: 1000.00,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.p.calculateSalary(); got != tt.want {
				t.Errorf("permanent.calculateSalary() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_contract_calculateSalary(t *testing.T) {
	tests := []struct {
		name string
		c    contract
		want float64
	}{
		{
			name: "Contract Employee",
			c: contract{
				empID:       1,
				basicPay:    600.00,
				variablePay: 350.00,
			},
			want: 950.00,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.calculateSalary(); got != tt.want {
				t.Errorf("contract.calculateSalary() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_freelance_calculateSalary(t *testing.T) {
	tests := []struct {
		name string
		f    freelance
		want float64
	}{
		{
			name: "Freelance Employee",
			f: freelance{
				empID:       1,
				ratePerHour: 100.00,
				hours:       10.00,
			},
			want: 1000.00,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.f.calculateSalary(); got != tt.want {
				t.Errorf("freelance.calculateSalary() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_totalExpenses(t *testing.T) {
	type args struct {
		salary []float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "Calculate Total Expenses in a month",
			args: args{[]float64{100.00, 100.00, 100.00}},
			want: 300.00,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := totalExpenses(tt.args.salary); got != tt.want {
				t.Errorf("totalExpenses() = %v, want %v", got, tt.want)
			}
		})
	}
}
