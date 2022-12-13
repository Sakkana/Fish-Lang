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
抽象语法树（AST）：用于源代码内部表示的数据结构，内部省略了许多细节，比如空格，分号，换行符等。\
该部分构建了一个 Fish 语言的语法分析器，根据前面词法分析器生成的 tokens 解析，并在递归求解的时候构建 AST 的实例。
语法分析主要有 `自上而下` 和 `自下而上` 两种，自上而下有多种变体，例如递归下降分析，Earley 分析，预测分析等。

该解释器使用 `递归下降` 的分析方法，从构造 AST 的根节点开始下降，
1. 解析 let 语句 
> let x = 123;\
> let y = a / b - 1;\
> let z = add(2, 3) - a;

let语句是最基础的一个语句，该语句使得值和变量名绑定。\
这个值可以是一个常量字面量，也可以是一个表达式。表达式会返回一个值，因此可以求值。\
第一部分实现了一个很破产的 let 语法解析功能，因为他直判断符合 `let x = ;` 的语法，字面量或者表达式的递归求解放在了后面。\
但不管怎么说，这个解释器可以定义变量啦！

3. 解析 return 语句
4. 解析表达式



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
