 error: aString 
"Alterando a mensagem de erro da imagem do Smalltalk"
	^Error new signal:'Você deve digitar um número inteiro maior que zero'

Object subclass: #CalcIMC
	instanceVariableNames: 'peso altura'
	classVariableNames: ''
	package: 'IMC'
  
	initialize 
		peso:= 90.
		altura:=1.90.

	calculoimc
  		(peso/(altura*altura))<0 ifTrue: [Transcript error: 'Validação número negativo'] 
		ifFalse: [Transcript show: ((peso/(altura*altura))asFloat ) asString ]
		


Transcript clear.
prova1:= CalcIMC new.
prova1 calculoimc .
Transcript cr.
prova1 instVarNamed: 'peso' put: ((UIManager default request: 'Digite o novo valor de peso') asNumber). "Nesta linha ocorre a alteração de valor
variável de instância peso em tempo de execução, utilizando reflection"
prova1 instVarNamed: 'altura' put: ((UIManager default request: 'Digite o novo valor de altura') asNumber)."Nesta linha ocorre a alteração de valor
variável de instância altura em tempo de execução, utilizando reflection"
Transcript show: 'Novo cálculo IMC :'.prova1 calculoimc .
