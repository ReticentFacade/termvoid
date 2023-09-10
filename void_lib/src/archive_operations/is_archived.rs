use std::path::Path;

pub fn is_archived(file_path: &Path) -> bool {
    // Get file's extension: 
    if let Some(ext) = file_path.extension() {
        // Convert the extension to a lowercase str (for case-insensitive comparison):
        let ext_str = ext.to_string_lossy().to_lowercase();

        // Check if the ext matches common archived formats
        match ext_str.as_str() {
            "zip" | "tar" | "gz" | "tgz" | "tar.gz" => true,
            _ => false,
        } 
    } else {
        // No extension:
        false 
    }
}

