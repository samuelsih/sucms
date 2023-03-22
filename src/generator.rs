use crate::config::Config;
use std::{
    io::{BufReader, BufRead}, process::exit,
};
use duct::cmd;

pub fn generate(config: Config) {
    generate_laravel_project(config.get_project_name(), config.want_in_new_folder());
}

fn generate_laravel_project(name: &str, want_in_new_folder: bool) {
    let cmd = if want_in_new_folder {
        cmd!("cmd", "/C", "composer", "create-project", "laravel/laravel", name)
    } else {
        cmd!("cmd", "/C", "composer", "create-project", "laravel/laravel", ".")
    };

    let reader = cmd.stderr_to_stdout().reader().unwrap();
    let lines = BufReader::new(reader).lines();

    for line in lines {
        match line {
            Ok(l) => println!("{}", l),
            Err(e) => {
                println!("{}", e);
                exit(1);
            }
        }
    }
}
