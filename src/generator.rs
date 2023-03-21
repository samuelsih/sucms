use crate::config::Config;
use std::{
    io::{BufReader, BufRead},
};
use duct::cmd;

pub fn generate(config: Config) {
    generate_laravel_project(config.get_project_name(), config.want_in_new_folder());
}

fn generate_laravel_project(name: &str, want_in_new_folder: bool) {
    if want_in_new_folder {
        let cmd = cmd!("cmd", "/C", "composer", "create-project", "laravel/laravel", name);
        let reader = cmd.stderr_to_stdout().reader().unwrap();
        let lines = BufReader::new(reader).lines();
    
        for line in lines {
            println!("{}", line.unwrap());
        }

        return;
    }

    let cmd = cmd!("cmd", "/C", "composer", "create-project", "laravel/laravel", ".");
    let reader = cmd.stderr_to_stdout().reader().unwrap();
    let lines = BufReader::new(reader).lines();

    for line in lines {
        println!("{}", line.unwrap());
    }
}
