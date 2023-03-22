use std::env;
use std::fs::File;
use std::process::exit;
use ctrlc;

mod config;
mod generator;
mod logger;
mod raw_yaml_data;

use config::Config;
use generator::generate;
use logger::log;
use raw_yaml_data::RawYAMLData;

fn main() {
    if let Err(err) = ctrlc::set_handler(move || {
        log("Cancelled!!!".to_string());
        exit(1);
    }) {
        log(format!("Can't handle Ctrl+C events: {}", err));
        exit(1);
    }

    let args: Vec<String> = env::args().collect();
    let filename = if args.len() < 2 {
        "config.yml"
    } else {
        &args[1]
    };

    // open file
    log(format!("Reading {}", filename));
    let yaml_file = open_file(filename);

    //read yaml file
    let raw: RawYAMLData = serde_yaml::from_reader(yaml_file).expect("Could not read yaml_file.");

    log("Parsing schema".to_string());
    let mut config = Config::new(raw.project_name, raw.create_new_folder);
    match config.extract_schema(raw.schema) {
        Ok(_) => (),
        Err(e) => log(e.to_string()),
    }

    generate(config);
}

fn open_file(filename: &str) -> File {
    let file_to_open = File::open(filename);

    let file = match file_to_open {
        Ok(file) => file,
        Err(error) => match error.kind() {
            other_error => {
                log(format!("Problem opening the file: {:?}", other_error));
                exit(1);
            }
        },
    };

    file
}
