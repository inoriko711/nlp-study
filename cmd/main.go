package main

import (
	"github.com/inoriko711/nlp-study/unixcommand"
	"github.com/inoriko711/nlp-study/warmup"
)

func main() {
	warmup.Reverse("stressed")
	warmup.Patatokukashi("パタトクカシーー")
	warmup.Patatokukashi2("パトカー", "タクシー")
	warmup.Pi("Now I need a drink, alcoholic of course, after the heavy lectures involving quantum mechanics.")
	warmup.ElementSymbol("Hi He Lied Because Boron Could Not Oxidize Fluorine. New Nations Might Also Sign Peace Security Clause. Arthur King Can.")
	warmup.Ngram("I am an NLPer")
	warmup.Aggregation("paraparaparadise", "paragraph")
	warmup.CreateSentence(12, "気温", 22.4)
	warmup.Cipher("You gods, will give us. Some faults to make us men.")
	warmup.Typoglycemia("I couldn't believe that I could actually understand what I was reading : the phenomenal power of the human mind .")
	unixcommand.CountLines()
	unixcommand.ReplaceTabsWithSpace()
	unixcommand.SaveByColumn()
}
