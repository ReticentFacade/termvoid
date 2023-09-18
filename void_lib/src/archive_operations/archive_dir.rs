use std::path::Path;
use zip_archive::{Archiver, Format};
#[allow(unused_imports)]
use std::os::raw::c_int;

#[repr(C)]
pub enum CResult {
    Ok,
    Err,
}

#[no_mangle]
pub extern "C" fn archive_dir(
    source_dir: *const std::os::raw::c_char,
    destination_dir: *const std::os::raw::c_char,
) -> CResult {
    // `unsafe` 'cause raw pointers obtained from C are considered unsafe in Rust
    let source_dir_str = unsafe { std::ffi::CStr::from_ptr(source_dir).to_str().unwrap() };
    let destination_dir_str = unsafe { std::ffi::CStr::from_ptr(destination_dir).to_str().unwrap() };

    let source_dir = Path::new(source_dir_str);
    let destination_dir = Path::new(destination_dir_str);

    let mut archiver = Archiver::new();
    archiver.push(source_dir);
    archiver.set_destination(destination_dir);
    archiver.set_format(Format::Zip);
    archiver.set_thread_count(4);

    match archiver.archive() {
        Ok(_) => CResult::Ok,
        Err(_) => CResult::Err,
    }
}
