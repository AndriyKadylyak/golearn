package main

import "testing"

func Test_convert(t *testing.T) {
	type args struct {
		twelveHourTime string
	}
	tests := []struct {
		name                   string
		args                   args
		wantTwentyFourHourTime string
	}{
		{
			name: "taste case1",
			args: args{
				twelveHourTime: "03:04:00 PM",
			},
			wantTwentyFourHourTime: "15:04:00",
		},
		{
			name: "taste case2",
			args: args{
				twelveHourTime: "11:04:00 PM",
			},
			wantTwentyFourHourTime: "23:04:00",
		},
		{
			name: "taste case2",
			args: args{
				twelveHourTime: "07:24:00 PM",
			},
			wantTwentyFourHourTime: "19:24:00",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotTwentyFourHourTime := convert(tt.args.twelveHourTime); gotTwentyFourHourTime != tt.wantTwentyFourHourTime {
				t.Errorf("convert() = %v, want %v", gotTwentyFourHourTime, tt.wantTwentyFourHourTime)
			}
		})
	}
}
