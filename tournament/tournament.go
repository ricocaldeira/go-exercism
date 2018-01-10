package tournament

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"sort"
	"strings"
)

/*
TeamRecord data type keeps information about a team's performance
*/
type TeamRecord struct {
	name string
	wins,
	losses,
	draws,
	points int
}

/*
SortableTeamRecord slice of sortable records
*/
type SortableTeamRecord []*TeamRecord

func (s SortableTeamRecord) Len() int      { return len(s) }
func (s SortableTeamRecord) Swap(i, j int) { s[i], s[j] = s[j], s[i] }
func (s SortableTeamRecord) Less(i, j int) bool {
	if s[i].points == s[j].points {
		return s[i].name < s[j].name
	}
	return s[i].points > s[j].points
}

/*
RecordKeeper contains a map from a teams name to their game record
set up for concurrent access
*/
type RecordKeeper map[string]*TeamRecord

/*
ValidateInput return a list of strings for each game
or an error if the input is invalid
*/
func ValidateInput(input string) ([]*string, error) {
	linesInInput := strings.Split(input, "\n")
	linesToRead := make([]*string, 0, 0)
	for _, line := range linesInInput {
		trimmedLine := strings.Replace(line, " \t\n\r", "", -1)
		if trimmedLine == "" || line[0:1] == "#" {
			continue
		}
		linesToRead = append(linesToRead, &trimmedLine)
	}
	if len(linesToRead) == 0 {
		return nil, errors.New("Invalid blank input")
	}
	for _, line := range linesToRead {
		statements := strings.Split(*line, ";")
		if len(statements) != 3 {
			return nil, errors.New("Invalid line")
		}
		gameResult := statements[2]
		if gameResult != "win" && gameResult != "loss" && gameResult != "draw" {
			return nil, errors.New("Invalid input in line")
		}
	}
	return linesToRead, nil
}

/*
MapLinesFromInputToTournamentRecord populates RecordKeeper with
data based on the results from the input
*/
func MapLinesFromInputToTournamentRecord(record RecordKeeper, lines []*string) {
	for _, line := range lines {
		line := strings.Split(*line, ";")
		team1 := line[0]
		team2 := line[1]
		result := line[2]
		if record[team1] == nil {
			record[team1] = &TeamRecord{name: team1}
		}
		if record[team2] == nil {
			record[team2] = &TeamRecord{name: team2}
		}
		switch result {
		case "win":
			record[team1].wins++
			record[team1].points += 3
			record[team2].losses++
		case "loss":
			record[team1].losses++
			record[team2].wins++
			record[team2].points += 3
		case "draw":
			record[team1].draws++
			record[team1].points++
			record[team2].draws++
			record[team2].points++
		}
	}
}

/*
SortResultsIntoSlice sort maps of team records to a list
sorted by score descending, name ascending is secondary sort
*/
func SortResultsIntoSlice(record RecordKeeper) SortableTeamRecord {
	recordList := make(SortableTeamRecord, 0, len(record))
	for _, record := range record {
		recordList = append(recordList, record)
	}
	sort.Sort(recordList)
	return recordList
}

/*
WriteResultsToString write tournament results to string
*/
func WriteResultsToString(recordList SortableTeamRecord) string {
	output := "Team                           | MP |  W |  D |  L |  P\n"
	for _, teamRecord := range recordList {
		matchesPlayed := teamRecord.wins + teamRecord.losses + teamRecord.draws
		output += fmt.Sprintf(
			"%-31s|%3d |%3d |%3d |%3d |%3d\n",
			teamRecord.name,
			matchesPlayed,
			teamRecord.wins,
			teamRecord.draws,
			teamRecord.losses,
			teamRecord.points,
		)
	}
	return output
}

/*
Tally results from tournament onto score card
*/
func Tally(reader io.Reader, writer io.Writer) error {
	readBuffer := new(bytes.Buffer)
	readBuffer.ReadFrom(reader)
	input := readBuffer.String()
	lines, err := ValidateInput(input)
	if err != nil {
		return err
	}
	record := RecordKeeper{}
	MapLinesFromInputToTournamentRecord(record, lines)
	recordList := SortResultsIntoSlice(record)
	output := WriteResultsToString(recordList)
	writer.Write([]byte(output))
	return nil
}
