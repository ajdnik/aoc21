package day10

const EmptyStack = -1
const EmptyRune = rune('E')

func Push(stack []rune, itm rune) []rune {
	return append(stack, itm)
}

func Pop(stack []rune) ([]rune, rune, bool) {
	sz := len(stack) - 1
	if sz == EmptyStack {
		return nil, EmptyRune, false
	}
	return stack[:sz], stack[sz], true
}

func ParseLine(line string) ([]rune, rune) {
	stack := []rune{}
	for _, char := range line {
		if char == '(' || char == '[' || char == '{' || char == '<' {
			stack = Push(stack, char)
			continue
		}
		var poped rune
		stack, poped, _ = Pop(stack)
		if poped == '(' && char != ')' {
			return stack, char
		}
		if poped == '[' && char != ']' {
			return stack, char
		}
		if poped == '{' && char != '}' {
			return stack, char
		}
		if poped == '<' && char != '>' {
			return stack, char
		}
	}
	return stack, EmptyRune

}
