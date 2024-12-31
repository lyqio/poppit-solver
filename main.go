package main
import "fmt"

const BOARD_SIZE = 6

type Pair[T, U any] struct {
    first T
    second U
}

type PoppitNode struct {
    winner       int
    player1      bool		// The current player at this node
    position     map[int]int      // The board position at this node e.g. 4 6 nodes, 2 4 nodes
    children     []*PoppitNode	// All child nodes that can be reached in a single move from this node
}

func copyMap(mp map[int]int) map[int]int{
    mp2 := make(map[int]int)

    for id, v := range mp {
	mp2[id] = v 
    }

    return mp2
}

func (p PoppitNode)hash() string {
    return fmt.Sprintf("%t %v", p.player1, p.position)
}

func options_1(node *PoppitNode) PoppitNode {
    new_pos := copyMap(node.position)
    new_pos[1]--

    child := PoppitNode {
	winner:  -1,
	player1: !node.player1,
	position: new_pos,
	children: nil,
    }

    return child
}

func options_2(node *PoppitNode) []PoppitNode {
    new_pos := copyMap(node.position)
    new_pos[2]--
    new_pos[1]++

    child := PoppitNode {
	winner:  -1,
	player1: !node.player1,
	position: new_pos,
	children: nil,
    }

    node.children = append(node.children, &child)

    new_pos2 := copyMap(node.position)
    new_pos2[2]--

    child2 := PoppitNode {
	winner:  -1,
	player1: !node.player1,
	position: new_pos2,
	children: nil,
    }

    return []PoppitNode{child, child2}
}

func options_3(node *PoppitNode) []PoppitNode {
    new_pos := copyMap(node.position)
    new_pos[3]--
    new_pos[2]++

    child1 := PoppitNode {
	winner: -1,
	player1: !node.player1,
	position: new_pos,
	children: nil,
    }

    new_pos2 := copyMap(node.position)
    new_pos2[3]--
    new_pos2[1] += 2

    child2 := PoppitNode {
	winner: -1,
	player1: !node.player1,
	position: new_pos2,
	children: nil,
    }

    new_pos3 := copyMap(node.position)
    new_pos3[3]--
    new_pos3[1]++

    child3 := PoppitNode {
	winner: -1,
	player1: !node.player1,
	position: new_pos3,
	children: nil,
    }

    new_pos4 := copyMap(node.position)
    new_pos4[3]--

    child4 := PoppitNode {
	winner: -1,
	player1: !node.player1,
	position: new_pos4,
	children: nil,
    }

    return []PoppitNode{child1, child2, child3, child4}
}

func generate_children(node *PoppitNode) {
    if node.position[1] > 0 {
	child := options_1(node)
	node.children = append(node.children, &child)
    }

    if node.position[2] > 0 {
	opts := options_2(node)
	for i := 0; i < len(opts); i++ {
	    node.children = append(node.children, &opts[i])
	}
    }

    if node.position[3] > 0 {
	opts := options_3(node)
	for i := 0; i < len(opts); i++ {
	    node.children = append(node.children, &opts[i])
	}
    }

    for i := 0; i < len(node.children); i++ {
	generate_children(node.children[i])
    }
}

func assign_children(node *PoppitNode) {
    // fmt.Print(node.position, node.player1, " ")
    
    // Return if this node has already been explored
    if node.winner != -1 {
	return
    }
    
    // The case we are at a terminal node
    if len(node.children) == 0 {
	// this runs when we are on an empty board
	// thus if we have reached here and it is our turn, we have won
	if node.player1 {
	    node.winner = 1
	} else {
	    node.winner = 2
	}

	// fmt.Println(node.winner)
	return
    }

    // If it is player 1s turn, then go through the children applying the following rules:
    //   1.  If there is at least 1 winning child node, then this node is a win for player 1
    //   2.  If all child nodes are losing nodes, then this node is a loss for player 1
    if node.player1 {
	has_winner := false
	for i := 0; i < len(node.children); i++{
	    assign_children(node.children[i])

	    if node.children[i].winner == 1 {
		has_winner = true
	    }
	}

	if has_winner {
	    node.winner = 1
	} else {
	    node.winner = 2
	}

    } else {
	// We apply the same rules if it isn't our turn, just instead, we do it opposite	
	has_loser := false
	for i := 0; i < len(node.children); i++ {
	    assign_children(node.children[i])

	    if node.children[i].winner == 2 {
		has_loser = true
	    }
	}

	if has_loser {
	    node.winner = 2
	} else {
	    node.winner = 1
	}
    } 

    // fmt.Println(node.winner)
}

func main() {
    start_position := PoppitNode {
	winner:   -1,
	player1:  true,
	position: map[int]int {
	    1: 0,
	    2: 2,
	    3: 0,
	    4: 0,
	    5: 0,
	    6: 0,
	},
	children: nil,
    } 

    known_positions := make(map[string]*PoppitNode)
    known_positions[start_position.hash()] = &start_position
    
    generate_children(&start_position)
    assign_children(&start_position)

    fmt.Println("\n\n\n\n")
    fmt.Printf("Winner is %d\n", start_position.winner)
    
//    for i := 0; i < len(start_position.children); i++ {
//	fmt.Printf("%d | ", start_position.children[i].winner)
//    }
}
