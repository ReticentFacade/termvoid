use std::ffi::OsStr;
use std::fs;
use std::io;
use std::path::{Path, PathBuf};

#[allow(dead_code)] // Suppress the "dead code" warning
pub fn find_file_path(start_dir: &Path, file_name: &str) -> io::Result<Option<PathBuf>> {
    let mut result = None;

    for entry in fs::read_dir(start_dir)? {
        let entry = entry?;
        let entry_path = entry.path();

        if entry_path.is_file() {
            if entry_path.file_name() == Some(OsStr::new(file_name)) {
                result = Some(entry_path);
                break;
            }
        } else if entry_path.is_dir() {
            if let Some(found_path) = find_file_path(&entry_path, file_name)? {
                result = Some(found_path);
                break;
            }
        }
    }

    Ok(result)
}

#[allow(dead_code)]
pub fn find_dir_path(start_dir: &Path, dir_name: &str) -> io::Result<Option<PathBuf>> {
    let mut result = None;

    for entry in fs::read_dir(start_dir)? {
        let entry = entry?;
        let entry_path = entry.path();

        if entry_path.is_dir() {
            if entry_path.file_name() == Some(OsStr::new(dir_name)) {
                result = Some(entry_path);
                break;
            } else if let Some(found_path) = find_dir_path(&entry_path, dir_name)? {
                result = Some(found_path);
                break;
            }
        }
    }
    Ok(result)
}
