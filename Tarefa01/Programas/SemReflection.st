
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
  
  valorAltura
    ^altura.
    
  valorAltura: medida
    altura:= medida
    
  valorPeso
    ^peso.
    
  valorPeso: pesagem
    altura:= pesagem
    
    
Transcript clear.
prova1:= CalcIMC new.
prova1 calculoimc .
Transcript cr.
prova1 valorAltura:((UIManager default request: 'Digite o novo valor de altura') asNumber).
prova1 valorPeso:((UIManager default request: 'Digite o novo valor de peso') asNumber).
prova1 calculoimc.


    
