Transcript clear.
n := (UIManager default request: 'Digite o último número número') asNumber.
p := (UIManager default request: 'Digite o primeiro número') asNumber.
a:=0.
Transcript show:'Do número ' , p asString, ' até o número ', n asString.
[n<=p]
whileFalse: [n isPrime ifTrue: [a:=a+1].n:=n-1 ].
Transcript show: ' existem ', a asString , ' número primos'
