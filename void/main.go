package main

/*
#cgo CFLAGS: -I../lib/void_lib
#cgo LDFLAGS: -L../lib/void_lib/target/release -lvoid_lib
#include "../lib/void_lib.h"
*/
import "C"

func main() {
	C.archive_dir(C.CString("../test_data/archive_dir/source_dir"), C.CString("../test_data/archive_dir/destination_dir/"))

	// C.decompress_file(C.CString("../test_data/archive_dir/destination_dir/source_dir.zip"), C.CString("../test_data/archive_dir/destination_dir/decompressed_file/"))
}
