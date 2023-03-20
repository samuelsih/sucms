use serde::{Deserialize, Serialize};
use serde_yaml::Value;
use std::fs::File;
use std::collections::HashMap;

#[derive(Debug, Serialize, Deserialize)]
struct Config {
    project_name: String,
    version: u32,
    schema: HashMap<String, Value>,
}

fn main() {
    let yaml_file = File
    ::open("file.yml")
    .expect("Could not open file.");
    
    let scrape_config: Config = serde_yaml
    ::from_reader(yaml_file)
    .expect("Could not read values.");

    for (key, value) in scrape_config.schema.iter() {
        println!("KEY: {} - VALUE: {:?}", key, value);
    }
}