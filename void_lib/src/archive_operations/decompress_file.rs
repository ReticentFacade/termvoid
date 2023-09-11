#[allow(unused_imports)]
use std::fs::File;
#[allow(unused_imports)]
use std::io::{self, Read, Write};
#[allow(unused_imports)]
use std::path::{Path, PathBuf};
use zip::ZipArchive;

pub fn decompress_file(archive_file: &str, destination_dir: &str) -> Result<(), String> {
    // Open the archive file
    let archive_file = File::open(archive_file).map_err(|e| format!("Failed to open archive file: {}", e))?;

    // Create a mutable ZipArchive from the archive file
    let mut zip_archive = ZipArchive::new(archive_file).map_err(|e| format!("Failed to create ZipArchive: {}", e))?;

    // Extract the contents to the destination directory
    for i in 0..zip_archive.len() {
        let mut entry = match zip_archive.by_index(i) {
            Ok(e) => e,
            Err(e) => return Err(format!("Failed to get archive entry: {}", e)),
        };

        let entry_path = PathBuf::from(destination_dir).join(entry.name().to_owned());

        // Ensure the parent directories exist
        if let Some(parent) = entry_path.parent() {
            std::fs::create_dir_all(parent).map_err(|e| format!("Failed to create directory: {}", e))?;
        }

        // Read and write the file
        let mut buffer = Vec::new();
        entry.read_to_end(&mut buffer).map_err(|e| format!("Failed to read file from archive: {}", e))?;

        // Write the file to the destination
        let mut output_file = File::create(&entry_path).map_err(|e| format!("Failed to create file: {}", e))?;
        output_file.write_all(&buffer).map_err(|e| format!("Failed to write file: {}", e))?;
    }

    Ok(())
}

// pub fn decompress_file(archive_file: &str, destination_dir: &str) -> Result<(), String> {
//     // Open the archive file
//     let archive_file = File::open(archive_file).map_err(|e| format!("Failed to open archive file: {}", e))?;

//     // Create a mutable ZipArchive from the archive file
//     let mut zip_archive = ZipArchive::new(archive_file).map_err(|e| format!("Failed to create ZipArchive: {}", e))?;

//     // Extract the contents to the destination directory
//     for i in 0..zip_archive.len() {
//         let entry = match zip_archive.by_index(i) {
//             Ok(e) => e,
//             Err(e) => return Err(format!("Failed to get archive entry: {}", e)),
//         };

//         let mut file = match entry {
//             Ok(f) => f,
//             Err(e) => return Err(format!("Failed to get archive entry: {}", e)),
//         };

//         let mut buffer = Vec::new();
//         file.read_to_end(&mut buffer).map_err(|e| format!("Failed to read file from archive: {}", e))?;

//         // Determine the path for the extracted file
//         let entry_path = PathBuf::from(destination_dir).join(entry.name());

//         // Ensure the parent directories exist
//         if let Some(parent) = entry_path.parent() {
//             std::fs::create_dir_all(parent).map_err(|e| format!("Failed to create directory: {}", e))?;
//         }

//         // Write the file to the destination
//         let mut output_file = File::create(&entry_path).map_err(|e| format!("Failed to create file: {}", e))?;
//         output_file.write_all(&buffer).map_err(|e| format!("Failed to write file: {}", e))?;
//     }

//     Ok(())
// }
