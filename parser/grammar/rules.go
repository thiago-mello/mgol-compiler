package grammar

import "strings"

var rules = [38]string{
	"P -> inicio V A",
	"V -> varinicio LV",
	"LV -> D LV",
	"LV -> varfim pt_v",
	"D -> TIPO L pt_v",
	"L -> id vir L",
	"L -> id",
	"TIPO -> inteiro",
	"TIPO -> real",
	"TIPO -> literal",
	"A -> ES A",
	"ES -> leia id pt_v",
	"ES -> escreva ARG pt_v",
	"ARG -> lit",
	"ARG -> num",
	"ARG -> id",
	"A -> CMD A",
	"CMD -> id rcb LD pt_v",
	"LD -> OPRD opm OPRD",
	"LD -> OPRD",
	"OPRD -> id",
	"OPRD -> num",
	"A -> COND A",
	"COND -> CAB CP",
	"CAB -> se ab_p EXP_R fc_p entao",
	"EXP_R -> OPRD opr OPRD",
	"CP -> ES CP",
	"CP -> CMD CP",
	"CP -> COND CP",
	"CP -> fimse",
	"A -> R A",
	"R  -> CABR CPR",
	"CABR  -> repita ab_p EXP_R fc_p",
	"CPR -> ES CPR",
	"CPR -> CMD CPR",
	"CPR -> COND CPR",
	"CPR -> fimrepita",
	"A -> fim",
}

type GrammarRule struct {
	Rule  string
	Left  string
	Right string
}

func GetRule(index int) GrammarRule {
	rule := rules[index]

	ruleSplit := strings.Split(rule, " -> ")

	return GrammarRule{
		Rule:  rule,
		Left:  ruleSplit[0],
		Right: ruleSplit[1],
	}
}
