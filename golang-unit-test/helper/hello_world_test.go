package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSubTest(t *testing.T) {
	t.Run("Khalid", func(t *testing.T) {
		result := HelloWorld("Khalid")
		// assert otomatis memanggil t.Fail jika gagal
		assert.Equal(t, "Hello Khalid", result, "Result must be 'Hello Khalid'")

		// kode ini akan dieksekusi karena t.Fail tetap melanjutkan test setelah gagal
		fmt.Println("TestHelloWorld with assert is done")
	})

	t.Run("Bilal", func(t *testing.T) {
		result := HelloWorld("Bilal")
		// require otomatis memanggil t.FailNow jika gagal
		require.Equal(t, "Hello Bilal", result, "Result must be 'Hello Bilal'")

		// kode ini tidak akan dieksekusi karena t.FailNow tidak melanjutkan test setelah gagal
		fmt.Println("TestHelloWorld with require is done")
	})
}

func TestMain(m *testing.M) {
	fmt.Println("Before unit test")
	m.Run()
	fmt.Println("After unit test")
}

func TestHelloWorldAssert(t *testing.T) {
	result := HelloWorld("Udin")
	// assert otomatis memanggil t.Fail jika gagal
	assert.Equal(t, "Hello Udin", result, "Result must be 'Hello Udin'")

	// kode ini akan dieksekusi karena t.Fail tetap melanjutkan test setelah gagal
	fmt.Println("TestHelloWorld with assert is done")
}

func TestHelloWorldRequire(t *testing.T) {
	result := HelloWorld("Udin")
	// require otomatis memanggil t.FailNow jika gagal
	require.Equal(t, "Hello Udin", result, "Result must be 'Hello Udin'")

	// kode ini tidak akan dieksekusi karena t.FailNow tidak melanjutkan test setelah gagal
	fmt.Println("TestHelloWorld with require is done")
}

func TestSkip(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("Can't run on windows")
	}
}

func TestHelloWorldAdam(t *testing.T) {
	result := HelloWorld("Adam")

	if result != "Hello Adam" {
		// jangan gunakan panic untuk menggagalkan unit test
		// panic("Result is not 'Hello Adam'")

		// gunakan t.Fail atau t.FailNow
		// t.Fail()

		// tapi lebih baik menggunakan t.Error atau t.Fatal
		// t.Error otomatis memanggil t.Fail, t.Fatal otomatis memanggil t.FailNow
		t.Error("Result must be 'Hello Adam'")

	}

	fmt.Println("TestHelloWorldAdam is done")
}

func TestHelloWorldDzikri(t *testing.T) {
	result := HelloWorld("Dzikri")

	if result != "Hello Dzikri" {
		// jangan gunakan panic untuk menggagalkan unit test
		// gunakan t.Fail atau t.FailNow

		// panic("Result is not 'Hello Dzikri'")
		// t.FailNow()

		// tapi lebih baik menggunakan t.Error atau t.Fatal
		// t.Error otomatis memanggil t.Fail, t.Fatal otomatis memanggil t.FailNow
		t.Fatal("Result must be 'Hello Adam'")
	}

	fmt.Println("TestHelloWorldDzikri is done")
}

func TestHelloWorldTable(t *testing.T) {
	tests := []struct {
		name string
		request string
		expected string
	} {
		{
			name: "Dzikri",
			request: "Dzikri",
			expected: "Hello Dzikri",
		},
		{
			name: "Budi",
			request: "Budi",
			expected: "Hello Budi",
		},
		{
			name: "Anwar",
			request: "Anwar",
			expected: "Hello Anwar",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := HelloWorld(test.request)
			assert.Equal(t, test.expected,result)
		})
	}
}

func BenchmarkHelloWorld(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Dzikri")
	}
}

func BenchmarkHelloCuy(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Cuy")
	}
}

func BenchmarkSub(b *testing.B) {
	b.Run("Dzikri", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Dzikri")
		}
	})
	
	b.Run("Cuy", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Cuy")
		}
	})
}

func BenchmarkTable(b *testing.B) {
	benchmarks := []struct{
		name string
		request string
	} {
		{
			name: "Dzikri",
			request: "Dzikri",
		},
		{
			name: "Cuy",
			request: "Cuy",
		},
		{
			name: "Kiki",
			request: "Kiki",
		},
		{
			name: "Abud",
			request: "Abud",
		},
		{
			name: "Udin",
			request: "Udin",
		},
	}

	for _, benchmark := range benchmarks {
		b.Run(benchmark.name, func(b *testing.B) {
			HelloWorld(benchmark.request)
		})
	}
}