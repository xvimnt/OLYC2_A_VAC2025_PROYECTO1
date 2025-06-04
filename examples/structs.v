module shapes

pub struct Point {
    pub mut x int
    pub mut y int
}

pub struct Rectangle {
    pub origin Point
    pub width int
    pub height int
}

pub fn (r Rectangle) area() int {
    return r.width * r.height
}

pub fn (mut p Point) move(dx int, dy int) {
    p.x += dx
    p.y += dy
}
