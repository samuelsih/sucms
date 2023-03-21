use std::{fs::File};
use std::env;
use std::process::exit;

mod raw_yaml_data;
mod config;
mod logger;
mod generator;

use raw_yaml_data::RawYAMLData;
use config::Config;
use logger::{Logger, Icons::LookingGlass, Icons::Sparkle};
use generator::generate;

fn main() {
    let mut logger = Logger::new();
    let args: Vec<String> = env::args().collect();
    let filename = if args.len() < 2 { "config.yml" } else { &args[1] };

    // open file
    logger.log(format!("Reading {}", filename), LookingGlass);
    let yaml_file = open_file(filename, &mut logger);
    
    //read yaml file
    let raw: RawYAMLData = serde_yaml
    ::from_reader(yaml_file)
    .expect("Could not read yaml_file.");

    let mut config = Config::new(raw.project_name, raw.version);
    match config.extract_schema(raw.schema) {
        Ok(_) => logger.log("parse schema success".to_string(), Sparkle),
        Err(e) => logger.log_err(e.to_string())
    }

    generate(config);
}

fn open_file(filename: &str, logger: &mut Logger) -> File {
    let file_to_open = File::open(filename);

    let file = match file_to_open {
        Ok(file) => file,
        Err(error) => match error.kind() {
            other_error => {
                logger.log_err(format!("Problem opening the file: {:?}", other_error));
                exit(1);
            }
        }
    };

    file
}