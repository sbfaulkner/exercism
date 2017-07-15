package tournament

import (
	"bufio"
	"fmt"
	"io"
	"sort"
	"strings"
)

const testVersion = 4

// Team holds a team name and their statistics for a Tournament
type Team struct {
	Name string
	MP,
	W,
	D,
	L,
	P uint
}

// RecordWin updates a Team's statistics for a win
func (team *Team) RecordWin() {
	team.MP++
	team.W++
	team.P += 3
}

// RecordDraw updates a Team's statistics for a draw
func (team *Team) RecordDraw() {
	team.MP++
	team.D++
	team.P++
}

// RecordLoss updates a Team's statistics for a loss
func (team *Team) RecordLoss() {
	team.MP++
	team.L++
}

// Stringify a Team
func (team Team) String() string {
	return fmt.Sprintf("{%s MP=%d W=%d D=%d L=%d P=%d}", team.Name, team.MP, team.W, team.D, team.L, team.P)
}

// Tournament represents a collection of Teams participating in a tournament
type Tournament map[string]*Team

// Tally reads match results and writes standings
func Tally(r io.Reader, w io.Writer) error {
	reader := bufio.NewScanner(r)

	tournament := Tournament{}

	for reader.Scan() {
		l := reader.Text()
		if l == "" || l[0] == '#' {
			continue
		}

		f := strings.Split(l, ";")
		if len(f) != 3 {
			return fmt.Errorf("tournament: invalid line - %q", l)
		}

		switch f[2] {
		case "win":
			tournament.RecordWin(f[0])
			tournament.RecordLoss(f[1])
		case "loss":
			tournament.RecordLoss(f[0])
			tournament.RecordWin(f[1])
		case "draw":
			tournament.RecordDraw(f[0])
			tournament.RecordDraw(f[1])
		default:
			return fmt.Errorf("tournament: invalid result - %q", f[2])
		}
	}

	o := bufio.NewWriter(w)

	o.WriteString(fmt.Sprintf("%-30s | %2s | %2s | %2s | %2s | %2s\n", "Team", "MP", "W", "D", "L", "P"))

	for _, t := range tournament.Standings() {
		o.WriteString(fmt.Sprintf("%-30s | %2d | %2d | %2d | %2d | %2d\n", t.Name, t.MP, t.W, t.D, t.L, t.P))
	}

	o.Flush()

	return nil
}

// Team finds a team within a tournament, adding it if necessary
func (tournament *Tournament) Team(team string) *Team {
	t := (*tournament)[team]

	if t == nil {
		t = &Team{Name: team}
		(*tournament)[team] = t
	}

	return t
}

// RecordWin records stats for a win by the team with the specified name
func (tournament *Tournament) RecordWin(team string) {
	tournament.Team(team).RecordWin()
}

// RecordDraw records stats for a draw by the team with the specified name
func (tournament *Tournament) RecordDraw(team string) {
	tournament.Team(team).RecordDraw()
}

// RecordLoss records stats for a loss by the team with the specified name
func (tournament *Tournament) RecordLoss(team string) {
	tournament.Team(team).RecordLoss()
}

// Standings returns a slice of the Teams in the tournament sorted by their points
func (tournament Tournament) Standings() []*Team {
	standings := make([]*Team, 0, len(tournament))

	for _, t := range tournament {
		standings = append(standings, t)
	}

	sort.Slice(standings, func(i, j int) bool {
		ti := standings[i]
		tj := standings[j]

		return ti.P > tj.P || ti.P == tj.P && ti.Name < tj.Name
	})

	return standings
}
