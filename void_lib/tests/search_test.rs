use void_lib::file_operations::search::file_search;
use std::fs::File;
use tempfile::tempdir;

#[test]
pub fn test_file_search_existing_file() {
    let temp_dir = tempdir().expect("Failed to create temporary directory");

    // Create a file with a known name in the temporary directory
    let file_name = "test_file.txt";
    let file_path = temp_dir.path().join(file_name);
    File::create(&file_path).expect("Failed to create test file");

    // Test when the file exists in the directory
    let result = file_search(temp_dir.path(), file_name);
    assert!(result.is_ok());
    assert_eq!(result.unwrap().unwrap(), file_path);
}

#[test]
pub fn test_file_search_non_existing_file() {
    let temp_dir = tempdir().expect("Failed to create temporary directory");

    // Test when the file does not exist in the directory
    let result = file_search(temp_dir.path(), "non_existent_file.txt");
    assert!(result.is_ok());
    assert!(result.unwrap().is_none());
}
