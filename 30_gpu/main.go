package main

/*
#cgo CFLAGS: -x objective-c
#cgo LDFLAGS: -framework Metal -framework Foundation
#import <Metal/Metal.h>
#import <Foundation/Foundation.h>
#include <stdlib.h>

void runMetalKernel(float* data, int count) {
    @autoreleasepool {

        id<MTLDevice> device = MTLCreateSystemDefaultDevice();
        if (!device) {
            NSLog(@"ERROR: No Metal device found");
            return;
        }
        NSLog(@"Using device: %@", device.name);

        NSError *error = nil;
        NSString *src = @"kernel void add_one(device float* data [[buffer(0)]], uint id [[thread_position_in_grid]]) { if (id < 1000) data[id] += 1; }";
        id<MTLLibrary> library = [device newLibraryWithSource:src options:nil error:&error];
        if (!library) {
            NSLog(@"ERROR compiling library: %@", error);
            return;
        }

        id<MTLFunction> func = [library newFunctionWithName:@"add_one"];
        if (!func) { NSLog(@"ERROR creating function"); return; }

        id<MTLComputePipelineState> pipeline = [device newComputePipelineStateWithFunction:func error:&error];
        if (!pipeline) { NSLog(@"ERROR creating pipeline: %@", error); return; }

        id<MTLCommandQueue> queue = [device newCommandQueue];
        id<MTLCommandBuffer> cmdBuf = [queue commandBuffer];
        id<MTLComputeCommandEncoder> encoder = [cmdBuf computeCommandEncoder];

        id<MTLBuffer> buffer = [device newBufferWithBytes:data length:count*sizeof(float) options:0];
        if (!buffer) { NSLog(@"ERROR creating buffer"); return; }

        [encoder setComputePipelineState:pipeline];
        [encoder setBuffer:buffer offset:0 atIndex:0];

        MTLSize threads = MTLSizeMake(count, 1, 1);
        MTLSize threadgroup = MTLSizeMake(1,1,1);

        NSLog(@"Dispatching %d threads", count);
        [encoder dispatchThreads:threads threadsPerThreadgroup:threadgroup];
        [encoder endEncoding];

        [cmdBuf commit];
        [cmdBuf waitUntilCompleted];

        NSLog(@"Copying buffer back");
        memcpy(data, [buffer contents], count*sizeof(float));
    }
}
*/
import "C"
import "fmt"

func main() {
    data := []C.float{0,1,2,3,4,5,6,7,8,9}
    fmt.Println("Input:", data)
    C.runMetalKernel(&data[0], C.int(len(data)))
    fmt.Println("Output:", data)
}
