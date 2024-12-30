package main
import "fmt"

const BOARD_SIZE = 6

type Pair[T, U any] struct {
    first T
    second U
}

type PoppitNode struct {
    player1 bool		// The current player at this node
    position map[int]int      // The board position at this node e.g. 4 6 nodes, 2 4 nodes
    children []*PoppitNode	// All child nodes that can be reached in a single move from this node
}

func (p PoppitNode)hash() string {
    return fmt.Sprintf("%t %v", p.player1, p.position)
}

func generate_children(node *PoppitNode) {
    if node.position[1] > 0 {
	new_pos := node.position
	new_pos[1]--

	child := PoppitNode {
	    player1: !node.player1,
	    position: new_pos,
	    children: nil,
	}

	node.children = append(node.children, &child)

    } else if node.position[2] > 0 {
	new_pos := node.position
	new_pos[2]--
	new_pos[1]++

	child := PoppitNode {
	    player1: !node.player1,
	    position: new_pos,
	    children: nil,
	}

	node.children = append(node.children, &child)

	new_pos2 := node.position
	new_pos2[2]--

	child2 := PoppitNode {
	    player1: !node.player1,
	    position: new_pos2,
	    children: nil,
	}

	node.children = append(node.children, &child2)
    }

    for i := 0; i < len(node.children); i++ {
	generate_children(node.children[i])
    }
}

func main() {
    start_position := PoppitNode {
	player1: true,
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

    fmt.Println(start_position)
    fmt.Println(*start_position.children[0], *start_position.children[1])
}
