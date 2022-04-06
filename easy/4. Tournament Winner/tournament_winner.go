package main

const HOME_TEAM_WON = 1

func TournamentWinner(competitions [][]string, results []int) string {
	winner_board := make(map[string]int)
	for i := 0; i < len(competitions); i++ {
		// by default we set as the away team the winner
		targetIdx := 1
		if results[i] == HOME_TEAM_WON {
			targetIdx = 0
		}
		curr_winner := competitions[i][targetIdx]
		winner_board[curr_winner] += 3
	}
	winner := ""
	curr_max := 0
	for k, v := range winner_board {
		if v > curr_max {
			curr_max = v
			winner = k
		}
	}
	return winner
}
