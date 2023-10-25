Icon Credits:
https://icon-icons.com/icon/bear-panda-animal-reddy/85225

Create build for windows:
goreleaser build --snapshot --single-target --clean

Create build for linux without needing to install goreleaser:
go build -o dist/quick-access-tray

Compiled binaries can be found in the dist folder.
