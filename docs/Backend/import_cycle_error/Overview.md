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
                            |   void    | {Also uses `github.com/joho/godotenv` (for accessing .env variables). Hence, the `import cycle error`}
                            |   (CLI)   |
                            +-----------+

```
----

# Folders Facing/Causing Similar Import Cycle Error:

1. [firebase/firebase.go](https://github.com/ReticentFacade/termvoid/blob/master/firebase/firebase.go)
2. [pkg/proto/service_grpc.pb.go](https://github.com/ReticentFacade/termvoid/blob/master/pkg/proto/service_grpc.pb.go)
3. [void/server/server.go](https://github.com/ReticentFacade/termvoid/blob/master/void/server/server.go)
