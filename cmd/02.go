package cmd

import "bufio"

func day02(input *bufio.Reader) error {
	calcScore := func(you, foe byte) uint {
		score := uint(you + 1)
		switch (3 + you - foe) % 3 {
		case 0:
			score += 3
		case 1:
			score += 6
		}
		return score
	}

	var score1, score2 uint
	s := bufio.NewScanner(input)
	for s.Scan() {
		str := s.Text()
		foe := str[0] - 'A'
		you := str[2] - 'X'
		score1 += calcScore(you, foe)

		you = (foe + ((you + 2) % 3)) % 3
		score2 += calcScore(you, foe)
	}

	partOne(score1)
	partTwo(score2)

	return nil
}
