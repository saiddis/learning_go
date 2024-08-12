package main

type PalindromeChecker interface {
	IsPalindrome() bool
	IsWordPalindrome() bool
}

type TextPalindromeChecker struct {
	s string
	Reverser
}

func (tpc TextPalindromeChecker) IsPalindrome() bool {
	reversed := tpc.Reverse()

	if tpc.s == reversed {
		return true
	}

	return false
}

func (tpc TextPalindromeChecker) IsWordPalindrome() bool {
	reversed := tpc.ReverseWords()

	if tpc.s == reversed {
		return true
	}

	return false
}

type Checker struct {
	PalindromeChecker
}

func NewChecker(pc PalindromeChecker) Checker {
	return Checker{
		pc,
	}
}
