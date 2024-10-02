package model

type SSScore struct {
	Score int8
	Shard [3][2]int8
}

type FFScore struct {
	Score  int8
	Cards  map[Card]bool
	FHouse House
}

func GetNewSSScore() SSScore {
	ssscore := SSScore{
		Score: 1,
		Shard: [3][2]int8{{-1, -1}, {-1, -1}, {-1, -1}},
	}
	return ssscore
}
