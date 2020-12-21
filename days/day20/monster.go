package day20

const (
	monsterTop    = `..................#.`
	monsterMiddle = `#....##....##....###`
	monsterBottom = `.#..#..#..#..#..#...`
	monsterLength = 15
)

// Monster is a monster.
type Monster struct {
	Top    string
	Middle string
	Bottom string
}
