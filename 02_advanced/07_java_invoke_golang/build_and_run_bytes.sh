#!/bin/bash

# 设置路径变量（根据实际环境修改）
JAVA_HOME="/Users/weijie.ma/Library/Java/JavaVirtualMachines/temurin-1.8.0_345/Contents/Home"
LIBRARY_PATH="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"

echo "LIBRARY_PATH is set to: $LIBRARY_PATH"

# ---------------------------------------------------------------
# Step 1: 生成 Go 代码文件
# ---------------------------------------------------------------
cat <<EOF > bytes_example.go
package main

/*
#include <stdlib.h>
#include <string.h>
*/
import "C"
import (
	"fmt"
	"unsafe"
)

//export ProcessByteArray2D
func ProcessByteArray2D(cArray **C.char, cLengths *C.int, numArrays C.int) {
	// 将 C 的二维数组转换为 Go 的 [][]byte
	goArray := make([][]byte, numArrays)
	ptr := (*[1<<30 - 1]*C.char)(unsafe.Pointer(cArray))[:numArrays:numArrays]
	lengths := (*[1<<30 - 1]C.int)(unsafe.Pointer(cLengths))[:numArrays:numArrays]

	for i := 0; i < int(numArrays); i++ {
		goArray[i] = C.GoBytes(unsafe.Pointer(ptr[i]), lengths[i])
		fmt.Printf("Go received byte array %d: %v\n", i, goArray[i])
	}
}

func main() {}
EOF

# ---------------------------------------------------------------
# Step 2: 生成 Java 代码文件
# ---------------------------------------------------------------
cat <<EOF > BytesJNI.java
import java.io.UnsupportedEncodingException;

public class BytesJNI {
    static {
        System.loadLibrary("bytesjni");
    }

    public native void processByteArray2D(byte[][] data);

    public static void main(String[] args) throws UnsupportedEncodingException {
        byte[][] data = {
            {0x00, 0x01},
            {0x02, 0x03}
        };
        byte[][] data2 = {
            hexStringToByteArray("0000"),
            "v282892hshshggsgsg".getBytes("UTF-8")
        };
        new BytesJNI().processByteArray2D(data);
        new BytesJNI().processByteArray2D(data2);
    }

    public static byte[] hexStringToByteArray(String hex) {
        if (hex == null || hex.length() % 2 != 0) {
            throw new IllegalArgumentException("Invalid hex string");
        }
        int len = hex.length();
        byte[] data = new byte[len / 2];
        for (int i = 0; i < len; i += 2) {
            data[i / 2] = (byte) ((Character.digit(hex.charAt(i), 16) << 4) +
                Character.digit(hex.charAt(i + 1), 16));
        }
        return data;
    }
}
EOF

# ---------------------------------------------------------------
# Step 3: 生成 C JNI 代码文件
# ---------------------------------------------------------------
cat <<EOF > BytesJNI.c
#include <jni.h>
#include <stdlib.h>
#include "BytesJNI.h"
#include "libbytes.h"

JNIEXPORT void JNICALL Java_BytesJNI_processByteArray2D(JNIEnv *env, jobject obj, jobjectArray jData) {
    // 获取二维数组长度
    int numArrays = (*env)->GetArrayLength(env, jData);

    // 分配内存保存每个字节数组的指针和长度
    char **cArrays = (char**)malloc(numArrays * sizeof(char*));
    int *cLengths = (int*)malloc(numArrays * sizeof(int));

    // 遍历每个字节数组
    for (int i = 0; i < numArrays; i++) {
        jbyteArray jba = (jbyteArray)(*env)->GetObjectArrayElement(env, jData, i);
        jsize len = (*env)->GetArrayLength(env, jba);
        jbyte *data = (*env)->GetByteArrayElements(env, jba, NULL);

        // 拷贝数据到 C 内存
        cArrays[i] = (char*)malloc(len);
        memcpy(cArrays[i], data, len);
        cLengths[i] = len;

        // 释放 Java 数组引用
        (*env)->ReleaseByteArrayElements(env, jba, data, JNI_ABORT);
    }

    // 调用 Go 函数
    ProcessByteArray2D(cArrays, cLengths, numArrays);

    // 释放内存
    for (int i = 0; i < numArrays; i++) free(cArrays[i]);
    free(cArrays);
    free(cLengths);
}
EOF

# ---------------------------------------------------------------
# Step 4: 编译 Go 代码生成共享库
# ---------------------------------------------------------------
echo "Step 1: 编译 Go 代码生成共享库"
read -p "Press Enter to continue..."
go build -o libbytes.dylib -buildmode=c-shared bytes_example.go

# ---------------------------------------------------------------
# Step 5: 编译 Java 代码生成头文件
# ---------------------------------------------------------------
echo "Step 2: 编译 Java 代码并生成 JNI 头文件"
read -p "Press Enter to continue..."
javac -h . BytesJNI.java

# ---------------------------------------------------------------
# Step 6: 编译 C JNI 代码生成共享库
# ---------------------------------------------------------------
echo "Step 3: 编译 C 代码生成共享库"
read -p "Press Enter to continue..."
gcc -shared -fpic -o libbytesjni.dylib \
    -I${JAVA_HOME}/include -I${JAVA_HOME}/include/darwin \
    BytesJNI.c -L. -lbytes

# ---------------------------------------------------------------
# Step 7: 运行 Java 程序
# ---------------------------------------------------------------
echo "Step 4: 运行 Java 程序"
read -p "Press Enter to continue..."
java -Djava.library.path=${LIBRARY_PATH} BytesJNI
