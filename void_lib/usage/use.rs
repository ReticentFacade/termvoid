use void_lib::file_operations::exists::{dir_exists, file_exists};
// use void_lib::file_operations::search::file_search;

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

    // let dirP = "~/";
    // let target_file_name = "test.md";
    // let result = file_search(dirP, target_file_name);
    // println!("Result: {:?}", result);
}
