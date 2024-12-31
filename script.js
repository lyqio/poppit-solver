
function add_grid(n) {
    let elements = []
    id = 0
    for (let i = 0; i < n; i++) {
	let row = document.createElement("div")
	let c = []
	row.className = `container${i+1}`
	for (let q = 0; q < n; q++) {
	    const div = document.createElement("div")
	    div.id = `${id}`
	    div.className = `circle${i+1}`
	    row.appendChild(div)
	    c.push(div)

	    id++
	}

	elements.push(c)
	document.body.appendChild(row)
    }

    return elements
}

// draw the grid to the screen
let elements = add_grid(6)
let board = [
    [1, 1, 0, 1, 1, 1],
    [1, 1, 0, 1, 1, 1],
    [1, 1, 0, 1, 1, 1],
    [1, 1, 0, 1, 1, 1],
    [1, 1, 0, 1, 1, 1],
    [1, 1, 0, 1, 1, 1]
]

function update_board(elem, board) {
    for (let i of elem) {
	for (let element of i) {
	    pos = Number(element.id)
	    var x = Math.floor(pos/6.0)
	    var y = pos % 6

	    if (board[x][y] != 1) {
		element.className = `d${element.className}`
	    }
	}
    }
}

update_board(elements, board)

function has(s, item) {
    for (let i of s) {
	all_match = true
	for (let q = 0; q < i.length; q++) {
	    if (i[q] != item[q]) {
		all_match = false
	    }
	}

	if (all_match) {
	    return true
	}
    }
    return false
}

function valid_spot(s, x, y) {
    if (s[0][0] != x) {
	return false
    }

    if (s.length == 3) {
	return false
    }

    let y_values = []
    for (let item of s) {
	y_values.push(item[1])
    }
    y_values.sort()

    if (y_values[0] != y+1 && y_values[y_values.length-1] != y-1) {
	return false
    }

    return true
}

let selected = [] 
function handle_click(e) {
    div_element = e.target
    pos = Number(div_element.id)

    var x = Math.floor(pos/6.0)
    var y = pos % 6

    if (has(selected, [x, y])) {
	return
    }

    if (board[x][y] != 1) {
	return
    }

    if (selected.length == 0) {
	selected.push([x, y])
	div_element.style.backgroundColor = "red"
	return
    }

    console.log("Checking ", x, y)
    if (valid_spot(selected, x, y)) {
	selected.push([x, y])
	div_element.style.backgroundColor = "red"
    }
}

for (let i = 0; i < elements.length; i++) {
    for (let q = 0; q < elements[i].length; q++) {
	elements[i][q].addEventListener("click", handle_click, true)
    }
}
