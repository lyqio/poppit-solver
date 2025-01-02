package main
import "fmt"
import "reflect"
import "os"
import "bufio"
import "strings"
import "strconv"

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

func make_node(node *PoppitNode, pos map[int]int) PoppitNode {
    return PoppitNode {
	winner: -1,
	player1: !node.player1,
	position: pos,
	children: nil,
    }
}

func options_1(node *PoppitNode) PoppitNode {
    new_pos := copyMap(node.position)
    new_pos[1]--

    child := make_node(node, new_pos)
    return child
}

func options_2(node *PoppitNode) []PoppitNode {
    new_pos := copyMap(node.position)
    new_pos[2]--
    new_pos[1]++
    child := make_node(node, new_pos)

    new_pos2 := copyMap(node.position)
    new_pos2[2]--
    child2 := make_node(node, new_pos2)

    return []PoppitNode{child, child2}
}

func options_3(node *PoppitNode) []PoppitNode {
    new_pos := copyMap(node.position)
    new_pos[3]--
    new_pos[2]++
    child1 := make_node(node, new_pos)

    new_pos2 := copyMap(node.position)
    new_pos2[3]--
    new_pos2[1] += 2
    child2 := make_node(node, new_pos2)

    new_pos3 := copyMap(node.position)
    new_pos3[3]--
    new_pos3[1]++
    child3 := make_node(node, new_pos3)

    new_pos4 := copyMap(node.position)
    new_pos4[3]--
    child4 := make_node(node, new_pos4)

    return []PoppitNode{child1, child2, child3, child4}
}

func options_4(node *PoppitNode) []PoppitNode {
    // 4 to 3
    // 4 to 1, 2

    // 4 to 2
    // 4 to 2 1

    // 4 to 1

    new_pos := copyMap(node.position)
    new_pos[4]--
    new_pos[3]++
    child1 := make_node(node, new_pos)

    new_pos2 := copyMap(node.position)
    new_pos2[4]--
    new_pos2[1]++
    new_pos2[2]++
    child2 := make_node(node, new_pos2)

    new_pos3 := copyMap(node.position)
    new_pos3[4]--
    new_pos3[2]++
    child3 := make_node(node, new_pos3)

    new_pos4 := copyMap(node.position)
    new_pos4[4]--
    new_pos4[1] += 2
    child4 := make_node(node, new_pos4)

    new_pos5 := copyMap(node.position)
    new_pos5[4]--
    new_pos5[1]++
    child5 := make_node(node, new_pos5)

    return []PoppitNode{child1, child2, child3, child4, child5}
}

func options_5(node *PoppitNode) []PoppitNode {
    // 5 to 4
    // 5 to 1, 3
    // 5 to 2x2
    
    // 5 to 3
    // 5 to 1, 2

    // 5 to 2
    // 5 to 2x1

    new_pos := copyMap(node.position)
    new_pos[5]--
    new_pos[4]++
    child1 := make_node(node, new_pos)

    new_pos2 := copyMap(node.position)
    new_pos2[5]--
    new_pos2[1]++
    new_pos2[3]++
    child2 := make_node(node, new_pos2)

    new_pos3 := copyMap(node.position)
    new_pos3[5]--
    new_pos3[2] += 2
    child3 := make_node(node, new_pos3)

    new_pos4 := copyMap(node.position)
    new_pos4[5]--
    new_pos4[3]++
    child4 := make_node(node, new_pos4)

    new_pos5 := copyMap(node.position)
    new_pos5[5]--
    new_pos5[1]++
    new_pos5[2]++
    child5 := make_node(node, new_pos5)

    new_pos6 := copyMap(node.position)
    new_pos6[5]--
    new_pos6[2]++
    child6 := make_node(node, new_pos6)

    new_pos7 := copyMap(node.position)
    new_pos7[5]--
    new_pos7[1] += 2
    child7 := make_node(node, new_pos7)

    return []PoppitNode{child1, child2, child3, child4, child5, child6, child7}
}

func options_6(node *PoppitNode) []PoppitNode {
    // 6 to 5
    // 6 to 4, 1
    // 6 to 3, 2

    // 6 to 4
    // 6 to 3, 1
    // 6 to 2x2

    // 6 to 3
    // 6 to 1, 2

    new_pos := copyMap(node.position)
    new_pos[6]--
    new_pos[5]++
    child1 := make_node(node, new_pos)

    new_pos2 := copyMap(node.position)
    new_pos2[6]--
    new_pos2[4]++
    new_pos2[1]++
    child2 := make_node(node, new_pos2)

    new_pos3 := copyMap(node.position)
    new_pos3[6]--
    new_pos3[3]++
    new_pos3[2]++
    child3 := make_node(node, new_pos3)

    new_pos4 := copyMap(node.position)
    new_pos4[6]--
    new_pos4[4]++
    child4 := make_node(node, new_pos4)

    new_pos5 := copyMap(node.position)
    new_pos5[6]--
    new_pos5[1]++
    new_pos5[3]++
    child5 := make_node(node, new_pos5)

    new_pos6 := copyMap(node.position)
    new_pos6[6]--
    new_pos6[2]++
    child6 := make_node(node, new_pos6)

    new_pos7 := copyMap(node.position)
    new_pos7[6]--
    new_pos7[3]++
    child7 := make_node(node, new_pos7)

    new_pos8 := copyMap(node.position)
    new_pos8[6]--
    new_pos8[1]++
    new_pos8[2]++
    child8 := make_node(node, new_pos8)

    return []PoppitNode{child1, child2, child3, child4, child5, child6, child7, child8}
}

var cache = make(map[string][]*PoppitNode)
func generate_children(node *PoppitNode) []*PoppitNode {
    val, in := cache[node.hash()]
    if in {
	node.children = val
	return val
    }

    children := []*PoppitNode{}

    if node.position[1] > 0 {
	child := options_1(node)
	children = append(children, &child)
    }

    if node.position[2] > 0 {
	opts := options_2(node)
	for i := 0; i < len(opts); i++ {
	    children = append(children, &opts[i])
	}
    }

    if node.position[3] > 0 {
	opts := options_3(node)
	for i := 0; i < len(opts); i++ {
	    children = append(children, &opts[i])
	}
    }

    if node.position[4] > 0 {
	opts := options_4(node)
	for i := 0; i < len(opts); i++ {
	    children = append(children, &opts[i])
	}
    }

    if node.position[5] > 0 {
	opts := options_5(node)
	for i := 0; i < len(opts); i++ {
	    children = append(children, &opts[i])
	}
    }

    if node.position[6] > 0 {
	opts := options_6(node)
	for i := 0; i < len(opts); i++ {
	    children = append(children, &opts[i])
	}

    }

    for i := 0; i < len(children); i++ {
	children[i].children = generate_children(children[i])
    }

    node.children = children
    cache[node.hash()] = children
    return children
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

func count_spots(board [][]int) map[int]int {
    output := map[int]int {
	1: 0,
	2: 0,
	3: 0,
	4: 0,
	5: 0,
	6: 0,
    }

    for i := 0; i < len(board); i++ {
	ln := 0
	for q := 0; q < len(board[i]); q++ {
	    if board[i][q] == 1 {
		ln++
	    } else {
		if ln != 0 {
		    output[ln]++
		    ln = 0
		}
	    }
	}
	if ln > -1 {
	    if ln != 0 {
		output[ln]++
		ln = 0
	    }
	}
    }

    return output
}

func copy2D(s1 [][]int) [][]int {
    o := [][]int{}
    for i := 0; i < len(s1); i++ {
	val := []int{}
	for q := 0; q < len(s1[i]); q++ {
	    val = append(val, s1[i][q])
	}
	o = append(o, val)
    }

    return o
}

func find_move(board [][]int, move PoppitNode) [][]int {
    for i := 0; i < len(board); i++ {
	for q := 0; q < len(board[i]); q++ {
	    if board[i][q] == 1 {
		nb := copy2D(board)

		fmt.Println(board)
		fmt.Println(nb)

		nb[i][q] = 0
		if reflect.DeepEqual(count_spots(nb), move.position){
		    return nb
		}
	    }
	}
    }

    for i := 0; i < len(board); i++ {
	for q := 1; q < len(board[i]); q++ {
	    if board[i][q] == 1 && board[i][q-1] == 1 {
		nb := copy2D(board)

		nb[i][q] = 0
		nb[i][q-1] = 0
		if reflect.DeepEqual(count_spots(nb), move.position){
		    return nb
		}
	    }
	}
    }

    for i := 0; i < len(board); i++ {
	for q := 2; q < len(board[i]); q++ {
	    if board[i][q] == 1 && board[i][q-1] == 1 && board[i][q-2] == 1{
		nb := copy2D(board)

		nb[i][q] = 0
		nb[i][q-1] = 0
		nb[i][q-2] = 0

		if reflect.DeepEqual(count_spots(nb), move.position){
		    return nb
		}
	    }
	}
    }

    return nil
}

func print_board(board [][]int) {
    for i := 0; i < len(board); i++ {
	for q := 0; q < len(board[i]); q++ {
	    fmt.Printf("%d ", board[i][q])
	}
	fmt.Println()
    }
}

func ai_move(pos *PoppitNode, board *[][]int) {
    fmt.Println("BEFORE:", pos.position)
    var move PoppitNode
    for i := 0; i < len(pos.children); i++ {
	if pos.children[i].winner == 1 {
	    move = *pos.children[i]
	}
    }

    fmt.Println("B:", move.position)

    b := find_move(*board, move)
    if b == nil {
	fmt.Println("ERR: Nil reached somehow")
	return
    }

    *board = b 
    *pos = move
    
    if pos.player1 {
	fmt.Println("player 1")
    } else {
	fmt.Println("player 2")
    }

    print_board(*board)
}

func user_move(pos *PoppitNode, board *[][]int) {
    scanner := bufio.NewScanner(os.Stdin) 
    scanner.Scan()
    text := scanner.Text()
    parts := strings.Split(text, " ")

    // input in form [x, y, len]
    x, _ := strconv.Atoi(parts[0])
    y, _ := strconv.Atoi(parts[1])
    ln, _ := strconv.Atoi(parts[2])

    b := *board
    for i := y; i < y+ln; i++ {
	b[x][i] = 0
    }

    *board = b

    state := count_spots(*board)
    var move PoppitNode
    for i := 0; i < len(pos.children); i++ {
 //	fmt.Println("C: ", *pos.children[i], " ", state)
	if reflect.DeepEqual(pos.children[i].position, state) {
	    move = *pos.children[i]
	}
    }
    *pos = move

    fmt.Println("A:", pos.position)

    if pos.player1 {
	fmt.Println("player 1")
    } else {
	fmt.Println("player 2")
    }
    print_board(*board)
}

func convert_board(link string) [][]int {
    bin, err := strconv.ParseUint(link, 16, 64)

    if err != nil {
	fmt.Println("ERROR IN CONVERSION")
    }

    b := strconv.FormatUint(bin, 2) 
    for b[0] == '0' {
	b = b[1:]
    }
    
    var board [][]int
    var t []int
    for i := 0; i < 36; i++ {
	if i % 6 == 0 {
	    board = append(board, t)
	    t = []int{}
	}
	
	a, _ := strconv.Atoi(string(b[i]))
	t = append(t, a)
    }

    board = board[1:]
    //fmt.Print("NEW BOARD: \n")
    print_board(board)
    return board
}

func frontend_user_move(pos *PoppitNode, board *[][]int) {
    for URL == "" {}    
    fmt.Println("RECIEVED: ", URL[1:])
    new_board := convert_board(URL[1:])

    fmt.Println("NEW BOARD: ")
    print_board(new_board)

    if reflect.DeepEqual(new_board, *board) {
	URL = ""
	frontend_user_move(pos, board)
	return
    }

    x := 0
    y := 0
    over := false
    b := *board
    for i := 0; i < len(b); i++ {
	for q := 0; q < len(b[i]); q++ {
	    if b[i][q] != new_board[i][q] {
		x = i
		y = q
		over = true
		break
	    } 
	}
	if over {
	    break
	}
    }

    ln := 1
    for i := y+1; i < len(b[x]) && b[x][i] != 1; i++ {
	ln++
    }

    fmt.Println("!()! ", x, y, ln)

    b = *board
    for i := y; i < y+ln; i++ {
	b[x][i] = 0
    }

    *board = b

    state := count_spots(*board)
    var move PoppitNode
    for i := 0; i < len(pos.children); i++ {
 //	fmt.Println("C: ", *pos.children[i], " ", state)
	if reflect.DeepEqual(pos.children[i].position, state) {
	    move = *pos.children[i]
	}
    }
    *pos = move

    fmt.Println("A:", pos.position)

    if pos.player1 {
	fmt.Println("player 1")
    } else {
	fmt.Println("player 2")
    }
    print_board(*board)

    URL = ""
}

func board_to_hex(board [][]int) string {
    bin := ""
    for i := 0; i < len(board); i++ {
	for q := 0; q < len(board[i]); q++ {
	    bin += string(board[i][q]+'0')
	}
    }
    
    // Add the fact it's player '0's turn although this doesn't matter and isn't used anywhere
    bin += "0"
    b, _ := strconv.ParseUint(bin, 2, 64)
    return fmt.Sprintf("%x", b)
}

func frontend_ai_move(pos *PoppitNode, board *[][]int) {
    fmt.Println("BEFORE:", pos.position)
    var move PoppitNode
    for i := 0; i < len(pos.children); i++ {
	if pos.children[i].winner == 1 {
	    move = *pos.children[i]
	}
    }

    fmt.Println("B:", move.position)

    b := find_move(*board, move)
    if b == nil {
	fmt.Println("ERR: Nil reached somehow")
	return
    }

    *board = b 
    *pos = move
    
    if pos.player1 {
	fmt.Println("player 1")
    } else {
	fmt.Println("player 2")
    }

    fmt.Println("AI BOARD:")
    print_board(*board)
    link := board_to_hex(*board)
    fmt.Println("AI LINK: ", link)
    CHANGE_URL = link
}

func play_game2() {
    board := [][]int{
	{1, 1, 1, 1, 1, 1},
	{1, 1, 1, 1, 1, 1},
	{1, 1, 1, 1, 1, 1},
	{1, 1, 1, 1, 1, 1},
	{1, 1, 1, 1, 1, 1},
	{1, 1, 1, 1, 1, 1},
    }

    pos := PoppitNode {
	winner: -1,
	player1: false,
	position: count_spots(board),
	children: nil,
    }

    fmt.Println("Generating nodes...")
    generate_children(&pos)
    fmt.Println("Assigning nodes...")
    assign_children(&pos)
    fmt.Println("Finished!")

    fmt.Printf("Player %d wins\n", pos.winner)
    print_board(board)

//    for i := 0; i < len(pos.children); i++ {
//	fmt.Printf("%d | ", pos.children[i].winner)
//    }

    for len(pos.children) > 0 {
	frontend_user_move(&pos, &board)
	// user_move(&pos, &board)
	frontend_ai_move(&pos, &board)
	// ai_move(&pos, &board)
    }
}

func play_game() {
    board := [][]int{
	{1, 1, 1, 1, 1, 1},
	{1, 1, 1, 1, 1, 1},
	{1, 1, 1, 1, 1, 1},
	{1, 1, 1, 1, 1, 1},
	{1, 1, 1, 1, 1, 1},
	{1, 1, 1, 1, 1, 1},
    }

    pos := PoppitNode {
	winner: -1,
	player1: true,
	position: count_spots(board),
	children: nil,
    }

    fmt.Println("Generating nodes...")
    generate_children(&pos)
    fmt.Println("Assigning nodes...")
    assign_children(&pos)
    fmt.Println("Finished!")

    fmt.Printf("Player %d wins\n", pos.winner)

//    for i := 0; i < len(pos.children); i++ {
//	fmt.Printf("%d | ", pos.children[i].winner)
//    }

    for len(pos.children) > 0 {
	ai_move(&pos, &board)
	user_move(&pos, &board)
    }
}

func main() {
    go run_app()
    play_game2()

    select{}
}
