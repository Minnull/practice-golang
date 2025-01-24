#!/bin/bash

# 设置绝对路径变量（请根据实际路径修改）
JAVA_HOME="/Users/weijie.ma/Library/Java/JavaVirtualMachines/temurin-1.8.0_345/Contents/Home"
LIBRARY_PATH="/Users/weijie.ma/Desktop/all/code/goland/mmdb-go/hello"

# 创建 Go 代码文件
cat <<EOF > hello.go
package main

import "C"
import "fmt"

//export SayHello
func SayHello(name *C.char) {
    fmt.Printf("Hello, %s from Go!\\n", C.GoString(name))
}

func main() {}
EOF

# 创建 Java 代码文件
cat <<EOF > HelloJNI.java
public class HelloJNI {
    static {
        System.loadLibrary("hellojni");
    }

    public native void sayHello(String name);

    public static void main(String[] args) {
        new HelloJNI().sayHello("World");
    }
}
EOF

# 提示用户执行步骤
echo "Step 1: 编译 Go 代码生成共享库"
read -p "Press Enter to continue..."
go build -o libhello.dylib -buildmode=c-shared hello.go

echo "Step 2: 编译 Java 代码并生成 JNI 头文件"
read -p "Press Enter to continue..."
javac -h . HelloJNI.java

# 创建 C/C++ 代码文件
cat <<EOF > HelloJNI.c
#include <jni.h>
#include <stdio.h>
#include "HelloJNI.h"
#include "libhello.h"

JNIEXPORT void JNICALL Java_HelloJNI_sayHello(JNIEnv *env, jobject obj, jstring name) {
    const char *nameStr = (*env)->GetStringUTFChars(env, name, 0);
    SayHello((char *)nameStr);
    (*env)->ReleaseStringUTFChars(env, name, nameStr);
}
EOF

echo "Step 3: 编译 C/C++ 代码生成共享库"
read -p "Press Enter to continue..."
gcc -shared -fpic -o libhellojni.dylib \
    -I${JAVA_HOME}/include -I${JAVA_HOME}/include/darwin \
    HelloJNI.c -L. -lhello

echo "Step 4: 运行 Java 程序"
read -p "Press Enter to continue..."
java -Djava.library.path=${LIBRARY_PATH} HelloJNI
