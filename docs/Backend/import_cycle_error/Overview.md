# File Structure:
```sh
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
+--------+
|  void  |
| (CLI)  | +------------------------+
+--------+          |
              +-----------+
              |           |
              +-----------+

```
----

# Folders Facing/Causing Import Cycle Error:

