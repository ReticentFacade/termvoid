use void_lib::file_operations::exists::file_exists;

#[test]
fn test_exists_existing_file() {
    // (Absolute) Path to an existing file: 
    let path = "/Users/Reticent/Desktop/Work/ProjectWorkspace/github/termvoid/void_lib/tests/file.txt";
    assert!(file_exists(path));
}

#[test] 
fn test_exists_non_existing_file() {
    // Path to a non-existing file:
    let path = "./non_existing_file.txt";
    // let path = "/Users/Reticent/Desktop/Work/ProjectWorkspace/github/termvoid/void_lib/tests/file.txt";
    
    // assert!(!exists(path));
    assert!(
        !file_exists(path),
        "\nThe file '{}' shouldn't exist, but it does!\n",
        (path)
    );
}

#[test]
fn test_dir_exists_existing_dir() {
    // Path to an existing directory:
    let path = "/Users/Reticent/Desktop/Work/ProjectWorkspace/github/termvoid/void_lib/tests";
    assert!(file_exists(path));
}

#[test]
fn test_dir_exists_non_existing_dir() {
    // Path to a non-existing directory:
    let path = "./non_existing_dir";
    assert!(
        !file_exists(path),
        "\nThe directory '{}' shouldn't exist, but it does!\n",
        (path)
    );
}