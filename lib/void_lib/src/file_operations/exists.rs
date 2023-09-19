use std::fs;

#[allow(dead_code)]  // Suppress the "dead code" warning
pub fn file_exists(path: &str) -> bool {
    // fs::metadata(path).is_ok()
    if let Ok(metadata) = fs::metadata(path) {
        metadata.is_file()
    } else {
        false
    }
}

pub fn dir_exists(path: &str) -> bool {
    if let Ok(metadata) = fs::metadata(path) {
        metadata.is_dir()
    } else {
        false
    }
}