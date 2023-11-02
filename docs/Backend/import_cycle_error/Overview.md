# Directory Structure:
```
+----------+
| Termvoid |
+----------+
    |
    |
    | 
+----------+
| firebase | --> Firebase/Storage initializer 
+----------+
    | 
+--------+
|  lib   | --> Rust Lib for file management (currently used; To be integrated later)
+--------+
    | 
+--------+
|  pkg   | --> `pkg/proto/` gRPC files
+--------+
    | 
+--------+
|  void  | --> CLI code
+--------+
```
----

# Flow Intended:

```
+----------------+
|   pkg/gRPC     | {Uses `github.com/joho/godotenv` (for accessing .env variables)}
|  (proto files) | 
+---------------+|+------------------------------------+       
                                  |
                        Proto files are imported in
                                  |
                            +-----------+
                            |    void   | {Also uses `github.com/joho/godotenv` (for accessing .env variables). Hence, the `import cycle error`}
                            |    (CLI)  |
                            +-----------+

```
----

# Folders Facing/Causing Similar Import Cycle Error:

```
firebase/firebase.go
```