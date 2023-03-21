use serde::{Deserialize, Serialize};
use serde_yaml::Value;

#[derive(Serialize, Deserialize)]
pub struct RawYAMLData {
    pub project_name: String,
    pub create_new_folder: bool,
    pub schema: Value,
}
