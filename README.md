# Fish-Lang
Writing an interpreter and compiler in Golang, supporting JavaScript-like grammar translation.

Owing to <a href="https://interpreterbook.com/">Writing An Interpreter In Go</a> and <a href="https://compilerbook.com/">Writing A Compiler In Go</a>.

My tutorial of fantastic formal language and automata & compiler principles.

### 1. 词法分析
不判别输入语句的语法正确性，仅仅将每一条输入语句拆解为 token，判别 token 的合法性。
该语法分析器包含两种类型的 token，每种又细分为多种子类：
* 字面量
  * 系统保留关键词，如 `if`, `else`, `return`, `fn` 等
  * 用户定义变量名，如 `a`, `result`, `Node` 等
* 符号
  * 运算符，如 `>`, `<`, `==` 等
  * 分割符，如 `(`， `)`， `{`， `}`， `;` 等



### 2. 语法分析

### 3. 求值

### 4. 解释器类型扩展

### 5. 宏
---
### 6. 编译器和虚拟机

### 7. 字节码

### 8. 编译表达式

### 9. 条件语句

### 10. 符号

### 11. 编译器类型扩展

### 12. 函数

### 13. 闭包
