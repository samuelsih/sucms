use console::style;

pub struct Logger {
    length: usize,
}

impl Logger {
    pub fn new() -> Logger {
        Logger { length: 0 }
    }

    pub fn log(&mut self, msg: String) {
        self.length += 1;
        println!("{} {}", style(format!("[{}] ", self.length)).bold(), msg);
    }   
}