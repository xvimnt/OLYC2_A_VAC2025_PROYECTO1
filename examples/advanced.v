module collections

// Generic stack implementation
pub struct Stack<T> {
    mut items []T
}

pub fn new_stack<T>() Stack<T> {
    return Stack<T>{
        items: []T{}
    }
}

pub fn (mut s Stack<T>) push(item T) {
    s.items << item
}

pub fn (mut s Stack<T>) pop() ?T {
    if s.items.len == 0 {
        return none
    }
    item := s.items[s.items.len - 1]
    s.items = s.items[..s.items.len - 1]
    return item
}

// Enum example
pub enum Color {
    red
    green
    blue
    custom_rgb = 100
}

// Interface example
pub interface Drawable {
    draw() string
}

// Match statement example
pub fn color_to_string(c Color) string {
    match c {
        .red => return "red"
        .green => return "green"
        .blue => return "blue"
        else => return "custom"
    }
}
