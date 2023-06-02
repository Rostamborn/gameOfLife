package main

type Point struct {
	X int
	Y int
}

type Cell struct {
	X              int
	Y              int
	State          bool
	LiveNeighbours []Cell
}

func (c *Cell) Die() {
	c.State = false
}

func (c *Cell) BeBorn() {
	c.State = true
}

func (c *Cell) CheckState() {
	if c.State {
		if len(c.LiveNeighbours) < 2 || 3 < len(c.LiveNeighbours) {
			c.Die()
		}
	} else {
		if len(c.LiveNeighbours) == 3 {
			c.BeBorn()
		}
	}
}

func (c *Cell) CheckForNeighbours(cells [][]Cell) {
    c.LiveNeighbours = nil
    
	points := []Point{{c.X, c.Y + 1}, {c.X, c.Y - 1}, {c.X - 1, c.Y},
		{c.X + 1, c.Y}, {c.X - 1, c.Y - 1}, {c.X - 1, c.Y + 1},
		{c.X + 1, c.Y - 1}, {c.X + 1, c.Y + 1},
	}
    for _, p := range points {
        if p.X < 0 || p.Y < 0 || len(cells) <= p.X || len(cells[0]) <= p.Y {
            continue
        }
        if cells[p.X][p.Y].State {
            c.LiveNeighbours = append(c.LiveNeighbours, cells[p.X][p.Y])
        }
    }
}
