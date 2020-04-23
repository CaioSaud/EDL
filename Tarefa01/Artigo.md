# **Smalltalk**


#### *Introdução*

Smalltalk-80, ou apenas, Smalltalk é a linguagem de programação que tornou popular o paradigma de orientação ao objeto.
A pureza e clareza de Smalltalk sobre esse conceito inspirou quase todas as outras arquiteturas das linguagens OOP.
  
Apesar da primeira versão de Smalltalk ter sido criada em 1971, essa linguagem ficou em segundo lugar na pesquisa da Stack
Overflow "most loved programming language" de 2017. Já em 2018 não ficou nem entre as top 26.

*Sumário*:
  * [Origens e Influências](#Origens-e-Influências)
  * [Classificações](#Classificações)
  * [Exemplos](#Exemplos)
    * [Exemplos Simples](#Exemplos-Simples)
    * [Exemplo Interessante](#Exemplo-Interessante)
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

Quanto a sua sintaxe, em Smalltalk apenas 6 palavras são de uso reservador true, false, nil, self, super, and thisContext. Como é considerada uma linguagem para crianças, assim como em portugês, na fim de uma linha de código é utilizado o "." (ponto final). A sintaxe em Smalltalk é bem diferente de outras linguagens de programação, utiliza-se sempre a ordem `<objeto recebedor>` `< mensagem >`.

Smalltalk é uma linguagem de propósito geral, sendo assim, ela pode ser usada em diversos tipos de problemas, como: machine
learning, app para desktop, entre outros.

## **Exemplos**

#### *Exemplos Simples*

  ~~~
  Transcript show: 'Hello, world!'
  ~~~
  *O comando Transcript é utilizado para abrir a janela Transcript, que é bastante utilizada para logs e prints de resultados (apenas em formato de texto).*
  ~~~~
  Transcript show: 2+3*8
  ~~~~
  *Em smalltalk as ordens das operações matemáticas são sempre feitas da esquerda para direita, independentemente da ordem de importancia. Sendo assim o resultado da equação acima é 40. Isso acontece pois as operações matemáticas são consideradas mensagens passadas a objetos. E a execução dessas mensagens é sempre feita da esquerda para a direita.*
  
  Exemplo de um jogo para adivinhar o número que o pc escolheu:
  ~~~
Transcript clear.
t :=(1 to: 10000) atRandom .
n:=1.
[(a:=(UIManager default request: 'Digite um número') asNumber) = t]
whileFalse: [Transcript show: (a<t  ifTrue: ['Meu número é maior']ifFalse: ['Meu número é menor']);cr. n:=n+1 ].
Transcript show:'O número era: ',t asString ,' e você acertou em: ', n asString, ' tentativas '
~~~

   Exemplo de programa para descobrir quantos números primos existem no intervalo selecionado:
  ~~~
  Transcript clear.
  n := (UIManager default request: 'Digite um número') asNumber.
  a:=0.
  Transcript show: 'Até o número ', n asString.
  [n<=1]
  whileFalse: [n isPrime ifTrue: [a:=a+1].n:=n-1 ].
  Transcript show: ' existem ', a asString , ' número primos'
  ~~~
  
## **Função de alta expressividade:**
  ### Image Based Language
  Smalltalk é uma linguagem baseada em imagem (image based). Normalmente, as linguagens de programação separam o código estático (funções, definições de classes), do código dinamico, ou em tempo de execução (como objetos). Quando executado, o código é carregado e sempre startado do zero, arquivos de configurações ou fontes de dados são recriados. Qualquer configuração do programa (ou do programador) não é salva e precisa ser reconfigurada a cada reinício.Quando você sai de um programa tradicional e salva e depois inicia de novo, você perde muita informação, como histórico do que foi alterado ou até mesmo a posição do cursor. Ao entrar no Microsoft Word, por exemplo, você vai perceber que vai demorar um pouco, isso acontece pois ele precisa reprocessar ou reaprender muitos conceitos antes de começar a se tornar utilizável.
  
  </p>Uma linguagem de programação baseada em imagem, como Smalltalk, não perde essas informações nem mesmo quando o computador é desligado. Quando se está programando em Smalltalk, o que acontece é que você está modificando uma imagem do ambiente de execução. A maneira que isso funciona é a aplicação pega todo o seu estado (variável, funcioções, tipos, etc) e sakva em um arquivo. Esse arquivo  chamado de arquivo imagem (image file). A próxima vez que o programa inciar, vai ser carregado do estado em que parou.<p>
  Algumas vantagens são óbvias como: As aplicaçes inciam mais rapidamente, mas é muito bom para o programador pois a implementação, compilação, execução e debugging estão integrados.
  
  E o motivo de não ser tão popular é que esse tipo de programação deixa o código mutio exposto em tempo de execução. Assim como não funciona bem com Sistema de controle de versão que programadores utilizam para salvar código. Támbem não é fácil se relacionar com software externos, pois a maioria das bibliotecas não utilizam as mesmas convenções da linguagens baseada em imagem.
  
  ### Reflexividade
  Em ciência da computação, reflexão computacional é a capacidade de um programa observar ou até mesmo modificar sua estrutura ou comportamento. Tipicamente, o termo se refere à reflexão dinâmica (em tempo de execução), embora muitas linguagens suportem reflexão estática (em tempo de compilação). A reflexão é mais comum em linguagens alto nível, como SmallTalk e menos comum em linguagens de mais baixo nível como o C.
  
  Embora sejam grandes os benefícios da reflexão, sobretudo num ambiente onde a flexibilidade seja muito necessária, o uso desse recurso também traz seus problemas. Entre eles está a maior complexidade do código. Como será visto nos exemplos, o código usando reflexão é mais complexo e, portanto, sujeito a erros. Também há a necessidade de um bom controle de versões: pequenas peças de código podem ser compiladas e disponibilizadas separadamente, dessa forma, um bom controle de versões faz-se necessário para garantir a compatibilidade de todo o sistema. Outro problema é a maior susceptibilidade a erros, não só devido a complexidade mas também à ausência de verificação sintática e semântica em tempo de compilação. Por fim, há uma redução do desempenho geral da aplicação, especialmente percebido em sistemas com diversas classes refletidas.
  
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
