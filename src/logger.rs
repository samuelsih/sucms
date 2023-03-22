use console::style;
use std::sync::atomic::{AtomicU8, Ordering};

static LENGTH: AtomicU8 = AtomicU8::new(0);

fn get_current_log_progress() -> u8 {
    return LENGTH.load(Ordering::Relaxed);
}

fn inc_log_progress() {
    let num = get_current_log_progress() + 1;
    LENGTH.store(num, Ordering::Relaxed);
} 

pub fn log(msg: String) {
    inc_log_progress();
    println!("{} {}", style(format!("[{}] ", get_current_log_progress())).bold(), msg);
}