use std::fmt::{Display, Result, Formatter};
use console::{Emoji, style};

static LOOKING_GLASS_EMOJI: Emoji<'_, '_> = Emoji("🔍  ", "");
static TRUCK_EMOJI: Emoji<'_, '_> = Emoji("🚚  ", "");
static CLIP_EMOJI: Emoji<'_, '_> = Emoji("🔗  ", "");
static PAPER_EMOJI: Emoji<'_, '_> = Emoji("📃  ", "");
static SPARKLE_EMOJI: Emoji<'_, '_> = Emoji("✨ ", "");
static ERROR_EMOJI: Emoji<'_, '_> = Emoji("❌  ", "");
#[derive(Clone, Copy)]
pub enum Icons {
    LookingGlass,
    TruckEmoji,
    Clip,
    Paper,
    Sparkle,
    Error
}

impl Display for Icons {
    fn fmt(&self, f: &mut Formatter<'_>) -> Result {
        match *self {
            Icons::LookingGlass => write!(f, "{}", LOOKING_GLASS_EMOJI),
            Icons::TruckEmoji => write!(f, "{}", TRUCK_EMOJI),
            Icons::Clip => write!(f, "{}", CLIP_EMOJI),
            Icons::Paper => write!(f, "{}", PAPER_EMOJI),
            Icons::Sparkle => write!(f, "{}", SPARKLE_EMOJI),
            Icons::Error => write!(f, "{}", ERROR_EMOJI),
        }
    }
}

pub struct Logger {
    length: usize,
}

impl Logger {
    pub fn new() -> Logger {
        Logger { length: 0 }
    }

    pub fn log(&mut self, msg: String, icon: Icons) {
        self.length += 1;
        println!("{} {}{}", icon, style(format!("[{}] ", self.length)).bold(), msg);
    }   

    pub fn log_err(&mut self, msg: String) {
        self.length += 1;
        println!("{} {}", Icons::Error, msg);
    }
}