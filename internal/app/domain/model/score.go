package model

type SSScore struct {
	Score int8
	Shard [3][2]int8
}

func GetNewSSScore() SSScore {
	ssscore := SSScore{
		Score: 1,
		Shard: [3][2]int8{{-1, -1}, {-1, -1}, {-1, -1}},
	}
	return ssscore
}
