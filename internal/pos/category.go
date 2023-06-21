package pos

type Category struct {
	Identifier rune
	Fields     []Field
}

func (c Category) IsEmpty() bool {
	return c.Identifier == 0
}

var (
	Adjective    = Category{'A', []Field{Type, Degree, Gender, Nmber, PossessorPerson, PossessorNumber}}
	AdPosition   = Category{'S', []Field{Type}}
	Adverb       = Category{'R', []Field{Type}}
	Conjunction  = Category{'C', []Field{Type}}
	Date         = Category{'W', []Field{}}
	Determiner   = Category{'D', []Field{Type, Person, Gender, Nmber, PossessorNumber}}
	Interjection = Category{'I', []Field{}}
	Noun         = Category{'N', []Field{Type, Gender, Nmber, NounClass, NounSubClass, Degree}}
	Number       = Category{'Z', []Field{Type}}
	Pronoun      = Category{'P', []Field{Type, Person, Gender, Nmber, Case, Polite}}
	Punctuation  = Category{'F', []Field{Type, PunctuationEnclose}}
	Verb         = Category{'V', []Field{Type, Mood, Tense, Person, Nmber, Gender}}
)

func GetCategory(c rune) Category {
	switch c {
	case Adjective.Identifier:
		return Adjective
	case Conjunction.Identifier:
		return Conjunction
	case Determiner.Identifier:
		return Determiner
	case Noun.Identifier:
		return Noun
	case Pronoun.Identifier:
		return Pronoun
	case Adverb.Identifier:
		return Adverb
	case AdPosition.Identifier:
		return AdPosition
	case Verb.Identifier:
		return Verb
	case Number.Identifier:
		return Number
	case Date.Identifier:
		return Date
	case Interjection.Identifier:
		return Interjection
	case Punctuation.Identifier:
		return Punctuation
	default:
		return Category{}
	}
}
