package tournament

import (
	"bufio"
	"fmt"
	"io"
	"sort"
	"strings"
)

const testVersion = 4

type Team struct {
	Name string
	MP,
	W,
	D,
	L,
	P uint
}

func (team *Team) RecordWin() {
	team.MP++
	team.W++
	team.P += 3
}

func (team *Team) RecordDraw() {
	team.MP++
	team.D++
	team.P++
}

func (team *Team) RecordLoss() {
	team.MP++
	team.L++
}

func (team Team) String() string {
	return fmt.Sprintf("{%s MP=%d W=%d D=%d L=%d P=%d}", team.Name, team.MP, team.W, team.D, team.L, team.P)
}

type Tournament map[string]*Team

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

func (tournament *Tournament) Team(team string) *Team {
	t := (*tournament)[team]

	if t == nil {
		t = &Team{Name: team}
		(*tournament)[team] = t
	}

	return t
}

func (tournament *Tournament) RecordWin(team string) {
	tournament.Team(team).RecordWin()
}

func (tournament *Tournament) RecordDraw(team string) {
	tournament.Team(team).RecordDraw()
}

func (tournament *Tournament) RecordLoss(team string) {
	tournament.Team(team).RecordLoss()
}

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
