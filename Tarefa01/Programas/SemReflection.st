
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
prova1 valorAltura: 2.
prova1 valorPeso: 85.
prova1 calculoimc.


    
