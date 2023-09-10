// use std::path::Path;
// #[allow(unused_imports)]
// use zip_archive::{Archiver, Format};

// fn decompress_file(archive_file: &str, destination_dir: &str) -> Result<(), String> {
//     let archive_file = Path::new(archive_file);
//     let destination_dir = Path::new(destination_dir);

//     let mut archiver = Archiver::new();
//     archiver.push(archive_file);
//     archive.set_destination(destination_dir);

//     match archiver.extract() {
//         Ok(_) => Ok(()),
//         Err(e) => Err(format!("Failed to decompress file: {}", e)),
//     }
// }