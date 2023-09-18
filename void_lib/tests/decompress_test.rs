use std::env;
use std::path::Path;
use std::ffi::CString;
use void_lib::archive_operations::decompress_file::{decompress_file, CResult};
use dotenv::dotenv;

#[test]
fn test_decompress_file() {
    // Loading environment variables from .env file:
    dotenv().ok();

    let archive_file = env::var("ARCHIVE_FILE").expect("ARCHIVE_FILE not set in .env");
    let destination_dir = env::var("DECOMPRESS_DEST").expect("DECOMPRESS_DEST not set in .env");

    let archive_file_cstr = CString::new(archive_file.clone()).expect("Failed to create CString");
    let destination_dir_cstr = CString::new(destination_dir.clone()).expect("Failed to create CString");

    let result = decompress_file(archive_file_cstr.as_ptr(), destination_dir_cstr.as_ptr());

    match result {
        CResult::Ok => {
            // Verify that the decompressed files exist
            let decompressed_file_path = Path::new(&destination_dir);

            assert!(decompressed_file_path.exists(), "Decompressed file does not exist");
        }
        CResult::Err => {
            panic!("Failed to decompress archive");
        }
    }
}
