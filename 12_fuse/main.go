// Copyright 2016 the Go-FUSE Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This program is the analogon of libfuse's hello.c, a a program that
// exposes a single file "file.txt" in the root directory.
package main

import (
	"context"
	"flag"
	"log"
	"syscall"
	"bufio"
	"fmt"
	"os"
	
//    "log"
//    "os"
    "os/exec"

	"github.com/hanwen/go-fuse/v2/fs"
	"github.com/hanwen/go-fuse/v2/fuse"
)

type HelloRoot struct {
	fs.Inode
}

func (r *HelloRoot) OnAdd(ctx context.Context) {

    cmd := exec.Command("sh","list.sh", "310", "320")
    cmd.Dir, _ = os.Getwd() 
    r1, _ := cmd.StdoutPipe()
    cmd.Stderr = cmd.Stdout
    done := make(chan struct{})
    scanner := bufio.NewScanner(r1)
    go func() {
        for scanner.Scan() {
            line := scanner.Text()
			fmt.Println(line)
			text := line
			ch := r.NewPersistentInode(
				ctx, &fs.MemRegularFile {
					Data: []byte(text),
					Attr: fuse.Attr{
						Mode: 0644,
					},
				}, fs.StableAttr {Ino: 2} )
			r.AddChild(text, ch, false)
        }
        done <- struct{}{}
    }()
    output := cmd.Start()
    fmt.Println(output)
}

func (r *HelloRoot) Getattr(ctx context.Context, fh fs.FileHandle, out *fuse.AttrOut) syscall.Errno {
	out.Mode = 0755
	return 0
}

var _ = (fs.NodeGetattrer)((*HelloRoot)(nil))
var _ = (fs.NodeOnAdder)((*HelloRoot)(nil))

func main() {
	debug := flag.Bool("debug", false, "print debug data")
	flag.Parse()
	if len(flag.Args()) < 1 {
		log.Fatal("Usage:\n  hello MOUNTPOINT")
	}
	opts := &fs.Options{}
	opts.Debug = *debug
	server, err := fs.Mount(flag.Arg(0), &HelloRoot{}, opts)
	if err != nil {
		log.Fatalf("Mount fail: %v\n", err)
	}
	fmt.Println(os.Getwd())
	server.Wait()
}
