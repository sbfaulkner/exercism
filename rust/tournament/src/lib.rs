use std::collections::HashMap;
use std::fmt;

struct Team {
    name: String,
    matches: u32,
    wins: u32,
    draws: u32,
    losses: u32,
    points: u32,
}

impl fmt::Display for Team {
    fn fmt(&self, f: &mut fmt::Formatter<'_>) -> fmt::Result {
        write!(
            f,
            "{:30} | {:2} | {:2} | {:2} | {:2} | {:2}",
            self.name, self.matches, self.wins, self.draws, self.losses, self.points
        )
    }
}

pub fn tally(match_results: &str) -> String {
    let mut teams = HashMap::new();

    for line in match_results.lines() {
        tally_match(&mut teams, line);
    }

    print_tally(&teams)
}

fn tally_match(teams: &mut HashMap<String, Team>, line: &str) {
    let (team1, team2, result) = parse_line(line);

    match result {
        "win" => {
            tally_team(teams, team1, 1, 0, 0);
            tally_team(teams, team2, 0, 0, 1);
        }
        "draw" => {
            tally_team(teams, team1, 0, 1, 0);
            tally_team(teams, team2, 0, 1, 0);
        }
        "loss" => {
            tally_team(teams, team1, 0, 0, 1);
            tally_team(teams, team2, 1, 0, 0);
        }
        _ => panic!("Invalid result: {}", line),
    }
}

fn parse_line(line: &str) -> (&str, &str, &str) {
    let parts: Vec<&str> = line.split(';').collect();
    (parts[0], parts[1], parts[2])
}

fn tally_team(teams: &mut HashMap<String, Team>, name: &str, wins: u32, draws: u32, losses: u32) {
    teams
        .entry(name.to_string())
        .and_modify(|team| {
            team.matches += 1;
            team.wins += wins;
            team.draws += draws;
            team.losses += losses;
            team.points += points(wins, draws);
        })
        .or_insert(Team {
            name: name.to_string(),
            matches: 1,
            wins: wins,
            draws: draws,
            losses: losses,
            points: points(wins, draws),
        });
}

fn points(wins: u32, draws: u32) -> u32 {
    wins * 3 + draws
}

fn print_tally(teams: &HashMap<String, Team>) -> String {
    let mut teams = teams.values().collect::<Vec<&Team>>();
    teams.sort_by(|a, b| b.points.cmp(&a.points).then(a.name.cmp(&b.name)));

    let mut table = vec!["Team                           | MP |  W |  D |  L |  P".to_string()];
    table.extend(teams.iter().map(|t| t.to_string()));
    table.join("\n")
}
