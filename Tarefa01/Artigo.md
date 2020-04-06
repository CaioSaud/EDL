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
<p>É uma lingaguem dinâmica. Smalltalk, assim como Java, na maioria das vezes é compilada em bytecode para depois ser interpretada por uma máquina virtual (just in time compilation)</p>
Smalltalk é uma linguagem de propósito geral, sendo assim, ela pode ser usada em diversos tipos de problemas, como: machine
learning, app para desktop, entre outros.

## **Exemplos**

#### *Exemplos Simples*

  ~~~
  Transcript show: 'Hello, world!'
  ~~~
  
  ~~~~
  Transcript show: 2+3*8
  ~~~~
  *Em smalltalk as ordens das operações matemáticas são sempre feitas da esquerda para direita, independentemente da ordem de importancia. Sendo assim o resultado da equação acima é 40.*
  
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
  Smalltalk é uma linguagem baseada em imagem (image based). Normalmente, as linguagens de programação separam o código estático (funções, definições de classes), do código dinamico, ou em tempo de execução (como objetos). Quando executado, o código é carregado e sempre startado do zero, arquivos de configurações ou fontes de dados são recriados. Qualquer configuração do programa (ou do programador) não é salva e precisa ser reconfigurada a cada reinício.Quando você sai de um programa tradicional e salva e depois inicia de novo, você perde muita informação, como histórico do que foi alterado ou até mesmo a posição do cursor. Ao entrar no Microsoft Word, por exemplo, você vai perceber que vai demorar um pouco, isso acontece pois ele precisa reprocessar ou reaprender muitos conceitos antes de começar a se tornar utilizável.
  </p>Uma linguagem de programação baseada em imagem, como Smalltalk, não perde essas informações nem mesmo quando o computador é desligado. Quando se está programando em Smalltalk, o que acontece é que você está modificando uma imagem do ambiente de execução. A maneira que isso funciona é a aplicação pega todo o seu estado (variável, funcioções, tipos, etc) e sakva em um arquivo. Esse arquivo  chamado de arquivo imagem (image file). A próxima vez que o programa inciar, vai ser carregado do estado em que parou.<p>
  Algumas vantagens são óbvias como: As aplicaçes inciam mais rapidamente, mas é muito bom para o programador pois a implementação, compilação, execução e debugging estão integrados.
  E o motivo de não ser tão popular é que esse tipo de programação deixa o código mutio exposto em tempo de execução. Assim como não funciona bem com Sistema de controle de versão que programadores utilizam para salvar código. Támbem não é fácil se relacionar com software externos, pois a maioria das bibliotecas não utilizam as mesmas convenções da linguagens baseada em imagem.
  
  

## **Bibliografia:**

+ https://hackernoon.com/back-to-the-future-with-smalltalk-57c68fab583a
+ https://insights.stackoverflow.com/survey/2017
+ http://worrydream.com/EarlyHistoryOfSmalltalk/#p12  
+ https://www.codeproject.com/Articles/1241904/Introduction-to-the-Smalltalk-Programming-Language
+ https://programming.dojo.net.nz/languages/smalltalk/index
+ https://www.ajibot.com/blog/image-based-programming
