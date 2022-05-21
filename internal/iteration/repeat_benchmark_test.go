package iteration

import "testing"

func BenchmarkTableDriven(b *testing.B) {

	benchmarks := []struct {
		name    string
		subject func(string) string
	}{
		{"Repeat", Repeat},
		{"RepeatWithBuilder", RepeatWithBuilder},
		{"RepeatWithStringsRepeat", RepeatWithStringsRepeat},
	}

	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				bm.subject("a")
			}
		})
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a")
	}
}

func BenchmarkRepeatWithBuilder(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RepeatWithBuilder("a")
	}
}
func BenchmarkRepeatWithStringsRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RepeatWithStringsRepeat("a")
	}
}
