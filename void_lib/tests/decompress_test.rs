use void_lib::archive_operations::decompress_file::decompress_file;
#[allow(unused_imports)]
use std::path::{Path, PathBuf};
use std::env;
use dotenv::dotenv;

#[test]
fn test_decompress_file() {
    // Loading environment variables from .env file:
    dotenv().ok();
    
    // Get the path to the archive file from the environment variable
    let archive_file = env::var("ARCHIVE_FILE").expect("ARCHIVE_FILE not set in .env");

    // Decompress the archive to the specified destination directory
    let destination_dir = env::var("DECOMPRESS_DEST").expect("DECOMPRESS_DEST not set in .env");

    decompress_file(&archive_file, &destination_dir).expect("Failed to decompress archive");

    // Verify that the decompressed files exist
    let decompressed_file_path = PathBuf::from(env::current_dir().unwrap()).join(&destination_dir);
    
    assert!(decompressed_file_path.exists(), "Decompressed file does not exist");
}
