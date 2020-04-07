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
  		Transcript show: (peso/(altura*altura)) asString 
		


Transcript clear.
prova1:= CalcIMC new.
prova1 calculoimc .
Transcript cr.
prova1 instVarNamed: 'peso' put: ((UIManager default request: 'Digite um número') asNumber). "Nesta linha ocorre a alteração de valor
variável de instância peso em tempo de execução, utilizando reflection"
caio instVarNamed: 'altura' put: ((UIManager default request: 'Digite o novo valor de altura') asNumber)."Nesta linha ocorre a alteração de valor
variável de instância altura em tempo de execução, utilizando reflection"
Transcript show: 'Novo cálculo IMC :'.caio calculoimc .
