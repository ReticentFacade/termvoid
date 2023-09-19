use void_lib::archive_operations::archive_dir::archive_dir;
use std::env;
use dotenv::dotenv;
// use void_lib::archive_operations::archive_dir::CResult;
#[repr(C)]
pub enum CResult {
    Ok,
    Err,
}

#[test]
pub fn test_archive_dir() {
    // Loading environment variables from .env file:
    dotenv().ok();
    
    let source_dir = env::var("SOURCE_DIR").expect("SOURCE_DIR not set in .env");
    let destination_dir = env::var("DESTINATION_DIR").expect("DESTINATION_DIR not set in .env");

    let result = archive_dir(source_dir.as_ptr() as *const i8, destination_dir.as_ptr() as *const i8);
    match result {
        CResult::Ok => assert!(true),
        CResult::Err => panic!("Error occurred while archiving")
    }
}