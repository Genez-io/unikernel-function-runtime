Binary size: 29K

1. Make package => `swift package init --type executable`
2. Insert code into `main.swift`
3. Build with statically linked swift stdlib => `swift build --static-swift-stdlib`
4. Add executable to `config.json`
5. Create image `ops image create -c config.json -i nanos_swift_bm`
6. Run firecracker
