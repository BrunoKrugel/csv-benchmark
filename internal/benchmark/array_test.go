package benchmark

import (
	"testing"

	"github.com/google/uuid"
)

func BenchmarkAppendIndexUUID(b *testing.B) {
	array := ArrayOfUUID{
		ID: []uuid.UUID{
			uuid.New(),
			uuid.New(),
			uuid.New(),
		},
	}

	for b.Loop() {
		AppendIndexUUID(array)
	}
}

// func BenchmarkAppendLimitUUID(b *testing.B) {
// 	array := ArrayOfUUID{
// 		ID: []uuid.UUID{
// 			uuid.New(),
// 			uuid.New(),
// 			uuid.New(),
// 		},
// 	}

// 	for b.Loop() {
// 		AppendLimitUUID(array)
// 	}
// }

// func BenchmarkAppendFullUUID(b *testing.B) {
// 	array := ArrayOfUUID{
// 		ID: []uuid.UUID{
// 			uuid.New(),
// 			uuid.New(),
// 			uuid.New(),
// 		},
// 	}

// 	for b.Loop() {
// 		AppendFullUUID(array)
// 	}
// }

func BenchmarkAppendIndexString(b *testing.B) {
	array := ArrayOfString{
		ID: []string{
			"e604f0c5-198e-43cd-a465-186278ef1807",
			"8241e6cf-2e6f-400a-8c1f-84acc3b33798",
			"33863527-af5b-434e-a800-563fe4f6bfa4",
		},
	}

	for b.Loop() {
		AppendIndexString(array)
	}
}

// func BenchmarkAppendLimitString(b *testing.B) {
// 	array := ArrayOfString{
// 		ID: []string{
// 			"e604f0c5-198e-43cd-a465-186278ef1807",
// 			"8241e6cf-2e6f-400a-8c1f-84acc3b33798",
// 			"33863527-af5b-434e-a800-563fe4f6bfa4",
// 		},
// 	}

// 	for b.Loop() {
// 		AppendLimitString(array)
// 	}
// }

// func BenchmarkAppendFullString(b *testing.B) {
// 	array := ArrayOfString{
// 		ID: []string{
// 			"e604f0c5-198e-43cd-a465-186278ef1807",
// 			"8241e6cf-2e6f-400a-8c1f-84acc3b33798",
// 			"33863527-af5b-434e-a800-563fe4f6bfa4",
// 		},
// 	}

// 	for b.Loop() {
// 		AppendFullString(array)
// 	}
// }
