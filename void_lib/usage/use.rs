use void_lib::file_operations::exists::{file_exists, dir_exists};

fn main() {
    let filePath = "./tests/file.txt";
    if exists::file_exists(filePath) {
        println!("The file '{}' exists! :D", filePath);
    } else {
        println!("The file '{}' doesn't exist! :(", filePath);
    }
    
    let dirPath = "./tests";
    if exists::dir_exists(dirPath) {
        println!("The directory '{}' exists! :D", dirPath);
    } else {
        println!("The directory '{}' doesn't exist! :(", dirPath);
    }
}