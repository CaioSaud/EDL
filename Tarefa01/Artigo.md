# **Smalltalk**


#### *Introdução*

Smalltalk-80, ou apenas, Smalltalk é a linguagem de programação que tornou popular o paradigma de orientação ao objeto.
A pureza e clareza de Smalltalk sobre esse conceito inspirou quase todas as outras arquiteturas das linguagens OOP. Por ser uma linguagem relativamente simples, Smalltalk é muito indicado para amadores em programação.

Umas das primeiras linguagens orientada a objetos foi Smalltalk, e o que explica o termo puramente orientada a objetos, é que em Smalltalk tudo é objeto, inclusive if, while e for não são comandos, são métodos.
  
Apesar da primeira versão de Smalltalk ter sido criada em 1971, essa linguagem ficou em segundo lugar na pesquisa da Stack
Overflow "most loved programming language" de 2017. Já em 2018 não ficou nem entre as top 26.

*Sumário*:
  * [Origens e Influências](#Origens-e-Influências)
  * [Classificações](#Classificações)
  * [Exemplos](#Exemplos)
    * [Exemplos Simples](#Exemplos-Simples)
    * [Exemplo Interessante](#Exemplo-Interessante)
  * [Funções de Alta Expressividade](#Funções-de-Alta-Expressividade)  
    * [Linguagem Baseada em Imagens](#Linguagem-Baseada-em-Imagens)
    * [Reflexividade](#Reflexividade)  
  * [Bibliografia](#Bibliografia)
    
    
#### *Origens e Influências*

Smalltalk-71 é uma linguagem de programação com uma rica história e um grande legado. Foi criado pela brilhante equipe de
Alan Kay, Dan Ingalls, and Adele Goldberg, na Xerox Palo Alto Research Center em 1971. Esse nome foi dado pois por ser 
algo tão simples, as pessoas iriam ficar surpreendidas caso fosse possível realizar algo de valor com a linguagem.
Smalltalk-80, também chamada apenas de Smalltalk,foi a primeira variante da linguagem disponível para fora do PARC.

Smalltalk para ser criada teve influências de FLEX, PLANNER, LOGO, META II, além de SIMULA.

Alguns feitos de Smalltalk:
  Linguagem pioneira no JIT (just in time) compilation.
  Introduziu ao mundo a linguem de máquina virtual, na qual Java e Ruby são baseadas.

#### *Classificações*

Smalltalk é uma liguagem "puramente" orientada a objeto. Isso significa que, diferente de C++ e Java, em Smalltalk não
existe diferença entre valores que são objetos e valores que são de tipos primitivos. Em smalltalk esses valores de tipos
primitivos, como inteiros e booleanos, também são considerados objetos.

  Para ser considerada uma linguagem puramente orientada a objeto a linguagem requer 7 caractrísticas:
  1. Encapsulamento (separar o programa em partes, o mais isolado possível).
  2. Herança (derivar uma nova classe de uma classe existente).
  3. Polimorfismo (permite ao desenvolvedor usar o mesmo elemento de formas diferentes).
  4. Abstração (trabalhar um objeto dentro da programação se preocupando somente com suas principais propriedades, sem se apegar a pontos acidentais).
  5. Todos os tipos predefinidos são objetos.
  6. Todos os tipos definidos pelo usuário são objetos.
  7. Todas as operações feitas em objeto são feitas através de métodos.
  
  Por exemplo, Java não é considerada puramente orientada a objeto pois não possui as características 5 e 7.
#### Possui tipos primitivos de dados:
  ~~~~
  int a = 13; 
  System.out.print(a);
  ~~~~
#### É possível se comunicar com objetos sem chamar seus métodos:
  ~~~~
  String s1 = "SMALL" + "TALK" ;
  ~~~~
  
<p>É uma lingaguem dinâmica. Smalltalk, assim como Java, na maioria das vezes é compilada em bytecode para depois ser interpretada por uma máquina virtual (just in time compilation)</p>

Quanto a sua sintaxe, em Smalltalk apenas 6 palavras são de uso reservador true, false, nil, self, super, and thisContext. Como é considerada uma linguagem para crianças, assim como em portugês, na fim de uma linha de código é utilizado o "." (ponto final).Em Smalltalk não é necessário a declaração de tipos de variáveis. A sintaxe em Smalltalk é bem diferente de outras linguagens de programação, utiliza-se sempre a ordem `<objeto recebedor>` `< mensagem >`. Quando você precisar passar várias mensagens para um objeto, é possível utilizar um ponto-e-vírgula separando as chamadas. Como no boco de código abaixo:
  ~~~~
  Transcript clear; show:‘Hello World’; cr.
  ~~~~
Acima é solicitado que o Transcript limpe seu console com o método clear, e com o método show é pedido para imprimir Hello World. Por fim, o método cr serve para pular uma linha (como "\n" em C).

Smalltalk é uma linguagem de propósito geral, sendo assim, ela pode ser usada em diversos tipos de problemas, como: machine
learning, app para desktop, entre outros.

## **Exemplos**

#### *Exemplos Simples*

  ~~~
  Transcript show: 'Hello, world!'
  ~~~
  *O objeto Transcript é utilizado para abrir um console, um lugar onde as mensagens são exibidas.É bastante utilizada para logs e prints de resultados (apenas em formato de texto).*
  ~~~~
  Transcript show: 2+3*8
  ~~~~
  *Em smalltalk as ordens das operações matemáticas são sempre feitas da esquerda para direita, independentemente da ordem de importancia. Sendo assim o resultado da equação acima é 40. Isso acontece pois as operações matemáticas são consideradas mensagens passadas a objetos. E a execução dessas mensagens é sempre feita da esquerda para a direita.*
  
  Exemplo de um jogo para adivinhar o número que o pc escolheu:
  ~~~
Transcript clear. "Limpa o console Transcript"
t :=(1 to: 10000) atRandom .
n:=1. "Atribui um valor para n"
[(a:=(UIManager default request: 'Digite um número') asNumber) = t] "Pede entrada para o usuário, e declara essa entrada como numero, e é o bloco para ser avaliado pelo whileFalse"
whileFalse: [Transcript show: (a<t  ifTrue: ['Meu número é maior']ifFalse: ['Meu número é menor']);cr. n:=n+1 ]. "Loop com a condição atrelada ao bloco anterior, realizando prints com condições atreladas ao bloco a<t"
Transcript show:'O número era: ',t asString ,' e você acertou em: ', n asString, ' tentativas ' "Print"
~~~

   Exemplo de programa para descobrir quantos números primos existem no intervalo selecionado:
  ~~~
  Transcript clear. "Limpa o console Transcript"
  n := (UIManager default request: 'Digite um número') asNumber. "Pede entrada para o usuário, e declara essa entrada como numero"
  a:=0. "Atribui um valor para a"
  Transcript show: 'Até o número ', n asString. "Print"
  [n<=1] "Bloco para ser avaliado pelo whileFalse"
  whileFalse: [n isPrime ifTrue: [a:=a+1].n:=n-1 ]. "Loop com a condição atrelada ao bloco anterior"
  Transcript show: ' existem ', a asString , ' número primos' "Print"
  ~~~
  
## **Funções de Alta Expressividade:**
  ### Linguagem Baseada em Imagens
  Smalltalk é uma linguagem baseada em imagem (image based). Normalmente, as linguagens de programação separam o código estático (funções, definições de classes), do código dinamico, ou em tempo de execução (como objetos). Quando executado, o código é carregado e sempre startado do zero, arquivos de configurações ou fontes de dados são recriados. Qualquer configuração do programa (ou do programador) não é salva e precisa ser reconfigurada a cada reinício.Quando você sai de um programa tradicional e salva e depois inicia de novo, você perde muita informação, como histórico do que foi alterado ou até mesmo a posição do cursor. Ao entrar no Microsoft Word, por exemplo, você vai perceber que vai demorar um pouco, isso acontece pois ele precisa reprocessar muitos conceitos antes de começar a se tornar utilizável.
  
  </p>Uma linguagem de programação baseada em imagem, como Smalltalk, não perde essas informações nem mesmo quando o computador é desligado. Quando se está programando em Smalltalk, o que acontece é que você está modificando uma imagem do ambiente de execução. A maneira que isso funciona é a aplicação pega todo o seu estado (variável, funções, tipos, etc) e salva em um arquivo. Esse arquivo  chamado de arquivo imagem (image file). A próxima vez que o programa inciar, vai ser carregado do estado em que parou.<p>
  Algumas vantagens são: As aplicaçes inciam mais rapidamente, mas é muito bom para o programador pois a implementação, compilação, execução e debugging estão integrados.
  
  E o motivo de não ser tão popular é que esse tipo de programação deixa o código mutio exposto em tempo de execução. Assim como não funciona bem com Sistema de controle de versão que programadores utilizam para salvar código. Támbem não é fácil se relacionar com software externos, pois a maioria das bibliotecas não utilizam as mesmas convenções da linguagens baseada em imagem.
  
  ### Reflexividade
  Em ciência da computação, reflexão computacional é a capacidade de um programa observar ou até mesmo modificar sua estrutura ou comportamento. Tipicamente, o termo se refere à reflexão dinâmica (em tempo de execução), embora muitas linguagens suportem reflexão estática (em tempo de compilação). A reflexão é mais comum em linguagens alto nível, como SmallTalk e menos comum em linguagens de mais baixo nível como o C.
  
  O primeiros computadores eram programados usando a linguagem assembly, que era ineremente reflexivel, pois essas arquiteturas podiam ser programadas usando instruções como dados e usando um código auto-modificável. Enquanto a programação foi indo para um nível mais alto com linguagens como Algol, Cobol, and Fortran, Pascal and C, essa habilidade de reflexão foi perdendo espaço até aparecerem linguagens de programação com reflexão incorporadas em seus sistemas de tipos (como SmallTalk e Java)
 
  Embora sejam grandes os benefícios da reflexão, sobretudo num ambiente onde a flexibilidade seja muito necessária, o uso desse recurso também traz seus problemas. Entre eles está a maior complexidade do código. Como será visto nos exemplos, o código usando reflexão é mais complexo e, portanto, sujeito a erros. Também há a necessidade de um bom controle de versões: pequenas peças de código podem ser compiladas e disponibilizadas separadamente, dessa forma, um bom controle de versões faz-se necessário para garantir a compatibilidade de todo o sistema. Outro problema é a maior susceptibilidade a erros, não só devido a complexidade mas também à ausência de verificação sintática e semântica em tempo de compilação. Por fim, há uma redução do desempenho geral da aplicação, especialmente percebido em sistemas com diversas classes refletidas.
  
  A reflexão pode ser utilizada para auto-otimização ou auto-modificação de um programa. Um sub-componente reflexivo de um programa monitorará a execução e poderá se otimizar ou se modificar de acordo com a função que o programa está resolvendo. Isso pode ser feito modificando a própria área de memória do programa, onde o código está armazenado.

A reflexão pode ser também utilizada para adaptar um determinado sistema dinamicamente à diferentes situações. Considere, por exemplo, uma aplicação que use uma classe X para comunicar-se com algum serviço. Agora, suponha que essa aplicação precise comunicar-se com um serviço diferente, usando a classe Y, que tem nomes de métodos diferentes. Sem reflexão, a aplicação teria de ser modificada e recompilada. Mas, se a reflexão for usada, isso pode ser evitado. A aplicação poderia conhecer os métodos da classe X e essa classe lhe diria que método era usado para que propósito. Assim, quando um novo serviço for usado, via classe Y, a aplicação também procuraria pelos métodos necessários e os utilizaria. Nenhuma modificação no código seria necessária. Nem mesmo o nome da nova classe deveria ser recompilado com a aplicação, uma vez que ele poderia estar armazenado em algum arquivo de configuração, ser verificado e ter a sua classe carregada em tempo de execução.
  
  Nos **Exemplos de Programas** será demonstrado a diferença de acessar uma variável de instância em tempo de execução por meio de reflexão e sem o uso da reflexão utilizando getter/setters.

  Em Smalltalk as varáveis de instância são privadas para a própria instância. Diferentemente de Java e C++, que permitem escolher esse tipo de encapsulamento para esse tipo de variável. Com isso nâo é possível duas instâncias da mesma class conseguir acessar as variáveis de instância uma da outra. Apenas fazendo uso de reflexão ou getter e setter.


## **Bibliografia:**

+ https://hackernoon.com/back-to-the-future-with-smalltalk-57c68fab583a
+ https://insights.stackoverflow.com/survey/2017
+ http://worrydream.com/EarlyHistoryOfSmalltalk/#p12  
+ https://www.codeproject.com/Articles/1241904/Introduction-to-the-Smalltalk-Programming-Language
+ https://programming.dojo.net.nz/languages/smalltalk/index
+ https://www.ajibot.com/blog/image-based-programming
+ http://pharo.gforge.inria.fr/PBE1/PBE1ch15.html
+ https://www.geeksforgeeks.org/java-not-purely-object-oriented-language/
