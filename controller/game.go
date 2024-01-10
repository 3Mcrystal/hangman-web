package hangmanweb

var Accents = map[string]string{
	"a": "aàâäÀÁÂÄÆ",
	"e": "eéèêëÉÈÊË",
	"i": "iïîÍÌÎÏ",
	"o": "oôöœÓÒÔÖØ",
	"u": "uùüûÚÙÛÜ",
	"c": "cçÇ",
}

type GameState struct {
	Total       []rune
	Current     []rune
	NewLetters  []int
	ErrorCount  int
	UsedLetters []string
	Tries       int
	Difficulty  string
}

func NewGame(seq string) *GameState {
	total := []rune{}
	for _, char := range seq {
		if char == 13 || char == 10 {
			continue
		}
		total = append(total, char)
	}
	res := GameState{Total: []rune(total)}

	for range res.Total {
		res.Current = append(res.Current, '_')
	}
	for j := 0; j < len(res.Total)/2-1; j++ {
		index := RandInt(len(total) - (j + 1))
		for i, char := range res.Total {
			if index == i && i < len(res.Current) {
				if res.Current[i] != '_' {
					index++
				} else {
					res.Current[i] = char
				}
			}
		}
	}
	res.ErrorCount = 0
	res.Tries = 0
	return &res
}

func (g *GameState) IsFinish() bool {
	if g.ErrorCount >= 10 {
		return true
	}
	for _, char := range g.Current {
		if char == '_' {
			return false
		}
	}
	return true
}

func (g *GameState) CompleteWord() {
	for index, el := range g.Total {
		g.Current[index] = el
	}
}

func (g *GameState) AddLetter(seq string) bool {
	valid := false
	for i, char := range g.Total {
		_, exists := Accents[seq]
		if exists {
			for _, accent := range Accents[seq] {
				if char == accent {
					if g.Current[i] == '_' {
						g.NewLetters = append(g.NewLetters, i)
						g.Current[i] = char
						valid = true
					}
				}
			}
		} else {
			if string(char) == seq && g.Current[i] == '_' {
				g.NewLetters = append(g.NewLetters, i)
				g.Current[i] = rune(seq[0])
				valid = true
			}
		}
	}
	return valid
}
