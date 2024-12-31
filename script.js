
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


function valid_spot(s, x, y) {
    if (s[0][0] != x) {
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

    if (selected.length == 0) {
	selected.push([x, y])
	div_element.style.backgroundColor = "red"
    }

    if (selected.includes([x, y])) {
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
