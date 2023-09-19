// use std::os::raw::c_char;
// #[allow(unused_imports)]
// use std::ffi::{CString, CStr};
// #[allow(unused_imports)]
// use std::ptr;
// use std::fs::File;
// #[allow(unused_imports)]
// use std::io::{self, Read, Write};
// use std::path::PathBuf;
// use zip::ZipArchive;

// #[repr(C)]
// pub enum CResult {
//     Ok,
//     Err,
// }

// #[no_mangle]
// pub extern "C" fn decompress_file(archive_file: *const c_char, destination_dir: *const c_char) -> CResult {
//     let archive_file_cstr = unsafe { CStr::from_ptr(archive_file) };
//     let destination_dir_cstr = unsafe { CStr::from_ptr(destination_dir) };

//     let archive_file_str = archive_file_cstr.to_str().expect("Invalid UTF-8 in archive_file");
//     let destination_dir_str = destination_dir_cstr.to_str().expect("Invalid UTF-8 in destination_dir");

//     match decompress(archive_file_str, destination_dir_str) {
//         Ok(_) => CResult::Ok,
//         Err(_) => CResult::Err,
//     }
// }

// fn decompress(archive_file: &str, destination_dir: &str) -> Result<(), String> {
//     // Open the archive file
//     let archive_file = File::open(archive_file).map_err(|e| format!("Failed to open archive file: {}", e))?;

//     // Create a mutable ZipArchive from the archive file
//     let mut zip_archive = ZipArchive::new(archive_file).map_err(|e| format!("Failed to create ZipArchive: {}", e))?;

//     // Extract the contents to the destination directory
//     for i in 0..zip_archive.len() {
//         let mut entry = match zip_archive.by_index(i) {
//             Ok(e) => e,
//             Err(e) => return Err(format!("Failed to get archive entry: {}", e)),
//         };

//         let entry_path = PathBuf::from(destination_dir).join(entry.name().to_owned());

//         // Ensure the parent directories exist
//         if let Some(parent) = entry_path.parent() {
//             std::fs::create_dir_all(parent).map_err(|e| format!("Failed to create directory: {}", e))?;
//         }

//         // Read and write the file
//         let mut buffer = Vec::new();
//         entry.read_to_end(&mut buffer).map_err(|e| format!("Failed to read file from archive: {}", e))?;

//         // Write the file to the destination
//         let mut output_file = File::create(&entry_path).map_err(|e| format!("Failed to create file: {}", e))?;
//         output_file.write_all(&buffer).map_err(|e| format!("Failed to write file: {}", e))?;
//     }

//     Ok(())
// }

extern crate libc;

#[allow(unused_imports)]
use libc::{c_char, c_int};

#[allow(unused_imports)]
use std::ffi::{CString, CStr};
use std::fs::File;

#[allow(unused_imports)]
use std::io::{self, Read, Write};
use std::path::PathBuf;
use zip::ZipArchive;

#[repr(C)]
pub enum CResult {
    Ok,
    Err,
}

#[no_mangle]
pub extern "C" fn decompress_file(
    archive_file: *const c_char,
    destination_dir: *const c_char,
) -> CResult {
    let archive_file_cstr = unsafe { CStr::from_ptr(archive_file) };
    let destination_dir_cstr = unsafe { CStr::from_ptr(destination_dir) };

    let archive_file_str = archive_file_cstr.to_str().expect("Invalid UTF-8 in archive_file");
    let destination_dir_str = destination_dir_cstr.to_str().expect("Invalid UTF-8 in destination_dir");

    match decompress(archive_file_str, destination_dir_str) {
        Ok(_) => CResult::Ok,
        Err(_) => CResult::Err,
    }
}

fn decompress(archive_file: &str, destination_dir: &str) -> Result<(), String> {
    let archive_file = File::open(archive_file).map_err(|e| format!("Failed to open archive file: {}", e))?;
    let mut zip_archive = ZipArchive::new(archive_file).map_err(|e| format!("Failed to create ZipArchive: {}", e))?;

    for i in 0..zip_archive.len() {
        let mut entry = match zip_archive.by_index(i) {
            Ok(e) => e,
            Err(e) => return Err(format!("Failed to get archive entry: {}", e)),
        };

        let entry_path = PathBuf::from(destination_dir).join(entry.name().to_owned());

        if let Some(parent) = entry_path.parent() {
            std::fs::create_dir_all(parent).map_err(|e| format!("Failed to create directory: {}", e))?;
        }

        let mut buffer = Vec::new();
        entry.read_to_end(&mut buffer).map_err(|e| format!("Failed to read file from archive: {}", e))?;

        let mut output_file = File::create(&entry_path).map_err(|e| format!("Failed to create file: {}", e))?;
        output_file.write_all(&buffer).map_err(|e| format!("Failed to write file: {}", e))?;
    }

    Ok(())
}
