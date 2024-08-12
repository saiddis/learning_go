package main

import "fmt"

func main() {
	fmt.Println("---Ex1---")
	str := MyString{
		s: "hello world",
	}
	sp := NewString(str)
	fmt.Println(sp.Length())
	fmt.Println(sp.WordCount())

	fmt.Println("---Ex2---")
	myfmtr := MyFormatter{
		s: "whats up",
	}

	fmtr := NewFormatter(myfmtr)
	fmt.Println(fmtr.ToUpper())

	fmt.Println("---Ex3---")
	var n = 3
	ip := IntPointer{
		n: &n,
	}
	p := NewPointer(ip)
	p.Decrement()
	p.Decrement()
	p.Increment()
	fmt.Println(*ip.n)

	fmt.Println("---Ex4---")
	tc := TextCleaner{
		s: " h e l l o ",
	}
	t := NewTrimmer(tc)
	fmt.Println(t.TrimSpaces())
	fmt.Println(t.RemoveSpaces())

	fmt.Println("---Ex5---")
	sj := StringJoiner{
		s: [2]string{"hello", "world"},
	}
	ah := NewArrayHandler(sj)
	fmt.Println(ah.Concat())
	fmt.Println(ah.Join(","))

	fmt.Println("---Ex6---")
	te := TextExtractor{
		s: "hello world",
	}
	e := NewExtractor(te)
	fmt.Println(e.Substr(4, 4))
	fmt.Println(e.Suffix(5))

	fmt.Println("---Ex7---")
	tr := TextReverser{
		s: "reverse me",
	}
	r := NewReverser(tr)
	fmt.Println(r.Reverse())
	fmt.Println(r.ReverseWords())

	fmt.Println("---Ex8---")
	ti := TextInserter{
		s: "hello",
	}
	i := NewInserter(ti)
	fmt.Println(i.InsertAfterWord("hello", "world"))
	fmt.Println(i.InsertAt(4, "go"))

	fmt.Println("---Ex9---")
	tpc := TextPalindromeChecker{
		"yoy",
		r,
	}
	ch := NewChecker(tpc)
	fmt.Println(ch.IsPalindrome())
	fmt.Println(ch.IsWordPalindrome())

	fmt.Println("---Ex10---")
	tdr := TextDuplicateRemover{
		"hello go hello world",
	}
	dh := NewDuplicateHandler(tdr)
	fmt.Println(dh.RemoveDuplicates())
	fmt.Println(dh.RemoveDuplicatesCaseInsensitive())

	fmt.Println("---Ex11---")
	var balance float64 = 100
	mb := MonthlyBudget{
		balance: &balance,
	}
	bh := NewBudgetHandler(mb)
	bh.AddExpense(1)
	bh.AddIncome(3)
	fmt.Println(*mb.balance)

	fmt.Println("---Ex12---")
	ex := Exchange{
		coefToUSD: 0.11,
		coefToEUR: 0.09,
	}
	mex := NewMoneyExchanger(ex)
	fmt.Println(mex.ToEUR(1000))
	fmt.Println(mex.ToUSD(1000))

	fmt.Println("---Ex13---")
	balance = 1000
	a := Account{
		balance: &balance,
	}
	ach := NewAccountHandler(a)
	ach.Deposit(120)
	ach.WithDraw(400)
	fmt.Println(*a.balance)

	fmt.Println("---Ex14---")
	stc := SimpleTaxCalculator{
		taxPercent: 13,
		vatPercent: 4,
	}
	th := NewTaxHandler(stc)
	fmt.Println(th.CalculateIncomeTax(100000))
	fmt.Println(th.CalculateVAT(100000))
}
