use serde::{Deserialize, Serialize};
use serde_yaml::Value;

#[derive(Serialize, Deserialize)]
pub struct RawYAMLData {
    pub project_name: String,
    pub version: u32,
    pub schema: Value,
}