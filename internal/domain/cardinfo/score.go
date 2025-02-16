package cardinfo

type SSScore struct {
	Score int8
	Shard [3][2]int8
}

type FFScore struct {
	Score  int8
	Cards  map[Card]bool
	FHouse House
}

type PPScore struct {
	Score      int8
	CardValues map[Value]int8
}

func GetNewSSScore() SSScore {
	ssscore := SSScore{
		Score: 1,
		Shard: [3][2]int8{{-1, -1}, {-1, -1}, {-1, -1}},
	}
	return ssscore
}

func GetNewFFScore() FFScore {
	return FFScore{
		Score: 1,
		Cards: map[Card]bool{},
	}
}

func GetNewPPScore() PPScore {
	ppscore := PPScore{
		Score:      1,
		CardValues: map[Value]int8{},
	}
	return ppscore
}
