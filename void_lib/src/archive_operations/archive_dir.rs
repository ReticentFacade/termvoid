use std::path::Path;
use zip_archive::{Archiver, Format};

#[allow(dead_code)]
pub fn archive_dir(source_dir: &str, destination_dir: &str) -> Result<(), String> {

    // Archive the directory using the default settings for FORMAT & THREAD_COUNT:
    archive_dir_without_defaults(source_dir, destination_dir, Format::Zip, 4)
}

#[allow(dead_code)]
fn archive_dir_without_defaults(source_dir: &str, destination_dir: &str, format: Format, thread_count: usize) -> Result<(), String> {
    let source_dir = Path::new(source_dir);
    let destination_dir = Path::new(destination_dir);

    let mut archiver = Archiver::new();
    archiver.push(source_dir);
    archiver.set_destination(destination_dir);
    archiver.set_format(format);
    archiver.set_thread_count(thread_count.try_into().unwrap());

    match archiver.archive() {
        Ok(_) => Ok(()),
        Err(e) => Err(format!("Failed to archive directory: {}", e)),
    }
}