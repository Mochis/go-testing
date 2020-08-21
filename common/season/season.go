package season

type Season int
const (
	Spring Season = iota
	Summer
	Autumn
	Winter
	almostSpring  // this constant is not exported,
	              // because first letter is not upper case
)

