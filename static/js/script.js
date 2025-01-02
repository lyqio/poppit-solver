
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

function remove_leading_zero(s) {
    while (s.length > 0 && s[0] == '0') {
	s.shift()
    }
    return s
}

function get_position() {
    if (document.URL.split('/')[3].length == 0) {
	location.replace("http://localhost:8080/1FFFFFFFFF")
	return [[
	    [1,1,1,1,1,1],
	    [1,1,1,1,1,1],
	    [1,1,1,1,1,1],
	    [1,1,1,1,1,1],
	    [1,1,1,1,1,1],
	    [1,1,1,1,1,1],
	], "1"]
    }

    let url = document.URL.split('/')[3]
    let bin = remove_leading_zero(parseInt(url.toString(), 16).toString(2))

    let player = bin[bin.length-1]

    let brd = []
    let t = []
    for (let i = 0; i < 36; i++) {
	if (i % 6 == 0 && t != []) {
	    brd.push(t)
	    t = []
	}

	t.push(parseInt(bin[i]))
    }
    brd.push(t)
    brd.shift()

    return [brd, player]
}

// draw the grid to the screen
let elements = add_grid(6)
console.log("HERE")
let board = []
let player = "1"
let c = get_position()
board = c[0]
player = c[1]

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
let elems = []
let last_colour = []
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
	elems.push(div_element)
	last_colour.push(div_element.style.backgroundColor)
	div_element.style.backgroundColor = "red"
	return
    }

    console.log("Checking ", x, y)
    if (valid_spot(selected, x, y)) {
	selected.push([x, y])
	elems.push(div_element)
	last_colour.push(div_element.style.backgroundColor)
	div_element.style.backgroundColor = "red"
    } else {
	selected = []
	for (let i = 0; i < elems.length; i++) {
	    elems[i].style.backgroundColor = last_colour[i]
	}
	elems = []
	last_colour = []
    }
}

function change_link() {
    let new_bin = ""
    
    for (let i of board) {
	for (let q of i) {
	    new_bin += q.toString()
	}
    }

    if (player == "1") {
	new_bin += "0"
    }
    else {
	new_bin += "1"
    }

    let link = remove_leading_zero(parseInt(new_bin.toString(), 2).toString(16))
    location.replace(`http://localhost:8080/${link}`)
}

function submit() {
    for (let s of selected) {
	board[s[0]][s[1]] = 0
    }

    change_link()
}

for (let i = 0; i < elements.length; i++) {
    for (let q = 0; q < elements[i].length; q++) {
	elements[i][q].addEventListener("click", handle_click, true)
    }
}
