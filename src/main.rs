use std::env;
use std::fs::File;
use std::process::exit;

mod config;
mod generator;
mod logger;
mod raw_yaml_data;

use config::Config;
use generator::generate;
use logger::Logger;
use raw_yaml_data::RawYAMLData;

fn main() {
    let mut logger = Logger::new();
    let args: Vec<String> = env::args().collect();
    let filename = if args.len() < 2 {
        "config.yml"
    } else {
        &args[1]
    };

    // open file
    logger.log(format!("Reading {}", filename));
    let yaml_file = open_file(filename, &mut logger);

    //read yaml file
    let raw: RawYAMLData = serde_yaml::from_reader(yaml_file).expect("Could not read yaml_file.");

    let mut config = Config::new(raw.project_name, raw.create_new_folder);
    match config.extract_schema(raw.schema) {
        Ok(_) => logger.log("parse schema success".to_string()),
        Err(e) => logger.log(e.to_string()),
    }

    generate(config);
}

fn open_file(filename: &str, logger: &mut Logger) -> File {
    let file_to_open = File::open(filename);

    let file = match file_to_open {
        Ok(file) => file,
        Err(error) => match error.kind() {
            other_error => {
                logger.log(format!("Problem opening the file: {:?}", other_error));
                exit(1);
            }
        },
    };

    file
}
