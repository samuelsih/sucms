use std::{fs::File};

mod raw_yaml_data;
mod config;

use raw_yaml_data::RawYAMLData;
use config::Config;

fn main() {
    let yaml_file = File
    ::open("file.yml")
    .expect("Could not open file.");
    
    let scrape_config: RawYAMLData = serde_yaml
    ::from_reader(yaml_file)
    .expect("Could not read yaml_file.");

    let mut config = Config::new(scrape_config.project_name, scrape_config.version);
    config.extract_schema(scrape_config.schema);

    println!("{:#?}", config);

}