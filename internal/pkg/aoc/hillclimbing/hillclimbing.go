package hillclimbing

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

func PartOne(input io.Reader) (string, error) {

	hmap, startPos, endPos := loadHeightMap(input)
	sb := newScoreboard(hmap.width, hmap.height, endPos, 1000)
	sb.start()
	defer sb.stop()

	pf := newPathFinder(hmap, sb, startPos, endPos)
	result := pf.start()

	shortestPathLength := <-result

	fmt.Println(sb.String())

	return strconv.FormatInt(int64(shortestPathLength), 10), nil
}

func PartTwo(input io.Reader) (string, error) {

	hmap, _, endPos := loadHeightMap(input)
	startPositions := hmap.getLowestPositions()

	shortestPathLength := 1000
	var leaderBoard *scoreboard

	for _, startPos := range startPositions {
		sb := newScoreboard(hmap.width, hmap.height, endPos, shortestPathLength)
		sb.start()

		pf := newPathFinder(hmap, sb, startPos, endPos)
		result := pf.start()

		r := <-result
		if r < shortestPathLength {
			shortestPathLength = r
			leaderBoard = sb
		}
		sb.stop()
	}

	fmt.Println(leaderBoard.String())

	return strconv.FormatInt(int64(shortestPathLength), 10), nil
}

func loadHeightMap(input io.Reader) (*heightmap, *position, *position) {
	scanner := bufio.NewScanner(input)
	scanner.Scan()
	row := []byte(scanner.Text())

	hmap := &heightmap{
		width:  len(row),
		height: 1,
		rows:   make([][]byte, 0, len(row)),
	}

	hmap.rows = append(hmap.rows, row)

	for scanner.Scan() {
		hmap.rows = append(hmap.rows, []byte(scanner.Text()))
		hmap.height++
	}

	hmap.positions = make([][]*position, hmap.height)

	for y := 0; y < hmap.height; y++ {
		hmap.positions[y] = make([]*position, hmap.width)

		for x := 0; x < hmap.width; x++ {
			hmap.positions[y][x] = newPos(x, y)
		}
	}

	var startPos *position
	var endPos *position

	for y := 0; y < hmap.height; y++ {
		for x := 0; x < hmap.width; x++ {
			if hmap.rows[y][x] == 'S' {
				hmap.rows[y][x] = 'a'
				startPos = hmap.positions[y][x]
				if endPos != nil {
					break
				}
			} else if hmap.rows[y][x] == 'E' {
				hmap.rows[y][x] = 'z'
				endPos = hmap.positions[y][x]
				if startPos != nil {
					break
				}
			}
		}
	}

	return hmap, startPos, endPos
}

type pathFinder struct {
	startPos *position
	endPos   *position
	visited  map[uint64]bool
	trail    []*position

	hmap *heightmap
	sb   *scoreboard

	result chan int
}

func newPathFinder(hmap *heightmap, sb *scoreboard, startPos, endPos *position, trail ...*position) *pathFinder {
	pf := &pathFinder{
		startPos: startPos,
		endPos:   endPos,
		visited:  map[uint64]bool{startPos.key: true},
		hmap:     hmap,
		sb:       sb,
		result:   make(chan int),
	}

	trailLength := len(trail)
	pf.trail = make([]*position, trailLength+1)

	if len(trail) > 0 {
		for step, p := range trail {
			pf.visited[p.key] = true
			pf.trail[step] = p
		}
	}

	pf.trail[trailLength] = startPos

	return pf
}

func (pf *pathFinder) start() <-chan int {
	go pf.run()
	return pf.result
}

func (pf *pathFinder) run() {
	finders := []*pathFinder{}

	result := 10000

	pf.hmap.forAllPossibleMovesFrom(pf.startPos, func(p *position) {
		if _, ok := pf.visited[p.key]; !ok {
			pf.visited[p.key] = true

			if pf.sb.test(p, len(pf.trail)) {
				if p.key == pf.endPos.key {
					r := len(pf.trail)
					if r < result {
						result = r
					}
					return
				}

				finders = append(finders, newPathFinder(pf.hmap, pf.sb, p, pf.endPos, pf.trail...))
			}
		}
	})

	chans := make([]<-chan int, len(finders))

	for idx, f := range finders {
		chans[idx] = f.start()
	}

	for _, c := range chans {
		r := <-c
		if r < result {
			result = r
		}
	}

	pf.result <- result
}

type heightmap struct {
	width     int
	height    int
	rows      [][]byte
	positions [][]*position
}

func (hm *heightmap) adjacent(pos *position, xoff, yoff int) *position {

	isOK := func(x, y int) bool {
		if x < 0 || x >= hm.width || y < 0 || y >= hm.height {
			return false
		}

		return true
	}

	if !isOK(int(pos.x)+xoff, int(pos.y)+yoff) {
		return nil
	}

	return hm.positions[int(pos.y)+yoff][int(pos.x)+xoff]
}

func (hm *heightmap) forAllPossibleMovesFrom(startPos *position, try func(p *position)) {
	if adj := hm.adjacent(startPos, -1, 0); hm.isValidMove(startPos, adj) {
		try(adj)
	}
	if adj := hm.adjacent(startPos, 0, -1); hm.isValidMove(startPos, adj) {
		try(adj)
	}
	if adj := hm.adjacent(startPos, 1, 0); hm.isValidMove(startPos, adj) {
		try(adj)
	}
	if adj := hm.adjacent(startPos, 0, 1); hm.isValidMove(startPos, adj) {
		try(adj)
	}
}

func (hm *heightmap) getLowestPositions() []*position {
	lowest := []*position{}
	for y := 0; y < hm.height; y++ {
		for x := 0; x < hm.width; x++ {
			if hm.rows[y][x] == 'a' {
				lowest = append(lowest, hm.positions[y][x])
			}
		}
	}
	return lowest
}

func (hm *heightmap) isValidMove(from, to *position) bool {
	if to == nil {
		return false
	}

	startHeight := int(hm.rows[from.y][from.x])
	endHeight := int(hm.rows[to.y][to.x])

	// Can the destination really be lower than the start?
	return endHeight < (startHeight + 2)
}

type scoreboard struct {
	width        int
	height       int
	endPos       *position
	trailLenghts map[uint64]int

	queue chan func()
	exit  bool
}

func newScoreboard(width, height int, endPos *position, maxPathLength int) *scoreboard {
	board := &scoreboard{
		width:        width,
		height:       height,
		endPos:       endPos,
		trailLenghts: map[uint64]int{endPos.key: maxPathLength},
		exit:         false,
		queue:        make(chan func()),
	}
	return board
}

func (sb *scoreboard) run() {
	// repeat until the queue is closed
	for fn := range sb.queue {
		if fn == nil {
			return
		}

		fn()
	}
}

func (sb *scoreboard) start() {
	go sb.run()
}

func (sb *scoreboard) stop() {
	// Create a result channel so that we can wait for completion
	resultChan := make(chan bool)

	sb.queue <- func() {
		// close the queue to signal the consumers that we are going out of business
		close(sb.queue)
		resultChan <- true
	}

	// blocking read until our action has been processed
	<-resultChan
}

func (sb *scoreboard) String() string {
	results := make([][]string, sb.height)
	for y := 0; y < sb.height; y++ {
		results[y] = make([]string, sb.width)
		for x := 0; x < sb.width; x++ {
			results[y][x] = "     "
		}
	}

	for key, score := range sb.trailLenghts {
		x := (key >> 32)
		y := (key << 32) >> 32

		results[y][x] = fmt.Sprintf("[%3d]", score)
	}

	output := ""

	for _, row := range results {
		output += (strings.Join(row, "") + "\n")
	}

	return output
}

func (sb *scoreboard) test(p *position, trailLength int) bool {
	// Create a result channel so that we can wait for completion
	resultChan := make(chan bool)

	sb.queue <- func() {
		result := false

		if sb.trailLenghts[sb.endPos.key] > trailLength {
			score, ok := sb.trailLenghts[p.key]
			if !ok {
				result = true
				sb.trailLenghts[p.key] = trailLength
			} else if score > trailLength {
				sb.trailLenghts[p.key] = trailLength
				result = true
			}
		}

		resultChan <- result
	}

	// blocking read until our action has been processed
	return <-resultChan
}

type position struct {
	x   uint32
	y   uint32
	key uint64
}

func newPos(x, y int) *position {
	p := &position{
		x:   uint32(x),
		y:   uint32(y),
		key: uint64(x),
	}

	p.key = (p.key << 32) | uint64(y)

	return p
}
