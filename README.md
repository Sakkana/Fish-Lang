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




### 3. 求值

### 4. interpreter 类型扩展

### 5. 宏

### 6. compiler & virtual machine

### 7. 字节码

### 8. 编译表达式

### 9. 条件语句

### 10. 符号

### 11. compiler 类型扩展

### 12. 函数

### 13. 闭包
