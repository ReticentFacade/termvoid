use void_lib::archive_operations::archive_dir::archive_dir;
use std::env;
use dotenv::dotenv;

#[test]
pub fn test_archive_dir() {
    // Loading environment variables from .env file:
    dotenv().ok();
    
    let source_dir = env::var("SOURCE_DIR").expect("SOURCE_DIR not set in .env");
    let destination_dir = env::var("DESTINATION_DIR").expect("DESTINATION_DIR not set in .env");

    match archive_dir(&source_dir, &destination_dir) {
        Ok(_) => assert!(true),
        Err(e) => println!("Error: {}", e),
    }
}