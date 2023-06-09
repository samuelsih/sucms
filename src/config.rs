use std::error::Error;

use serde_yaml::Value;

#[derive(Debug)]
pub struct Config {
    project_name: String,
    create_new_folder: bool,
    schema: Vec<(String, Vec<(String, String)>)>,
}

impl Config {
    pub fn new(project_name: String, create_new_folder: bool) -> Config {
        Config {
            project_name,
            create_new_folder,
            schema: Vec::new(),
        }
    }

    pub fn get_project_name(&self) -> &str {
        &self.project_name
    }

    pub fn want_in_new_folder(&self) -> bool {
        self.create_new_folder
    }

    pub fn extract_schema(&mut self, value: Value) -> Result<(), Box<dyn Error>> {
        let schema = value.as_mapping().expect("cannot extract raw yaml spec");

        for (key, value) in schema.iter() {
            let specs = value.as_mapping().expect("Bad schema config");
            let tablename = key.as_str().expect("can't read tablename");
            let mut columns = Vec::<(String, String)>::new();

            for (col, datatype) in specs.iter() {
                columns.push((match_type(col), match_type(datatype)));
            }

            self.schema.push((tablename.to_string(), columns));
        }

        Ok(())
    }
}

fn match_type(value: &Value) -> String {
    match value {
        Value::Bool(v) => {
            return if *v {
                "true".to_string()
            } else {
                "false".to_string()
            }
        }
        Value::Number(v) => v.to_string(),
        Value::String(v) => v.clone(),
        _ => "".to_string(),
    }
}
