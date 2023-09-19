use std::path::Path;
use void_lib::archive_operations::is_archived::is_archived;

#[test]
pub fn test_is_archived_is() {
    let file_path = Path::new("test.zip");
    if is_archived(&file_path) {
        println!("{} is an archive! 😬", file_path.display());
    } else {
        println!("{} is NOT an archive! 😬", file_path.display())
    }
}