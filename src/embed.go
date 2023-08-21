//go:generate sh -c "cd ../frontend; npm install -g pnpm; pnpm i; pnpm build"
package main

import _ "embed"

//go:embed embed/Forum.xdc
var xdcContent []byte
