### Exercício 1: Filtro de Idades com Structs e Métodos

**Objetivo**: Criar uma função que filtra pessoas por faixa etária usando structs e métodos.

1. Defina uma struct `Pessoa` com pelo menos dois campos: `Nome` e `Idade`.
2. Implemente um método `DentroDaFaixaEtaria(min, max int) bool` na struct `Pessoa` que retorna `true` se a idade da pessoa estiver entre `min` e `max`, inclusive.
3. Crie uma função que recebe uma slice de `Pessoa` e a faixa etária (`min` e `max`), retornando uma nova slice apenas com as pessoas dentro da faixa etária especificada.