use std::ffi::OsStr;
use std::fs;
use std::io;
use std::path::{Path, PathBuf};

pub fn file_search<P>(directory: P, target_file_name: &str) -> io::Result<Option<PathBuf>>
where
    P: AsRef<Path>,
{
    let directory = directory.as_ref();
    for entry in fs::read_dir(directory)? {
        let entry = entry?;
        let entry_path = entry.path();

        if entry_path.is_file() {
            if entry_path.file_name() == Some(OsStr::new(target_file_name)) {
                let result = Some(entry_path);
                println!("Found file: {:?}", result);
                return Ok(result);
            }
        } else if entry_path.is_dir() {
            if let Some(found_path) = file_search(&entry_path, target_file_name)? {
                let result = Some(found_path);
                println!("Found file: {:?}", result);

                return Ok(result);
            }
        }
    }
    let result = None;
    println!("No match found! 🙄 ");
    Ok(result)
}
