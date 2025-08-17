package main

import (
	"bytes"
	"errors"
	"reflect"
	"testing"
)

func TestHello(t *testing.T) {
	// got := Hello()
	// want := "Hello, world"

	// if got != want {
	// 	t.Errorf("got %q want %q", got, want)
	// }

	// got := Hello("Chris")
	// want := "Hello, Chris"

	// if got != want {
	//     t.Errorf("got '%q' want '%q'", got, want)
	// }

	// t.Run("say hello world when an empty string in supplied", func(t *testing.T) {
	// 	got := Hello("")
	// 	want := "Hello, world"

	// 	if got != want {
	// 		t.Errorf("got %q want %q", got, want)
	// 	}
	// })

	// 抽出assertCorrectMessage
	assertCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Chris", "")
		want := "Hello, Chris"
		assertCorrectMessage(t, got, want)
	})

	t.Run("empty string defaults to 'world'", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, world"
		assertCorrectMessage(t, got, want)
	})

	t.Run("saying hello in Spanish", func(t *testing.T) {
		got := Hello("Elodie", "Spanish")
		want := "Hola, Elodie"
		assertCorrectMessage(t, got, want)
	})

}

func TestAdder(t *testing.T) {
	sum := Add(2, 2)
	expected := 4

	if sum != expected {
		t.Errorf("expected '%d' but got '%d'", expected, sum)
	}
}

func TestRepeat(t *testing.T) {
	repeated := Repeat("a")
	expected := "aaaaa"

	if repeated != expected {
		t.Errorf("expected '%q' but got '%q'", expected, repeated)
	}
}

func TestSum(t *testing.T) {
	t.Run("collection of 5 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}
		got := Sum(numbers)
		expected := 15

		if got != expected {
			t.Errorf("expected '%d' but got '%d'", expected, got)
		}
	})

	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		got := Sum(numbers)
		want := 6

		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})
}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2}, []int{0, 9})
	want := []int{3, 9}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestSumAllTails(t *testing.T) {
	// got := SumAllTails([]int{1, 2}, []int{0, 9})
	// want := []int{2, 9}

	// if !reflect.DeepEqual(got, want) {
	// 	t.Errorf("got %v want %v", got, want)
	// }

	checkSums := func(t *testing.T, got, want []int) {
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	}

	t.Run("make the sums of some slices", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{0, 9})
		want := []int{2, 9}
		checkSums(t, got, want)
	})

	t.Run("safely sum empty slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{3, 4, 5})
		want := []int{0, 9}
		checkSums(t, got, want)
	})
}

// func TestPerimeter(t *testing.T) {
// 	rectangle := Rectangle{10.0, 10.0}
//     got := Perimeter(rectangle)
//     want := 40.0

//     if got != want {
//         t.Errorf("got %.2f want %.2f", got, want)
//     }
// }

// func TestArea(t *testing.T) {
// 	rectangle := Rectangle{12.0, 6.0}
//     got := Area(rectangle)
//     want := 72.0

//     if got != want {
//         t.Errorf("got %.2f want %.2f", got, want)
//     }
// }

func TestArea(t *testing.T) {
	// checkArea := func(t *testing.T, shape Shape, want float64) {
	//     t.Helper()
	//     got := shape.Area()
	//     if got != want {
	//         t.Errorf("got %.2f want %.2f", got, want)
	//     }
	// }

	// t.Run("rectangles", func(t *testing.T) {
	//     rectangle := Rectangle{12, 6}
	//     got := rectangle.Area()
	//     want := 72.0

	//     if got != want {
	//         t.Errorf("got %.2f want %.2f", got, want)
	//     }
	// })

	// t.Run("circles", func(t *testing.T) {
	//     circle := Circle{10}
	//     got := circle.Area()
	//     want := 314.1592653589793

	//     if got != want {
	//         t.Errorf("got %.2f want %.2f", got, want)
	//     }
	// })

	// t.Run("rectangles", func(t *testing.T) {
	//     rectangle := Rectangle{12, 6}
	//     checkArea(t, rectangle, 72.0)
	// })

	// t.Run("circles", func(t *testing.T) {
	//     circle := Circle{10}
	//     checkArea(t, circle, 314.1592653589793)
	// })

	areaTests := []struct {
		name    string
		shape   Shape
		hasArea float64
	}{
		{name: "Rectangle", shape: Rectangle{12, 6}, hasArea: 72.0},
		{name: "Circle", shape: Circle{10}, hasArea: 314.1592653589793},
		{name: "Triangle", shape: Triangle{12, 6}, hasArea: 36.0},
	}

	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()
			if got != tt.hasArea {
				t.Errorf("%v got %.2f want %.2f", tt.shape, got, tt.hasArea)
			}
		})
	}
}

func TestWallet(t *testing.T) {
	assertBalance := func(t *testing.T, wallet Wallet, want Bitcoin) {
		got := wallet.Balance()
		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	}

	assertError := func(t *testing.T, got error, want error) {
		if got == nil {
			t.Error("didn't get an error but wanted one")
		}

		if !errors.Is(got, want) {
			t.Errorf("got '%s', want '%s'", got, want)
		}
	}

	assertNoError := func(t *testing.T, got error) {
		if got != nil {
			t.Fatal("got an error but didnt want one")
		}
	}

	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))
		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("Withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}
		err := wallet.Withdraw(Bitcoin(10))

		assertBalance(t, wallet, Bitcoin(10))
		assertNoError(t, err)
	})

	t.Run("Withdraw insufficient funds", func(t *testing.T) {
		startingBalance := Bitcoin(20)
		wallet := Wallet{balance: startingBalance}
		err := wallet.Withdraw(Bitcoin(100))

		assertBalance(t, wallet, startingBalance)
		assertError(t, err, InsufficientFundsError)
	})
}

func TestSearch(t *testing.T) {
	assertStrings := func(t *testing.T, got, want string) {
		t.Helper()

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	assertError := func(t *testing.T, got error, want error) {
		t.Helper()

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	assertDefinition := func(t *testing.T, dictionary Dictionary, got, want string) {
		t.Helper()

		got, err := dictionary.Search("test")
		if err != nil {
			t.Fatal(err)
		}

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	dictionary := Dictionary{"test": "this is just a test"}

	t.Run("Search", func(t *testing.T) {
		dictionary := map[string]string{"test": "this is just a test"}
		got := Search(dictionary, "test")
		want := "this is just a test"
		assertStrings(t, got, want)
	})

	//t.Run("Search with dictionary", func(t *testing.T) {
	//	got := dictionary.Search("test")
	//	want := "this is just a test"
	//	assertStrings(t, got, want)
	//})

	t.Run("known word", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		want := "this is just a test"

		assertStrings(t, got, want)
	})

	t.Run("unknown word", func(t *testing.T) {
		_, got := dictionary.Search("unknown")
		assertError(t, got, ErrNotFound)
	})

	t.Run("Add keyword", func(t *testing.T) {
		dictionary := Dictionary{}
		word := "test"
		definition := "this is just a test"
		err := dictionary.Add(word, definition)

		assertError(t, err, nil)
		assertDefinition(t, dictionary, word, definition)
	})

	t.Run("Existing word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		newDefinition := "new definition"

		dictionary := Dictionary{word: definition}
		err := dictionary.Add(word, definition)

		assertError(t, err, ErrWordExists)
		assertDefinition(t, dictionary, word, definition)

		err = dictionary.Update(word, newDefinition)
		assertError(t, err, nil)
		assertDefinition(t, dictionary, word, newDefinition)
	})

	t.Run("New word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dictionary := Dictionary{}

		err := dictionary.Update(word, definition)
		assertError(t, err, ErrWordDoesNotExist)
	})

	t.Run("delete word", func(t *testing.T) {
		word := "test"
		dictionary := Dictionary{word: "test definition"}

		dictionary.Delete(word)

		_, err := dictionary.Search(word)
		if !errors.Is(err, ErrNotFound) {
			t.Errorf("got %q want ErrNotFound", err)
		}
	})

}

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Chris")

	got := buffer.String()
	want := "Hello, Chris"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

type SpySleeper struct {
	Calls int
}

func (s *SpySleeper) Sleep() {
	s.Calls++
}

type CountdownOperationsSpy struct {
	Calls []string
}

const write = "write"
const sleep = "sleep"

func (s *CountdownOperationsSpy) Sleep() {
	s.Calls = append(s.Calls, sleep)
}

func (s *CountdownOperationsSpy) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	return
}

func TestCountdown(t *testing.T) {
	t.Run("prints 3 to Go!", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		spySleeper := &SpySleeper{}

		Countdown(buffer, spySleeper)

		got := buffer.String()
		want := `3
2
1
GO!`

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}

		if spySleeper.Calls != 4 {
			t.Errorf("got %d calls", spySleeper.Calls)
		}
	})

	t.Run("sleep after every print", func(t *testing.T) {
		spySleepPrinter := &CountdownOperationsSpy{}
		Countdown(spySleepPrinter, spySleepPrinter)

		want := []string{
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
		}

		if !reflect.DeepEqual(want, spySleepPrinter.Calls) {
			t.Errorf("wanted calls %v got %v", want, spySleepPrinter.Calls)
		}
	})
}
