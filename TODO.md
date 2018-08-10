# TODO

## Grammar
* Add VariableDeclarationStatement
  This simplifies the function body to an ast.BlockStatement, and makes parsing variable declarations cleaner.

## Scanner
* Define separate tokens for literals True/False?
* Define separate tokens for types Int/Bool/Void?

## Parser/AST
* Try to group parse errors together
  * Consume tokens until what is probably the end of the expected AST node?
    This increases the chance that the rest of the file will parse correctly.
  * Show only the first error on any line?
* Separate AST node types for builtin types Int/Bool/Void?
* Pretty printer:
  * Print comments at the correct positions
  * Configurable style?

## Semantic analysis
* Entirely

## Code generation
* Entirely
