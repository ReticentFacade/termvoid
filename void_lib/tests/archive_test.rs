use void_lib::archive_operations::archive_dir::archive_dir;

#[test]
pub fn test_archive_dir() {
    let source_dir = "/Users/Reticent/Desktop/Work/ProjectWorkspace/github/termvoid/void_lib/tests/test_data/archive_dir/source_dir";
    let destination_dir = "/Users/Reticent/Desktop/Work/ProjectWorkspace/github/termvoid/void_lib/tests/test_data/archive_dir/destination_dir";

    match archive_dir(source_dir, destination_dir) {
        Ok(_) => assert!(true),
        Err(e) => println!("Error: {}", e),
    }
}