use void_lib::file_operations::location::find_file_path;
use void_lib::file_operations::location::find_dir_path;
use std::fs::{create_dir, File};
// use std::io::Write;
// use std::path::PathBuf;
use tempfile::tempdir;

#[allow(dead_code)]

#[test]
fn test_find_file_path_existing_file() {
    // Create a temporary directory for testing
    let temp_dir = tempdir().expect("Failed to create temporary directory");
    let temp_dir_path = temp_dir.path();

    // Create a file with a known name in the temporary directory
    let file_name = "test_file.txt";
    let file_path = temp_dir_path.join(file_name);
    File::create(&file_path).expect("Failed to create test file");

    // Test when the file exists in the directory
    let result = find_file_path(temp_dir_path, file_name);
    assert!(result.is_ok());
    assert_eq!(result.unwrap().unwrap(), file_path);
}

#[test]
fn test_find_file_path_non_existing_file() {
    // Create a temporary directory for testing
    let temp_dir = tempdir().expect("Failed to create temporary directory");
    let temp_dir_path = temp_dir.path();

    // Test when the file does not exist in the directory
    let result = find_file_path(temp_dir_path, "non_existent_file.txt");
    assert!(result.is_ok());
    assert!(result.unwrap().is_none());
}

#[test]
fn test_find_dir_path_existing_dir() {
    // Create a temporary directory for testing
    let temp_dir = tempdir().expect("Failed to create temporary directory");
    let temp_dir_path = temp_dir.path();

    // Create a subdirectory with a known name in the temporary directory
    let dir_name = "test_dir";
    let sub_dir_path = temp_dir_path.join(dir_name);
    create_dir(&sub_dir_path).expect("Failed to create test directory");

    // Test when the directory exists in the parent directory
    let result = find_dir_path(temp_dir_path, dir_name);
    assert!(result.is_ok());
    assert_eq!(result.unwrap().unwrap(), sub_dir_path);
}

#[test]
fn test_find_dir_path_non_existing_dir() {
    // Create a temporary directory for testing
    let temp_dir = tempdir().expect("Failed to create temporary directory");
    let temp_dir_path = temp_dir.path();

    // Test when the directory does not exist in the parent directory
    let result = find_dir_path(temp_dir_path, "non_existent_dir");
    assert!(result.is_ok());
    assert!(result.unwrap().is_none());
}
