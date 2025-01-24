## 简介

本项目演示了如何使用 Java 的 JNI（Java Native Interface）来调用 Go 语言编写的函数。通过这个示例，你可以了解如何在 Java 和 Go 之间进行跨语言调用。

## 脚本功能（Macos运行版本）

`build_and_run.sh` 是一个自动化脚本，完整模拟了 Java 调用 Go 代码的整个过程。它包括以下步骤：

1. 创建 Go 代码文件，并编译生成共享库。
2. 创建 Java 代码文件，并编译生成 JNI 头文件。
3. 创建 C/C++ 代码文件，并编译生成共享库。
4. 运行 Java 程序，调用 Go 代码。

这个脚本是一个完整的演示，展示了从代码编写到执行的全过程。

## 使用说明

1. 将脚本内容保存到 `build_and_run.sh` 文件中。
2. 确保脚本具有执行权限，可以通过以下命令设置权限：
   ```bash
   chmod +x build_and_run.sh
   ```
3. 运行脚本：
   ```bash
   ./build_and_run.sh
   ```
4. 每个步骤完成后，按回车键继续到下一个步骤。
5. 请根据你的实际路径修改 `JAVA_HOME` 和 `LIBRARY_PATH` 变量，以确保正确的编译和运行环境。