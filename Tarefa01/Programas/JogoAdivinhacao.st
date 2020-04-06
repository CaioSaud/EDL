Transcript clear.
t :=(1 to: 10000) atRandom .
n:=1.
[(a:=(UIManager default request: 'Digite um número') asNumber) = t]
whileFalse: [Transcript show: (a<t  ifTrue: ['Meu número é maior']ifFalse: ['Meu número é menor']);cr. n:=n+1 ].
Transcript show:'O número era: ',t asString ,' e você acertou em: ', n asString, ' tentativas '
