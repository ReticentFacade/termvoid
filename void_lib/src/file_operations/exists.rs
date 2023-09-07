use std::fs;

#[allow(dead_code)]  // Suppress the "dead code" warning
pub fn file_exists(path: &str) -> bool {
    fs::metadata(path).is_ok()
}

pub fn dir_exists(path: &str) -> bool {
    fs::metadata(path).is_ok()
}